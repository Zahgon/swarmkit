package watch

import (
	"fmt"
	"time"

	events "github.com/docker/go-events"
)

// ErrSinkTimeout is returned from the Write method when a sink times out.
var ErrSinkTimeout = fmt.Errorf("timeout exceeded, tearing down sink")

// timeoutSink is a sink that wraps another sink with a timeout. If the
// embedded sink fails to complete a Write operation within the specified
// timeout, the Write operation of the timeoutSink fails.
type timeoutSink struct {
	timeout time.Duration
	sink    events.Sink
}

func (s timeoutSink) Write(event events.Event) error { _ = "STUB: not implemented"; return nil }

func (s timeoutSink) Close() error { _ = "STUB: not implemented"; return nil }

// dropErrClosed is a sink that suppresses ErrSinkClosed from Write, to avoid
// debug log messages that may be confusing. It is possible that the queue
// will try to write an event to its destination channel while the queue is
// being removed from the broadcaster. Since the channel is closed before the
// queue, there is a narrow window when this is possible. In some event-based
// dropping events when a sink is removed from a broadcaster is a problem, but
// for the usage in this watch package that's the expected behavior.
type dropErrClosed struct {
	sink events.Sink
}

func (s dropErrClosed) Write(event events.Event) error { _ = "STUB: not implemented"; return nil }

func (s dropErrClosed) Close() error { _ = "STUB: not implemented"; return nil }

// dropErrClosedChanGen is a ChannelSinkGenerator for dropErrClosed sinks wrapping
// unbuffered channels.
type dropErrClosedChanGen struct{}

func (s *dropErrClosedChanGen) NewChannelSink() (events.Sink, *events.Channel) {
	_ = "STUB: not implemented"
	return *new(events.Sink), nil
}

// TimeoutDropErrChanGen is a ChannelSinkGenerator that creates a channel,
// wrapped by the dropErrClosed sink and a timeout.
type TimeoutDropErrChanGen struct {
	timeout time.Duration
}

// NewChannelSink creates a new sink chain of timeoutSink->dropErrClosed->Channel
func (s *TimeoutDropErrChanGen) NewChannelSink() (events.Sink, *events.Channel) {
	_ = "STUB: not implemented"
	return *new(events.Sink), nil
}

// NewTimeoutDropErrSinkGen returns a generator of timeoutSinks wrapping dropErrClosed
// sinks, wrapping unbuffered channel sinks.
func NewTimeoutDropErrSinkGen(timeout time.Duration) ChannelSinkGenerator {
	_ = "STUB: not implemented"
	return *new(ChannelSinkGenerator)
}
