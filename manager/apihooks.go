package manager

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/allocator/networkallocator"
)

func (m *Manager) networkAllocator() networkallocator.NetworkAllocator {
	_ = "STUB: not implemented"
	return *new(networkallocator.NetworkAllocator)
}

func (m *Manager) OnGetNetwork(ctx context.Context, n *api.Network, appdataTypeURL string, appdata []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) OnListNetworks(ctx context.Context, networks []*api.Network, appdataTypeURL string, appdata []byte) error {
	_ = "STUB: not implemented"
	return nil
}
