// Package connectionbroker is a layer on top of remotes that returns
// a gRPC connection to a manager. The connection may be a local connection
// using a local socket such as a UNIX socket.
package connectionbroker

import (
	"sync"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/remotes"
	"google.golang.org/grpc"
)

// Broker is a simple connection broker. It can either return a fresh
// connection to a remote manager selected with weighted randomization, or a
// local gRPC connection to the local manager.
type Broker struct {
	mu        sync.Mutex
	remotes   remotes.Remotes
	localConn *grpc.ClientConn
}

// New creates a new connection broker.
func New(remotes remotes.Remotes) *Broker { _ = "STUB: not implemented"; return nil }

// SetLocalConn changes the local gRPC connection used by the connection broker.
func (b *Broker) SetLocalConn(localConn *grpc.ClientConn) { _ = "STUB: not implemented"; return }

// Select a manager from the set of available managers, and return a connection.
func (b *Broker) Select(dialOpts ...grpc.DialOption) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SelectRemote chooses a manager from the remotes, and returns a TCP
// connection.
func (b *Broker) SelectRemote(dialOpts ...grpc.DialOption) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// gRPC dialer connects to proxy first. Provide a custom dialer here avoid that.
// TODO(anshul) Add an option to configure this.

// Remotes returns the remotes interface used by the broker, so the caller
// can make observations or see weights directly.
func (b *Broker) Remotes() remotes.Remotes {
	_ = "STUB: not implemented"

	// Conn is a wrapper around a gRPC client connection.
	return *new(remotes.Remotes)
}

type Conn struct {
	*grpc.ClientConn
	isLocal bool
	remotes remotes.Remotes
	peer    api.Peer
}

// Peer returns the peer for this Conn.
func (c *Conn) Peer() api.Peer {
	_ = "STUB: not implemented"

	// Close closes the client connection if it is a remote connection. It also
	// records a positive experience with the remote peer if success is true,
	// otherwise it records a negative experience. If a local connection is in use,
	// Close is a noop.
	return *new(api.Peer)
}

func (c *Conn) Close(success bool) error { _ = "STUB: not implemented"; return nil }
