// Package transport provides grpc transport layer for raft.
// All methods are non-blocking.
package transport

import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/pkg/errors"
	"go.etcd.io/raft/v3"
	"go.etcd.io/raft/v3/raftpb"
)

// ErrIsNotFound indicates that peer was never added to transport.
var ErrIsNotFound = errors.New("peer not found")

// Raft is interface which represents Raft API for transport package.
type Raft interface {
	ReportUnreachable(id uint64)
	ReportSnapshot(id uint64, status raft.SnapshotStatus)
	IsIDRemoved(id uint64) bool
	UpdateNode(id uint64, addr string)

	NodeRemoved()
}

// Config for Transport
type Config struct {
	HeartbeatInterval time.Duration
	SendTimeout       time.Duration
	Credentials       credentials.TransportCredentials
	RaftID            string

	Raft
}

// Transport is structure which manages remote raft peers and sends messages
// to them.
type Transport struct {
	config *Config

	unknownc chan raftpb.Message

	mu      sync.Mutex
	peers   map[uint64]*peer
	stopped bool

	ctx    context.Context
	cancel context.CancelFunc
	done   chan struct{}

	deferredConns map[*grpc.ClientConn]*time.Timer
}

// New returns new Transport with specified Config.
func New(cfg *Config) *Transport { _ = "STUB: not implemented"; return nil }

func (t *Transport) run(ctx context.Context) { _ = "STUB: not implemented"; return }

// Stop stops transport and waits until it finished
func (t *Transport) Stop() { _ = "STUB: not implemented"; return }

// Send sends raft message to remote peers.
func (t *Transport) Send(m raftpb.Message) error { _ = "STUB: not implemented"; return nil }

// we need to process messages to unknown peers in separate goroutine
// to not block sender

// AddPeer adds new peer with id and address addr to Transport.
// If there is already peer with such id in Transport it will return error if
// address is different (UpdatePeer should be used) or nil otherwise.
func (t *Transport) AddPeer(id uint64, addr string) error { _ = "STUB: not implemented"; return nil }

// RemovePeer removes peer from Transport and wait for it to stop.
func (t *Transport) RemovePeer(id uint64) error { _ = "STUB: not implemented"; return nil }

// store connection and timer for cleaning up on stop

// UpdatePeer updates peer with new address. It replaces connection immediately.
func (t *Transport) UpdatePeer(id uint64, addr string) error { _ = "STUB: not implemented"; return nil }

// UpdatePeerAddr updates peer with new address, but delays connection creation.
// New address won't be used until first failure on old address.
func (t *Transport) UpdatePeerAddr(id uint64, addr string) error {
	_ = "STUB: not implemented"
	return nil
}

// PeerConn returns raw grpc connection to peer.
func (t *Transport) PeerConn(id uint64) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PeerAddr returns address of peer with id.
func (t *Transport) PeerAddr(id uint64) (string, error) { _ = "STUB: not implemented"; return "", nil }

// HealthCheck checks health of particular peer.
func (t *Transport) HealthCheck(ctx context.Context, id uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Active returns true if node was recently active and false otherwise.
func (t *Transport) Active(id uint64) bool { _ = "STUB: not implemented"; return false }

// LongestActive returns the ID of the peer that has been active for the longest
// length of time.
func (t *Transport) LongestActive() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// longestActive returns the peer that has been active for the longest length of
// time.
func (t *Transport) longestActive() (*peer, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Transport) dial(addr string) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// gRPC dialer connects to proxy first. Provide a custom dialer here avoid that.
// TODO(anshul) Add an option to configure this.

// TODO(dperny): this changes the max received message size for outgoing
// client connections. this means if the server sends a message larger than
// this, we will still accept and unmarshal it. i'm unsure what the
// potential consequences are of setting this to be effectively unbounded,
// so after docker/swarmkit#2774 is fixed, we should remove this option

func (t *Transport) withContext(ctx context.Context) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

func (t *Transport) resolvePeer(ctx context.Context, id uint64) (*peer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) sendUnknownMessage(ctx context.Context, m raftpb.Message) error {
	_ = "STUB: not implemented"
	return nil
}
