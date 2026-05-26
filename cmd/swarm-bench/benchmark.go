package main

import (
	"context"
	"time"

	"github.com/moby/swarmkit/v2/api"
)

// Config holds the benchmarking configuration.
type Config struct {
	Count   uint64
	Manager string
	IP      string
	Port    int
	Unit    time.Duration
}

// Benchmark represents a benchmark session.
type Benchmark struct {
	cfg       *Config
	collector *Collector
}

// NewBenchmark creates a new benchmark session with the given configuration.
func NewBenchmark(cfg *Config) *Benchmark { _ = "STUB: not implemented"; return nil }

// Run starts the benchmark session and waits for it to be completed.
func (b *Benchmark) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Periodically print stats.

func (b *Benchmark) spec() *api.ServiceSpec { _ = "STUB: not implemented"; return nil }

func (b *Benchmark) launch(ctx context.Context) (*api.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
