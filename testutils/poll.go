package testutils

import (
	"time"

	"code.cloudfoundry.org/clock/fakeclock"
)

// PollFuncWithTimeout is used to periodically execute a check function, it
// returns error after timeout.
func PollFuncWithTimeout(clockSource *fakeclock.FakeClock, f func() error, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// PollFunc is like PollFuncWithTimeout with timeout=10s.
func PollFunc(clockSource *fakeclock.FakeClock, f func() error) error {
	_ = "STUB: not implemented"
	return nil
}
