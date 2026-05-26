package agent

import (
	"context"
	"sync"

	"github.com/moby/swarmkit/v2/api"
)

// StatusReporter receives updates to task status. Method may be called
// concurrently, so implementations should be goroutine-safe.
type StatusReporter interface {
	UpdateTaskStatus(ctx context.Context, taskID string, status *api.TaskStatus) error
}

// Reporter receives update to both task and volume status.
type Reporter interface {
	StatusReporter
	ReportVolumeUnpublished(ctx context.Context, volumeID string) error
}

type statusReporterFunc func(ctx context.Context, taskID string, status *api.TaskStatus) error

func (fn statusReporterFunc) UpdateTaskStatus(ctx context.Context, taskID string, status *api.TaskStatus) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:unused // currently only used in tests.
type volumeReporterFunc func(ctx context.Context, volumeID string) error

//nolint:unused // currently only used in tests.
func (fn volumeReporterFunc) ReportVolumeUnpublished(ctx context.Context, volumeID string) error {
	_ = "STUB: not implemented"
	return nil

	//nolint:unused // currently only used in tests.
}

type statusReporterCombined struct {
	statusReporterFunc
	volumeReporterFunc
}

// statusReporter creates a reliable StatusReporter that will always succeed.
// It handles several tasks at once, ensuring all statuses are reported.
//
// The reporter will continue reporting the current status until it succeeds.
type statusReporter struct {
	reporter Reporter
	statuses map[string]*api.TaskStatus
	// volumes is a set of volumes which are to be reported unpublished.
	volumes map[string]struct{}
	mu      sync.Mutex
	cond    sync.Cond
	closed  bool
}

func newStatusReporter(ctx context.Context, upstream Reporter) *statusReporter {
	_ = "STUB: not implemented"
	return nil
}

func (sr *statusReporter) UpdateTaskStatus(_ context.Context, taskID string, status *api.TaskStatus) error {
	_ = "STUB: not implemented"
	return nil
}

// ignore old updates

func (sr *statusReporter) ReportVolumeUnpublished(_ context.Context, volumeID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (sr *statusReporter) Close() error { _ = "STUB: not implemented"; return nil }

func (sr *statusReporter) run(ctx context.Context) { _ = "STUB: not implemented"; return }

// released during wait, below.

// TODO(stevvooe): Add support here for waiting until all
// statuses are flushed before shutting down.

// delete the entry, while trying to send.

// reporter might be closed during UpdateTaskStatus call

// place it back in the map, if not there, allowing us to pick
// the value if a new one came in when we were sending the last
// update.

// reporter might be closed during ReportVolumeUnpublished call
