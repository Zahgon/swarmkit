package volumeenforcer

import (
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

// VolumeEnforcer is a component, styled off of the ConstraintEnforcer, that
// watches for updates to Volumes, and shuts down tasks if those Volumes are
// being drained.
type VolumeEnforcer struct {
	store    *store.MemoryStore
	stopChan chan struct{}
	doneChan chan struct{}
}

func New(s *store.MemoryStore) *VolumeEnforcer { _ = "STUB: not implemented"; return nil }

func (ve *VolumeEnforcer) Run() { _ = "STUB: not implemented"; return }

func (ve *VolumeEnforcer) Stop() { _ = "STUB: not implemented"; return }

func (ve *VolumeEnforcer) rejectNoncompliantTasks(v *api.Volume) { _ = "STUB: not implemented"; return }

// ignore the error, it only happens if you pass an invalid find by

// skip any tasks we know are already shut down or shutting
// down. Do this before we open the transaction. This saves us
// copying volumeTasks while still avoiding unnecessary
// transactions. we will still need to check again once we
// start the transaction against the latest version of the
// task.

// another check for task liveness.

// as documented in the ConstraintEnforcer:
//
// We set the observed state to
// REJECTED, rather than the desired
// state. Desired state is owned by the
// orchestrator, and setting it directly
// will bypass actions such as
// restarting the task on another node
// (if applicable).
