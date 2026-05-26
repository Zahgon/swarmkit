package restart

import (
	"container/list"
	"context"
	"sync"
	"time"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/orchestrator"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

const defaultOldTaskTimeout = time.Minute

type restartedInstance struct {
	timestamp time.Time
}

type instanceRestartInfo struct {
	// counter of restarts for this instance.
	totalRestarts uint64
	// Linked list of restartedInstance structs. Only used when
	// Restart.MaxAttempts and Restart.Window are both
	// nonzero.
	restartedInstances *list.List
	// Why is specVersion in this structure and not in the map key? While
	// putting it in the key would be a very simple solution, it wouldn't
	// be easy to clean up map entries corresponding to old specVersions.
	// Making the key version-agnostic and clearing the value whenever the
	// version changes avoids the issue of stale map entries for old
	// versions.
	specVersion api.Version
}

type delayedStart struct {
	// cancel is called to cancel the delayed start.
	cancel func()
	doneCh chan struct{}

	// waiter is set to true if the next restart is waiting for this delay
	// to complete.
	waiter bool
}

// SupervisorInterface is an interface implemented by the Supervisor. It exists
// to make testing easier, by allowing the restart supervisor to be mocked or
// faked where desired.
type SupervisorInterface interface {
	Restart(context.Context, store.Tx, *api.Cluster, *api.Service, api.Task) error
	UpdatableTasksInSlot(context.Context, orchestrator.Slot, *api.Service) orchestrator.Slot
	RecordRestartHistory(orchestrator.SlotTuple, *api.Task)
	DelayStart(context.Context, store.Tx, *api.Task, string, time.Duration, bool) <-chan struct{}
	StartNow(store.Tx, string) error
	Cancel(string)
	CancelAll()
	ClearServiceHistory(string)
}

// Supervisor initiates and manages restarts. It's responsible for
// delaying restarts when applicable.
type Supervisor struct {
	mu               sync.Mutex
	store            *store.MemoryStore
	delays           map[string]*delayedStart
	historyByService map[string]map[orchestrator.SlotTuple]*instanceRestartInfo
	TaskTimeout      time.Duration
}

// NewSupervisor creates a new RestartSupervisor.
func NewSupervisor(store *store.MemoryStore) *Supervisor { _ = "STUB: not implemented"; return nil }

func (r *Supervisor) waitRestart(ctx context.Context, oldDelay *delayedStart, cluster *api.Cluster, taskID string) {
	_ = "STUB: not implemented"
	// Wait for the last restart delay to elapse.
	return
}

// Start the next restart

// Restart initiates a new task to replace t if appropriate under the service's
// restart policy.
func (r *Supervisor) Restart(ctx context.Context, tx store.Tx, cluster *api.Cluster, service *api.Service, t api.Task) error {
	_ = "STUB: not implemented"
	// TODO(aluzzardi): This function should not depend on `service`.
	return nil
}

// Is the old task still in the process of restarting? If so, wait for
// its restart delay to elapse, to avoid tight restart loops (for
// example, when the image doesn't exist).

// Sanity check: was the task shut down already by a separate call to
// Restart? If so, we must avoid restarting it, because this will create
// an extra task. This should never happen unless there is a bug.

// Restart delay is not applied to drained nodes

// Normally we wait for the old task to stop running, but we skip this
// if the old task is already dead or the node it's assigned to is down.

// shouldRestart returns true if a task should be restarted according to the
// restart policy.
func (r *Supervisor) shouldRestart(ctx context.Context, t *api.Task, service *api.Service) bool {
	_ = "STUB: not implemented"
	// TODO(aluzzardi): This function should not depend on `service`.
	// There are 3 possible restart policies.
	return false
}

// we will be restarting, we just need to do a few more checks.
// however, if the task belongs to a job, then we will treat
// RestartOnAny the same as RestartOnFailure, as it would be
// nonsensical to restart completed jobs.

// it'd be nice to put a fallthrough here, but we can't fallthrough
// from inside of an if statement.

// we won't restart if the task is in TaskStateCompleted, as this is a
// not a failed state -- it indicates that the task exited with 0

// RestartOnNone means we just don't restart, ever

// Slot is not meaningful for "global" tasks, so they need to be
// indexed by NodeID.

// Prefer the manager's timestamp over the agent's, since manager
// clocks are more trustworthy.

// It's safe to call TimestampFromProto with a nil timestamp

// Disregard any restarts that happened before the lookback window,
// and remove them from the linked list since they will no longer
// be relevant to figuring out if tasks should be restarted going
// forward.

// Ignore restarts that didn't happen before the task we're looking at.

// UpdatableTasksInSlot returns the set of tasks that should be passed to the
// updater from this slot, or an empty slice if none should be.  An updatable
// slot has either at least one task that with desired state <= RUNNING, or its
// most recent task has stopped running and should not be restarted. The latter
// case is for making sure that tasks that shouldn't normally be restarted will
// still be handled by rolling updates when they become outdated.  There is a
// special case for rollbacks to make sure that a rollback always takes the
// service to a converged state, instead of ignoring tasks with the original
// spec that stopped running and shouldn't be restarted according to the
// restart policy.
func (r *Supervisor) UpdatableTasksInSlot(ctx context.Context, slot orchestrator.Slot, service *api.Service) orchestrator.Slot {
	_ = "STUB: not implemented"
	return *new(orchestrator.Slot)
}

// Find most recent task

// RecordRestartHistory updates the historyByService map to reflect the restart
// of restartedTask.
func (r *Supervisor) RecordRestartHistory(tuple orchestrator.SlotTuple, replacementTask *api.Task) {
	_ = "STUB: not implemented"
	return
}

// No limit on the number of restarts, so no need to record
// history.

// This task has a different SpecVersion from the one we're
// tracking. Most likely, the service was updated. Past failures
// shouldn't count against the new service definition, so clear
// the history for this instance.

// it's okay to call TimestampFromProto with a nil argument

// DelayStart starts a timer that moves the task from READY to RUNNING once:
// - The restart delay has elapsed (if applicable)
// - The old task that it's replacing has stopped running (or this times out)
// It must be called during an Update transaction to ensure that it does not
// miss events. The purpose of the store.Tx argument is to avoid accidental
// calls outside an Update transaction.
func (r *Supervisor) DelayStart(ctx context.Context, _ store.Tx, oldTask *api.Task, newTaskID string, delay time.Duration, waitStop bool) <-chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

// Note that this channel read should only block for a very
// short time, because we cancelled the existing delay and
// that should cause it to stop immediately.

// Wait for either the old task to complete, or the old task's
// node to become unavailable.

// Wait for the delay to elapse, if one is specified.

// StartNow moves the task into the RUNNING state so it will proceed to start
// up.
func (r *Supervisor) StartNow(tx store.Tx, taskID string) error {
	_ = "STUB: not implemented"
	return nil
}

// only tasks belonging to jobs will have a JobIteration, so this can be
// used to distinguish whether this is a job task without looking at the
// service.

// Cancel cancels a pending restart.
func (r *Supervisor) Cancel(taskID string) { _ = "STUB: not implemented"; return }

// CancelAll aborts all pending restarts
func (r *Supervisor) CancelAll() { _ = "STUB: not implemented"; return }

// ClearServiceHistory forgets restart history related to a given service ID.
func (r *Supervisor) ClearServiceHistory(serviceID string) { _ = "STUB: not implemented"; return }
