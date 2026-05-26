package scheduler

import (
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

// the scheduler package does double duty -- in addition to choosing nodes, it
// must also choose volumes. this is because volumes are fungible, and can be
// scheduled to several nodes, and used by several tasks. we should endeavor to
// spread tasks across volumes, like we spread nodes. on the positive side,
// unlike nodes, volumes are not heirarchical. that is, we don't need to
// spread across multiple levels of a tree, only a flat set.

// volumeSet is the set of all volumes currently managed
type volumeSet struct {
	// volumes is a mapping of volume IDs to volumeInfo
	volumes map[string]volumeInfo
	// byGroup is a mapping from a volume group name to a set of volumes in
	// that group
	byGroup map[string]map[string]struct{}
	// byName is a mapping of volume names to swarmkit volume IDs.
	byName map[string]string
}

// volumeUsage contains information about the usage of a Volume by a specific
// task.
type volumeUsage struct {
	nodeID   string
	readOnly bool
}

// volumeInfo contains scheduler information about a given volume
type volumeInfo struct {
	volume *api.Volume
	tasks  map[string]volumeUsage
	// nodes is a set of nodes a volume is in use on. it maps a node ID to a
	// reference count for how many tasks are using the volume on that node.
	nodes map[string]int
}

func newVolumeSet() *volumeSet { _ = "STUB: not implemented"; return nil }

// getVolume returns the volume object for the given ID as stored in the
// volumeSet, or nil if none exists.
//
//nolint:unused // TODO(thaJeztah) this is currently unused: is it safe to remove?
func (vs *volumeSet) getVolume(id string) *api.Volume { _ = "STUB: not implemented"; return nil }

func (vs *volumeSet) addOrUpdateVolume(v *api.Volume) { _ = "STUB: not implemented"; return }

// if the volume already exists in the set, then only update the volume
// object, not the tasks map.

//nolint:unused // only used in tests.
func (vs *volumeSet) removeVolume(volumeID string) { _ = "STUB: not implemented"; return }

// if the volume exists in the set, look up its group ID and remove it
// from the byGroup mapping as well

// chooseTaskVolumes selects a set of VolumeAttachments for the task on the
// given node. it expects that the node was already validated to have the
// necessary volumes, but it will return an error if a full set of volumes is
// not available.
func (vs *volumeSet) chooseTaskVolumes(task *api.Task, nodeInfo *NodeInfo) ([]*api.VolumeAttachment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// we'll reserve volumes in this loop, but release all of our reservations
// before we finish. the caller will need to call reserveTaskVolumes after
// calling this function
// TODO(dperny): this is probably not optimal

// TODO(dperny): handle non-container tasks

// TODO(dperny): return structured error types, instead of
// error strings

// reserveTaskVolumes identifies all volumes currently in use on a task and
// marks them in the volumeSet as in use.
func (vs *volumeSet) reserveTaskVolumes(task *api.Task) { _ = "STUB: not implemented"; return }

// we shouldn't need to handle non-container tasks because those tasks
// won't have any entries in task.Volumes.

func (vs *volumeSet) reserveVolume(volumeID, taskID, nodeID string, readOnly bool) {
	_ = "STUB: not implemented"
	return
}

// TODO(dperny): don't just return nothing.

// increment the reference count for this node.

func (vs *volumeSet) releaseVolume(volumeID, taskID string) { _ = "STUB: not implemented"; return }

// if the volume isn't in the set, no action to take.

// decrement the reference count for this task's node

// this is probably an unnecessarily high level of caution, but make
// sure we don't go below zero on node count.

// freeVolumes finds volumes that are no longer in use on some nodes, and
// updates them to be unpublished from those nodes.
//
// TODO(dperny): this is messy and has a lot of overhead. it should be reworked
// to something more streamlined.
func (vs *volumeSet) freeVolumes(batch *store.Batch) error { _ = "STUB: not implemented"; return nil }

// when we are freeing a volume, we may update more than one of the
// volume's PublishStatuses. this means we can't simply put the
// Update call inside of the if statement; we need to know if we've
// changed anything once we've checked *all* of the statuses.

// isVolumeAvailableOnNode checks if a volume satisfying the given mount is
// available on the given node.
//
// Returns the ID of the volume available, or an empty string if no such volume
// is found.
func (vs *volumeSet) isVolumeAvailableOnNode(mount *api.Mount, node *NodeInfo) string {
	_ = "STUB: not implemented"
	return ""

	// first, discern whether we're looking for a group or a volume
	// try trimming off the "group:" prefix. if the resulting string is
	// different from the input string (meaning something has been trimmed),
	// then this volume is actually a volume group.
}

// if there are no volumes of this group specified, then no volume
// meets the moutn criteria.

// iterate through all ids in the group, checking if any one meets the
// spec.

// if it's not a group, it's a name. resolve the volume name to its ID

// checkVolume checks if an individual volume with the given ID can be placed
// on the given node.
func (vs *volumeSet) checkVolume(id string, info *NodeInfo, readOnly bool) bool {
	_ = "STUB: not implemented"
	return false

	// first, check if the volume's availability is even Active. If not. no
	// reason to bother with anything further.
}

// get the node topology for this volume

// get the topology for this volume's driver on this node

// check if the volume is available on this node. a volume's
// availability on a node depends on its accessible topology, how it's
// already being used, and how this task intends to use it.

// if the volume is not in use on this node already, then it can't
// be used here.

// even if the volume is currently on this node, or it has multi-node
// access, the volume sharing needs to be compatible.

// if the volume sharing is none, then the volume cannot be
// used by another task

// if the mount is not ReadOnly, and the volume has a writer, then
// we this volume does not work.

// if the volume sharing is read-only, then the Mount must also
// be read-only

// then, do the quick check of whether this volume is in the topology.  if
// the volume has an AccessibleTopology, and it does not lie within the
// node's topology, then this volume won't fit.

// hasWriter is a helper function that returns true if at least one task is
// using this volume not in read-only mode.
func hasWriter(info volumeInfo) bool { _ = "STUB: not implemented"; return false }
