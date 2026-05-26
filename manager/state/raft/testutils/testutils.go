package testutils

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

	"code.cloudfoundry.org/clock/fakeclock"
	"github.com/docker/go-events"
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/ca"
	cautils "github.com/moby/swarmkit/v2/ca/testutils"
	"github.com/moby/swarmkit/v2/manager/state/raft"
	"github.com/moby/swarmkit/v2/manager/state/store"
	etcdraft "go.etcd.io/raft/v3"
	"go.etcd.io/raft/v3/raftpb"
	"google.golang.org/grpc"
)

// TestNode represents a raft test node
type TestNode struct {
	*raft.Node
	Server         *grpc.Server
	Listener       *WrappedListener
	SecurityConfig *ca.SecurityConfig
	Address        string
	StateDir       string
	cancel         context.CancelFunc
	KeyRotator     *SimpleKeyRotator
}

// Leader is wrapper around real Leader method to suppress error.
// TODO: tests should use Leader method directly.
func (n *TestNode) Leader() uint64 { _ = "STUB: not implemented"; return 0 }

func (n *TestNode) Status() etcdraft.Status {
	_ = "STUB: not implemented"
	return *new(etcdraft.Status)
}

func (n *TestNode) Run(ctx context.Context) { _ = "STUB: not implemented"; return }

func (n *TestNode) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (n *TestNode) ProposeValue(ctx context.Context, actions []api.StoreAction, cb func()) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *TestNode) MemoryStore() *store.MemoryStore { _ = "STUB: not implemented"; return nil }

func (n *TestNode) ReadyForProposals() bool { _ = "STUB: not implemented"; return false }

func (n *TestNode) GetMemberlist() map[uint64]*api.RaftMember {
	_ = "STUB: not implemented"
	return nil
}

