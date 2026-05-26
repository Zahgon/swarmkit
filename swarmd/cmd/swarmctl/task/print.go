package task

import (
	"github.com/moby/swarmkit/swarmd/cmd/swarmctl/common"
	"github.com/moby/swarmkit/v2/api"
)

type tasksBySlot []*api.Task

func (t tasksBySlot) Len() int { _ = "STUB: not implemented"; return 0 }

func (t tasksBySlot) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (t tasksBySlot) Less(i, j int) bool {
	_ = "STUB: not implemented"
	// Sort by slot.
	return false
}

// If same slot, sort by most recent.

// Print prints a list of tasks.
func Print(tasks []*api.Task, all bool, res *common.Resolver) { _ = "STUB: not implemented"; return }
