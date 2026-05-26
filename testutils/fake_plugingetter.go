package testutils

import (
	"net"

	"github.com/moby/swarmkit/v2/node/plugin"
)

const DockerCSIPluginNodeCap = "csinode"
const DockerCSIPluginControllerCap = "csicontroller"

type FakePluginGetter struct {
	Plugins map[string]*FakePlugin
}

var _ plugin.Getter = &FakePluginGetter{}

func (f *FakePluginGetter) Get(name, capability string) (plugin.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(plugin.Plugin), nil
}

// GetAllManagedPluginsByCap returns all of the fake's plugins. If capability
// is anything other than DockerCSIPluginCap, it returns nothing.
func (f *FakePluginGetter) GetAllManagedPluginsByCap(capability string) []plugin.Plugin {
	_ = "STUB: not implemented"
	return nil
}

type FakePlugin struct {
	PluginName string
	PluginAddr net.Addr
	Scope      string
}

var _ plugin.AddrPlugin = &FakePlugin{}

func (f *FakePlugin) Name() string { _ = "STUB: not implemented"; return "" }

func (f *FakePlugin) ScopedPath(path string) string { _ = "STUB: not implemented"; return "" }

func (f *FakePlugin) Client() plugin.Client { _ = "STUB: not implemented"; return *new(plugin.Client) }

func (f *FakePlugin) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }
