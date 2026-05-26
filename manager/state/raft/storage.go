package raft

import (
	"context"

	"github.com/docker/go-metrics"
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/state/raft/storage"
	"go.etcd.io/raft/v3"
	"go.etcd.io/raft/v3/raftpb"
)

var (
	// Snapshot create latency timer.
	snapshotLatencyTimer metrics.Timer
)

func init() {
	ns := metrics.NewNamespace("swarm", "raft", nil)
	snapshotLatencyTimer = ns.NewTimer("snapshot_latency",
		"Raft snapshot create latency.")
	metrics.Register(ns)
}

func (n *Node) readFromDisk(ctx context.Context) (*raftpb.Snapshot, storage.WALData, error) {
	_ = "STUB: not implemented"
	return nil, *new(storage.WALData), nil
}

// bootstraps a node's raft store from the raft logs and snapshots on disk
func (n *Node) loadAndStart(ctx context.Context, forceNewCluster bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Read logs to fully catch up store

// All members that are no longer part of the cluster must be added to
// the removed list right away, so that we don't try to connect to them
// before processing the configuration change entries, which could make
// us get stuck.

// discard the previously uncommitted entries

// force append the configuration change entries

// All members that are being removed as part of the
// force-new-cluster process must be added to the
// removed list right away, so that we don't try to
// connect to them before processing the configuration
// change entries, which could make us get stuck.

// force commit newly appended entries

func (n *Node) newRaftLogs(nodeID string) (raft.Peer, error) {
	_ = "STUB: not implemented"
	return *new(raft.Peer), nil
}

func (n *Node) triggerSnapshot(ctx context.Context, raftConfig api.RaftConfig) {
	_ = "STUB: not implemented"
	return
}

// buffered in case Shutdown is called during the snapshot

// Deferred latency capture.

// Wait for the goroutine to establish a read transaction, to make
// sure it sees the state as of this moment.

func (n *Node) clusterSnapshot(data []byte) (api.ClusterSnapshot, error) {
	_ = "STUB: not implemented"
	return *new(api.ClusterSnapshot), nil
}
