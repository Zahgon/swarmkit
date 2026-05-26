package logbroker

import (
	"context"
	"errors"
	"sync"

	"github.com/docker/go-events"
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/state/store"
	"github.com/moby/swarmkit/v2/watch"
)

var (
	errAlreadyRunning = errors.New("broker is already running")
	errNotRunning     = errors.New("broker is not running")
)

type logMessage struct {
	*api.PublishLogsMessage
	completed bool
	err       error
}

// LogBroker coordinates log subscriptions to services and tasks. Clients can
// publish and subscribe to logs channels.
//
// Log subscriptions are pushed to the work nodes by creating log subscription
// tasks. As such, the LogBroker also acts as an orchestrator of these tasks.
type LogBroker struct {
	mu                sync.RWMutex
	logQueue          *watch.Queue
	subscriptionQueue *watch.Queue

	registeredSubscriptions map[string]*subscription
	subscriptionsByNode     map[string]map[*subscription]struct{}

	pctx      context.Context
	cancelAll context.CancelFunc

	store *store.MemoryStore
}

// New initializes and returns a new LogBroker
func New(store *store.MemoryStore) *LogBroker { _ = "STUB: not implemented"; return nil }

// Start starts the log broker
func (lb *LogBroker) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Stop stops the log broker
func (lb *LogBroker) Stop() error { _ = "STUB: not implemented"; return nil }

func validateSelector(selector *api.LogSelector) error { _ = "STUB: not implemented"; return nil }

func (lb *LogBroker) newSubscription(selector *api.LogSelector, options *api.LogSubscriptionOptions) *subscription {
	_ = "STUB: not implemented"
	return nil
}

func (lb *LogBroker) getSubscription(id string) *subscription {
	_ = "STUB: not implemented"
	return nil
}

func (lb *LogBroker) registerSubscription(subscription *subscription) {
	_ = "STUB: not implemented"
	return
}

// Mark nodes that won't receive the message as done.

// otherwise, add the subscription to the node's subscriptions list

func (lb *LogBroker) unregisterSubscription(subscription *subscription) {
	_ = "STUB: not implemented"
	return
}

// remove the subscription from all of the nodes

// but only if a node exists

// watchSubscriptions grabs all current subscriptions and notifies of any
// subscription change for this node.
//
// Subscriptions may fire multiple times and the caller has to protect against
// dupes.
func (lb *LogBroker) watchSubscriptions(nodeID string) ([]*subscription, chan events.Event, func()) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Watch for subscription changes for this node.

// Grab current subscriptions.

func (lb *LogBroker) subscribe(id string) (chan events.Event, func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lb *LogBroker) publish(log *api.PublishLogsMessage) { _ = "STUB: not implemented"; return }

// markDone wraps (*Subscription).Done() so that the removal of the sub from
// the node's subscription list is possible
func (lb *LogBroker) markDone(sub *subscription, nodeID string, err error) {
	_ = "STUB: not implemented"
	return
}

// remove the subscription from the node's subscription list, if it exists

// mark the sub as done

// SubscribeLogs creates a log subscription and streams back logs
func (lb *LogBroker) SubscribeLogs(request *api.SubscribeLogsRequest, stream api.Logs_SubscribeLogsServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (lb *LogBroker) nodeConnected(nodeID string) { _ = "STUB: not implemented"; return }

func (lb *LogBroker) nodeDisconnected(nodeID string) { _ = "STUB: not implemented"; return }

// ListenSubscriptions returns a stream of matching subscriptions for the current node
func (lb *LogBroker) ListenSubscriptions(_ *api.ListenSubscriptionsRequest, stream api.LogBroker_ListenSubscriptionsServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Start by sending down all active subscriptions.

// Send down new subscriptions.

// Avoid sending down the same subscription multiple times

// PublishLogs publishes log messages for a given subscription
func (lb *LogBroker) PublishLogs(stream api.LogBroker_PublishLogsServer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// if we have a close message, close out the subscription

// Mark done and then set to nil so if we error after this point,
// we don't try to close again in the defer

// Make sure logs are emitted using the right Node ID to avoid impersonation.
