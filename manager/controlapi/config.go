package controlapi

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
)

// MaxConfigSize is the maximum byte length of the `Config.Spec.Data` field.
const MaxConfigSize = 1000 * 1024 // 1000KB

// assumes spec is not nil
func configFromConfigSpec(spec *api.ConfigSpec) *api.Config { _ = "STUB: not implemented"; return nil }

// GetConfig returns a `GetConfigResponse` with a `Config` with the same
// id as `GetConfigRequest.ConfigID`
// - Returns `NotFound` if the Config with the given id is not found.
// - Returns `InvalidArgument` if the `GetConfigRequest.ConfigID` is empty.
// - Returns an error if getting fails.
func (s *Server) GetConfig(_ context.Context, request *api.GetConfigRequest) (*api.GetConfigResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateConfig updates a Config referenced by ConfigID with the given ConfigSpec.
// - Returns `NotFound` if the Config is not found.
// - Returns `InvalidArgument` if the ConfigSpec is malformed or anything other than Labels is changed
// - Returns an error if the update fails.
func (s *Server) UpdateConfig(ctx context.Context, request *api.UpdateConfigRequest) (*api.UpdateConfigResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if the Name is different than the current name, or the config is non-nil and different
// than the current config

// We only allow updating Labels

// ListConfigs returns a `ListConfigResponse` with a list all non-internal `Config`s being
// managed, or all configs matching any name in `ListConfigsRequest.Names`, any
// name prefix in `ListConfigsRequest.NamePrefixes`, any id in
// `ListConfigsRequest.ConfigIDs`, or any id prefix in `ListConfigsRequest.IDPrefixes`.
// - Returns an error if listing fails.
func (s *Server) ListConfigs(_ context.Context, request *api.ListConfigsRequest) (*api.ListConfigsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// return all configs that match either any of the names or any of the name prefixes (why would you give both?)

// filter by label

// CreateConfig creates and returns a `CreateConfigResponse` with a `Config` based
// on the provided `CreateConfigRequest.ConfigSpec`.
//   - Returns `InvalidArgument` if the `CreateConfigRequest.ConfigSpec` is malformed,
//     or if the config data is too long or contains invalid characters.
//   - Returns an error if the creation fails.
func (s *Server) CreateConfig(ctx context.Context, request *api.CreateConfigRequest) (*api.CreateConfigResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the store will handle name conflicts

// RemoveConfig removes the config referenced by `RemoveConfigRequest.ID`.
// - Returns `InvalidArgument` if `RemoveConfigRequest.ID` is empty.
// - Returns `NotFound` if the a config named `RemoveConfigRequest.ID` is not found.
// - Returns `ConfigInUse` if the config is currently in use
// - Returns an error if the deletion fails.
func (s *Server) RemoveConfig(ctx context.Context, request *api.RemoveConfigRequest) (*api.RemoveConfigResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if the config exists

// Check if any services currently reference this config, return error if so

func validateConfigSpec(spec *api.ConfigSpec) error { _ = "STUB: not implemented"; return nil }
