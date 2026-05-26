package exec

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
)

// Controller controls execution of a task.
type Controller interface {
	// Update the task definition seen by the controller. Will return
	// ErrTaskUpdateFailed if the provided task definition changes fields that
	// cannot be changed.
	//
	// Will be ignored if the task has exited.
	Update(ctx context.Context, t *api.Task) error

	// Prepare the task for execution. This should ensure that all resources
	// are created such that a call to start should execute immediately.
	Prepare(ctx context.Context) error

	// Start the target and return when it has started successfully.
	Start(ctx context.Context) error

	// Wait blocks until the target has exited.
	Wait(ctx context.Context) error

	// Shutdown requests to exit the target gracefully.
	Shutdown(ctx context.Context) error

	// Terminate the target.
	Terminate(ctx context.Context) error

	// Remove all resources allocated by the controller.
	Remove(ctx context.Context) error

	// Close closes any ephemeral resources associated with controller instance.
	Close() error
}

// ControllerLogs defines a component that makes logs accessible.
//
// Can usually be accessed on a controller instance via type assertion.
type ControllerLogs interface {
	// Logs will write publisher until the context is cancelled or an error
	// occurs.
	Logs(ctx context.Context, publisher LogPublisher, options api.LogSubscriptionOptions) error
}

// LogPublisher defines the protocol for receiving a log message.
type LogPublisher interface {
	Publish(ctx context.Context, message api.LogMessage) error
}

// LogPublisherFunc implements publisher with just a function.
type LogPublisherFunc func(ctx context.Context, message api.LogMessage) error

// Publish calls the wrapped function.
func (fn LogPublisherFunc) Publish(ctx context.Context, message api.LogMessage) error {
	_ = "STUB: not implemented"
	return nil

	// LogPublisherProvider defines the protocol for receiving a log publisher
}

type LogPublisherProvider interface {
	Publisher(ctx context.Context, subscriptionID string) (LogPublisher, func(), error)
}

// ContainerStatuser reports status of a container.
//
// This can be implemented by controllers or error types.
type ContainerStatuser interface {
	// ContainerStatus returns the status of the target container, if
	// available. When the container is not available, the status will be nil.
	ContainerStatus(ctx context.Context) (*api.ContainerStatus, error)
}

// PortStatuser reports status of ports which are allocated by the executor
type PortStatuser interface {
	// PortStatus returns the status on a list of PortConfigs
	// which are managed at the host level by the controller.
	PortStatus(ctx context.Context) (*api.PortStatus, error)
}

// Resolve attempts to get a controller from the executor and reports the
// correct status depending on the tasks current state according to the result.
//
// Unlike Do, if an error is returned, the status should still be reported. The
// error merely reports the failure at getting the controller.
func Resolve(ctx context.Context, task *api.Task, executor Executor) (Controller, *api.TaskStatus, error) {
	_ = "STUB: not implemented"
	return *new(Controller), nil, nil
}

// depending on the tasks state, a failed controller resolution has varying
// impact. The following expresses that impact.

// before the task has been started, we consider it a rejection.
// if task is running, consider the task has failed
// otherwise keep the existing state

// we always want to proceed to accepted when we resolve the controller

// Do progresses the task state using the controller performing a single
// operation on the controller. The return TaskStatus should be marked as the
// new state of the task.
//
// The returned status should be reported and placed back on to task
// before the next call. The operation can be cancelled by creating a
// cancelling context.
//
// Errors from the task controller will reported on the returned status. Any
// errors coming from this function should not be reported as related to the
// individual task.
//
// If ErrTaskNoop is returned, it means a second call to Do will result in no
// change. If ErrTaskDead is returned, calls to Do will no longer result in any
// action.
func Do(ctx context.Context, task *api.Task, ctlr Controller) (*api.TaskStatus, error) {
	_ = "STUB: not implemented"
	return nil,

		// stay in the current state.
		nil
}

// while we retry on all errors, this allows us to explicitly declare
// retry cases.

// transition moves the task to the next state.

// containerStatus exitCode keeps track of whether or not we've set it in
// this particular method. Eventually, we assemble this as part of a defer.

// returned when a fatal execution of the task is fatal. In this case, we
// proceed to a terminal error state and set the appropriate fields.
//
// Common checks for the nature of an error should be included here. If the
// error is determined not to be fatal for the task,

// make sure we've set the *correct* exit code

// still reported on temporary

// only at this point do we consider the error fatal to the task.

// NOTE(stevvooe): The following switch dictates the terminal failure
// state based on the state in which the failure was encountered.

// below, we have several callbacks that are run after the state transition
// is completed.

// extract the container status from the container, if supported.

// only do this if in an active state

// collect this, if we haven't

// at this point, things have gone fairly wrong. Remain positive
// and let's get something out the door.

// copy it over.

// at this point, we *must* have a containerStatus.

// this branch bounds the largest state achievable in the agent as SHUTDOWN, which
// is exactly the correct behavior for the agent.

// way beyond desired state, pause

// the following states may proceed past desired state.

// The following represent "pause" states. We can only proceed when the
// desired state is beyond our current state.

// terminal states

func logStateChange(ctx context.Context, desired, previous, next api.TaskState) {
	_ = "STUB: not implemented"
	return
}

func contextDoneError(err error) bool { _ = "STUB: not implemented"; return false }
