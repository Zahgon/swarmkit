package agent

import (
	"github.com/moby/swarmkit/v2/agent/exec"
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/node/plugin"
)

type dependencyManager struct {
	secrets exec.SecretsManager
	configs exec.ConfigsManager
	volumes exec.VolumesManager
}

// NewDependencyManager creates a dependency manager object that wraps
// objects which provide access to various dependency types.
func NewDependencyManager(pg plugin.Getter) exec.DependencyManager {
	_ = "STUB: not implemented"
	return *new(exec.DependencyManager)
}

func (d *dependencyManager) Secrets() exec.SecretsManager {
	_ = "STUB: not implemented"
	return *new(exec.SecretsManager)
}

func (d *dependencyManager) Configs() exec.ConfigsManager {
	_ = "STUB: not implemented"
	return *new(exec.ConfigsManager)
}

func (d *dependencyManager) Volumes() exec.VolumesManager {
	_ = "STUB: not implemented"
	return *new(exec.VolumesManager)
}

type dependencyGetter struct {
	secrets exec.SecretGetter
	configs exec.ConfigGetter
	volumes exec.VolumeGetter
}

func (d *dependencyGetter) Secrets() exec.SecretGetter {
	_ = "STUB: not implemented"
	return *new(exec.SecretGetter)
}

func (d *dependencyGetter) Configs() exec.ConfigGetter {
	_ = "STUB: not implemented"
	return *new(exec.ConfigGetter)
}

func (d *dependencyGetter) Volumes() exec.VolumeGetter {
	_ = "STUB: not implemented"

	// Restrict provides getters that only allows access to the dependencies
	// referenced by the task.
	return *new(exec.VolumeGetter)
}

func Restrict(dependencies exec.DependencyManager, t *api.Task) exec.DependencyGetter {
	_ = "STUB: not implemented"
	return *new(exec.DependencyGetter)
}
