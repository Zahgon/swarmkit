package dispatcher

import (
	"sync"
	"time"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/dispatcher/heartbeat"
)

const rateLimitCount = 3

type registeredNode struct {
	SessionID  string
	Heartbeat  *heartbeat.Heartbeat
	Registered time.Time
	Attempts   int
	Node       *api.Node
	Disconnect chan struct{} // signal to disconnect
	mu         sync.Mutex
}

// checkSessionID determines if the SessionID has changed and returns the
// appropriate GRPC error code.
//
// This may not belong here in the future.
func (rn *registeredNode) checkSessionID(sessionID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Before each message send, we need to check the nodes sessionID hasn't
// changed. If it has, we will the stream and make the node
// re-register.

type nodeStore struct {
	periodChooser                *periodChooser
	gracePeriodMultiplierNormal  time.Duration
	gracePeriodMultiplierUnknown time.Duration
	rateLimitPeriod              time.Duration
	nodes                        map[string]*registeredNode
	mu                           sync.RWMutex
}

func newNodeStore(hbPeriod, hbEpsilon time.Duration, graceMultiplier int, rateLimitPeriod time.Duration) *nodeStore {
	_ = "STUB: not implemented"
	return nil
}

func (s *nodeStore) updatePeriod(hbPeriod, hbEpsilon time.Duration, gracePeriodMultiplier int) {
	_ = "STUB: not implemented"
	return
}

func (s *nodeStore) Len() int { _ = "STUB: not implemented"; return 0 }

func (s *nodeStore) AddUnknown(n *api.Node, expireFunc func()) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckRateLimit returns error if node with specified id is allowed to re-register
// again.
func (s *nodeStore) CheckRateLimit(id string) error { _ = "STUB: not implemented"; return nil }

// Add adds new node and returns it, it replaces existing without notification.
func (s *nodeStore) Add(n *api.Node, expireFunc func()) *registeredNode {
	_ = "STUB: not implemented"
	return nil
}

// session ID is local to the dispatcher.

func (s *nodeStore) Get(id string) (*registeredNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *nodeStore) GetWithSession(id, sid string) (*registeredNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *nodeStore) Heartbeat(id, sid string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// base period for node

func (s *nodeStore) Delete(id string) *registeredNode { _ = "STUB: not implemented"; return nil }

func (s *nodeStore) Disconnect(id string) { _ = "STUB: not implemented"; return }

// Clean removes all nodes and stops their heartbeats.
// It's equivalent to invalidate all sessions.
func (s *nodeStore) Clean() { _ = "STUB: not implemented"; return }
