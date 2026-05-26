package manager

// IsStateDirty returns true if any objects have been added to raft which make
// the state "dirty". Currently, the existence of any object other than the
// default cluster or the local node implies a dirty state.
func (m *Manager) IsStateDirty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Check Nodes and Clusters fields.

// Use reflection to check that other fields don't have values. This
// lets us implement a whitelist-type approach, where we don't need to
// remember to add individual types here.

// One of the other data types has an entry
