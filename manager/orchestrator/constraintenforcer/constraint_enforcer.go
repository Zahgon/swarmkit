package constraintenforcer

import (
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

// ConstraintEnforcer watches for updates to nodes and shuts down tasks that no
// longer satisfy scheduling constraints or resource limits.
type ConstraintEnforcer struct {
	store    *store.MemoryStore
	stopChan chan struct{}
	doneChan chan struct{}
}

// New creates a new ConstraintEnforcer.
func New(store *store.MemoryStore) *ConstraintEnforcer { _ = "STUB: not implemented"; return nil }

// Run is the ConstraintEnforcer's main loop.
func (ce *ConstraintEnforcer) Run() { _ = "STUB: not implemented"; return }

func (ce *ConstraintEnforcer) rejectNoncompliantTasks(node *api.Node) {
	_ = "STUB: not implemented"
	// If the availability is "drain", the orchestrator will
	// shut down all tasks.
	// If the availability is "pause", we shouldn't touch
	// the tasks on this node.
	return
}

// Deduplicate service IDs using the services map. It's okay for the
// values to be nil for now, we will look them up from the store next.

// TODO(aaronl): The set of tasks removed will be
// nondeterministic because it depends on the order of
// the slice returned from FindTasks. We could do
// a separate pass over the tasks for each type of
// resource, and sort by the size of the reservation
// to remove the most resource-intensive tasks.

// Ensure that the node still satisfies placement constraints.
// NOTE: If the task is associacted with a service then we must use the
// constraints from the current service spec rather than the
// constraints from the task spec because they may be outdated. This
// will happen if the service was previously updated in a way which
// only changes the placement constraints and the node matched the
// placement constraints both before and after that update. In the case
// of such updates, the tasks are not considered "dirty" and are not
// restarted but it will mean that the task spec's placement
// constraints are outdated. Consider this example:
// - A service is created with no constraints and a task is scheduled
//   to a node.
// - The node is updated to add a label, this doesn't affect the task
//   on that node because it has no constraints.
// - The service is updated to add a node label constraint which
//   matches the label which was just added to the node. The updater
//   does not shut down the task because the only the constraints have
//   changed and the node still matches the updated constraints.
// - The node is updated to remove the node label. The node no longer
//   satisfies the placement constraints of the service, so the task
//   should be shutdown. However, the task's spec still has the
//   original and outdated constraints (that are still satisfied by
//   the node). If we used those original constraints then the task
//   would incorrectly not be removed. This is why the constraints
//   from the service spec should be used instead.

// This task is associated with a service, so we use the service's
// current placement constraints.

// This task is not associated with a service (or the service no
// longer exists), so we use the placement constraints from the
// original task spec.

// Ensure that the task assigned to the node
// still satisfies the resource limits.

// Ensure that the task assigned to the node
// still satisfies the available generic resources

// Type change or no longer available

// We set the observed state to
// REJECTED, rather than the desired
// state. Desired state is owned by the
// orchestrator, and setting it directly
// will bypass actions such as
// restarting the task on another node
// (if applicable).

// Stop stops the ConstraintEnforcer and waits for the main loop to exit.
func (ce *ConstraintEnforcer) Stop() { _ = "STUB: not implemented"; return }
