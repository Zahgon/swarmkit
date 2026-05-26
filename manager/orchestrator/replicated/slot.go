package replicated

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/orchestrator"
)

type slotsByRunningState []orchestrator.Slot

func (is slotsByRunningState) Len() int      { _ = "STUB: not implemented"; return 0 }
func (is slotsByRunningState) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Less returns true if the first task should be preferred over the second task,
// all other things being equal in terms of node balance.
func (is slotsByRunningState) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Use Slot number as a tie-breaker to prefer to remove tasks in reverse
// order of Slot number. This would help us avoid unnecessary master
// migration when scaling down a stateful service because the master
// task of a stateful service is usually in a low numbered Slot.

type slotWithIndex struct {
	slot orchestrator.Slot

	// index is a counter that counts this task as the nth instance of
	// the service on its node. This is used for sorting the tasks so that
	// when scaling down we leave tasks more evenly balanced.
	index int
}

type slotsByIndex []slotWithIndex

func (is slotsByIndex) Len() int      { _ = "STUB: not implemented"; return 0 }
func (is slotsByIndex) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (is slotsByIndex) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// updatableAndDeadSlots returns two maps of slots. The first contains slots
// that have at least one task with a desired state above NEW and lesser or
// equal to RUNNING, or a task that shouldn't be restarted. The second contains
// all other slots with at least one task.
func (r *Orchestrator) updatableAndDeadSlots(ctx context.Context, service *api.Service) (map[uint64]orchestrator.Slot, map[uint64]orchestrator.Slot, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// SlotTuple returns a slot tuple for the replicated service task.
func (r *Orchestrator) SlotTuple(t *api.Task) orchestrator.SlotTuple {
	_ = "STUB: not implemented"
	return *new(orchestrator.SlotTuple)
}
