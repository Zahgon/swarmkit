package plugin

import (
	"context"
	"sync"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/node/plugin"
)

const (
	// DockerCSIPluginCap is the capability name of the plugins we use with the
	// PluginGetter to get only the plugins we need. The full name of the
	// plugin interface is "docker.csinode/1.0". This gets only plugins with
	// Node capabilities.
	DockerCSIPluginCap = "csinode"
)

// Manager manages the multiple CSI plugins that may be in use on the
// node. Manager should be thread-safe.
type Manager interface {
	// Get gets the plugin with the given name
	Get(name string) (NodePlugin, error)

	// NodeInfo returns the NodeCSIInfo for every active plugin.
	NodeInfo(ctx context.Context) ([]*api.NodeCSIInfo, error)
}

type pluginManager struct {
	plugins   map[string]NodePlugin
	pluginsMu sync.Mutex

	// newNodePluginFunc usually points to NewNodePlugin. However, for testing,
	// NewNodePlugin can be swapped out with a function that creates fake node
	// plugins
	newNodePluginFunc func(string, plugin.AddrPlugin, SecretGetter) NodePlugin

	// secrets is a SecretGetter for use by node plugins.
	secrets SecretGetter

	pg plugin.Getter
}

func NewManager(pg plugin.Getter, secrets SecretGetter) Manager {
	_ = "STUB: not implemented"
	return *new(Manager)
}

func (pm *pluginManager) Get(name string) (NodePlugin, error) {
	_ = "STUB: not implemented"
	return *new(NodePlugin), nil
}

func (pm *pluginManager) NodeInfo(ctx context.Context) ([]*api.NodeCSIInfo, error) {
	_ = "STUB: not implemented"
	// TODO(dperny): do not acquire this lock for the duration of the the
	// function call. that's too long and too blocking.
	return nil, nil
}

// first, we should make sure all of the plugins are initialized. do this
// by looking up all the current plugins with DockerCSIPluginCap.

// TODO(dperny): use this opportunity to drop plugins that we're
// tracking but which no longer exist.

// we don't actually need the plugin returned, we just need it loaded
// as a side effect.

// skip any plugin that returns an error

// getPlugin looks up the plugin with the specified name. Loads the plugin if
// not yet loaded.
//
// pm.pluginsMu must be obtained before calling this method.
func (pm *pluginManager) getPlugin(name string) (NodePlugin, error) {
	_ = "STUB: not implemented"
	return *new(NodePlugin), nil
}
