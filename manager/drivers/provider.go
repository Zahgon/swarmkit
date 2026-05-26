package drivers

import (
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/node/plugin"
)

// DriverProvider provides external drivers
type DriverProvider struct {
	pluginGetter plugin.Getter
}

// New returns a new driver provider
func New(pluginGetter plugin.Getter) *DriverProvider { _ = "STUB: not implemented"; return nil }

// NewSecretDriver creates a new driver for fetching secrets
func (m *DriverProvider) NewSecretDriver(driver *api.Driver) (*SecretDriver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Search for the specified plugin
