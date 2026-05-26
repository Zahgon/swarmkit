package testutils

import (
	"context"
	"sync"
	"testing"

	"github.com/moby/swarmkit/v2/agent/exec"
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/ca"
)

// TestExecutor is executor for integration tests
type TestExecutor struct {
	mu   sync.Mutex
	desc *api.NodeDescription
}

// Describe just returns empty NodeDescription.
func (e *TestExecutor) Describe(_ context.Context) (*api.NodeDescription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Configure does nothing.
func (e *TestExecutor) Configure(_ context.Context, _ *api.Node) error {
	_ = "STUB: not implemented"

	// SetNetworkBootstrapKeys does nothing.
	return nil
}

func (e *TestExecutor) SetNetworkBootstrapKeys([]*api.EncryptionKey) error {
	_ = "STUB: not implemented"

	// Controller returns TestController.
	return nil
}

func (e *TestExecutor) Controller(_ *api.Task) (exec.Controller, error) {
	_ = "STUB: not implemented"
	return *new(exec.Controller), nil
}

// UpdateNodeDescription sets the node description on the test executor
func (e *TestExecutor) UpdateNodeDescription(newDesc *api.NodeDescription) {
	_ = "STUB: not implemented"
	return
}

// TestController is dummy channel based controller for tests.
type TestController struct {
	ch        chan struct{}
	closeOnce sync.Once
}

// Update does nothing.
func (t *TestController) Update(_ context.Context, _ *api.Task) error {
	_ = "STUB: not implemented"

	// Prepare does nothing.
	return nil
}

func (t *TestController) Prepare(_ context.Context) error {
	_ = "STUB: not implemented"

	// Start does nothing.
	return nil
}

func (t *TestController) Start(_ context.Context) error {
	_ = "STUB: not implemented"

	// Wait waits on internal channel.
	return nil
}

func (t *TestController) Wait(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Shutdown closes internal channel
func (t *TestController) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

// Terminate closes internal channel if it wasn't closed before.
func (t *TestController) Terminate(_ context.Context) error { _ = "STUB: not implemented"; return nil }

// Remove does nothing.
func (t *TestController) Remove(_ context.Context) error {
	_ = "STUB: not implemented"

	// Close does nothing.
	return nil
}

func (t *TestController) Close() error { _ = "STUB: not implemented"; return nil }

// SessionHandler is an injectable function that can be used handle session requests
type SessionHandler func(*api.SessionRequest, api.Dispatcher_SessionServer) error

// MockDispatcher is a fake dispatcher that one agent at a time can connect to
type MockDispatcher struct {
	mu             sync.Mutex
	sessionCh      chan *api.SessionMessage
	openSession    *api.SessionRequest
	closedSessions []*api.SessionRequest
	sessionHandler SessionHandler

	Addr string
}

// UpdateTaskStatus is not implemented
func (m *MockDispatcher) UpdateTaskStatus(context.Context, *api.UpdateTaskStatusRequest) (*api.UpdateTaskStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MockDispatcher) UpdateVolumeStatus(context.Context, *api.UpdateVolumeStatusRequest) (*api.UpdateVolumeStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil,

		// Tasks keeps an open stream until canceled
		nil
}

func (m *MockDispatcher) Tasks(_ *api.TasksRequest, stream api.Dispatcher_TasksServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Assignments keeps an open stream until canceled
func (m *MockDispatcher) Assignments(_ *api.AssignmentsRequest, stream api.Dispatcher_AssignmentsServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Heartbeat always successfully heartbeats
func (m *MockDispatcher) Heartbeat(context.Context, *api.HeartbeatRequest) (*api.HeartbeatResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Session allows a session to be established, and sends the node info
func (m *MockDispatcher) Session(r *api.SessionRequest, stream api.Dispatcher_SessionServer) error {
	_ = "STUB: not implemented"
	return nil
}

// only overwrite session if it hasn't changed

// send the initial message first

// GetSessions return all the established and closed sessions
func (m *MockDispatcher) GetSessions() (*api.SessionRequest, []*api.SessionRequest) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SessionMessageChannel returns a writable channel to inject session messages
func (m *MockDispatcher) SessionMessageChannel() chan<- *api.SessionMessage {
	_ = "STUB: not implemented"
	return nil

	// SetSessionHandler lets you inject a custom function to handle session requests
}

func (m *MockDispatcher) SetSessionHandler(s SessionHandler) { _ = "STUB: not implemented"; return }

// NewMockDispatcher starts and returns a mock dispatcher instance that can be connected to
func NewMockDispatcher(t *testing.T, secConfig *ca.SecurityConfig, local bool) (*MockDispatcher, func()) {
	_ = "STUB: not implemented"
	return nil, nil
}
