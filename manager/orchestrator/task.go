package orchestrator

import (
	google_protobuf "github.com/gogo/protobuf/types"
	"github.com/moby/swarmkit/v2/api"
)

// NewTask creates a new task.
func NewTask(cluster *api.Cluster, service *api.Service, slot uint64, nodeID string) *api.Task {
	_ = "STUB: not implemented"
	return nil
}

// use the log driver specific to the task, if we have it.

// pick up the cluster default, if available.
// nil is okay here.

// In global mode we also set the NodeID

// RestartCondition returns the restart condition to apply to this task.
func RestartCondition(task *api.Task) api.RestartPolicy_RestartCondition {
	_ = "STUB: not implemented"
	return *new(api.RestartPolicy_RestartCondition)
}

// IsTaskDirty determines whether a task matches the given service's spec and
// if the given node satisfies the placement constraints.
// Returns false if the spec version didn't change,
// only the task placement constraints changed and the assigned node
// satisfies the new constraints, or the service task spec and the endpoint spec
// didn't change at all.
// Returns true otherwise.
// Note: for non-failed tasks with a container spec runtime that have already
// pulled the required image (i.e., current state is between READY and
// RUNNING inclusively), the value of the `PullOptions` is ignored.
func IsTaskDirty(s *api.Service, t *api.Task, n *api.Node) bool {
	_ = "STUB: not implemented"
	// If the spec version matches, we know the task is not dirty. However,
	// if it does not match, that doesn't mean the task is dirty, since
	// only a portion of the spec is included in the comparison.
	return false
}

// Make a deep copy of the service and task spec for the comparison.

// Task is not dirty if the placement constraints alone changed
// and the node currently assigned can satisfy the changed constraints.

// For non-failed tasks with a container spec runtime that have already
// pulled the required image (i.e., current state is between READY and
// RUNNING inclusively), ignore the value of the `PullOptions` field by
// setting the copied service to have the same PullOptions value as the
// task. A difference in only the `PullOptions` field should not cause
// a running (or ready to run) task to be considered 'dirty' when we
// handle updates.
// See https://github.com/docker/swarmkit/issues/971

// Ignore PullOpts if the task is desired to be in a "runnable" state
// and its last known current state is between READY and RUNNING in
// which case we know that the task either successfully pulled its
// container image or didn't need to.

// Modify the service's container spec.

// Checks if the current assigned node matches the Placement.Constraints
// specified in the task spec for Updater.newService.
func nodeMatches(s *api.Service, n *api.Node) bool { _ = "STUB: not implemented"; return false }

// IsTaskDirtyPlacementConstraintsOnly checks if the Placement field alone
// in the spec has changed.
func IsTaskDirtyPlacementConstraintsOnly(serviceTaskSpec api.TaskSpec, t *api.Task) bool {
	_ = "STUB: not implemented"
	// Compare the task placement constraints.
	return false
}

// Update spec placement to only the fields
// other than the placement constraints in the spec.

// InvalidNode is true if the node is nil, down, or drained
func InvalidNode(n *api.Node) bool { _ = "STUB: not implemented"; return false }

func taskTimestamp(t *api.Task) *google_protobuf.Timestamp { _ = "STUB: not implemented"; return nil }

// TasksByTimestamp sorts tasks by applied timestamp if available, otherwise
// status timestamp.
type TasksByTimestamp []*api.Task

// Len implements the Len method for sorting.
func (t TasksByTimestamp) Len() int {
	_ = "STUB: not implemented"

	// Swap implements the Swap method for sorting.
	return 0
}

func (t TasksByTimestamp) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Less implements the Less method for sorting.
func (t TasksByTimestamp) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
