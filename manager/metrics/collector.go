package metrics

import (
	"context"

	"github.com/docker/go-events"
	metrics "github.com/docker/go-metrics"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

var (
	ns = metrics.NewNamespace("swarm", "manager", nil)

	// counts of the various objects in swarmkit
	nodesMetric metrics.LabeledGauge
	tasksMetric metrics.LabeledGauge

	// none of these objects have state, so they're just regular gauges
	servicesMetric metrics.Gauge
	networksMetric metrics.Gauge
	secretsMetric  metrics.Gauge
	configsMetric  metrics.Gauge
)

func init() {
	nodesMetric = ns.NewLabeledGauge("nodes", "The number of nodes", "", "state")
	tasksMetric = ns.NewLabeledGauge("tasks", "The number of tasks in the cluster object store", metrics.Total, "state")
	servicesMetric = ns.NewGauge("services", "The number of services in the cluster object store", metrics.Total)
	networksMetric = ns.NewGauge("networks", "The number of networks in the cluster object store", metrics.Total)
	secretsMetric = ns.NewGauge("secrets", "The number of secrets in the cluster object store", metrics.Total)
	configsMetric = ns.NewGauge("configs", "The number of configs in the cluster object store", metrics.Total)

	resetMetrics()

	metrics.Register(ns)
}

// Collector collects swarmkit metrics
type Collector struct {
	store *store.MemoryStore

	// stopChan signals to the state machine to stop running.
	stopChan chan struct{}
	// doneChan is closed when the state machine terminates.
	doneChan chan struct{}
}

// NewCollector creates a new metrics collector
func NewCollector(store *store.MemoryStore) *Collector { _ = "STUB: not implemented"; return nil }

// Run contains the collector event loop
func (c *Collector) Run(_ context.Context) error { _ = "STUB: not implemented"; return nil }

// Stop stops the collector.
func (c *Collector) Stop() { _ = "STUB: not implemented"; return }

// Clean the metrics on exit.

// resetMetrics resets all metrics to their default (base) value
func resetMetrics() { _ = "STUB: not implemented"; return }

// handleEvent handles a single incoming cluster event.
func (c *Collector) handleEvent(event events.Event) { _ = "STUB: not implemented"; return }

func (c *Collector) handleNodeEvent(event events.Event) { _ = "STUB: not implemented"; return }

// Skip updates if nothing changed.

func (c *Collector) handleTaskEvent(event events.Event) { _ = "STUB: not implemented"; return }

// Skip updates if nothing changed.

func (c *Collector) handleServiceEvent(event events.Event) { _ = "STUB: not implemented"; return }

func (c *Collector) handleNetworkEvent(event events.Event) { _ = "STUB: not implemented"; return }

func (c *Collector) handleSecretsEvent(event events.Event) { _ = "STUB: not implemented"; return }

func (c *Collector) handleConfigsEvent(event events.Event) { _ = "STUB: not implemented"; return }
