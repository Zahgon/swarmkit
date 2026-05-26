package equality

import (
	"github.com/moby/swarmkit/v2/api"
)

// TasksEqualStable returns true if the tasks are functionally equal, ignoring status,
// version and other superfluous fields.
//
// This used to decide whether or not to propagate a task update to a controller.
func TasksEqualStable(a, b *api.Task) bool {
	_ = "STUB: not implemented"
	// shallow copy
	return false
}

// TaskStatusesEqualStable compares the task status excluding timestamp fields.
func TaskStatusesEqualStable(a, b *api.TaskStatus) bool { _ = "STUB: not implemented"; return false }

// RootCAEqualStable compares RootCAs, excluding join tokens, which are randomly generated
func RootCAEqualStable(a, b *api.RootCA) bool { _ = "STUB: not implemented"; return false }

// ExternalCAsEqualStable compares lists of external CAs and determines whether they are equal.
func ExternalCAsEqualStable(a, b []*api.ExternalCA) bool {
	_ = "STUB: not implemented"
	// because DeepEqual will treat an empty list and a nil list differently, we want to manually check this first
	return false
}

// The assumption is that each individual api.ExternalCA within both lists are created from deserializing from a
// protobuf, so no special affordances are made to treat a nil map and empty map in the Options field of an
// api.ExternalCA as equivalent.
