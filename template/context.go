package template

import (
	"github.com/moby/swarmkit/v2/agent/exec"
	"github.com/moby/swarmkit/v2/api"
)

// Platform holds information about the underlying platform of the node
type Platform struct {
	Architecture string
	OS           string
}

// Context defines the strict set of values that can be injected into a
// template expression in SwarmKit data structure.
// NOTE: Be very careful adding any fields to this structure with types
// that have methods defined on them. The template would be able to
// invoke those methods.
type Context struct {
	Service struct {
		ID     string
		Name   string
		Labels map[string]string
	}

	Node struct {
		ID       string
		Hostname string
		Platform Platform
	}

	Task struct {
		ID   string
		Name string
		Slot string

		// NOTE(stevvooe): Why no labels here? Tasks don't actually have labels
		// (from a user perspective). The labels are part of the container! If
		// one wants to use labels for templating, use service labels!
	}
}

// NewContext returns a new template context from the data available in the
// task and the node where it is scheduled to run.
// The provided context can then be used to populate runtime values in a
// ContainerSpec.
func NewContext(n *api.NodeDescription, t *api.Task) (ctx Context) {
	_ = "STUB: not implemented"
	return *new(Context)
}

// Add node information to context only if we have them available

// fall back to node id for slot when there is no slot

// Expand treats the string s as a template and populates it with values from
// the context.
func (ctx *Context) Expand(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// PayloadContext provides a context for expanding a config or secret payload.
// NOTE: Be very careful adding any fields to this structure with types
// that have methods defined on them. The template would be able to
// invoke those methods.
type PayloadContext struct {
	Context

	t                 *api.Task
	restrictedSecrets exec.SecretGetter
	restrictedConfigs exec.ConfigGetter
	sensitive         bool
}

func (ctx *PayloadContext) secretGetter(target string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ctx *PayloadContext) configGetter(target string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ctx *PayloadContext) envGetter(variable string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// NewPayloadContextFromTask returns a new template context from the data
// available in the task and the node where it is scheduled to run.
// This context also provides access to the configs
// and secrets that the task has access to. The provided context can then
// be used to populate runtime values in a templated config or secret.
func NewPayloadContextFromTask(node *api.NodeDescription, t *api.Task, dependencies exec.DependencyGetter) (ctx PayloadContext) {
	_ = "STUB: not implemented"
	return *new(PayloadContext)
}

// Expand treats the string s as a template and populates it with values from
// the context.
func (ctx *PayloadContext) Expand(s string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
