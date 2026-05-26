package replicated

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/orchestrator/restart"
	"github.com/moby/swarmkit/v2/manager/orchestrator/update"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

// An Orchestrator runs a reconciliation loop to create and destroy
// tasks as necessary for the replicated services.
type Orchestrator struct {
	store *store.MemoryStore

	reconcileServices map[string]*api.Service
	restartTasks      map[string]struct{}

	// stopChan signals to the state machine to stop running.
	stopChan chan struct{}
	// doneChan is closed when the state machine terminates.
	doneChan chan struct{}

	updater  *update.Supervisor
	restarts *restart.Supervisor

	cluster *api.Cluster // local cluster instance
}

// NewReplicatedOrchestrator creates a new replicated Orchestrator.
func NewReplicatedOrchestrator(store *store.MemoryStore) *Orchestrator {
	_ = "STUB: not implemented"
	return nil
}

// Run contains the orchestrator event loop. It runs until Stop is called.
func (r *Orchestrator) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Watch changes to services and tasks

// Balance existing services and drain initial tasks attached to invalid
// nodes

// TODO(stevvooe): Use ctx to limit running time of operation.

// Stop stops the orchestrator.
func (r *Orchestrator) Stop() { _ = "STUB: not implemented"; return }

func (r *Orchestrator) tick(ctx context.Context) {
	_ = "STUB: not implemented"
	// tickTasks must be called first, so we respond to task-level changes
	// before performing service reconciliation.
	return
}
