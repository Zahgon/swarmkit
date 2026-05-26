package raft

import (
	"context"
	"time"

	"github.com/moby/swarmkit/v2/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// dial returns a grpc client connection
func dial(addr string, _ string, creds credentials.TransportCredentials, timeout time.Duration) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	// gRPC dialer connects to proxy first. Provide a custom dialer here avoid that.
	return nil, nil
}

// Register registers the node raft server
func Register(server *grpc.Server, node *Node) { _ = "STUB: not implemented"; return }

// WaitForLeader waits until node observe some leader in cluster. It returns
// error if ctx was cancelled before leader appeared.
func WaitForLeader(ctx context.Context, n *Node) error { _ = "STUB: not implemented"; return nil }

// WaitForCluster waits until node observes that the cluster wide config is
// committed to raft. This ensures that we can see and serve informations
// related to the cluster.
func WaitForCluster(ctx context.Context, n *Node) (cluster *api.Cluster, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
