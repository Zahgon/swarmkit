package scheduler

import (
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/constraint"
)

// Filter checks whether the given task can run on the given node.
// A filter may only operate
type Filter interface {
	// SetTask returns true when the filter is enabled for a given task
	// and assigns the task to the filter. It returns false if the filter
	// isn't applicable to this task.  For instance, a constraints filter
	// would return `false` if the task doesn't contain any constraints.
	SetTask(*api.Task) bool

	// Check returns true if the task assigned by SetTask can be scheduled
	// into the given node. This function should not be called if SetTask
	// returned false.
	Check(*NodeInfo) bool

	// Explain what a failure of this filter means
	Explain(nodes int) string
}

// ReadyFilter checks that the node is ready to schedule tasks.
type ReadyFilter struct {
}

// SetTask returns true when the filter is enabled for a given task.
func (f *ReadyFilter) SetTask(_ *api.Task) bool {
	_ = "STUB: not implemented"

	// Check returns true if the task can be scheduled into the given node.
	return false
}

func (f *ReadyFilter) Check(n *NodeInfo) bool { _ = "STUB: not implemented"; return false }

// Explain returns an explanation of a failure.
func (f *ReadyFilter) Explain(nodes int) string { _ = "STUB: not implemented"; return "" }

// ResourceFilter checks that the node has enough resources available to run
// the task.
type ResourceFilter struct {
	reservations *api.Resources
}

// SetTask returns true when the filter is enabled for a given task.
func (f *ResourceFilter) SetTask(t *api.Task) bool { _ = "STUB: not implemented"; return false }

// Check returns true if the task can be scheduled into the given node.
func (f *ResourceFilter) Check(n *NodeInfo) bool { _ = "STUB: not implemented"; return false }

// Explain returns an explanation of a failure.
func (f *ResourceFilter) Explain(nodes int) string { _ = "STUB: not implemented"; return "" }

// PluginFilter checks that the node has a specific volume plugin installed
type PluginFilter struct {
	t *api.Task
}

func referencesVolumePlugin(mount api.Mount) bool { _ = "STUB: not implemented"; return false }

// SetTask returns true when the filter is enabled for a given task.
func (f *PluginFilter) SetTask(t *api.Task) bool { _ = "STUB: not implemented"; return false }

// Check returns true if the task can be scheduled into the given node.
// TODO(amitshukla): investigate storing Plugins as a map so it can be easily probed
func (f *PluginFilter) Check(n *NodeInfo) bool { _ = "STUB: not implemented"; return false }

// If the node is not running Engine, plugins are not
// supported.

// Get list of plugins on the node

// Check if all volume plugins required by task are installed on node

// Check if all network plugins required by task are installed on node

// It's possible that the LogDriver object does not carry a name, just some
// configuration options. In that case, the plugin filter shouldn't fail to
// schedule the task

// If there are no log driver types in the list at all, most likely this is
// an older daemon that did not report this information. In this case don't filter

// pluginExistsOnNode returns true if the (pluginName, pluginType) pair is present in nodePlugins
func (f *PluginFilter) pluginExistsOnNode(pluginType string, pluginName string, nodePlugins []api.PluginDescription) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

// This does not use the reference package to avoid the
// overhead of parsing references as part of the scheduling
// loop. This is okay only because plugin names are a very
// strict subset of the reference grammar that is always
// name:tag.

// Explain returns an explanation of a failure.
func (f *PluginFilter) Explain(nodes int) string { _ = "STUB: not implemented"; return "" }

// ConstraintFilter selects only nodes that match certain labels.
type ConstraintFilter struct {
	constraints []constraint.Constraint
}

// SetTask returns true when the filter is enable for a given task.
func (f *ConstraintFilter) SetTask(t *api.Task) bool { _ = "STUB: not implemented"; return false }

// constraints have been validated at controlapi
// if in any case it finds an error here, treat this task
// as constraint filter disabled.

// Check returns true if the task's constraint is supported by the given node.
func (f *ConstraintFilter) Check(n *NodeInfo) bool { _ = "STUB: not implemented"; return false }

// Explain returns an explanation of a failure.
func (f *ConstraintFilter) Explain(nodes int) string { _ = "STUB: not implemented"; return "" }

// PlatformFilter selects only nodes that run the required platform.
type PlatformFilter struct {
	supportedPlatforms []*api.Platform
}

// SetTask returns true when the filter is enabled for a given task.
func (f *PlatformFilter) SetTask(t *api.Task) bool { _ = "STUB: not implemented"; return false }

// copy the platform information

// Check returns true if the task can be scheduled into the given node.
func (f *PlatformFilter) Check(n *NodeInfo) bool {
	_ = "STUB: not implemented"
	// if the supportedPlatforms field is empty, then either it wasn't
	// provided or there are no constraints
	return false
}

// check if the platform for the node is supported

func (f *PlatformFilter) platformEqual(imgPlatform, nodePlatform api.Platform) bool {
	_ = "STUB: not implemented"
	// normalize "x86_64" architectures to "amd64"
	return false
}

// normalize "aarch64" architectures to "arm64"

// Explain returns an explanation of a failure.
func (f *PlatformFilter) Explain(nodes int) string { _ = "STUB: not implemented"; return "" }

// HostPortFilter checks that the node has a specific port available.
type HostPortFilter struct {
	t *api.Task
}

// SetTask returns true when the filter is enabled for a given task.
func (f *HostPortFilter) SetTask(t *api.Task) bool { _ = "STUB: not implemented"; return false }

// Check returns true if the task can be scheduled into the given node.
func (f *HostPortFilter) Check(n *NodeInfo) bool { _ = "STUB: not implemented"; return false }

// Explain returns an explanation of a failure.
func (f *HostPortFilter) Explain(nodes int) string { _ = "STUB: not implemented"; return "" }

// MaxReplicasFilter selects only nodes that does not exceed max replicas per node.
type MaxReplicasFilter struct {
	t *api.Task
}

// SetTask returns true when max replicas per node filter > 0 for a given task.
func (f *MaxReplicasFilter) SetTask(t *api.Task) bool { _ = "STUB: not implemented"; return false }

// Check returns true if there is less active (assigned or pre-assigned) tasks for this service on current node than set to MaxReplicas limit
func (f *MaxReplicasFilter) Check(n *NodeInfo) bool { _ = "STUB: not implemented"; return false }

// Explain returns an explanation of a failure.
func (f *MaxReplicasFilter) Explain(_ int) string { _ = "STUB: not implemented"; return "" }

type VolumesFilter struct {
	vs *volumeSet
	t  *api.Task

	// requestedVolumes is a set of volumes requested by the task. This can
	// include either volume names or volume groups. Volume groups, as in the
	// Mount.Source field, are prefixed with "group:"
	requestedVolumes []*api.Mount
}

func (f *VolumesFilter) SetTask(t *api.Task) bool {
	_ = "STUB: not implemented"
	// if there is no volume Manager, skip this filter always
	return false
}

// reset requestedVolumes every time we set a task, so we don't
// accidentally append to the last task's set of requested volumes.

// t should never be nil, but we should ensure that it is not just in case
// we make mistakes in the future.

// hasCSI will be set true if one of the mounts is a CSI-type mount.

func (f *VolumesFilter) Check(nodeInfo *NodeInfo) bool { _ = "STUB: not implemented"; return false }

func (f *VolumesFilter) Explain(nodes int) string { _ = "STUB: not implemented"; return "" }
