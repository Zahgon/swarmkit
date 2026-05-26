package genericresource

import (
	"github.com/moby/swarmkit/v2/api"
)

// Claim assigns GenericResources to a task by taking them from the
// node's GenericResource list and storing them in the task's available list
func Claim(nodeAvailableResources, taskAssigned *[]*api.GenericResource,
	taskReservations []*api.GenericResource) error {
	_ = "STUB: not implemented"
	return nil
}

// Select the resources

// ClaimResources adds the specified resources to the task's list
// and removes them from the node's generic resource list
func ClaimResources(nodeAvailableResources, taskAssigned *[]*api.GenericResource,
	resSelected []*api.GenericResource) {
	_ = "STUB: not implemented"
	return
}

func selectNodeResources(nodeRes []*api.GenericResource,
	tr *api.DiscreteGenericResource) ([]*api.GenericResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reclaim adds the resources taken by the task to the node's store
func Reclaim(nodeAvailableResources *[]*api.GenericResource, taskAssigned, nodeRes []*api.GenericResource) error {
	_ = "STUB: not implemented"
	return nil
}

func reclaimResources(nodeAvailableResources *[]*api.GenericResource, taskAssigned []*api.GenericResource) error {
	_ = "STUB: not implemented"
	// The node could have been updated
	return nil
}

// If the resource went down to 0 it's no longer in the
// available list

// Type change

// Type change

// sanitize checks that nodeAvailableResources does not add resources unknown
// to the nodeSpec (nodeRes) or goes over the integer bound specified
// by the spec.
// Note this is because the user is able to update a node's resources
func sanitize(nodeRes []*api.GenericResource, nodeAvailableResources *[]*api.GenericResource) {
	_ = "STUB: not implemented"
	// - We add the sanitized resources at the end, after
	// having removed the elements from the list
	return
}

// - When a set changes to a Discrete we also need
// to make sure that we don't add the Discrete multiple
// time hence, the need of a map to remember that

// Returns true if the element is in nodeRes and "sane"
// Returns false if the element isn't in nodeRes and "sane" and the element(s) that should be replacing it
func sanitizeResource(nodeRes []*api.GenericResource, res *api.GenericResource) (ok bool, nrs []*api.GenericResource) {
	_ = "STUB: not implemented"
	return false, nil
}

// Type change or removed: reset

// Type change: reset

// Amount change: reset

// Type change

// Type change: reset

// Removed
