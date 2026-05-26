package raft

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"code.cloudfoundry.org/clock"
	"github.com/docker/go-events"
	"github.com/docker/go-metrics"
	"github.com/gogo/protobuf/proto"
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/state"
	"github.com/moby/swarmkit/v2/manager/state/raft/membership"
	"github.com/moby/swarmkit/v2/manager/state/raft/storage"
	"github.com/moby/swarmkit/v2/manager/state/raft/transport"
	"github.com/moby/swarmkit/v2/manager/state/store"
	"github.com/moby/swarmkit/v2/watch"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/pkg/v3/idutil"
	"go.etcd.io/raft/v3"
	"go.etcd.io/raft/v3/raftpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	// ErrNoRaftMember is thrown when the node is not yet part of a raft cluster
	ErrNoRaftMember = errors.New("raft: node is not yet part of a raft cluster")
	// ErrConfChangeRefused is returned when there is an issue with the configuration change
	ErrConfChangeRefused = errors.New("raft: propose configuration change refused")
	// ErrApplyNotSpecified is returned during the creation of a raft node when no apply method was provided
	ErrApplyNotSpecified = errors.New("raft: apply method was not specified")
	// ErrSetHardState is returned when the node fails to set the hard state
	ErrSetHardState = errors.New("raft: failed to set the hard state for log append entry")
	// ErrStopped is returned when an operation was submitted but the node was stopped in the meantime
	ErrStopped = errors.New("raft: failed to process the request: node is stopped")
	// ErrLostLeadership is returned when an operation was submitted but the node lost leader status before it became committed
	ErrLostLeadership = errors.New("raft: failed to process the request: node lost leader status")
	// ErrRequestTooLarge is returned when a raft internal message is too large to be sent
	ErrRequestTooLarge = errors.New("raft: raft message is too large and can't be sent")
	// ErrCannotRemoveMember is thrown when we try to remove a member from the cluster but this would result in a loss of quorum
	ErrCannotRemoveMember = errors.New("raft: member cannot be removed, because removing it may result in loss of quorum")
	// ErrNoClusterLeader is thrown when the cluster has no elected leader
	ErrNoClusterLeader = errors.New("raft: no elected cluster leader")
	// ErrMemberUnknown is sent in response to a message from an
	// unrecognized peer.
	ErrMemberUnknown = errors.New("raft: member unknown")

	// work around lint
	lostQuorumMessage = "The swarm does not have a leader. It's possible that too few managers are online. Make sure more than half of the managers are online."
	errLostQuorum     = errors.New(lostQuorumMessage)

	// Timer to capture ProposeValue() latency.
	proposeLatencyTimer metrics.Timer
)

// LeadershipState indicates whether the node is a leader or follower.
type LeadershipState int

const (
	// IsLeader indicates that the node is a raft leader.
	IsLeader LeadershipState = iota
	// IsFollower indicates that the node is a raft follower.
	IsFollower

	// lostQuorumTimeout is the number of ticks that can elapse with no
	// leader before LeaderConn starts returning an error right away.
	lostQuorumTimeout = 10
)

// EncryptionKeys are the current and, if necessary, pending DEKs with which to
// encrypt raft data
type EncryptionKeys struct {
	CurrentDEK []byte
	PendingDEK []byte
}

// EncryptionKeyRotator is an interface to find out if any keys need rotating.
type EncryptionKeyRotator interface {
	GetKeys() EncryptionKeys
	UpdateKeys(EncryptionKeys) error
	NeedsRotation() bool
	RotationNotify() chan struct{}
}

