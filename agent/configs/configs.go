package configs

import (
	"sync"

	"github.com/moby/swarmkit/v2/agent/exec"
	"github.com/moby/swarmkit/v2/api"
)

// configs is a map that keeps all the currently available configs to the agent
// mapped by config ID.
type configs struct {
	mu sync.RWMutex
	m  map[string]*api.Config
}

// NewManager returns a place to store configs.
func NewManager() exec.ConfigsManager { _ = "STUB: not implemented"; return *new(exec.ConfigsManager) }

// Get returns a config by ID.  If the config doesn't exist, returns nil.
func (r *configs) Get(configID string) (*api.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add adds one or more configs to the config map.
func (r *configs) Add(configs ...api.Config) { _ = "STUB: not implemented"; return }

// Remove removes one or more configs by ID from the config map. Succeeds
// whether or not the given IDs are in the map.
func (r *configs) Remove(configs []string) { _ = "STUB: not implemented"; return }

// Reset removes all the configs.
func (r *configs) Reset() { _ = "STUB: not implemented"; return }

// taskRestrictedConfigsProvider restricts the ids to the task.
type taskRestrictedConfigsProvider struct {
	configs   exec.ConfigGetter
	configIDs map[string]struct{} // allow list of config ids
}

func (sp *taskRestrictedConfigsProvider) Get(configID string) (*api.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Restrict provides a getter that only allows access to the configs
// referenced by the task.
func Restrict(configs exec.ConfigGetter, t *api.Task) exec.ConfigGetter {
	_ = "STUB: not implemented"
	return *new(exec.ConfigGetter)
}
