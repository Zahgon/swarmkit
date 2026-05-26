package agent

import (
	"context"
	"sync"

	"github.com/moby/swarmkit/v2/agent/exec"
	"github.com/moby/swarmkit/v2/api"
)

// taskManager manages all aspects of task execution and reporting for an agent
// through state management.
type taskManager struct {
	task     *api.Task
	ctlr     exec.Controller
	reporter StatusReporter

	updateq chan *api.Task

	shutdown     chan struct{}
	shutdownOnce sync.Once
	closed       chan struct{}
	closeOnce    sync.Once
}

func newTaskManager(ctx context.Context, task *api.Task, ctlr exec.Controller, reporter StatusReporter) *taskManager {
	_ = "STUB: not implemented"
	return nil
}

// Update the task data.
func (tm *taskManager) Update(ctx context.Context, task *api.Task) error {
	_ = "STUB: not implemented"
	return nil
}

// Close shuts down the task manager, blocking until it is closed.
func (tm *taskManager) Close() error { _ = "STUB: not implemented"; return nil }

func (tm *taskManager) Logs(ctx context.Context, options api.LogSubscriptionOptions, publisher exec.LogPublisher) {
	_ = "STUB: not implemented"
	return
}

// no logs available

func (tm *taskManager) run(ctx context.Context) { _ = "STUB: not implemented"; return }

// cancel all child operations on exit.

// true if the task was updated.

// closure  picks up current value of cancel.

// prime the pump

// always check for shutdown before running.

// a little questionable
// ignore run request and handle shutdown

// Several variables need to be snapshotted for the closure below.
// fork for the closure
// clone the task before dispatch

// capture state of update for goroutine

// before we do anything, update the task for the controller.
// always update the controller before running.

// always report the status if we get one back. This
// returns to the manager loop, then reports the status
// upstream.

// not opctx, since that may have been cancelled.

// This branch is always executed when an operations completes. The
// goal is to decide whether or not we re-dispatch the operation.

// re-enable the shutdown branch
// no dispatch if we are in shutdown.

// wait till getting pumped via update.

// TODO(stevvooe): Add exponential backoff with random jitter
// here. For now, this backoff is enough to keep the task
// manager from running away with the CPU.

// repump this branch, with no err

// no log in this case

// ignore the update

// overwrite our status, as it is canonical.

// we have accepted the task update

// cancel outstanding if necessary.

// If this channel op fails, it means there is already a
// message on the run queue.

// cancel outstanding operation.

// subtle: after a cancellation, we want to avoid busy wait
// here. this gets renabled in the errs branch and we'll come
// back around and try shutdown again.
// turn off this branch until op proceeds
// wait until operation actually exits.

// disable everything, and prepare for closing.
