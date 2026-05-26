package dockerexec

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	engineapi "github.com/docker/docker/client"
	"github.com/moby/swarmkit/v2/agent/exec"
	"github.com/moby/swarmkit/v2/api"
)

// containerAdapter conducts remote operations for a container. All calls
// are mostly naked calls to the client API, seeded with information from
// containerConfig.
type containerAdapter struct {
	client    engineapi.APIClient
	container *containerConfig
	secrets   exec.SecretGetter
}

func newContainerAdapter(client engineapi.APIClient, nodeDescription *api.NodeDescription, task *api.Task, secrets exec.SecretGetter) (*containerAdapter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func noopPrivilegeFn() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (c *containerConfig) imagePullOptions() types.ImagePullOptions {
	_ = "STUB: not implemented"
	return *new(types.ImagePullOptions)
}

// if the image needs to be pulled, the auth config will be retrieved and updated

func (c *containerAdapter) pullImage(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// limit pull progress logs unless the status changes

// if we have progress details, we have everything we need

// first, log the image and status

// then, if we have progress, log the progress

// sometimes, we get no useful information at all, and add no fields

// if the final stream object contained an error, return it

func (c *containerAdapter) createNetworks(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *containerAdapter) removeNetworks(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *containerAdapter) create(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *containerAdapter) start(ctx context.Context) error {
	_ = "STUB: not implemented"
	// TODO(nishanttotla): Consider adding checkpoint handling later
	return nil
}

func (c *containerAdapter) inspect(ctx context.Context) (types.ContainerJSON, error) {
	_ = "STUB: not implemented"
	return *new(types.ContainerJSON), nil
}

// events issues a call to the events API and returns a channel with all
// events. The stream of events can be shutdown by cancelling the context.
//
// A chan struct{} is returned that will be closed if the event processing
// fails and needs to be restarted.
func (c *containerAdapter) events(ctx context.Context) (<-chan events.Message, <-chan struct{}, error) {
	_ = "STUB: not implemented"
	// TODO(stevvooe): Move this to a single, global event dispatch. For
	// now, we create a connection per container.
	return nil, nil, nil
}

// TODO(stevvooe): For long running tasks, it is likely that we will have
// to restart this under failure.

// exit

func (c *containerAdapter) shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Default stop grace period to 10s.
	// TODO(thaJeztah): this should probably not set a default and leave it to the daemon to pick
	// a default (which could be in the container's config or daemon config)
	// see https://github.com/docker/cli/commit/86c30e6a0d153c2a99d9d12385179b22e8b7b935
	return nil
}

func (c *containerAdapter) terminate(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *containerAdapter) remove(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *containerAdapter) createVolumes(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Create plugin volumes that are embedded inside a Mount
	return nil
}

// we create volumes when there is a volume driver available volume options

// TODO(amitshukla): Today, volume create through the engine api does not return an error
// when the named volume with the same parameters already exists.
// It returns an error if the driver name is different - that is a valid error

func (c *containerAdapter) logs(ctx context.Context, options api.LogSubscriptionOptions) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// See protobuf documentation for details of how this works.

// empty == all

// TODO(mrjana/stevvooe): There is no proper error code for network not found
// error in engine-api. Resort to string matching until engine-api is fixed.

func isActiveEndpointError(err error) bool { _ = "STUB: not implemented"; return false }

func isNetworkExistError(err error, name string) bool { _ = "STUB: not implemented"; return false }

func isContainerCreateNameConflict(err error) bool { _ = "STUB: not implemented"; return false }

func isUnknownContainer(err error) bool { _ = "STUB: not implemented"; return false }

func isStoppedContainer(err error) bool { _ = "STUB: not implemented"; return false }
