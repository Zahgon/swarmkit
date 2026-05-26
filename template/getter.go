package template

import (
	"github.com/moby/swarmkit/v2/agent/exec"
	"github.com/moby/swarmkit/v2/api"
)

type templatedSecretGetter struct {
	dependencies exec.DependencyGetter
	t            *api.Task
	node         *api.NodeDescription
}

// NewTemplatedSecretGetter returns a SecretGetter that evaluates templates.
func NewTemplatedSecretGetter(dependencies exec.DependencyGetter, t *api.Task, node *api.NodeDescription) exec.SecretGetter {
	_ = "STUB: not implemented"
	return *new(exec.SecretGetter)
}

func (t templatedSecretGetter) Get(secretID string) (*api.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TemplatedConfigGetter is a ConfigGetter with an additional method to expose
// whether a config contains sensitive data.
type TemplatedConfigGetter interface {
	exec.ConfigGetter

	// GetAndFlagSecretData returns the interpolated config, and also
	// returns true if the config has been interpolated with data from a
	// secret. In this case, the config should be handled specially and
	// should not be written to disk.
	GetAndFlagSecretData(configID string) (*api.Config, bool, error)
}

type templatedConfigGetter struct {
	dependencies exec.DependencyGetter
	t            *api.Task
	node         *api.NodeDescription
}

// NewTemplatedConfigGetter returns a ConfigGetter that evaluates templates.
func NewTemplatedConfigGetter(dependencies exec.DependencyGetter, t *api.Task, node *api.NodeDescription) TemplatedConfigGetter {
	_ = "STUB: not implemented"
	return *new(TemplatedConfigGetter)
}

func (t templatedConfigGetter) Get(configID string) (*api.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t templatedConfigGetter) GetAndFlagSecretData(configID string) (*api.Config, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

type templatedDependencyGetter struct {
	secrets exec.SecretGetter
	configs TemplatedConfigGetter
	volumes exec.VolumeGetter
}

// NewTemplatedDependencyGetter returns a DependencyGetter that evaluates templates.
func NewTemplatedDependencyGetter(dependencies exec.DependencyGetter, t *api.Task, node *api.NodeDescription) exec.DependencyGetter {
	_ = "STUB: not implemented"
	return *new(exec.DependencyGetter)
}

func (t templatedDependencyGetter) Secrets() exec.SecretGetter {
	_ = "STUB: not implemented"
	return *new(exec.SecretGetter)
}

func (t templatedDependencyGetter) Configs() exec.ConfigGetter {
	_ = "STUB: not implemented"
	return *new(exec.ConfigGetter)
}

func (t templatedDependencyGetter) Volumes() exec.VolumeGetter {
	_ = "STUB: not implemented"
	// volumes are not templated, but we include that call (and pass it
	// straight through to the underlying getter) in order to fulfill the
	// DependencyGetter interface.
	return *new(exec.VolumeGetter)
}
