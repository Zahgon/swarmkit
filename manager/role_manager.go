package manager

import (
	"context"
	"time"

	"code.cloudfoundry.org/clock"
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/state/raft"
	"github.com/moby/swarmkit/v2/manager/state/raft/membership"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

const (
	// roleReconcileInterval is how often to retry removing a node, if a reconciliation or
	// removal failed
	roleReconcileInterval = 5 * time.Second

	// removalTimeout is how long to wait before a raft member removal fails to be applied
	// to the store
	removalTimeout = 5 * time.Second
)

// roleManager reconciles the raft member list with desired role changes.
type roleManager struct {
	ctx    context.Context
	cancel func()

	store    *store.MemoryStore
	raft     *raft.Node
	doneChan chan struct{}

	// pendingReconciliation contains changed nodes that have not yet been reconciled in
	// the raft member list.
	pendingReconciliation map[string]*api.Node

	// pendingRemoval contains the IDs of nodes that have been deleted - if these correspond
	// to members in the raft cluster, those members need to be removed from raft
	pendingRemoval map[string]struct{}

	// leave this nil except for tests which need to inject a fake time source
	clocksource clock.Clock
}

// newRoleManager creates a new roleManager.
func newRoleManager(store *store.MemoryStore, raftNode *raft.Node) *roleManager {
	_ = "STUB: not implemented"
	return nil
}

// getTicker returns a ticker based on the configured clock source
func (rm *roleManager) getTicker(interval time.Duration) clock.Ticker {
	_ = "STUB: not implemented"
	return *new(clock.Ticker)
}

// Run is roleManager's main loop.  On startup, it looks at every node object in the cluster and
// attempts to reconcile the raft member list with all the nodes' desired roles.  If any nodes
// need to be demoted or promoted, it will add them to a reconciliation queue, and if any raft
// members' node have been deleted, it will add them to a removal queue.

// These queues are processed immediately, and any nodes that failed to be processed are
// processed again in the next reconciliation interval, so that nodes will hopefully eventually
// be reconciled.  As node updates come in, any promotions or demotions are also added to the
// reconciliation queue and reconciled.  As node removals come in, they are added to the removal
// queue to be removed from the raft cluster.

// Removal from a raft cluster is idempotent (and it's the only raft cluster change that will occur
// during reconciliation or removal), so it's fine if a node is in both the removal and reconciliation
// queues.

// The ctx param is only used for logging.
func (rm *roleManager) Run(ctx context.Context) { _ = "STUB: not implemented"; return }

// ticker and tickerCh are used to time the reconciliation interval, which will
// periodically attempt to re-reconcile nodes that failed to reconcile the first
// time through

// Assume all raft members have been deleted from the cluster, until the node list
// tells us otherwise.  We can make this assumption because the node object must
// exist first before the raft member object.

// Background life-cycle for a manager: it joins the cluster, getting a new TLS
// certificate. To get a TLS certificate, it makes an RPC call to the CA server,
// which on successful join adds its information to the cluster node list and
// eventually generates a TLS certificate for it. Once it has a TLS certificate,
// it can contact the other nodes, and makes an RPC call to request to join the
// raft cluster.  The node it contacts will add the node to the raft membership.

// if the node exists, we don't want it removed from the raft membership cluster
// necessarily

// reconcile each existing node

// If any reconciliations or member removals failed, we want to try again, so
// make sure that we start the ticker so we can try again and again every
// roleReconciliationInterval seconds until the queues are both empty.

// If any reconciliations or member removals failed, we want to try again, so
// make sure that we start the ticker so we can try again and again every
// roleReconciliationInterval seconds until the queues are both empty.

// evictRemovedNode evicts a removed node from the raft cluster membership.  This is to cover an edge case in which
// a node might have been removed, but somehow the role was not reconciled (possibly a demotion and a removal happened
// in rapid succession before the raft membership configuration went through).
func (rm *roleManager) evictRemovedNode(ctx context.Context, nodeID string) {
	_ = "STUB: not implemented"
	// Check if the member still exists in the membership
	return
}

// We first try to remove the raft node from the raft cluster.  On the next tick, if the node
// has been removed from the cluster membership, we then delete it from the removed list

// removeMember removes a member from the raft cluster membership
func (rm *roleManager) removeMember(ctx context.Context, member *membership.Member) {
	_ = "STUB: not implemented"
	// Quorum safeguard - quorum should have been checked before a node was allowed to be demoted, but if in the
	// intervening time some other node disconnected, removing this node would result in a loss of cluster quorum.
	// We leave it
	return
}

// TODO(aaronl): Retry later

// Don't use rmCtx, because we expect to lose
// leadership, which will cancel this context.

// TODO(aaronl): Retry later

// reconcileRole looks at the desired role for a node, and if it is being demoted or promoted, updates the
// node role accordingly.   If the node is being demoted, it also removes the node from the raft cluster membership.
func (rm *roleManager) reconcileRole(ctx context.Context, node *api.Node) {
	_ = "STUB: not implemented"
	return
}

// Nothing to do.

// Promotion can proceed right away.

// Check for node in memberlist

// We first try to remove the raft node from the raft cluster.  On the next tick, if the node
// has been removed from the cluster membership, we then update the store to reflect the fact
// that it has been successfully demoted, and if that works, remove it from the pending list.

// Stop stops the roleManager and waits for the main loop to exit.
func (rm *roleManager) Stop() { _ = "STUB: not implemented"; return }
