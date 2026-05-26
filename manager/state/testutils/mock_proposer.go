package testutils

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/state"
)

// MockProposer is a simple proposer implementation for use in tests.
type MockProposer struct {
	index   uint64
	changes []state.Change
}

// ProposeValue propagates a value. In this mock implementation, it just stores
// the value locally.
func (mp *MockProposer) ProposeValue(_ context.Context, storeAction []api.StoreAction, cb func()) error {
	_ = "STUB: not implemented"
	return nil
}

// GetVersion returns the current version.
func (mp *MockProposer) GetVersion() *api.Version { _ = "STUB: not implemented"; return nil }

// ChangesBetween returns changes after "from" up to and including "to".
func (mp *MockProposer) ChangesBetween(from, to api.Version) ([]state.Change, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
