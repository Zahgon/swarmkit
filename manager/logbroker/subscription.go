package logbroker

import (
	"context"
	"sync"

	events "github.com/docker/go-events"
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/state/store"
	"github.com/moby/swarmkit/v2/watch"
)

type subscription struct {
	mu sync.RWMutex
	wg sync.WaitGroup

	store   *store.MemoryStore
	message *api.SubscriptionMessage
	changed *watch.Queue

	ctx    context.Context
	cancel context.CancelFunc

	errors       []error
	nodes        map[string]struct{}
	pendingTasks map[string]struct{}
}

func newSubscription(store *store.MemoryStore, message *api.SubscriptionMessage, changed *watch.Queue) *subscription {
	_ = "STUB: not implemented"
	return nil
}

func (s *subscription) follow() bool { _ = "STUB: not implemented"; return false }

func (s *subscription) Contains(nodeID string) bool { _ = "STUB: not implemented"; return false }

func (s *subscription) Nodes() []string { _ = "STUB: not implemented"; return nil }

func (s *subscription) Run(ctx context.Context) { _ = "STUB: not implemented"; return }

func (s *subscription) Stop() { _ = "STUB: not implemented"; return }

func (s *subscription) Wait(_ context.Context) <-chan struct{} {
	_ = "STUB: not implemented"
	// Follow subscriptions never end
	return nil
}

func (s *subscription) Done(nodeID string, err error) { _ = "STUB: not implemented"; return }

func (s *subscription) Err() error { _ = "STUB: not implemented"; return nil }

func (s *subscription) Close() { _ = "STUB: not implemented"; return }

func (s *subscription) Closed() bool { _ = "STUB: not implemented"; return false }

func (s *subscription) match() { _ = "STUB: not implemented"; return }

// if we're not following, don't add tasks that aren't running yet

func (s *subscription) watch(ch <-chan events.Event) error { _ = "STUB: not implemented"; return nil }

// this mutex does not have a deferred unlock, because there is work
// we need to do after we release it.

// Un-allocated task.

// if we try to call Publish before we release the lock, we can end
// up in a situation where the receiver is trying to acquire a read
// lock on it. it's hard to explain.
