package controlapi

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

func validateNodeSpec(spec *api.NodeSpec) error { _ = "STUB: not implemented"; return nil }

// GetNode returns a Node given a NodeID.
// - Returns `InvalidArgument` if NodeID is not provided.
// - Returns `NotFound` if the Node is not found.
func (s *Server) GetNode(_ context.Context, request *api.GetNodeRequest) (*api.GetNodeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func filterNodes(candidates []*api.Node, filters ...func(*api.Node) bool) []*api.Node {
	_ = "STUB: not implemented"
	return nil
}

// ListNodes returns a list of all nodes.
func (s *Server) ListNodes(_ context.Context, request *api.ListNodesRequest) (*api.ListNodesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add in manager information on nodes that are managers

// UpdateNode updates a Node referenced by NodeID with the given NodeSpec.
// - Returns `NotFound` if the Node is not found.
// - Returns `InvalidArgument` if the NodeSpec is malformed.
// - Returns an error if the update fails.
func (s *Server) UpdateNode(_ context.Context, request *api.UpdateNodeRequest) (*api.UpdateNodeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Demotion sanity checks.

// Check for manager entries in Store.

// Check for node in memberlist

// Quorum safeguard

func orphanNodeTasks(tx store.Tx, nodeID string) error {
	_ = "STUB: not implemented"
	// when a node is deleted, all of its tasks are irrecoverably removed.
	// additionally, the Dispatcher can no longer be relied on to update the
	// task status. Therefore, when the node is removed, we must additionally
	// move all of its assigned tasks to the Orphaned state, so that their
	// resources can be cleaned up.
	return nil
}

// this operation must occur within the same transaction boundary. If
// we cannot accomplish this task orphaning in the same transaction, we
// could crash or die between transactions and not get a chance to do
// this. however, in cases were there is an exceptionally large number
// of tasks for a node, this may cause the transaction to exceed the
// max message size.
//
// therefore, we restrict updating to only tasks in a non-terminal
// state. Tasks in a terminal state do not need to be updated.

// RemoveNode removes a Node referenced by NodeID with the given NodeSpec.
// - Returns NotFound if the Node is not found.
// - Returns FailedPrecondition if the Node has manager role (and is part of the memberlist) or is not shut down.
// - Returns InvalidArgument if NodeID or NodeVersion is not valid.
// - Returns an error if the delete fails.
func (s *Server) RemoveNode(_ context.Context, request *api.RemoveNodeRequest) (*api.RemoveNodeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// lookup the cluster

// Set an expiry time for this RemovedNode if a certificate
// exists and can be parsed.