// Node represents the Raft Node useful
// configuration.
type Node struct {
	raftNode  raft.Node
	cluster   *membership.Cluster
	transport *transport.Transport

	raftStore           *raft.MemoryStorage
	memoryStore         *store.MemoryStore
	Config              *raft.Config
	opts                NodeOptions
	reqIDGen            *idutil.Generator
	wait                *wait
	campaignWhenAble    bool
	signalledLeadership uint32
	isMember            uint32
	bootstrapMembers    []*api.RaftMember

	// waitProp waits for all the proposals to be terminated before
	// shutting down the node.
	waitProp sync.WaitGroup

	confState       raftpb.ConfState
	appliedIndex    uint64
	snapshotMeta    raftpb.SnapshotMetadata
	writtenWALIndex uint64

	ticker clock.Ticker
	doneCh chan struct{}
	// RemovedFromRaft notifies about node deletion from raft cluster
	RemovedFromRaft chan struct{}
	cancelFunc      func()

	removeRaftOnce      sync.Once
	leadershipBroadcast *watch.Queue

	// used to coordinate shutdown
	// Lock should be used only in stop(), all other functions should use RLock.
	stopMu sync.RWMutex
	// used for membership management checks
	membershipLock sync.Mutex
	// synchronizes access to n.opts.Addr, and makes sure the address is not
	// updated concurrently with JoinAndStart.
	addrLock sync.Mutex

	snapshotInProgress chan raftpb.SnapshotMetadata
	asyncTasks         sync.WaitGroup

	// stopped chan is used for notifying grpc handlers that raft node going
	// to stop.
	stopped chan struct{}

	raftLogger     *storage.EncryptedRaftLogger
	keyRotator     EncryptionKeyRotator
	rotationQueued bool
	clearData      bool

	// waitForAppliedIndex stores the index of the last log that was written using
	// an raft DEK during a raft DEK rotation, so that we won't finish a rotation until
	// a snapshot covering that index has been written encrypted with the new raft DEK
	waitForAppliedIndex uint64
	ticksWithNoLeader   uint32
}

// NodeOptions provides node-level options.
type NodeOptions struct {
	// ID is the node's ID, from its certificate's CN field.
	ID string
	// Addr is the address of this node's listener
	Addr string
	// ForceNewCluster defines if we have to force a new cluster
	// because we are recovering from a backup data directory.
	ForceNewCluster bool
	// JoinAddr is the cluster to join. May be an empty string to create
	// a standalone cluster.
	JoinAddr string
	// ForceJoin tells us to join even if already part of a cluster.
	ForceJoin bool
	// Config is the raft config.
	Config *raft.Config
	// StateDir is the directory to store durable state.
	StateDir string
	// TickInterval interval is the time interval between raft ticks.
	TickInterval time.Duration
	// ClockSource is a Clock interface to use as a time base.
	// Leave this nil except for tests that are designed not to run in real
	// time.
	ClockSource clock.Clock
	// SendTimeout is the timeout on the sending messages to other raft
	// nodes. Leave this as 0 to get the default value.
	SendTimeout    time.Duration
	TLSCredentials credentials.TransportCredentials
	KeyRotator     EncryptionKeyRotator
	// DisableStackDump prevents Run from dumping goroutine stacks when the
	// store becomes stuck.
	DisableStackDump bool

	// FIPS specifies whether the raft encryption should be FIPS compliant
	FIPS bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
	ns := metrics.NewNamespace("swarm", "raft", nil)
	proposeLatencyTimer = ns.NewTimer("transaction_latency", "Raft transaction latency.")
	metrics.Register(ns)
}

// NewNode generates a new Raft node
func NewNode(opts NodeOptions) *Node { _ = "STUB: not implemented"; return nil }

// IsIDRemoved reports if member with id was removed from cluster.
// Part of transport.Raft interface.
func (n *Node) IsIDRemoved(id uint64) bool { _ = "STUB: not implemented"; return false }

// NodeRemoved signals that node was removed from cluster and should stop.
// Part of transport.Raft interface.
func (n *Node) NodeRemoved() { _ = "STUB: not implemented"; return }

// ReportSnapshot reports snapshot status to underlying raft node.
// Part of transport.Raft interface.
func (n *Node) ReportSnapshot(id uint64, status raft.SnapshotStatus) {
	_ = "STUB: not implemented"
	return
}

// ReportUnreachable reports to underlying raft node that member with id is
// unreachable.
// Part of transport.Raft interface.
func (n *Node) ReportUnreachable(id uint64) { _ = "STUB: not implemented"; return }