func (n *TestNode) SubscribeLeadership() (q chan events.Event, cancel func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AdvanceTicks advances the raft state machine fake clock
func AdvanceTicks(clockSource *fakeclock.FakeClock, ticks int) {
	_ = "STUB: not implemented"
	// A FakeClock timer won't fire multiple times if time is advanced
	// more than its interval.
	return
}

// WaitForCluster waits until leader will be one of specified nodes
func WaitForCluster(t *testing.T, clockSource *fakeclock.FakeClock, nodes map[uint64]*TestNode) {
	_ = "STUB: not implemented"
	return
}

// Don't raise error just because test machine is running slowly

// WaitForPeerNumber waits until peers in cluster converge to specified number
func WaitForPeerNumber(t *testing.T, clockSource *fakeclock.FakeClock, nodes map[uint64]*TestNode, count int) {
	_ = "STUB: not implemented"
	return
}

// WrappedListener disables the Close method to make it possible to reuse a
// socket. close must be called to release the socket.
type WrappedListener struct {
	net.Listener
	acceptConn chan net.Conn
	acceptErr  chan error
	closed     chan struct{}
}

// NewWrappedListener creates a new wrapped listener to register the raft server
func NewWrappedListener(l net.Listener) *WrappedListener { _ = "STUB: not implemented"; return nil }

// grpc closes multiple times

// Accept connections

// Accept accepts new connections on a wrapped listener
func (l *WrappedListener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	// closure must take precedence over taking a connection
	// from the channel
	return *new(net.Conn), nil
}

// Close notifies that the listener can't accept any more connections
func (l *WrappedListener) Close() error { _ = "STUB: not implemented"; return nil }

// CloseListener closes the underlying listener
func (l *WrappedListener) CloseListener() error { _ = "STUB: not implemented"; return nil }

// RecycleWrappedListener creates a new wrappedListener that uses the same
// listening socket as the supplied wrappedListener.
func RecycleWrappedListener(old *WrappedListener) *WrappedListener {
	_ = "STUB: not implemented"
	return nil
}

// grpc closes multiple times

// SimpleKeyRotator does some DEK rotation
type SimpleKeyRotator struct {
	mu                 sync.Mutex
	rotateCh           chan struct{}
	updateFunc         func() error
	overrideNeedRotate *bool
	raft.EncryptionKeys
}

// GetKeys returns the current set of keys
func (s *SimpleKeyRotator) GetKeys() raft.EncryptionKeys {
	_ = "STUB: not implemented"
	return *new(raft.EncryptionKeys)
}

// NeedsRotation returns whether we need to rotate
func (s *SimpleKeyRotator) NeedsRotation() bool { _ = "STUB: not implemented"; return false }

// UpdateKeys updates the current encryption keys
func (s *SimpleKeyRotator) UpdateKeys(newKeys raft.EncryptionKeys) error {
	_ = "STUB: not implemented"
	return nil
}

// RotationNotify returns the rotation notification channel
func (s *SimpleKeyRotator) RotationNotify() chan struct{} {
	_ = "STUB: not implemented"

	// QueuePendingKey lets us rotate the key
	return nil
}

func (s *SimpleKeyRotator) QueuePendingKey(key []byte) { _ = "STUB: not implemented"; return }

// SetUpdateFunc enables you to inject an error when updating keys
func (s *SimpleKeyRotator) SetUpdateFunc(updateFunc func() error) {
	_ = "STUB: not implemented"
	return
}

// SetNeedsRotation enables you to inject a value for NeedsRotation
func (s *SimpleKeyRotator) SetNeedsRotation(override *bool) { _ = "STUB: not implemented"; return }

// NewSimpleKeyRotator returns a basic EncryptionKeyRotator
func NewSimpleKeyRotator(keys raft.EncryptionKeys) *SimpleKeyRotator {
	_ = "STUB: not implemented"
	return nil
}

var _ raft.EncryptionKeyRotator = NewSimpleKeyRotator(raft.EncryptionKeys{})

// NewNode creates a new raft node to use for tests
func NewNode(t *testing.T, clockSource *fakeclock.FakeClock, tc *cautils.TestCA, opts ...raft.NodeOptions) *TestNode {
	_ = "STUB: not implemented"
	return nil
}

// NewInitNode creates a new raft node initiating the cluster
// for other members to join
func NewInitNode(t *testing.T, tc *cautils.TestCA, raftConfig *api.RaftConfig, opts ...raft.NodeOptions) (*TestNode, *fakeclock.FakeClock) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Wait for the node to become the leader.

// NewJoinNode creates a new raft node joining an existing cluster
func NewJoinNode(t *testing.T, clockSource *fakeclock.FakeClock, join string, tc *cautils.TestCA, opts ...raft.NodeOptions) *TestNode {
	_ = "STUB: not implemented"
	return nil
}

// CopyNode returns a copy of a node
func CopyNode(_ *testing.T, clockSource *fakeclock.FakeClock, oldNode *TestNode, forceNewCluster bool, kr *SimpleKeyRotator) (*TestNode, context.Context) {
	_ = "STUB: not implemented"
	return nil, *new(context.Context)
}

// RestartNode restarts a raft test node
func RestartNode(t *testing.T, clockSource *fakeclock.FakeClock, oldNode *TestNode, forceNewCluster bool) *TestNode {
	_ = "STUB: not implemented"
	return nil
}

// NewRaftCluster creates a new raft cluster with 3 nodes for testing
func NewRaftCluster(t *testing.T, tc *cautils.TestCA, config ...*api.RaftConfig) (map[uint64]*TestNode, *fakeclock.FakeClock) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddRaftNode adds an additional raft test node to an existing cluster
func AddRaftNode(t *testing.T, clockSource *fakeclock.FakeClock, nodes map[uint64]*TestNode, tc *cautils.TestCA, opts ...raft.NodeOptions) {
	_ = "STUB: not implemented"
	return
}

// TeardownCluster destroys a raft cluster used for tests
func TeardownCluster(nodes map[uint64]*TestNode) { _ = "STUB: not implemented"; return }

// ShutdownNode shuts down a raft test node and deletes the content
// of the state directory
func ShutdownNode(node *TestNode) { _ = "STUB: not implemented"; return }

// ShutdownRaft shutdowns only raft part of node.
func (n *TestNode) ShutdownRaft() { _ = "STUB: not implemented"; return }

// CleanupNonRunningNode frees resources associated with a node which is not
// running.
func CleanupNonRunningNode(node *TestNode) { _ = "STUB: not implemented"; return }

// Leader determines who is the leader amongst a set of raft nodes
// belonging to the same cluster
func Leader(nodes map[uint64]*TestNode) *TestNode { _ = "STUB: not implemented"; return nil }

// ProposeValue proposes a value to a raft test cluster
func ProposeValue(t *testing.T, raftNode *TestNode, time time.Duration, nodeID ...string) (*api.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CheckValue checks that the value has been propagated between raft members
func CheckValue(t *testing.T, clockSource *fakeclock.FakeClock, raftNode *TestNode, createdNode *api.Node) {
	_ = "STUB: not implemented"
	return
}

// CheckNoValue checks that there is no value replicated on nodes, generally
// used to test the absence of a leader
func CheckNoValue(t *testing.T, clockSource *fakeclock.FakeClock, raftNode *TestNode) {
	_ = "STUB: not implemented"
	return
}

// CheckValuesOnNodes checks that all the nodes in the cluster have the same
// replicated data, generally used to check if a node can catch up with the logs
// correctly
func CheckValuesOnNodes(t *testing.T, clockSource *fakeclock.FakeClock, checkNodes map[uint64]*TestNode, ids []string, values []*api.Node) {
	_ = "STUB: not implemented"
	return
}

// GetAllValuesOnNode returns all values on this node
func GetAllValuesOnNode(t *testing.T, clockSource *fakeclock.FakeClock, raftNode *TestNode) ([]string, []*api.Node) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewSnapshotMessage creates and returns a raftpb.Message of type MsgSnap
// where the snapshot data is of the given size and the value of each byte
// is (index of the byte) % 256.
func NewSnapshotMessage(from, to uint64, size int) *raftpb.Message {
	_ = "STUB: not implemented"
	return nil
}

// Include the snapshot size in the Index field for testing.

// VerifySnapshot verifies that the snapshot data where each byte is
// of the value (index % sizeof(byte)).
func VerifySnapshot(raftMsg *raftpb.Message) bool { _ = "STUB: not implemented"; return false }
