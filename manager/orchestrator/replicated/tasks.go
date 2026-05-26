package replicated

import (
	"context"

	"github.com/docker/go-events"
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

// This file provides task-level orchestration. It observes changes to task
// and node state and kills/recreates tasks if necessary. This is distinct from
// service-level reconciliation, which observes changes to services and creates
// and/or kills tasks to match the service definition.

func (r *Orchestrator) initTasks(ctx context.Context, readTx store.ReadTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Orchestrator) handleTaskEvent(ctx context.Context, event events.Event) {
	_ = "STUB: not implemented"
	return
}

func (r *Orchestrator) tickTasks(ctx context.Context) { _ = "STUB: not implemented"; return }

// TODO(aaronl): optimistic update?

// Restart task if applicable

func (r *Orchestrator) restartTasksByNodeID(ctx context.Context, nodeID string) {
	_ = "STUB: not implemented"
	return
}

func (r *Orchestrator) handleNodeChange(ctx context.Context, n *api.Node) {
	_ = "STUB: not implemented"
	return
}

// handleTaskChange defines what orchestrator does when a task is updated by agent.
func (r *Orchestrator) handleTaskChange(_ context.Context, t *api.Task) {
	_ = "STUB: not implemented"
	// If we already set the desired state past TaskStateRunning, there is no
	// further action necessary.
	return
}

// FixTask validates a task with the current cluster settings, and takes
// action to make it conformant. it's called at orchestrator initialization.
func (r *Orchestrator) FixTask(_ context.Context, batch *store.Batch, t *api.Task) {
	_ = "STUB: not implemented"
	// If we already set the desired state past TaskStateRunning, there is no
	// further action necessary.
	return
}