// SetAddr provides the raft node's address. This can be used in cases where
// opts.Addr was not provided to NewNode, for example when a port was not bound
// until after the raft node was created.
func (n *Node) SetAddr(ctx context.Context, addr string) error {
	_ = "STUB: not implemented"
	return nil
}

// If the raft node is running, submit a configuration change
// with the new address.

// TODO(aaronl): Currently, this node must be the leader to
// submit this configuration change. This works for the initial
// use cases (single-node cluster late binding ports, or calling
// SetAddr before joining a cluster). In the future, we may want
// to support having a follower proactively change its remote
// address.

// WithContext returns context which is cancelled when parent context cancelled
// or node is stopped.
func (n *Node) WithContext(ctx context.Context) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

func (n *Node) initTransport() { _ = "STUB: not implemented"; return }

// JoinAndStart joins and starts the raft server
func (n *Node) JoinAndStart(ctx context.Context) (err error) { _ = "STUB: not implemented"; return nil }

// to shutdown transport

// Snapshot never returns an error

// lastIndex always returns nil as an error

// override the module field entirely, since etcd/raft is not exactly a submodule

// restore from snapshot

// First member in the cluster, self-assign ID

// join to existing cluster

func (n *Node) joinCluster(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// DefaultNodeConfig returns the default config for a
// raft node that can be modified and customized
func DefaultNodeConfig() *raft.Config { _ = "STUB: not implemented"; return nil }

// Recommended value in etcd/raft is 10 x (HeartbeatTick).
// Lower values were seen to have caused instability because of
// frequent leader elections when running on flakey networks.

// DefaultRaftConfig returns a default api.RaftConfig.
func DefaultRaftConfig() api.RaftConfig { _ = "STUB: not implemented"; return *new(api.RaftConfig) }

// Recommended value in etcd/raft is 10 x (HeartbeatTick).
// Lower values were seen to have caused instability because of
// frequent leader elections when running on flakey networks.

// MemoryStore returns the memory store that is kept in sync with the raft log.
func (n *Node) MemoryStore() *store.MemoryStore { _ = "STUB: not implemented"; return nil }

func (n *Node) done() { _ = "STUB: not implemented"; return }

// ClearData tells the raft node to delete its WALs, snapshots, and keys on
// shutdown.
func (n *Node) ClearData() {
	_ = "STUB: not implemented"

	// Run is the main loop for a Raft node, it goes along the state machine,
	// acting on the messages received from other Raft nodes in the cluster.
	//
	// Before running the main loop, it first starts the raft node based on saved
	// cluster state. If no saved state exists, it starts a single-node cluster.
	return
}

func (n *Node) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Delete WAL and snapshots, since they are no longer
// usable.

// clear out the DEKs

// Flag that indicates if this manager node is *currently* the raft leader.

// Save entries to storage

// If the memory store lock has been held for too long,
// transferring leadership is an easy way to break out of it.

// if the message is a snapshot, before we send it, we should
// overwrite the original ConfState from the snapshot with the
// current one

// Send raft messages to peers

// Apply snapshot to memory store. The snapshot
// was applied to the raft store in
// saveToStorage.

// Load the snapshot data into the store

// If we cease to be the leader, we must cancel any
// proposals that are currently waiting for a quorum to
// acknowledge them. It is still possible for these to
// become committed, but if that happens we will apply
// them as any follower would.

// It is important that we cancel these proposals before
// calling processCommitted, so processCommitted does
// not deadlock.

// It is important that we set n.signalledLeadership to 0
// before calling n.wait.cancelAll. When a new raft
// request is registered, it checks n.signalledLeadership
// afterwards, and cancels the registration if it is 0.
// If cancelAll was called first, this call might run
// before the new request registers, but
// signalledLeadership would be set after the check.
// Setting signalledLeadership before calling cancelAll
// ensures that if a new request is registered during
// this transition, it will either be cancelled by
// cancelAll, or by its own check of signalledLeadership.

// Node just became a leader.

// Process committed entries

// in case the previous attempt to update the key failed

// Trigger a snapshot every once in awhile

// If all the entries in the log have become
// committed, broadcast our leadership status.

// Advance the state machine

// On the first startup, or if we are the only
// registered member after restoring from the state,
// campaign to be the leader.

// there was a key rotation that took place before while the snapshot
// was in progress - we have to take another snapshot and encrypt with the new key

// There are 2 separate checks:  rotationQueued, and n.needsSnapshot().
// We set rotationQueued so that when we are notified of a rotation, we try to
// do a snapshot as soon as possible.  However, if there is an error while doing
// the snapshot, we don't want to hammer the node attempting to do snapshots over
// and over.  So if doing a snapshot fails, wait until the next entry comes in to
// try again.

func (n *Node) restoreFromSnapshot(ctx context.Context, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *Node) needsSnapshot(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// we want to wait for the last index written with the old DEK to be committed, else a snapshot taken
// may have an index less than the index of a WAL written with an old DEK.  We want the next snapshot
// written with the new key to supercede any WAL written with an old DEK.

// if there is already a snapshot at this index or higher, bump the wait index up to 1 higher than the current
// snapshot index, because the rotation cannot be completed until the next snapshot

func (n *Node) maybeMarkRotationFinished(ctx context.Context) { _ = "STUB: not implemented"; return }

// this means we tried to rotate - so finish the rotation

func (n *Node) getCurrentRaftConfig() api.RaftConfig {
	_ = "STUB: not implemented"
	return *new(api.RaftConfig)
}

// Cancel interrupts all ongoing proposals, and prevents new ones from
// starting. This is useful for the shutdown sequence because it allows
// the manager to shut down raft-dependent services that might otherwise
// block on shutdown if quorum isn't met. Then the raft node can be completely
// shut down once no more code is using it.
func (n *Node) Cancel() {
	_ = "STUB: not implemented"

	// Done returns channel which is closed when raft node is fully stopped.
	return
}

func (n *Node) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (n *Node) stop(ctx context.Context) { _ = "STUB: not implemented"; return }

// TODO(stevvooe): Handle ctx.Done()

// isLeader checks if we are the leader or not, without the protection of lock
func (n *Node) isLeader() bool { _ = "STUB: not implemented"; return false }

// IsLeader checks if we are the leader or not, with the protection of lock
func (n *Node) IsLeader() bool { _ = "STUB: not implemented"; return false }

// leader returns the id of the leader, without the protection of lock and
// membership check, so it's caller task.
func (n *Node) leader() uint64 { _ = "STUB: not implemented"; return 0 }

// Leader returns the id of the leader, with the protection of lock
func (n *Node) Leader() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadyForProposals returns true if the node has broadcasted a message
// saying that it has become the leader. This means it is ready to accept
// proposals.
func (n *Node) ReadyForProposals() bool { _ = "STUB: not implemented"; return false }

func (n *Node) caughtUp() bool {
	_ = "STUB: not implemented"
	// obnoxious function that always returns a nil error
	return false
}

// Join asks to a member of the raft to propose
// a configuration change and add us as a member thus
// beginning the log replication process. This method
// is called from an aspiring member to an existing member
func (n *Node) Join(ctx context.Context, req *api.JoinRequest) (*api.JoinResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// can't stop the raft node while an async RPC is in progress

// If the joining node sent an address like 0.0.0.0:4242, automatically
// determine its actual address based on the GRPC connection. This
// avoids the need for a prospective member to know its own address.

// We do not bother submitting a configuration change for the
// new member if we can't contact it back using its address

// If the peer is already a member of the cluster, we will only update
// its information, not add it as a new member. Adding it again would
// cause the quorum to be computed incorrectly.

// Find a unique ID for the joining member.

func (n *Node) joinResponse(raftID uint64) *api.JoinResponse { _ = "STUB: not implemented"; return nil }

// checkHealth tries to contact an aspiring member through its advertised address
// and checks if its raft server is running.
func (n *Node) checkHealth(ctx context.Context, addr string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// addMember submits a configuration change to add a new member on the raft cluster.
func (n *Node) addMember(ctx context.Context, addr string, raftID uint64, nodeID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Wait for a raft round to process the configuration change

// updateNodeBlocking runs synchronous job to update node address in whole cluster.
func (n *Node) updateNodeBlocking(ctx context.Context, id uint64, addr string) error {
	_ = "STUB: not implemented"
	return nil
}

// Wait for a raft round to process the configuration change

// UpdateNode submits a configuration change to change a member's address.
func (n *Node) UpdateNode(id uint64, addr string) { _ = "STUB: not implemented"; return }

// spawn updating info in raft in background to unblock transport

// Leave asks to a member of the raft to remove
// us from the raft cluster. This method is called
// from a member who is willing to leave its raft
// membership to an active member of the raft
func (n *Node) Leave(ctx context.Context, req *api.LeaveRequest) (*api.LeaveResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CanRemoveMember checks if a member can be removed from
// the context of the current node.
func (n *Node) CanRemoveMember(id uint64) bool { _ = "STUB: not implemented"; return false }

// reachable managers after removal

// Local node from where the remove is issued

func (n *Node) removeMember(ctx context.Context, id uint64) error {
	_ = "STUB: not implemented"
	// can't stop the raft node while an async RPC is in progress
	return nil
}

// TransferLeadership attempts to transfer leadership to a different node,
// and wait for the transfer to happen.
func (n *Node) TransferLeadership(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// RemoveMember submits a configuration change to remove a member from the raft cluster
// after checking if the operation would not result in a loss of quorum.
func (n *Node) RemoveMember(ctx context.Context, id uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// processRaftMessageLogger is used to lazily create a logger for
// ProcessRaftMessage. Usually nothing will be logged, so it is useful to avoid
// formatting strings and allocating a logger when it won't be used.
func (n *Node) processRaftMessageLogger(ctx context.Context, msg *api.ProcessRaftMessageRequest) *logrus.Entry {
	_ = "STUB: not implemented"
	return nil
}

//nolint:unused // currently unused, but should be used again; see TODO in Node.ProcessRaftMessage
func (n *Node) reportNewAddress(ctx context.Context, id uint64) error {
	_ = "STUB: not implemented"
	// too early
	return nil
}

// Don't know the address of the peer yet, so can't report an
// update.

// StreamRaftMessage is the server endpoint for streaming Raft messages.
// It accepts a stream of raft messages to be processed on this raft member,
// returning a StreamRaftMessageResponse when processing of the streamed
// messages is complete.
// It is called from the Raft leader, which uses it to stream messages
// to this raft member.
// A single stream corresponds to a single raft message,
// which may be disassembled and streamed by the sender
// as individual messages. Therefore, each of the messages
// received by the stream will have the same raft message type and index.
// Currently, only messages of type raftpb.MsgSnap can be disassembled, sent
// and received on the stream.
func (n *Node) StreamRaftMessage(stream api.Raft_StreamRaftMessageServer) error {
	_ = "STUB: not implemented"
	// recvdMsg is the current messasge received from the stream.
	// assembledMessage is where the data from recvdMsg is appended to.
	return nil
}

// First message index.

// Initialized the message to be used for assembling
// the raft message.

// For all message types except raftpb.MsgSnap,
// we don't expect more than a single message
// on the stream so we'll get an EOF on the next Recv()
// and go on to process the received message.

// Verify raft message index.

// Verify that multiple message received on a stream
// can only be of type raftpb.MsgSnap.

// Append the received snapshot data.

// We should have the complete snapshot. Verify and process.

// Translate the response of ProcessRaftMessage() from
// ProcessRaftMessageResponse to StreamRaftMessageResponse if needed.

// ProcessRaftMessage calls 'Step' which advances the
// raft state machine with the provided message on the
// receiving node
func (n *Node) ProcessRaftMessage(ctx context.Context, msg *api.ProcessRaftMessageRequest) (*api.ProcessRaftMessageResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Don't process the message if this comes from
// a node in the remove set

// TODO(aaronl): Address changes are temporarily disabled.
// See https://github.com/docker/docker/issues/30455.
// This should be reenabled in the future with additional
// safeguards (perhaps storing multiple addresses per node).
// if err := n.reportNewAddress(ctx, msg.Message.From); err != nil {
//	log.G(ctx).WithError(err).Errorf("failed to report new address of %x to transport", msg.Message.From)
// }

// Reject vote requests from unreachable peers

// We don't accept forwarded proposals. Our
// current architecture depends on only the leader
// making proposals, so in-flight proposals can be
// guaranteed not to conflict.

// can't stop the raft node while an async RPC is in progress

// ResolveAddress returns the address reaching for a given node ID.
func (n *Node) ResolveAddress(ctx context.Context, msg *api.ResolveAddressRequest) (*api.ResolveAddressResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *Node) getLeaderConn() (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LeaderConn returns current connection to cluster leader or raftselector.ErrIsLeader
// if current machine is leader.
func (n *Node) LeaderConn(ctx context.Context) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// registerNode registers a new node on the cluster memberlist
func (n *Node) registerNode(node *api.RaftMember) error { _ = "STUB: not implemented"; return nil }

// Member already exists

// If the address is different from what we thought it was,
// update it. This can happen if we just joined a cluster
// and are adding ourself now with the remotely-reachable
// address.

// Avoid opening a connection to the local node

// ProposeValue calls Propose on the underlying raft library(etcd/raft) and waits
// on the commit log action before returning a result
func (n *Node) ProposeValue(ctx context.Context, storeAction []api.StoreAction, cb func()) error {
	_ = "STUB: not implemented"
	return nil
}

// GetVersion returns the sequence information for the current raft round.
func (n *Node) GetVersion() *api.Version { _ = "STUB: not implemented"; return nil }

// ChangesBetween returns the changes starting after "from", up to and
// including "to". If these changes are not available because the log
// has been compacted, an error will be returned.
func (n *Node) ChangesBetween(from, to api.Version) ([]state.Change, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// never returns error

// SubscribePeers subscribes to peer updates in cluster. It sends always full
// list of peers.
func (n *Node) SubscribePeers() (q chan events.Event, cancel func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetMemberlist returns the current list of raft members in the cluster.
func (n *Node) GetMemberlist() map[uint64]*api.RaftMember { _ = "STUB: not implemented"; return nil }

// Status returns status of underlying etcd.Node.
func (n *Node) Status() raft.Status { _ = "STUB: not implemented"; return *new(raft.Status) }

// GetMemberByNodeID returns member information based
// on its generic Node ID.
func (n *Node) GetMemberByNodeID(nodeID string) *membership.Member {
	_ = "STUB: not implemented"
	return nil
}

// GetNodeIDByRaftID returns the generic Node ID of a member given its raft ID.
// It returns ErrMemberUnknown if the raft ID is unknown.
func (n *Node) GetNodeIDByRaftID(raftID uint64) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// this is the only possible error value that should be returned; the
// manager code depends on this. if you need to add more errors later, make
// sure that you update the callers of this method accordingly

// IsMember checks if the raft node has effectively joined
// a cluster of existing members.
func (n *Node) IsMember() bool { _ = "STUB: not implemented"; return false }

// Saves a log entry to our Store
func (n *Node) saveToStorage(
	ctx context.Context,
	raftConfig *api.RaftConfig,
	hardState raftpb.HardState,
	entries []raftpb.Entry,
	snapshot raftpb.Snapshot,
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// processInternalRaftRequest proposes a value to be appended to the raft log.
// It calls Propose() on etcd/raft, which calls back into the raft FSM,
// which then sends a message to each of the participating nodes
// in the raft group to apply a log entry and then waits for it to be applied
// on this node. It will block until the this node:
// 1. Gets the necessary replies back from the participating nodes and also performs the commit itself, or
// 2. There is an error, or
// 3. Until the raft node finalizes all the proposals on node shutdown.
func (n *Node) processInternalRaftRequest(ctx context.Context, r *api.InternalRaftRequest, cb func()) (proto.Message, error) {
	_ = "STUB: not implemented"
	return *new(proto.Message), nil
}

// This must be derived from the context which is cancelled by stop()
// to avoid a deadlock on shutdown.

// Do this check after calling register to avoid a race.

// Wait notification channel was closed. This should only happen if the wait was cancelled.

// If we can read from the channel, wait item was triggered. Otherwise it was cancelled.

// if channel is closed, wait item was canceled, otherwise it was triggered

// configure sends a configuration change through consensus and
// then waits for it to be applied to the server. It will block
// until the change is performed or there is an error.
func (n *Node) configure(ctx context.Context, cc raftpb.ConfChange) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *Node) processCommitted(ctx context.Context, entry raftpb.Entry) error {
	_ = "STUB: not implemented"
	// Process a normal entry
	return nil
}

// Process a configuration change (add/remove node)

func (n *Node) processEntry(ctx context.Context, entry raftpb.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// There was no wait on this ID, meaning we don't have a
// transaction in progress that would be committed to the
// memory store by the "trigger" call. This could mean that:
// 1. Startup is in progress, and the raft WAL is being parsed,
// processed and applied to the store, or
// 2. Either a different node wrote this to raft,
// or we wrote it before losing the leader
// position and cancelling the transaction. This entry still needs
// to be committed since other nodes have already committed it.
// Create a new transaction to commit this entry.

// It should not be possible for processInternalRaftRequest
// to be running in this situation, but out of caution we
// cancel any current invocations to avoid a deadlock.
// TODO(anshul) This call is likely redundant, remove after consideration.

func (n *Node) processConfChange(ctx context.Context, entry raftpb.Entry) {
	_ = "STUB: not implemented"
	return
}

// applyAddNode is called when we receive a ConfChange
// from a member in the raft cluster, this adds a new
// node to the existing raft cluster
func (n *Node) applyAddNode(cc raftpb.ConfChange) error { _ = "STUB: not implemented"; return nil }

// ID must be non zero

// applyUpdateNode is called when we receive a ConfChange from a member in the
// raft cluster which update the address of an existing node.
func (n *Node) applyUpdateNode(_ context.Context, cc raftpb.ConfChange) error {
	_ = "STUB: not implemented"
	return nil
}

// applyRemoveNode is called when we receive a ConfChange
// from a member in the raft cluster, this removes a node
// from the existing raft cluster
func (n *Node) applyRemoveNode(ctx context.Context, cc raftpb.ConfChange) (err error) {
	_ = "STUB: not implemented"
	// If the node from where the remove is issued is
	// a follower and the leader steps down, Campaign
	// to be the leader.
	return nil
}

// wait for the commit ack to be sent before closing connection

// SubscribeLeadership returns channel to which events about leadership change
// will be sent in form of raft.LeadershipState. Also cancel func is returned -
// it should be called when listener is no longer interested in events.
func (n *Node) SubscribeLeadership() (q chan events.Event, cancel func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

// createConfigChangeEnts creates a series of Raft entries (i.e.
// EntryConfChange) to remove the set of given IDs from the cluster. The ID
// `self` is _not_ removed, even if present in the set.
// If `self` is not inside the given ids, it creates a Raft entry to add a
// default member with the given `self`.
func createConfigChangeEnts(ids []uint64, self uint64, term, index uint64) []raftpb.Entry {
	_ = "STUB: not implemented"
	return nil
}

// getIDs returns an ordered set of IDs included in the given snapshot and
// the entries. The given snapshot/entries can contain two kinds of
// ID-related entry:
// - ConfChangeAddNode, in which case the contained ID will be added into the set.
// - ConfChangeRemoveNode, in which case the contained ID will be removed from the set.
func getIDs(snap *raftpb.Snapshot, ents []raftpb.Entry) []uint64 {
	_ = "STUB: not implemented"
	return nil
}

// do nothing

func (n *Node) reqTimeout() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// stackDump outputs the runtime stack to os.StdErr.
//
// It is based on Moby's stack.Dump(); https://github.com/moby/moby/blob/471fd27709777d2cce3251129887e14e8bb2e0c7/pkg/stack/stackdump.go#L41-L57
func stackDump() { _ = "STUB: not implemented"; return }
