package main

import (
	"context"
	"io"
	"net"
	"time"

	"github.com/rcrowley/go-metrics"
)

// Collector waits for tasks to phone home while collecting statistics.
type Collector struct {
	t  metrics.Timer
	ln net.Listener
}

// Listen starts listening on a TCP port. Tasks have to connect to this address
// once they come online.
func (c *Collector) Listen(port int) error { _ = "STUB: not implemented"; return nil }

// Collect blocks until `count` tasks phoned home.
func (c *Collector) Collect(ctx context.Context, count uint64) { _ = "STUB: not implemented"; return }

// Stats prints various statistics related to the collection.
func (c *Collector) Stats(w io.Writer, unit time.Duration) { _ = "STUB: not implemented"; return }

// NewCollector creates and returns a collector.
func NewCollector() *Collector { _ = "STUB: not implemented"; return nil }
