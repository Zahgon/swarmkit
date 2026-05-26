package taskinit

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/orchestrator"
	"github.com/moby/swarmkit/v2/manager/orchestrator/restart"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

// InitHandler defines orchestrator's action to fix tasks at start.
type InitHandler interface {
	IsRelatedService(service *api.Service) bool
	FixTask(ctx context.Context, batch *store.Batch, t *api.Task)
	SlotTuple(t *api.Task) orchestrator.SlotTuple
}

// CheckTasks fixes tasks in the store before orchestrator runs. The previous leader might
// not have finished processing their updates and left them in an inconsistent state.
func CheckTasks(ctx context.Context, s *store.MemoryStore, readTx store.ReadTx, initHandler InitHandler, startSupervisor restart.SupervisorInterface) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(aluzzardi): We should NOT retrieve the service here.

// Service was deleted

// handle task updates from agent which should have been triggered by task update events

// desired state ready is a transient state that it should be started.
// however previous leader may not have started it, retry start here

// TODO(aluzzardi): This is shady as well. We should have a more generic condition.

// Start now

// Find the most current spec version. That's the only one
// we care about for the purpose of reconstructing restart
// history.

// Create a new slice with just the current spec version tasks.

// Sort by creation timestamp

// All up-to-date tasks in this instance except the first one
// should be considered restarted.

type tasksByCreationTimestamp []*api.Task

func (t tasksByCreationTimestamp) Len() int { _ = "STUB: not implemented"; return 0 }

func (t tasksByCreationTimestamp) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (t tasksByCreationTimestamp) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
