package genericresource

import (
	"github.com/moby/swarmkit/v2/api"
)

// NewSet creates a set object
func NewSet(key string, vals ...string) []*api.GenericResource {
	_ = "STUB: not implemented"
	return nil
}

// NewString creates a String resource
func NewString(key, val string) *api.GenericResource { _ = "STUB: not implemented"; return nil }

// NewDiscrete creates a Discrete resource
func NewDiscrete(key string, val int64) *api.GenericResource { _ = "STUB: not implemented"; return nil }

// GetResource returns resources from the "resources" parameter matching the kind key
func GetResource(kind string, resources []*api.GenericResource) []*api.GenericResource {
	_ = "STUB: not implemented"
	return nil
}

// ConsumeNodeResources removes "res" from nodeAvailableResources
func ConsumeNodeResources(nodeAvailableResources *[]*api.GenericResource, res []*api.GenericResource) {
	_ = "STUB: not implemented"
	return
}

// If this wasn't the right element then
// we need to continue

// Returns true if the element is to be removed from the list
func remove(na, r *api.GenericResource) bool { _ = "STUB: not implemented"; return false }

// Type change, ignore

// Type change, ignore

// not the right item, ignore
