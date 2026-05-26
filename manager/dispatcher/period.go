package dispatcher

import (
	"math/rand"
	"time"
)

type periodChooser struct {
	period  time.Duration
	epsilon time.Duration
	rand    *rand.Rand
}

func newPeriodChooser(period, eps time.Duration) *periodChooser {
	_ = "STUB: not implemented"
	return nil
}

func (pc *periodChooser) Choose() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
