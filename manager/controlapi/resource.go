package controlapi

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
)

// CreateResource returns a `CreateResourceResponse` after creating a `Resource` based
// on the provided `CreateResourceRequest.Resource`.
//   - Returns `InvalidArgument` if the `CreateResourceRequest.Resource` is malformed,
//     or if the config data is too long or contains invalid characters.
//   - Returns an error if the creation fails.
func (s *Server) CreateResource(ctx context.Context, request *api.CreateResourceRequest) (*api.CreateResourceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// finally, validate that Kind is not an emptystring. We know that creating
// with Kind as empty string should fail at the store level, but to make
// errors clearer, special case this.

// GetResource returns a `GetResourceResponse` with a `Resource` with the same
// id as `GetResourceRequest.Resource`
// - Returns `NotFound` if the Resource with the given id is not found.
// - Returns `InvalidArgument` if the `GetResourceRequest.Resource` is empty.
// - Returns an error if getting fails.
func (s *Server) GetResource(_ context.Context, request *api.GetResourceRequest) (*api.GetResourceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoveResource removes the `Resource` referenced by `RemoveResourceRequest.ResourceID`.
// - Returns `InvalidArgument` if `RemoveResourceRequest.ResourceID` is empty.
// - Returns `NotFound` if the a resource named `RemoveResourceRequest.ResourceID` is not found.
// - Returns an error if the deletion fails.
func (s *Server) RemoveResource(_ context.Context, request *api.RemoveResourceRequest) (*api.RemoveResourceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListResources returns a `ListResourcesResponse` with a list of `Resource`s stored in the raft store,
// or all resources matching any name in `ListConfigsRequest.Names`, any
// name prefix in `ListResourcesRequest.NamePrefixes`, any id in
// `ListResourcesRequest.ResourceIDs`, or any id prefix in `ListResourcesRequest.IDPrefixes`.
// - Returns an error if listing fails.
func (s *Server) ListResources(_ context.Context, request *api.ListResourcesRequest) (*api.ListResourcesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// andKind is set to true if the Extension filter is not the only filter
// being used. If this is the case, we do not have to compare by strings,
// which could be slow.

// if we're filtering on Extensions, then set this to true. If Kind is
// the _only_ kind of filter, we'll set this to false below.

// NOTE(dperny): currently, filtering using store.ByKind would apply an
// Or operation, which means that filtering by kind would return a
// union. However, for Kind filters, we actually want the
// _intersection_; that is, _only_ objects of the specified kind. we
// could dig into the db code to figure out how to write and use an
// efficient And combinator, but I don't have the time nor expertise to
// do so at the moment. instead, we'll filter by kind after the fact.
// however, if there are no other kinds of filters, and we're ONLY
// listing by Kind, we can set that to be the only filter.

// filter by label and extension

// UpdateResource updates the resource with the given `UpdateResourceRequest.Resource.Id` using the given `UpdateResourceRequest.Resource` and returns a `UpdateResourceResponse`.
// - Returns `NotFound` if the Resource with the given `UpdateResourceRequest.Resource.Id` is not found.
// - Returns `InvalidArgument` if the UpdateResourceRequest.Resource.Id` is empty.
// - Returns an error if updating fails.
func (s *Server) UpdateResource(_ context.Context, request *api.UpdateResourceRequest) (*api.UpdateResourceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// only alter the payload if the
