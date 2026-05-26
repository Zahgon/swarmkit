package csi

import (
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

// SecretProvider is an interface for retrieving secrets to use with CSI calls.
type SecretProvider interface {
	// GetSecret returns the secret with the given ID, or nil if not found.
	GetSecret(id string) *api.Secret
}

type secretProvider struct {
	s *store.MemoryStore
}

func NewSecretProvider(s *store.MemoryStore) SecretProvider {
	_ = "STUB: not implemented"
	return *new(SecretProvider)
}

// GetSecret returns the secret with the given ID, or nil if not found.
//
// This method accesses the store, and so should not be called from inside
// another store transaction
func (p *secretProvider) GetSecret(id string) *api.Secret { _ = "STUB: not implemented"; return nil }
