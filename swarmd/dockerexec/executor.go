package dockerexec

import (
	"context"
	"sync"

	engineapi "github.com/docker/docker/client"
	"github.com/moby/swarmkit/v2/agent/exec"
	"github.com/moby/swarmkit/v2/api"
)

type executor struct {
	client           engineapi.APIClient
	secrets          exec.SecretsManager
	genericResources []*api.GenericResource
	mutex            sync.Mutex // This mutex protects the following node field
	node             *api.NodeDescription
}

// NewExecutor returns an executor from the docker client.
func NewExecutor(client engineapi.APIClient, genericResources []*api.GenericResource) exec.Executor {
	_ = "STUB: not implemented"
	return *new(exec.Executor)
}

// Describe returns the underlying node description from the docker client.
func (e *executor) Describe(ctx context.Context) (*api.NodeDescription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// add v1 plugins to 'plugins'

// Add builtin driver "overlay" (the only builtin multi-host driver) to
// the plugin list by default.

// retrieve v2 plugins

// add v2 plugins to 'plugins'

// parse []string labels into a map[string]string

// this will take the last value in the list for a given key
// ideally, one shouldn't assign multiple values to the same key

// Save the node information in the executor field

func (e *executor) Configure(_ context.Context, _ *api.Node) error {
	_ = "STUB: not implemented"

	// Controller returns a docker container controller.
	return nil
}

func (e *executor) Controller(t *api.Task) (exec.Controller, error) {
	_ = "STUB: not implemented"
	// Get the node description from the executor field
	return *new(exec.Controller), nil
}

func (e *executor) SetNetworkBootstrapKeys([]*api.EncryptionKey) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *executor) Secrets() exec.SecretsManager {
	_ = "STUB: not implemented"
	return *new(exec.SecretsManager)
}

type sortedPlugins []api.PluginDescription

func (sp sortedPlugins) Len() int { _ = "STUB: not implemented"; return 0 }

func (sp sortedPlugins) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (sp sortedPlugins) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
