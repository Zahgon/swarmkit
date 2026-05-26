// Package naming centralizes the naming of SwarmKit objects.
package naming

import (
	"errors"

	"github.com/moby/swarmkit/v2/api"
)

var (
	errUnknownRuntime = errors.New("unrecognized runtime")
)

// Task returns the task name from Annotations.Name,
// and, in case Annotations.Name is missing, fallback
// to construct the name from other information.
func Task(t *api.Task) string { _ = "STUB: not implemented"; return "" }

// if set, use the container Annotations.Name field, set in the orchestrator.

// when no slot id is assigned, we assume that this is node-bound task.

// fallback to service.instance.id.

// TODO(stevvooe): Consolidate "Hostname" style validation here.

// Runtime returns the runtime name from a given spec.
func Runtime(t api.TaskSpec) (string, error) { _ = "STUB: not implemented"; return "", nil }
