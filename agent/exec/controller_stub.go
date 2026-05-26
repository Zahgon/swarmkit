package exec

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
)

// StubController implements the Controller interface,
// but allows you to specify behaviors for each of its methods.
type StubController struct {
	Controller
	UpdateFn    func(ctx context.Context, t *api.Task) error
	PrepareFn   func(ctx context.Context) error
	StartFn     func(ctx context.Context) error
	WaitFn      func(ctx context.Context) error
	ShutdownFn  func(ctx context.Context) error
	TerminateFn func(ctx context.Context) error
	RemoveFn    func(ctx context.Context) error
	CloseFn     func() error
	calls       map[string]int
}

// NewStubController returns an initialized StubController
func NewStubController() *StubController { _ = "STUB: not implemented"; return nil }

// If function A calls updateCountsForSelf,
// The callCount[A] value will be incremented
func (sc *StubController) called() { _ = "STUB: not implemented"; return }

// longName looks like 'github.com/moby/swarmkit/agent/exec.(*StubController).Prepare:1'

// Update is part of the Controller interface
func (sc *StubController) Update(ctx context.Context, t *api.Task) error {
	_ = "STUB: not implemented"
	return nil
}

// Prepare is part of the Controller interface
func (sc *StubController) Prepare(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Start is part of the Controller interface
func (sc *StubController) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Wait is part of the Controller interface
func (sc *StubController) Wait(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Shutdown is part of the Controller interface
func (sc *StubController) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Terminate is part of the Controller interface
func (sc *StubController) Terminate(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove is part of the Controller interface
func (sc *StubController) Remove(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Close is part of the Controller interface
func (sc *StubController) Close() error { _ = "STUB: not implemented"; return nil }
