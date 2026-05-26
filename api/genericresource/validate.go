package genericresource

import (
	"github.com/moby/swarmkit/v2/api"
)

// ValidateTask validates that the task only uses integers
// for generic resources
func ValidateTask(resources *api.Resources) error { _ = "STUB: not implemented"; return nil }

// HasEnough returns true if node can satisfy the task's GenericResource request
func HasEnough(nodeRes []*api.GenericResource, taskRes *api.GenericResource) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// HasResource checks if there is enough "res" in the "resources" argument
func HasResource(res *api.GenericResource, resources []*api.GenericResource) bool {
	_ = "STUB: not implemented"
	return false
}
