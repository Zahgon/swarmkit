package watch

import (
	"context"
	"sync"
	"time"

	"github.com/docker/go-events"
)

// ChannelSinkGenerator is a constructor of sinks that eventually lead to a
// channel.
type ChannelSinkGenerator interface {
	NewChannelSink() (events.Sink, *events.Channel)
}

// Queue is the structure used to publish events and watch for them.
type Queue struct {
	sinkGen ChannelSinkGenerator
	// limit is the max number of items to be held in memory for a watcher
	limit       uint64
	mu          sync.Mutex
	broadcast   *events.Broadcaster
	cancelFuncs map[events.Sink]func()

	// closeOutChan indicates whether the watchers' channels should be closed
	// when a watcher queue reaches its limit or when the Close method of the
	// sink is called.
	closeOutChan bool
}

// NewQueue creates a new publish/subscribe queue which supports watchers.
// The channels that it will create for subscriptions will have the buffer
// size specified by buffer.
func NewQueue(options ...func(*Queue) error) *Queue {
	_ = "STUB: not implemented"
	// Create a queue with the default values
	return nil
}

// WithTimeout returns a functional option for a queue that sets a write timeout
func WithTimeout(timeout time.Duration) func(*Queue) error { _ = "STUB: not implemented"; return nil }

// WithCloseOutChan returns a functional option for a queue whose watcher
// channel is closed when no more events are expected to be sent to the watcher.
func WithCloseOutChan() func(*Queue) error { _ = "STUB: not implemented"; return nil }

// WithLimit returns a functional option for a queue with a max size limit.
func WithLimit(limit uint64) func(*Queue) error { _ = "STUB: not implemented"; return nil }

// Watch returns a channel which will receive all items published to the
// queue from this point, until cancel is called.
func (q *Queue) Watch() (eventq chan events.Event, cancel func()) {
	_ = "STUB: not implemented"
	return nil,

		// WatchContext returns a channel where all items published to the queue will
		// be received. The channel will be closed when the provided context is
		// cancelled.
		nil
}

func (q *Queue) WatchContext(ctx context.Context) (eventq chan events.Event) {
	_ = "STUB: not implemented"
	return nil
}

// CallbackWatch returns a channel which will receive all events published to
// the queue from this point that pass the check in the provided callback
// function. The returned cancel function will stop the flow of events and
// close the channel.
func (q *Queue) CallbackWatch(matcher events.Matcher) (eventq chan events.Event, cancel func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the output channel shouldn't be closed and the queue is limitless,
// there's no need for an additional goroutine.

// Close the output channel if the ChannelSink is Done for any
// reason. This can happen if the cancelFunc is called
// externally or if it has been closed by a wrapper sink, such
// as the TimeoutSink.

// Close the output channel and tear down the Queue if the
// LimitQueue becomes full.

// CallbackWatchContext returns a channel where all items published to the queue will
// be received. The channel will be closed when the provided context is
// cancelled.
func (q *Queue) CallbackWatchContext(ctx context.Context, matcher events.Matcher) (eventq chan events.Event) {
	_ = "STUB: not implemented"
	return nil
}

// Publish adds an item to the queue.
func (q *Queue) Publish(item events.Event) { _ = "STUB: not implemented"; return }

// Close closes the queue and frees the associated resources.
func (q *Queue) Close() error {
	_ = "STUB: not implemented"
	// Make sure all watchers have been closed to avoid a deadlock when
	// closing the broadcaster.
	return nil
}
