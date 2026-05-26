package replicated

import (
	"context"

	"github.com/docker/go-events"
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/orchestrator"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

// This file provices service-level orchestration. It observes changes to
// services and creates and destroys tasks as necessary to match the service
// specifications. This is different from task-level orchestration, which
// responds to changes in individual tasks (or nodes which run them).

func (r *Orchestrator) initCluster(readTx store.ReadTx) error {
	_ = "STUB: not implemented"
	return nil
}

// we'll just pick it when it is created.

func (r *Orchestrator) initServices(readTx store.ReadTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Orchestrator) handleServiceEvent(ctx context.Context, event events.Event) {
	_ = "STUB: not implemented"
	return
}

func (r *Orchestrator) tickServices(ctx context.Context) { _ = "STUB: not implemented"; return }

func (r *Orchestrator) resolveService(_ context.Context, task *api.Task) *api.Service {
	_ = "STUB: not implemented"
	return nil
}

// reconcile decides what actions must be taken depending on the number of
// specificed slots and actual running slots. If the actual running slots are
// fewer than what is requested, it creates new tasks. If the actual running
// slots are more than requested, then it decides which slots must be removed
// and sets desired state of those tasks to REMOVE (the actual removal is handled
// by the task reaper, after the agent shuts the tasks down).
func (r *Orchestrator) reconcile(ctx context.Context, service *api.Service) {
	_ = "STUB: not implemented"
	return
}

// Update all current tasks then add missing tasks

// Update up to N tasks then remove the extra

// Preferentially remove tasks on the nodes that have the most
// copies of this service, to leave a more balanced result.

// First sort tasks such that tasks which are currently running
// (in terms of observed state) appear before non-running tasks.
// This will cause us to prefer to remove non-running tasks, all
// other things being equal in terms of node balance.

// Assign each task an index that counts it as the nth copy of
// of the service on its node (1, 2, 3, ...), and sort the
// tasks by this counter value.

// for all slots that we are removing, we set the desired state of those tasks
// to REMOVE. Then, the agent is responsible for shutting them down, and the
// task reaper is responsible for actually removing them from the store after
// shutdown.

// Simple update, no scaling - update all tasks.

func (r *Orchestrator) addTasks(ctx context.Context, batch *store.Batch, service *api.Service, runningSlots map[uint64]orchestrator.Slot, deadSlots map[uint64]orchestrator.Slot, count uint64) {
	_ = "STUB: not implemented"
	return
}

// Find a slot number that is missing a running task

// setTasksDesiredState sets the desired state for all tasks for the given slots to the
// requested state
func (r *Orchestrator) setTasksDesiredState(ctx context.Context, batch *store.Batch, slots []orchestrator.Slot, newDesiredState api.TaskState) {
	_ = "STUB: not implemented"
	return
}

// time travel is not allowed. if the current desired state is
// above the one we're trying to go to we can't go backwards.
// we have nothing to do and we should skip to the next task

// log a warning, though. we shouln't be trying to rewrite
// a state to an earlier state

// update desired state

// log an error if we get one

func (r *Orchestrator) deleteTasksMap(ctx context.Context, batch *store.Batch, slots map[uint64]orchestrator.Slot) {
	_ = "STUB: not implemented"
	return
}

func (r *Orchestrator) deleteTask(ctx context.Context, batch *store.Batch, t *api.Task) {
	_ = "STUB: not implemented"
	return
}

// IsRelatedService returns true if the service should be governed by this orchestrator
func (r *Orchestrator) IsRelatedService(service *api.Service) bool {
	_ = "STUB: not implemented"
	return false
}
