package global

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/orchestrator"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

// restartSupervisor is an interface representing the methods from the
// restart.SupervisorInterface that are actually needed by the reconciler. This
// more limited interface allows us to write a less ugly fake for unit testing.
type restartSupervisor interface {
	Restart(context.Context, store.Tx, *api.Cluster, *api.Service, api.Task) error
}

// Reconciler is an object that manages reconciliation of global jobs. It is
// blocking and non-asynchronous, for ease of testing. It implements the
// Reconciler interface from the orchestrator package above it, and the
// taskinit.InitHandler interface.
type Reconciler struct {
	store *store.MemoryStore

	restart restartSupervisor
}

// NewReconciler creates a new global job reconciler.
func NewReconciler(store *store.MemoryStore, restart restartSupervisor) *Reconciler {
	_ = "STUB: not implemented"
	return nil
}

// ReconcileService reconciles one global job service.
func (r *Reconciler) ReconcileService(id string) error { _ = "STUB: not implemented"; return nil }

// we need to first get the latest iteration of the service, its tasks, and
// the nodes in the cluster.

// getting tasks with FindTasks should only return an error if we've
// made a mistake coding; there's no user-input or even reasonable
// system state that can cause it. If it returns an error, we'll just
// panic and crash.

// same as with FindTasks

// the service may be nil if the service has been deleted before we entered
// the View.

// we need to compute the constraints on the service so we know which nodes
// to schedule it on

// constraint.Parse does return an error, but we don't need to check
// it, because it was already checked when the service was created or
// updated.

// instead of having a big ugly multi-line boolean expression in the
// if-statement, we'll have several if-statements, and bail out of
// this loop iteration with continue if the node is not acceptable

// if a node is invalid, we should remove any tasks that might be on it

// you can append to a nil slice and get a non-nil slice, which is
// pretty slick.

// now, we have a list of all nodes that match constraints. it's time to
// match running tasks to the nodes. we need to identify all nodes that
// need new tasks, which is any node that doesn't have a task of this job
// iteration. trade some space for some time by building a node ID to task
// ID mapping, so that we're just doing 2x linear operation, instead of a
// quadratic operation.

// additionally, while we're iterating through tasks, if any of those tasks
// are failed, we'll hand them to the restart supervisor to handle

// and if there are any tasks belonging to old job iterations, set them to
// be removed

// match all tasks belonging to this job iteration which are in desired
// state completed, including failed tasks. We only want to create
// tasks for nodes on which there are no existing tasks.

// we already know the task is desired to be executing (because its
// desired state is Completed). Check here to see if it's already
// failed, so we can restart it

// first, create any new tasks required.

// check if there is a task for this node ID. If not, then we need
// to create one.

// if the node does not already have a running or completed
// task, create a task for this node.

// then, restart any tasks that are failed

// get the latest version of the task for the restart

// if it's deleted, nothing to do

// if it's not still desired to be running, then don't restart
// it.

// Finally, restart it
// TODO(dperny): pass in context to ReconcileService, so we can
// pass it in here.

// TODO(dperny): probably should log like in the other
// orchestrators instead of returning here.

// remove tasks that need to be removed

// finally, shut down any tasks on invalid nodes

// if the task is still desired to be running, and is still
// actually, running, then it still needs to be shut down.

// IsRelatedService returns true if the task is a global job. This method
// fulfills the taskinit.InitHandler interface. Because it is just a wrapper
// around a well-tested function call, it has no tests of its own.
func (r *Reconciler) IsRelatedService(service *api.Service) bool {
	_ = "STUB: not implemented"
	return false
}

// FixTask validates that a task is compliant with the rest of the cluster
// state, and fixes it if it's not. This covers some main scenarios:
//
//   - The node that the task is running on is now paused or drained. we do not
//     need to check if the node still meets constraints -- that is the purview
//     of the constraint enforcer.
//   - The task has failed and needs to be restarted.
//
// This implements the FixTask method of the taskinit.InitHandler interface.
func (r *Reconciler) FixTask(_ context.Context, batch *store.Batch, t *api.Task) {
	_ = "STUB: not implemented"
	// tasks already desired to be shut down need no action.
	return
}

// if the node is no longer a valid node for this task, we need to shut
// it down

// we will reconcile all services after fixing the tasks, so we don't
// need to restart tasks right now; we'll do so after this.

// SlotTuple returns a slot tuple representing this task. It implements the
// taskinit.InitHandler interface.
func (r *Reconciler) SlotTuple(t *api.Task) orchestrator.SlotTuple {
	_ = "STUB: not implemented"
	return *new(orchestrator.SlotTuple)
}
