package heartbeat

import (
	"time"
)

// Heartbeat is simple way to track heartbeats.
type Heartbeat struct {
	timeout int64
	timer   *time.Timer
}

// New creates new Heartbeat with specified duration. timeoutFunc will be called
// if timeout for heartbeat is expired. Note that in case of timeout you need to
// call Beat() to reactivate Heartbeat.
func New(timeout time.Duration, timeoutFunc func()) *Heartbeat {
	_ = "STUB: not implemented"
	return nil
}

// Beat resets internal timer to zero. It also can be used to reactivate
// Heartbeat after timeout.
func (hb *Heartbeat) Beat() { _ = "STUB: not implemented"; return }

// Update updates internal timeout to d. It does not do Beat.
func (hb *Heartbeat) Update(d time.Duration) { _ = "STUB: not implemented"; return }

// Stop stops Heartbeat timer.
func (hb *Heartbeat) Stop() { _ = "STUB: not implemented"; return }
