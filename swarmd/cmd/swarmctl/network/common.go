package network

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
)

// GetNetwork tries to query for a network as an ID and if it can't be
// found tries to query as a name. If the name query returns exactly
// one entry then it is returned to the caller. Otherwise an error is
// returned.
func GetNetwork(ctx context.Context, c api.ControlClient, input string) (*api.Network, error) {
	_ = "STUB: not implemented"
	// GetService to match via full ID.
	return nil, nil
}

// If any error (including NotFound), ListServices to match via full name.

// ResolveServiceNetworks takes a service spec and resolves network names to network IDs.
func ResolveServiceNetworks(ctx context.Context, c api.ControlClient, spec *api.ServiceSpec) error {
	_ = "STUB: not implemented"
	return nil
}
