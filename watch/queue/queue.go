package queue

import (
	"container/list"
	"fmt"
	"sync"

	"github.com/docker/go-events"
)

// ErrQueueFull is returned by a Write operation when that Write causes the
// queue to reach its size limit.
var ErrQueueFull = fmt.Errorf("queue closed due to size limit")

// LimitQueue accepts all messages into a queue for asynchronous consumption by
// a sink until an upper limit of messages is reached. When that limit is
// reached, the entire Queue is Closed. It is thread safe but the
// sink must be reliable or events will be dropped.
// If a size of 0 is provided, the LimitQueue is considered limitless.
type LimitQueue struct {
	dst        events.Sink
	events     *list.List
	limit      uint64
	cond       *sync.Cond
	mu         sync.Mutex
	closed     bool
	full       chan struct{}
	fullClosed bool
}

// NewLimitQueue returns a queue to the provided Sink dst.
func NewLimitQueue(dst events.Sink, limit uint64) *LimitQueue {
	_ = "STUB: not implemented"
	return nil
}

// Write accepts the events into the queue, only failing if the queue has
// been closed or has reached its size limit.
func (eq *LimitQueue) Write(event events.Event) error { _ = "STUB: not implemented"; return nil }

// If the limit has been reached, don't write the event to the queue,
// and close the Full channel. This notifies listeners that the queue
// is now full, but the sink is still permitted to consume events. It's
// the responsibility of the listener to decide whether they want to
// live with dropped events or whether they want to Close() the
// LimitQueue

// signal waiters

// Full returns a channel that is closed when the queue becomes full for the
// first time.
func (eq *LimitQueue) Full() chan struct{} {
	_ = "STUB: not implemented"

	// Close shuts down the event queue, flushing all events
	return nil
}

func (eq *LimitQueue) Close() error { _ = "STUB: not implemented"; return nil }

// set the closed flag

// signal flushes queue
// wait for signal from last flush

// run is the main goroutine to flush events to the target sink.
func (eq *LimitQueue) run() { _ = "STUB: not implemented"; return }

// nil block means event queue is closed.

// TODO(aaronl): Dropping events could be bad depending
// on the application. We should have a way of
// communicating this condition. However, logging
// at a log level above debug may not be appropriate.
// Eventually, go-events should not use logrus at all,
// and should bubble up conditions like this through
// error values.

// Len returns the number of items that are currently stored in the queue and
// not consumed by its sink.
func (eq *LimitQueue) Len() int { _ = "STUB: not implemented"; return 0 }

func (eq *LimitQueue) String() string { _ = "STUB: not implemented"; return "" }

// next encompasses the critical section of the run loop. When the queue is
// empty, it will block on the condition. If new data arrives, it will wake
// and return a block. When closed, a nil slice will be returned.
func (eq *LimitQueue) next() events.Event { _ = "STUB: not implemented"; return *new(events.Event) }
