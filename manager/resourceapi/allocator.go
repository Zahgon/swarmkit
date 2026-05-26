package resourceapi

import (
	"context"
	"errors"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

var (
	errInvalidArgument = errors.New("invalid argument")
)

// ResourceAllocator handles resource allocation of cluster entities.
type ResourceAllocator struct {
	store *store.MemoryStore
}

// New returns an instance of the allocator
func New(store *store.MemoryStore) *ResourceAllocator { _ = "STUB: not implemented"; return nil }

// AttachNetwork allows the node to request the resources
// allocation needed for a network attachment on the specific node.
// - Returns `InvalidArgument` if the Spec is malformed.
// - Returns `NotFound` if the Network is not found.
// - Returns `PermissionDenied` if the Network is not manually attachable.
// - Returns an error if the creation fails.
func (ra *ResourceAllocator) AttachNetwork(ctx context.Context, request *api.AttachNetworkRequest) (*api.AttachNetworkResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: Add Network attachment.

// DetachNetwork allows the node to request the release of
// the resources associated to the network attachment.
// - Returns `InvalidArgument` if attachment ID is not provided.
// - Returns `NotFound` if the attachment is not found.
// - Returns an error if the deletion fails.
func (ra *ResourceAllocator) DetachNetwork(ctx context.Context, request *api.DetachNetworkRequest) (*api.DetachNetworkResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
