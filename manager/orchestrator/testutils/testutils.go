package testutils

import (
	"testing"

	"github.com/docker/go-events"
	"github.com/moby/swarmkit/v2/api"
)

// EnsureRuns takes a closure and runs it in a goroutine, blocking until the
// goroutine has had an opportunity to run. It returns a channel which will be
// closed when the provided closure exits.
func EnsureRuns(closure func()) <-chan struct{} { _ = "STUB: not implemented"; return nil }

// WatchTaskCreate waits for a task to be created.
func WatchTaskCreate(t *testing.T, watch chan events.Event) *api.Task {
	_ = "STUB: not implemented"
	return nil
}

// WatchTaskUpdate waits for a task to be updated.
func WatchTaskUpdate(t *testing.T, watch chan events.Event) *api.Task {
	_ = "STUB: not implemented"
	return nil
}

// WatchTaskDelete waits for a task to be deleted.
func WatchTaskDelete(t *testing.T, watch chan events.Event) *api.Task {
	_ = "STUB: not implemented"
	return nil
}

// WatchShutdownTask fails the test if the next event is not a task having its
// desired state changed to Shutdown.
func WatchShutdownTask(t *testing.T, watch chan events.Event) *api.Task {
	_ = "STUB: not implemented"
	return nil
}

// Expect fails the test if the next event is not one of the specified events.
func Expect(t *testing.T, watch chan events.Event, specifiers ...api.Event) {
	_ = "STUB: not implemented"
	return
}

// FatalStack logs the stacks of all goroutines and immediately fails the test.
func FatalStack(t *testing.T, msg string, args ...interface{}) { _ = "STUB: not implemented"; return }
