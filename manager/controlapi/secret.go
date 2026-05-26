package controlapi

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
)

// assumes spec is not nil
func secretFromSecretSpec(spec *api.SecretSpec) *api.Secret { _ = "STUB: not implemented"; return nil }

// GetSecret returns a `GetSecretResponse` with a `Secret` with the same
// id as `GetSecretRequest.SecretID`
// - Returns `NotFound` if the Secret with the given id is not found.
// - Returns `InvalidArgument` if the `GetSecretRequest.SecretID` is empty.
// - Returns an error if getting fails.
func (s *Server) GetSecret(_ context.Context, request *api.GetSecretRequest) (*api.GetSecretResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// clean the actual secret data so it's never returned

// UpdateSecret updates a Secret referenced by SecretID with the given SecretSpec.
// - Returns `NotFound` if the Secret is not found.
// - Returns `InvalidArgument` if the SecretSpec is malformed or anything other than Labels is changed
// - Returns an error if the update fails.
func (s *Server) UpdateSecret(ctx context.Context, request *api.UpdateSecretRequest) (*api.UpdateSecretResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if the Name is different than the current name, or the secret is non-nil and different
// than the current secret

// We only allow updating Labels

// WARN: we should never return the actual secret data here. We need to redact the private fields first.

// ListSecrets returns a `ListSecretResponse` with a list all non-internal `Secret`s being
// managed, or all secrets matching any name in `ListSecretsRequest.Names`, any
// name prefix in `ListSecretsRequest.NamePrefixes`, any id in
// `ListSecretsRequest.SecretIDs`, or any id prefix in `ListSecretsRequest.IDPrefixes`.
// - Returns an error if listing fails.
func (s *Server) ListSecrets(_ context.Context, request *api.ListSecretsRequest) (*api.ListSecretsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// return all secrets that match either any of the names or any of the name prefixes (why would you give both?)

// strip secret data from the secret, filter by label, and filter out all internal secrets

// clean the actual secret data so it's never returned

// CreateSecret creates and returns a `CreateSecretResponse` with a `Secret` based
// on the provided `CreateSecretRequest.SecretSpec`.
//   - Returns `InvalidArgument` if the `CreateSecretRequest.SecretSpec` is malformed,
//     or if the secret data is too long or contains invalid characters.
//   - Returns an error if the creation fails.
func (s *Server) CreateSecret(ctx context.Context, request *api.CreateSecretRequest) (*api.CreateSecretResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check that the requested driver is valid

// the store will handle name conflicts

// clean the actual secret data so it's never returned

// RemoveSecret removes the secret referenced by `RemoveSecretRequest.ID`.
// - Returns `InvalidArgument` if `RemoveSecretRequest.ID` is empty.
// - Returns `NotFound` if the a secret named `RemoveSecretRequest.ID` is not found.
// - Returns `SecretInUse` if the secret is currently in use
// - Returns an error if the deletion fails.
func (s *Server) RemoveSecret(ctx context.Context, request *api.RemoveSecretRequest) (*api.RemoveSecretResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if the secret exists

// Check if any services currently reference this secret, return error if so

func validateSecretSpec(spec *api.SecretSpec) error { _ = "STUB: not implemented"; return nil }

// Check if secret driver is defined

// Ensure secret driver has a name
