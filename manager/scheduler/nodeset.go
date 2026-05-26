package scheduler

import (
	"errors"

	"github.com/moby/swarmkit/v2/api"
)

var errNodeNotFound = errors.New("node not found in scheduler dataset")

type nodeSet struct {
	nodes map[string]NodeInfo // map from node id to node info
}

func (ns *nodeSet) alloc(n int) { _ = "STUB: not implemented"; return }

// nodeInfo returns the NodeInfo struct for a given node identified by its ID.
func (ns *nodeSet) nodeInfo(nodeID string) (NodeInfo, error) {
	_ = "STUB: not implemented"
	return *new(NodeInfo), nil
}

// addOrUpdateNode sets the number of tasks for a given node. It adds the node
// to the set if it wasn't already tracked.
func (ns *nodeSet) addOrUpdateNode(n NodeInfo) {
	_ = "STUB: not implemented"

	// updateNode sets the number of tasks for a given node. It ignores the update
	// if the node isn't already tracked in the set.
	return
}

func (ns *nodeSet) updateNode(n NodeInfo) { _ = "STUB: not implemented"; return }

func (ns *nodeSet) remove(nodeID string) { _ = "STUB: not implemented"; return }

func (ns *nodeSet) tree(serviceID string, preferences []*api.PlacementPreference, maxAssignments int, meetsConstraints func(*NodeInfo) bool, nodeLess func(*NodeInfo, *NodeInfo) bool) decisionTree {
	_ = "STUB: not implemented"
	return *new(decisionTree)
}

// Only spread is supported so far

// TODO(aaronl): Support other items from constraint
// syntax like node ID, hostname, os/arch, etc?

// If value is still uninitialized, the value used for
// the node at this level of the tree is "". This makes
// sure that the tree structure is not affected by
// which properties nodes have and don't have.
