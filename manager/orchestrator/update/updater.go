package update

import (
	"context"
	"sync"
	"time"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/orchestrator"
	"github.com/moby/swarmkit/v2/manager/orchestrator/restart"
	"github.com/moby/swarmkit/v2/manager/state/store"
	"github.com/moby/swarmkit/v2/watch"
)

// Supervisor supervises a set of updates. It's responsible for keeping track of updates,
// shutting them down and replacing them.
type Supervisor struct {
	store    *store.MemoryStore
	restarts *restart.Supervisor
	updates  map[string]*Updater
	l        sync.Mutex
}

// NewSupervisor creates a new UpdateSupervisor.
func NewSupervisor(store *store.MemoryStore, restartSupervisor *restart.Supervisor) *Supervisor {
	_ = "STUB: not implemented"
	return nil
}

// Update starts an Update of `slots` belonging to `service` in the background
// and returns immediately. Each slot contains a group of one or more tasks
// occupying the same slot (replicated service) or node (global service). There
// may be more than one task per slot in cases where an update is in progress
// and the new task was started before the old one was shut down. If an update
// for that service was already in progress, it will be cancelled before the
// new one starts.
func (u *Supervisor) Update(ctx context.Context, cluster *api.Cluster, service *api.Service, slots []orchestrator.Slot) {
	_ = "STUB: not implemented"
	return
}

// There's already an update working towards this goal.

// CancelAll cancels all current updates.
func (u *Supervisor) CancelAll() { _ = "STUB: not implemented"; return }

// Updater updates a set of tasks to a new version.
type Updater struct {
	store      *store.MemoryStore
	watchQueue *watch.Queue
	restarts   *restart.Supervisor

	cluster    *api.Cluster
	newService *api.Service

	updatedTasks   map[string]time.Time // task ID to creation time
	updatedTasksMu sync.Mutex

	// stopChan signals to the state machine to stop running.
	stopChan chan struct{}
	// doneChan is closed when the state machine terminates.
	doneChan chan struct{}
}

// NewUpdater creates a new Updater.
func NewUpdater(store *store.MemoryStore, restartSupervisor *restart.Supervisor, cluster *api.Cluster, newService *api.Service) *Updater {
	_ = "STUB: not implemented"
	return nil
}

// Cancel cancels the current update immediately. It blocks until the cancellation is confirmed.
func (u *Updater) Cancel() { _ = "STUB: not implemented"; return }

// Run starts the update and returns only once its complete or cancelled.
func (u *Updater) Run(ctx context.Context, slots []orchestrator.Slot) {
	_ = "STUB: not implemented"
	return
}

// If the update is in a PAUSED state, we should not do anything.

// Abort immediately if all tasks are clean.

// If there's no update in progress, we are starting one.

// TODO(aluzzardi): We could try to optimize unlimited parallelism by performing updates in a single
// goroutine using a batch transaction.

// Start the workers.

// Ignore tasks we have already seen as failures.

// If this failed/completed task is one that we
// created as part of this update, we should
// follow the failure action.

// Never roll back a rollback

// Wait for a worker to pick up the task or abort the update, whichever comes first.

// if a delay is set we need to monitor for a period longer than the delay
// otherwise we will leave the monitorLoop before the task is done delaying

// Keep watching for task failures for one more monitoringPeriod,
// before declaring the update complete.

// TODO(aaronl): Potentially roll back the service if not enough tasks
// have reached RUNNING by this point.

func (u *Updater) worker(ctx context.Context, queue <-chan orchestrator.Slot, updateConfig *api.UpdateConfig) {
	_ = "STUB: not implemented"
	return
}

// Do we have a task with the new spec in desired state = RUNNING?
// If so, all we have to do to complete the update is remove the
// other tasks. Or if we have a task with the new spec that has
// desired state < RUNNING, advance it to running and remove the
// other tasks.

func (u *Updater) updateTask(ctx context.Context, slot orchestrator.Slot, updated *api.Task, order api.UpdateConfig_UpdateOrder) error {
	_ = "STUB: not implemented"
	// Kick off the watch before even creating the updated task. This is in order to avoid missing any event.
	return nil
}

// Create an empty entry for this task, so the updater knows a failure
// should count towards the failure count. The timestamp is added
// if/when the task reaches RUNNING.

// Atomically create the updated task and bring down the old one.

// Wait for the new task to come up.
// TODO(aluzzardi): Consider adding a timeout here.

func (u *Updater) useExistingTask(ctx context.Context, slot orchestrator.Slot, existing *api.Task) error {
	_ = "STUB: not implemented"
	return nil
}

// removeOldTasks shuts down the given tasks and returns one of the tasks that
// was shut down, or an error.
func (u *Updater) removeOldTasks(_ context.Context, batch *store.Batch, removeTasks []*api.Task) (*api.Task, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (u *Updater) isTaskDirty(t *api.Task) bool { _ = "STUB: not implemented"; return false }

func (u *Updater) isSlotDirty(slot orchestrator.Slot) bool { _ = "STUB: not implemented"; return false }

func (u *Updater) startUpdate(ctx context.Context, serviceID string) {
	_ = "STUB: not implemented"
	return
}

func (u *Updater) pauseUpdate(ctx context.Context, serviceID, message string) {
	_ = "STUB: not implemented"
	return
}

// The service was updated since we started this update

func (u *Updater) rollbackUpdate(ctx context.Context, serviceID, message string) {
	_ = "STUB: not implemented"
	return
}

// The service was updated since we started this update

func (u *Updater) completeUpdate(ctx context.Context, serviceID string) {
	_ = "STUB: not implemented"
	return
}

// The service was changed since we started this update
