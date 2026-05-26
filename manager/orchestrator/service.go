package orchestrator

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

// IsReplicatedService checks if a service is a replicated service.
func IsReplicatedService(service *api.Service) bool {
	_ = "STUB: not implemented"
	// service nil validation is required as there are scenarios
	// where service is removed from store
	return false
}

// IsGlobalService checks if the service is a global service.
func IsGlobalService(service *api.Service) bool { _ = "STUB: not implemented"; return false }

// IsReplicatedJob returns true if the service is a replicated job.
func IsReplicatedJob(service *api.Service) bool { _ = "STUB: not implemented"; return false }

// IsGlobalJob returns true if the service is a global job.
func IsGlobalJob(service *api.Service) bool { _ = "STUB: not implemented"; return false }

// SetServiceTasksRemove sets the desired state of tasks associated with a service
// to REMOVE, so that they can be properly shut down by the agent and later removed
// by the task reaper.
func SetServiceTasksRemove(ctx context.Context, s *store.MemoryStore, service *api.Service) {
	_ = "STUB: not implemented"
	return
}

// the task may have changed for some reason in the meantime
// since we read it out, so we need to get from the store again
// within the boundaries of a transaction

// in case the task is deleted

// time travel is not allowed. if the current desired state is
// above the one we're trying to go to we can't go backwards.
// we have nothing to do and we should skip to the next task

// log a warning, though. we shouln't be trying to rewrite
// a state to an earlier state

// update desired state to REMOVE
