package scheduler

import (
	"github.com/moby/swarmkit/v2/api"
)

var (
	defaultFilters = []Filter{
		// Always check for readiness first.
		&ReadyFilter{},
		&ResourceFilter{},
		&PluginFilter{},
		&ConstraintFilter{},
		&PlatformFilter{},
		&HostPortFilter{},
		&MaxReplicasFilter{},
	}
)

type checklistEntry struct {
	f       Filter
	enabled bool

	// failureCount counts the number of nodes that this filter failed
	// against.
	failureCount int
}

type checklistByFailures []checklistEntry

func (c checklistByFailures) Len() int           { _ = "STUB: not implemented"; return 0 }
func (c checklistByFailures) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (c checklistByFailures) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Pipeline runs a set of filters against nodes.
type Pipeline struct {
	// checklist is a slice of filters to run
	checklist []checklistEntry
}

// NewPipeline returns a pipeline with the default set of filters.
func NewPipeline() *Pipeline { _ = "STUB: not implemented"; return nil }

// Process a node through the filter pipeline.
// Returns true if all filters pass, false otherwise.
func (p *Pipeline) Process(n *NodeInfo) bool { _ = "STUB: not implemented"; return false }

// Immediately stop on first failure.

func (p *Pipeline) AddFilter(f Filter) { _ = "STUB: not implemented"; return }

// SetTask sets up the filters to process a new task. Once this is called,
// Process can be called repeatedly to try to assign the task various nodes.
func (p *Pipeline) SetTask(t *api.Task) { _ = "STUB: not implemented"; return }

// Explain returns a string explaining why a task could not be scheduled.
func (p *Pipeline) Explain() string { _ = "STUB: not implemented"; return "" }

// Sort from most failures to least
