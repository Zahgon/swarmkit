package template

import (
	"github.com/moby/swarmkit/v2/agent/exec"
	"github.com/moby/swarmkit/v2/api"
)

// ExpandContainerSpec expands templated fields in the runtime using the task
// state and the node where it is scheduled to run.
// Templating is all evaluated on the agent-side, before execution.
//
// Note that these are projected only on runtime values, since active task
// values are typically manipulated in the manager.
func ExpandContainerSpec(n *api.NodeDescription, t *api.Task) (*api.ContainerSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// For now, we only allow templating of string-based mount fields

func expandMounts(ctx Context, mounts []api.Mount) ([]api.Mount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func expandMap(ctx Context, m map[string]string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func expandEnv(ctx Context, values []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func expandPayload(ctx *PayloadContext, payload []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExpandSecretSpec expands the template inside the secret payload, if any.
// Templating is evaluated on the agent-side.
func ExpandSecretSpec(s *api.Secret, node *api.NodeDescription, t *api.Task, dependencies exec.DependencyGetter) (*api.SecretSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExpandConfigSpec expands the template inside the config payload, if any.
// Templating is evaluated on the agent-side.
func ExpandConfigSpec(c *api.Config, node *api.NodeDescription, t *api.Task, dependencies exec.DependencyGetter) (*api.ConfigSpec, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}
