package membership

import (
	"errors"
	"sync"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/watch"
	"go.etcd.io/raft/v3/raftpb"
)

var (
	// ErrIDExists is thrown when a node wants to join the existing cluster but its ID already exists
	ErrIDExists = errors.New("membership: can't add node to cluster, node id is a duplicate")
	// ErrIDRemoved is thrown when a node tries to perform an operation on an existing cluster but was removed
	ErrIDRemoved = errors.New("membership: node was removed during cluster lifetime")
	// ErrIDNotFound is thrown when we try an operation on a member that does not exist in the cluster list
	ErrIDNotFound = errors.New("membership: member not found in cluster list")
	// ErrConfigChangeInvalid is thrown when a configuration change we received looks invalid in form
	ErrConfigChangeInvalid = errors.New("membership: ConfChange type should be either AddNode, RemoveNode or UpdateNode")
	// ErrCannotUnmarshalConfig is thrown when a node cannot unmarshal a configuration change
	ErrCannotUnmarshalConfig = errors.New("membership: cannot unmarshal configuration change")
	// ErrMemberRemoved is thrown when a node was removed from the cluster
	ErrMemberRemoved = errors.New("raft: member was removed from the cluster")
)

// Cluster represents a set of active
// raft Members
type Cluster struct {
	mu      sync.RWMutex
	members map[uint64]*Member

	// removed contains the list of removed Members,
	// those ids cannot be reused
	removed map[uint64]bool

	PeersBroadcast *watch.Queue
}

// Member represents a raft Cluster Member
type Member struct {
	*api.RaftMember
}

// NewCluster creates a new Cluster neighbors list for a raft Member.
func NewCluster() *Cluster {
	_ = "STUB: not implemented"
	// TODO(abronan): generate Cluster ID for federation
	return nil
}

// Members returns the list of raft Members in the Cluster.
func (c *Cluster) Members() map[uint64]*Member { _ = "STUB: not implemented"; return nil }

// Removed returns the list of raft Members removed from the Cluster.
func (c *Cluster) Removed() []uint64 { _ = "STUB: not implemented"; return nil }

// GetMember returns informations on a given Member.
func (c *Cluster) GetMember(id uint64) *Member { _ = "STUB: not implemented"; return nil }

func (c *Cluster) broadcastUpdate() { _ = "STUB: not implemented"; return }

// AddMember adds a node to the Cluster Memberlist.
func (c *Cluster) AddMember(member *Member) error { _ = "STUB: not implemented"; return nil }

// RemoveMember removes a node from the Cluster Memberlist, and adds it to
// the removed list.
func (c *Cluster) RemoveMember(id uint64) error { _ = "STUB: not implemented"; return nil }

// UpdateMember updates member address.
func (c *Cluster) UpdateMember(id uint64, m *api.RaftMember) error {
	_ = "STUB: not implemented"
	return nil
}

// Should never happen; this is a sanity check

// nothing to do

// ClearMember removes a node from the Cluster Memberlist, but does NOT add it
// to the removed list.
func (c *Cluster) ClearMember(id uint64) error { _ = "STUB: not implemented"; return nil }

func (c *Cluster) clearMember(id uint64) error { _ = "STUB: not implemented"; return nil }

// IsIDRemoved checks if a Member is in the remove set.
func (c *Cluster) IsIDRemoved(id uint64) bool { _ = "STUB: not implemented"; return false }

// Clear resets the list of active Members and removed Members.
func (c *Cluster) Clear() { _ = "STUB: not implemented"; return }

// ValidateConfigurationChange takes a proposed ConfChange and
// ensures that it is valid.
func (c *Cluster) ValidateConfigurationChange(cc raftpb.ConfChange) error {
	_ = "STUB: not implemented"
	return nil
}
