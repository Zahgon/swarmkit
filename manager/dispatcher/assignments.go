package dispatcher

import (
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/drivers"
	"github.com/moby/swarmkit/v2/manager/state/store"
	"github.com/sirupsen/logrus"
)

type typeAndID struct {
	id      string
	objType api.ResourceType
}

type assignmentSet struct {
	nodeID   string
	dp       *drivers.DriverProvider
	tasksMap map[string]*api.Task
	// volumesMap keeps track of the VolumePublishStatus of the given volumes.
	// this tells us both which volumes are assigned to the node, and what the
	// last known VolumePublishStatus was, so we can understand if we need to
	// send an update.
	volumesMap map[string]*api.VolumePublishStatus
	// tasksUsingDependency tracks both tasks and volumes using a given
	// dependency. this works because the ID generated for swarm comes from a
	// large enough space that it is reliably astronomically unlikely that IDs
	// will ever collide.
	tasksUsingDependency map[typeAndID]map[string]struct{}
	changes              map[typeAndID]*api.AssignmentChange
	log                  *logrus.Entry
}

func newAssignmentSet(nodeID string, log *logrus.Entry, dp *drivers.DriverProvider) *assignmentSet {
	_ = "STUB: not implemented"
	return nil
}

func assignSecret(a *assignmentSet, readTx store.ReadTx, mapKey typeAndID, t *api.Task) {
	_ = "STUB: not implemented"
	return
}

// If the secret should not be reused for other tasks, give it a unique ID
// for the task to allow different values for different tasks.

// Give the secret a new ID and mark it as internal

// Create a new mapKey with the new ID and insert it into the
// dependencies map for the task.  This will make the changes map
// contain an entry with the new ID rather than the original one.

func assignConfig(a *assignmentSet, readTx store.ReadTx, mapKey typeAndID) {
	_ = "STUB: not implemented"
	return
}

func (a *assignmentSet) addTaskDependencies(readTx store.ReadTx, t *api.Task) {
	_ = "STUB: not implemented"
	// first, we go through all ResourceReferences, which give us the necessary
	// information about which secrets and configs are in use.
	return
}

// if there are no tasks using this dependency yet, then we can assign
// it.

// otherwise, we don't need to add a new assignment. we just need to
// track the fact that another task is now using this dependency.

// This checks for the presence of each task in the dependency map for the
// secret. This is currently only done for secrets since the other types of
// dependencies do not support driver plugins. Arguably, the same task would
// not have the same secret as a dependency more than once, but this check
// makes sure the task only gets the secret assigned once.

func (a *assignmentSet) releaseDependency(mapKey typeAndID, assignment *api.Assignment, taskID string) bool {
	_ = "STUB: not implemented"
	return false
}

// No tasks are using the dependency anymore

// releaseTaskDependencies needs a store transaction because volumes have
// associated Secrets which need to be released.
func (a *assignmentSet) releaseTaskDependencies(_ store.ReadTx, t *api.Task) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *assignmentSet) addOrUpdateTask(readTx store.ReadTx, t *api.Task) bool {
	_ = "STUB: not implemented"
	// We only care about tasks that are ASSIGNED or higher.
	return false
}

// States ASSIGNED and below are set by the orchestrator/scheduler,
// not the agent, so tasks in these states need to be sent to the
// agent even if nothing else has changed.

// this update should not trigger a task change for the agent

// If this task got updated to a final state, let's release
// the dependencies that are being used by the task

// If releasing the dependencies caused us to
// remove something from the assignment set,
// mark one modification.

// If this task wasn't part of the assignment set before, and it's <= RUNNING
// add the dependencies it references to the assignment.
// Task states > RUNNING are worker reported only, are never created in
// a > RUNNING state.

// addOrUpdateVolume tracks a Volume assigned to a node.
func (a *assignmentSet) addOrUpdateVolume(readTx store.ReadTx, v *api.Volume) bool {
	_ = "STUB: not implemented"
	return false
}

// if there is no publishStatus for this Volume on this Node, or if the
// Volume has not yet been published to this node, then we do not need to
// track this assignment.

// check if we are already tracking this volume, and what its old status
// is. if the states are identical, then we don't have any update to make.

// if the volume has already been confirmed as unpublished, we can stop
// tracking it and remove its dependencies.

// we can call assignSecret with task being nil, but it does mean
// that any secret that uses a driver will not work. we'll call
// that a limitation of volumes for now.

// volumes are sent to nodes as VolumeAssignments. This is because a node
// needs node-specific information (the PublishContext from
// ControllerPublishVolume).

// assignmentChange is the whole assignment without the action, which we
// will set next

// if we're in state PENDING_NODE_UNPUBLISH, we actually need to send a
// remove message. we do this every time, even if the node never got the
// first add assignment. This is because the node might not know that it
// has a volume published; for example, the node may be restarting, and
// the in-memory store does not have knowledge of the volume.

func (a *assignmentSet) removeVolume(_ store.ReadTx, v *api.Volume) bool {
	_ = "STUB: not implemented"
	return false
}

// if the volume does exists, we can release its secrets

// we don't need to add a removal message. the removal of the
// VolumeAssignment will have already happened.

func (a *assignmentSet) removeTask(readTx store.ReadTx, t *api.Task) bool {
	_ = "STUB: not implemented"
	return false
}

// Release the dependencies being used by this task.
// Ignoring the return here. We will always mark this as a
// modification, since a task is being removed.

func (a *assignmentSet) message() api.AssignmentsMessage {
	_ = "STUB: not implemented"
	return *new(api.AssignmentsMessage)
}

// The the set of changes is reinitialized to prepare for formation
// of the next message.

// secret populates the secret value from raft store. For external secrets, the value is populated
// from the secret driver. The function returns: a secret object; an indication of whether the value
// is to be reused across tasks; and an error if the secret is not found in the store, if the secret
// driver responds with one or if the payload does not pass validation.
func (a *assignmentSet) secret(readTx store.ReadTx, task *api.Task, secretID string) (*api.Secret, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// Assign the secret
