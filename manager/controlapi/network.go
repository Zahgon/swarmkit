package controlapi

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
)

func validateIPAMConfiguration(ipamConf *api.IPAMConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) validateIPAM(ipam *api.IPAMOptions) error {
	_ = "STUB: not implemented"

	// It is ok to not specify any IPAM configurations. We
	// will choose good defaults.
	return nil
}

func (s *Server) validateNetworkSpec(spec *api.NetworkSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateNetwork creates and returns a Network based on the provided NetworkSpec.
// - Returns `InvalidArgument` if the NetworkSpec is malformed.
// - Returns an error if the creation fails.
func (s *Server) CreateNetwork(_ context.Context, request *api.CreateNetworkRequest) (*api.CreateNetworkResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(mrjana): Consider using `Name` as a primary key to handle
// duplicate creations. See #65

// GetNetwork returns a Network given a NetworkID.
// - Returns `InvalidArgument` if NetworkID is not provided.
// - Returns `NotFound` if the Network is not found.
func (s *Server) GetNetwork(ctx context.Context, request *api.GetNetworkRequest) (*api.GetNetworkResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoveNetwork removes a Network referenced by NetworkID.
// - Returns `InvalidArgument` if NetworkID is not provided.
// - Returns `NotFound` if the Network is not found.
// - Returns an error if the deletion fails.
func (s *Server) RemoveNetwork(_ context.Context, request *api.RemoveNetworkRequest) (*api.RemoveNetworkResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) removeNetwork(id string) error { _ = "STUB: not implemented"; return nil }

func (s *Server) removeIngressNetwork(id string) error { _ = "STUB: not implemented"; return nil }

func filterNetworks(candidates []*api.Network, filters ...func(*api.Network) bool) []*api.Network {
	_ = "STUB: not implemented"
	return nil
}

// ListNetworks returns a list of all networks.
func (s *Server) ListNetworks(ctx context.Context, request *api.ListNetworksRequest) (*api.ListNetworksResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
