package transport

import (
	"context"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/moby/swarmkit/v2/api"
	"go.etcd.io/raft/v3/raftpb"
)

const (
	// GRPCMaxMsgSize is the max allowed gRPC message size for raft messages.
	GRPCMaxMsgSize = 4 << 20
)

type peer struct {
	id uint64

	tr *Transport

	msgc chan raftpb.Message

	ctx    context.Context
	cancel context.CancelFunc
	done   chan struct{}

	mu      sync.Mutex
	cc      *grpc.ClientConn
	addr    string
	newAddr string

	active       bool
	becameActive time.Time
}

func newPeer(id uint64, addr string, tr *Transport) (*peer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *peer) send(m raftpb.Message) (err error) { _ = "STUB: not implemented"; return nil }

func (p *peer) update(addr string) error { _ = "STUB: not implemented"; return nil }

func (p *peer) updateAddr(addr string) error { _ = "STUB: not implemented"; return nil }

func (p *peer) conn() *grpc.ClientConn { _ = "STUB: not implemented"; return nil }

func (p *peer) address() string { _ = "STUB: not implemented"; return "" }

func (p *peer) resolveAddr(ctx context.Context, id uint64) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Returns the raft message struct size (not including the payload size) for the given raftpb.Message.
// The payload is typically the snapshot or append entries.
func raftMessageStructSize(m *raftpb.Message) int { _ = "STUB: not implemented"; return 0 }

// Returns the max allowable payload based on MaxRaftMsgSize and
// the struct size for the given raftpb.Message.
func raftMessagePayloadSize(m *raftpb.Message) int { _ = "STUB: not implemented"; return 0 }

// Split a large raft message into smaller messages.
// Currently this means splitting the []Snapshot.Data into chunks whose size
// is dictacted by MaxRaftMsgSize.
func splitSnapshotData(_ context.Context, m *raftpb.Message) []api.StreamRaftMessageRequest {
	_ = "STUB: not implemented"
	return nil
}

// get the size of the data to be split.

// Get the max payload size.

// split the snapshot into smaller messages.

// Clone Snapshot so that re-slicing Snapshot.Data below
// does not mutate m.Snapshot.Data through the shared pointer.

// sub-slice for this snapshot chunk.

// add message to the list of messages to be sent.

// Function to check if this message needs to be split to be streamed
// (because it is larger than GRPCMaxMsgSize).
// Returns true if the message type is MsgSnap
// and size larger than MaxRaftMsgSize.
func needsSplitting(m *raftpb.Message) bool { _ = "STUB: not implemented"; return false }

func (p *peer) sendProcessMessage(ctx context.Context, m raftpb.Message) error {
	_ = "STUB: not implemented"
	// These lines used to be in the code, but they've been removed. I'm
	// leaving them in in a comment just in case they cause some unforeseen
	// breakage later, to show why they were removed.
	//
	// ctx, cancel := context.WithTimeout(ctx, p.tr.config.SendTimeout)
	// defer cancel()
	//
	// Basically, these lines created a timeout that applied not to each chunk
	// of a streaming message, but to the whole streaming process. With a
	// sufficiently large raft log, the bandwidth on some connections can not
	// physically be enough to fit within the default 2 second timeout.
	// Further, it seems that because of some gRPC magic, the timeout was
	// getting propagated to the stream *server*, meaning it wasn't even the
	// sender timing out, it was the receiver.
	//
	// It should be fine to remove this timeout. The whole purpose of this
	// method is to send very large raft messages that could take several
	// seconds to send.
	return nil
}

// This is a bootleg watchdog timer. If the timer elapses without something
// being written to the bump channel, it will cancel the context.
//
// We use this because the operations on this stream *must* either time out
// or succeed for raft to function correctly. We can't just time out the
// whole operation, because of the reasons stated above. But we also only
// set the context once, when we create the stream, and so can't set an
// individual timeout for each stream operation.
//
// By doing it as this watchdog-type structure, we can time out individual
// operations by canceling the context on our own terms.

// Split the message if needed.
// Currently only supported for MsgSnap.

// Stream

// If the send succeeds, bump the watchdog timer.

// Finished sending all the messages.
// Close and receive response.

// Try doing a regular rpc if the receiver doesn't support streaming.

// Handle errors.

func healthCheckConn(ctx context.Context, cc *grpc.ClientConn) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *peer) healthCheck(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (p *peer) setActive() { _ = "STUB: not implemented"; return }

func (p *peer) setInactive() { _ = "STUB: not implemented"; return }

func (p *peer) activeTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (p *peer) drain() error { _ = "STUB: not implemented"; return nil }

// all messages proceeded

func (p *peer) handleAddressChange(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// there is possibility of race if host changing address too fast, but
// it's unlikely and eventually thing should be settled

func (p *peer) run(ctx context.Context) { _ = "STUB: not implemented"; return }

// at this point we can be sure that nobody will write to msgc

// we do not propagate context here, because this operation should be finished
// or timed out for correct raft work.

func (p *peer) stop() { _ = "STUB: not implemented"; return }
