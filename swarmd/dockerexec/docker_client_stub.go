package dockerexec

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

// StubAPIClient implements the client.APIClient interface, but allows
// you to specify the behavior of each of the methods.
type StubAPIClient struct {
	client.APIClient
	calls              map[string]int
	ContainerCreateFn  func(_ context.Context, config *container.Config, hostConfig *container.HostConfig, networking *network.NetworkingConfig, platform *v1.Platform, containerName string) (container.CreateResponse, error)
	ContainerInspectFn func(_ context.Context, containerID string) (types.ContainerJSON, error)
	ContainerKillFn    func(_ context.Context, containerID, signal string) error
	ContainerRemoveFn  func(_ context.Context, containerID string, options types.ContainerRemoveOptions) error
	ContainerStartFn   func(_ context.Context, containerID string, options types.ContainerStartOptions) error
	ContainerStopFn    func(_ context.Context, containerID string, options container.StopOptions) error
	ImagePullFn        func(_ context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error)
	EventsFn           func(_ context.Context, options types.EventsOptions) (<-chan events.Message, <-chan error)
}

// NewStubAPIClient returns an initialized StubAPIClient
func NewStubAPIClient() *StubAPIClient { _ = "STUB: not implemented"; return nil }

// If function A calls updateCountsForSelf,
// The callCount[A] value will be incremented
func (sa *StubAPIClient) called() { _ = "STUB: not implemented"; return }

// longName looks like 'github.com/moby/swarmkit/agent/exec.(*StubController).Prepare:1'

// ContainerCreate is part of the APIClient interface
func (sa *StubAPIClient) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networking *network.NetworkingConfig, platform *v1.Platform, containerName string) (container.CreateResponse, error) {
	_ = "STUB: not implemented"
	return *new(container.CreateResponse), nil
}

// ContainerInspect is part of the APIClient interface
func (sa *StubAPIClient) ContainerInspect(ctx context.Context, containerID string) (types.ContainerJSON, error) {
	_ = "STUB: not implemented"
	return *new(types.ContainerJSON), nil
}

// ContainerKill is part of the APIClient interface
func (sa *StubAPIClient) ContainerKill(ctx context.Context, containerID, signal string) error {
	_ = "STUB: not implemented"
	return nil
}

// ContainerRemove is part of the APIClient interface
func (sa *StubAPIClient) ContainerRemove(ctx context.Context, containerID string, options types.ContainerRemoveOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// ContainerStart is part of the APIClient interface
func (sa *StubAPIClient) ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// ContainerStop is part of the APIClient interface
func (sa *StubAPIClient) ContainerStop(ctx context.Context, containerID string, options container.StopOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// ImagePull is part of the APIClient interface
func (sa *StubAPIClient) ImagePull(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// Events is part of the APIClient interface
func (sa *StubAPIClient) Events(ctx context.Context, options types.EventsOptions) (<-chan events.Message, <-chan error) {
	_ = "STUB: not implemented"
	return nil, nil
}
