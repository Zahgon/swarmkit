package global

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/constraint"
	"github.com/moby/swarmkit/v2/manager/orchestrator"
	"github.com/moby/swarmkit/v2/manager/orchestrator/restart"
	"github.com/moby/swarmkit/v2/manager/orchestrator/update"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

type globalService struct {
	*api.Service

	// Compiled constraints
	constraints []constraint.Constraint
}

// Orchestrator runs a reconciliation loop to create and destroy tasks as
// necessary for global services.
type Orchestrator struct {
	store *store.MemoryStore
	// nodes is the set of non-drained nodes in the cluster, indexed by node ID
	nodes map[string]*api.Node
	// globalServices has all the global services in the cluster, indexed by ServiceID
	globalServices map[string]globalService
	restartTasks   map[string]struct{}

	// stopChan signals to the state machine to stop running.
	stopChan chan struct{}
	// doneChan is closed when the state machine terminates.
	doneChan chan struct{}

	updater  *update.Supervisor
	restarts *restart.Supervisor

	cluster *api.Cluster // local instance of the cluster
}

// NewGlobalOrchestrator creates a new global Orchestrator
func NewGlobalOrchestrator(store *store.MemoryStore) *Orchestrator {
	_ = "STUB: not implemented"
	return nil
}

func (g *Orchestrator) initTasks(ctx context.Context, readTx store.ReadTx) error {
	_ = "STUB: not implemented"
	return nil
}

// Run contains the global orchestrator event loop
func (g *Orchestrator) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Watch changes to services and tasks

// lookup the cluster

// just pick up the cluster when it is created.

// Get list of nodes

// Lookup global services

// fix tasks in store before reconciliation loop

// TODO(stevvooe): Use ctx to limit running time of operation.

// delete the service from service map

// FixTask validates a task with the current cluster settings, and takes
// action to make it conformant to node state and service constraint
// it's called at orchestrator initialization
func (g *Orchestrator) FixTask(ctx context.Context, batch *store.Batch, t *api.Task) {
	_ = "STUB: not implemented"
	return
}

// if a task's DesiredState has past running, the task has been processed

// if the node no longer valid, remove the task

// restart a task if it fails

// handleTaskChange defines what orchestrator does when a task is updated by agent
func (g *Orchestrator) handleTaskChange(_ context.Context, t *api.Task) {
	_ = "STUB: not implemented"
	return
}

// if a task's DesiredState has passed running, it
// means the task has been processed

// if a task has passed running, restart it

// Stop stops the orchestrator.
func (g *Orchestrator) Stop() { _ = "STUB: not implemented"; return }

func (g *Orchestrator) foreachTaskFromNode(ctx context.Context, node *api.Node, cb func(context.Context, *store.Batch, *api.Task)) {
	_ = "STUB: not implemented"
	return
}

// Global orchestrator only removes tasks from globalServices

func (g *Orchestrator) reconcileServices(ctx context.Context, serviceIDs []string) {
	_ = "STUB: not implemented"
	return
}

// nodeID -> task list

// Keep all runnable instances of this service,
// and instances that were not be restarted due
// to restart policy but may be updated if the
// service spec changed.

// the node is paused, so we won't add or update
// any tasks

// this node needs to run 1 copy of the task

// Remove any tasks assigned to nodes not found in g.nodes.
// These must be associated with nodes that are drained, or
// nodes that no longer exist.

// updateNode updates g.nodes based on the current node value
func (g *Orchestrator) updateNode(node *api.Node) { _ = "STUB: not implemented"; return }

// updateService updates g.globalServices based on the current service value
func (g *Orchestrator) updateService(service *api.Service) { _ = "STUB: not implemented"; return }

// reconcileOneNode checks all global services on one node
func (g *Orchestrator) reconcileOneNode(ctx context.Context, node *api.Node) {
	_ = "STUB: not implemented"
	return
}

// the node is paused, so we won't add or update tasks

// tasks by service

// Keep all runnable instances of this service,
// and instances that were not be restarted due
// to restart policy but may be updated if the
// service spec changed.

// If task is out of date, update it. This can happen
// on node reconciliation if, for example, we pause a
// node, update the service, and then activate the node
// later.

// We don't use g.updater here for two reasons:
// - This is not a rolling update. Since it was not
//   triggered directly by updating the service, it
//   should not observe the rolling update parameters
//   or show status in UpdateStatus.
// - Calling Update cancels any current rolling updates
//   for the service, such as one triggered by service
//   reconciliation.

func (g *Orchestrator) tickTasks(ctx context.Context) { _ = "STUB: not implemented"; return }

func (g *Orchestrator) shutdownTask(ctx context.Context, batch *store.Batch, t *api.Task) {
	_ = "STUB: not implemented"
	// set existing task DesiredState to TaskStateShutdown
	// TODO(aaronl): optimistic update?
	return
}

func (g *Orchestrator) addTask(ctx context.Context, batch *store.Batch, service *api.Service, nodeID string) {
	_ = "STUB: not implemented"
	return
}

func (g *Orchestrator) shutdownTasks(ctx context.Context, batch *store.Batch, tasks []*api.Task) {
	_ = "STUB: not implemented"
	return
}

func (g *Orchestrator) deleteTask(ctx context.Context, batch *store.Batch, t *api.Task) {
	_ = "STUB: not implemented"
	return
}

// IsRelatedService returns true if the service should be governed by this orchestrator
func (g *Orchestrator) IsRelatedService(service *api.Service) bool {
	_ = "STUB: not implemented"
	return false
}

// SlotTuple returns a slot tuple for the global service task.
func (g *Orchestrator) SlotTuple(t *api.Task) orchestrator.SlotTuple {
	_ = "STUB: not implemented"
	return *new(orchestrator.SlotTuple)
}
