package controlapi

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
)

// CreateExtension creates an `Extension` based on the provided `CreateExtensionRequest.Extension`
// and returns a `CreateExtensionResponse`.
//   - Returns `InvalidArgument` if the `CreateExtensionRequest.Extension` is malformed,
//     or fails validation.
//   - Returns an error if the creation fails.
func (s *Server) CreateExtension(ctx context.Context, request *api.CreateExtensionRequest) (*api.CreateExtensionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetExtension returns a `GetExtensionResponse` with a `Extension` with the same
// id as `GetExtensionRequest.extension_id`
// - Returns `NotFound` if the Extension with the given id is not found.
// - Returns `InvalidArgument` if the `GetExtensionRequest.extension_id` is empty.
// - Returns an error if the get fails.
func (s *Server) GetExtension(_ context.Context, request *api.GetExtensionRequest) (*api.GetExtensionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoveExtension removes the extension referenced by `RemoveExtensionRequest.ID`.
// - Returns `InvalidArgument` if `RemoveExtensionRequest.extension_id` is empty.
// - Returns `NotFound` if the an extension named `RemoveExtensionRequest.extension_id` is not found.
// - Returns an error if the deletion fails.
func (s *Server) RemoveExtension(ctx context.Context, request *api.RemoveExtensionRequest) (*api.RemoveExtensionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if the extension exists

// Check if any resources of this type present in the store, return error if so

// Number of resources for an extension could be quite large.
// Show a limited number of resources for debugging.
