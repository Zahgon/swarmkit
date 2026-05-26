package secrets

import (
	"sync"

	"github.com/moby/swarmkit/v2/agent/exec"
	"github.com/moby/swarmkit/v2/api"
)

// secrets is a map that keeps all the currently available secrets to the agent
// mapped by secret ID.
type secrets struct {
	mu sync.RWMutex
	m  map[string]*api.Secret
}

// NewManager returns a place to store secrets.
func NewManager() exec.SecretsManager { _ = "STUB: not implemented"; return *new(exec.SecretsManager) }

// Get returns a secret by ID.  If the secret doesn't exist, returns nil.
func (s *secrets) Get(secretID string) (*api.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add adds one or more secrets to the secret map.
func (s *secrets) Add(secrets ...api.Secret) { _ = "STUB: not implemented"; return }

// Remove removes one or more secrets by ID from the secret map.  Succeeds
// whether or not the given IDs are in the map.
func (s *secrets) Remove(secrets []string) { _ = "STUB: not implemented"; return }

// Reset removes all the secrets.
func (s *secrets) Reset() { _ = "STUB: not implemented"; return }

// taskRestrictedSecretsProvider restricts the ids to the task.
type taskRestrictedSecretsProvider struct {
	secrets   exec.SecretGetter
	secretIDs map[string]struct{} // allow list of secret ids
	taskID    string              // ID of the task the provider restricts for
}

func (sp *taskRestrictedSecretsProvider) Get(secretID string) (*api.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// First check if the secret is available with the task specific ID, which is the concatenation
// of the original secret ID and the task ID with a dot in between.
// That is the case when a secret driver has returned DoNotReuse == true for a secret value.

// Otherwise, which is the default case, the secret is retrieved by its original ID.

// For all intents and purposes, the rest of the flow should deal with the original secret ID.

// Restrict provides a getter that only allows access to the secrets
// referenced by the task.
func Restrict(secrets exec.SecretGetter, t *api.Task) exec.SecretGetter {
	_ = "STUB: not implemented"
	return *new(exec.SecretGetter)
}
