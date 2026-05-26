package store

import "github.com/moby/swarmkit/v2/api"

// By is an interface type passed to Find methods. Implementations must be
// defined in this package.
type By interface {
	// isBy allows this interface to only be satisfied by certain internal
	// types.
	isBy()
}

type byAll struct{}

func (a byAll) isBy() {
	_ = "STUB: not implemented"

	// All is an argument that can be passed to find to list all items in the
	// set.
	return
}

var All byAll

type byNamePrefix string

func (b byNamePrefix) isBy() {
	_ = "STUB: not implemented"

	// ByNamePrefix creates an object to pass to Find to select by query.
	return
}

func ByNamePrefix(namePrefix string) By { _ = "STUB: not implemented"; return *new(By) }

type byIDPrefix string

func (b byIDPrefix) isBy() {
	_ = "STUB: not implemented"

	// ByIDPrefix creates an object to pass to Find to select by query.
	return
}

func ByIDPrefix(idPrefix string) By { _ = "STUB: not implemented"; return *new(By) }

type byName string

func (b byName) isBy() {
	_ = "STUB: not implemented"

	// ByName creates an object to pass to Find to select by name.
	return
}

func ByName(name string) By { _ = "STUB: not implemented"; return *new(By) }

type byService string

func (b byService) isBy() { _ = "STUB: not implemented"; return }

type byRuntime string

func (b byRuntime) isBy() {
	_ = "STUB: not implemented"

	// ByRuntime creates an object to pass to Find to select by runtime.
	return
}

func ByRuntime(runtime string) By {
	_ = "STUB: not implemented"
	return *

	// ByServiceID creates an object to pass to Find to select by service.
	new(By)
}

func ByServiceID(serviceID string) By { _ = "STUB: not implemented"; return *new(By) }

type byNode string

func (b byNode) isBy() {
	_ = "STUB: not implemented"

	// ByNodeID creates an object to pass to Find to select by node.
	return
}

func ByNodeID(nodeID string) By { _ = "STUB: not implemented"; return *new(By) }

type bySlot struct {
	serviceID string
	slot      uint64
}

func (b bySlot) isBy() {
	_ = "STUB: not implemented"

	// BySlot creates an object to pass to Find to select by slot.
	return
}

func BySlot(serviceID string, slot uint64) By { _ = "STUB: not implemented"; return *new(By) }

type byDesiredState api.TaskState

func (b byDesiredState) isBy() {
	_ = "STUB: not implemented"

	// ByDesiredState creates an object to pass to Find to select by desired state.
	return
}

func ByDesiredState(state api.TaskState) By { _ = "STUB: not implemented"; return *new(By) }

type byTaskState api.TaskState

func (b byTaskState) isBy() {
	_ = "STUB: not implemented"

	// ByTaskState creates an object to pass to Find to select by task state.
	return
}

func ByTaskState(state api.TaskState) By { _ = "STUB: not implemented"; return *new(By) }

type byRole api.NodeRole

func (b byRole) isBy() {
	_ = "STUB: not implemented"

	// ByRole creates an object to pass to Find to select by role.
	return
}

func ByRole(role api.NodeRole) By { _ = "STUB: not implemented"; return *new(By) }

type byMembership api.NodeSpec_Membership

func (b byMembership) isBy() {
	_ = "STUB: not implemented"

	// ByMembership creates an object to pass to Find to select by Membership.
	return
}

func ByMembership(membership api.NodeSpec_Membership) By {
	_ = "STUB: not implemented"
	return *new(By)
}

type byReferencedNetworkID string

func (b byReferencedNetworkID) isBy() {
	_ = "STUB: not implemented"

	// ByReferencedNetworkID creates an object to pass to Find to search for a
	// service or task that references a network with the given ID.
	return
}

func ByReferencedNetworkID(networkID string) By { _ = "STUB: not implemented"; return *new(By) }

type byReferencedSecretID string

func (b byReferencedSecretID) isBy() {
	_ = "STUB: not implemented"

	// ByReferencedSecretID creates an object to pass to Find to search for a
	// service or task that references a secret with the given ID.
	return
}

func ByReferencedSecretID(secretID string) By { _ = "STUB: not implemented"; return *new(By) }

type byReferencedConfigID string

func (b byReferencedConfigID) isBy() {
	_ = "STUB: not implemented"

	// ByReferencedConfigID creates an object to pass to Find to search for a
	// service or task that references a config with the given ID.
	return
}

func ByReferencedConfigID(configID string) By { _ = "STUB: not implemented"; return *new(By) }

type byVolumeAttachment string

func (b byVolumeAttachment) isBy() {
	_ = "STUB: not implemented"

	// ByVolumeAttachment creates an object to pass to Find to search for a Task
	// that has been assigned the given ID.
	return
}

func ByVolumeAttachment(volumeID string) By { _ = "STUB: not implemented"; return *new(By) }

type byKind string

func (b byKind) isBy() {
	_ = "STUB: not implemented"

	// ByKind creates an object to pass to Find to search for a Resource of a
	// particular kind.
	return
}

func ByKind(kind string) By { _ = "STUB: not implemented"; return *new(By) }

type byCustom struct {
	objType string
	index   string
	value   string
}

func (b byCustom) isBy() {
	_ = "STUB: not implemented"

	// ByCustom creates an object to pass to Find to search a custom index.
	return
}

func ByCustom(objType, index, value string) By { _ = "STUB: not implemented"; return *new(By) }

type byCustomPrefix struct {
	objType string
	index   string
	value   string
}

func (b byCustomPrefix) isBy() {
	_ = "STUB: not implemented"

	// ByCustomPrefix creates an object to pass to Find to search a custom index by
	// a value prefix.
	return
}

func ByCustomPrefix(objType, index, value string) By { _ = "STUB: not implemented"; return *new(By) }

// ByVolumeGroup creates an object to pass to Find to search for volumes
// belonging to a particular group.
func ByVolumeGroup(group string) By { _ = "STUB: not implemented"; return *new(By) }

type byVolumeGroup string

func (b byVolumeGroup) isBy() {
	_ = "STUB: not implemented"

	// ByDriver creates an object to pass to Find to search for objects using a
	// specific driver.
	return
}

func ByDriver(driver string) By { _ = "STUB: not implemented"; return *new(By) }

type byDriver string

func (b byDriver) isBy() { _ = "STUB: not implemented"; return }
