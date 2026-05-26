package csi

import (
	"context"
	"sync"
	"time"

	"github.com/moby/swarmkit/v2/agent/csi/plugin"
	"github.com/moby/swarmkit/v2/agent/exec"
	"github.com/moby/swarmkit/v2/api"
	mobyplugin "github.com/moby/swarmkit/v2/node/plugin"
	"github.com/moby/swarmkit/v2/volumequeue"
)

const csiCallTimeout = 15 * time.Second

// volumeState keeps track of the state of a volume on this node.
type volumeState struct {
	// volume is the actual VolumeAssignment for this volume
	volume *api.VolumeAssignment
	// remove is true if the volume is to be removed, or false if it should be
	// active.
	remove bool
	// removeCallback is called when the volume is successfully removed.
	removeCallback func(id string)
}

// volumes is a map that keeps all the currently available volumes to the agent
// mapped by volume ID.
type volumes struct {
	// mu guards access to the volumes map.
	mu sync.RWMutex

	// volumes is a mapping of volume ID to volumeState
	volumes map[string]volumeState

	// plugins is the Manager, which provides translation to the CSI RPCs
	plugins plugin.Manager

	// pendingVolumes is a VolumeQueue which manages which volumes are
	// processed and when.
	pendingVolumes *volumequeue.VolumeQueue
}

// NewManager returns a place to store volumes.
func NewManager(pg mobyplugin.Getter, secrets exec.SecretGetter) exec.VolumesManager {
	_ = "STUB: not implemented"
	return *new(exec.VolumesManager)
}

// retryVolumes runs in a goroutine to retry failing volumes.
func (r *volumes) retryVolumes() { _ = "STUB: not implemented"; return }

// this case occurs when the Stop method has been called on
// pendingVolumes, and means that we should pack up and exit.

// tryVolume synchronously tries one volume. it puts the volume back into the
// queue if the attempt fails.
func (r *volumes) tryVolume(ctx context.Context, id string, attempt uint) {
	_ = "STUB: not implemented"
	return
}

// create a sub-context with a timeout. because we can only process one
// volume at a time, if we rely on the server-side or default timeout, we
// may be waiting a very long time for a particular volume to fail.
//
// TODO(dperny): there is almost certainly a more intelligent way to do
// this. For example, we could:
//
//   * Change code such that we can service volumes managed by different
//     plugins at the same time.
//   * Take longer timeouts when we don't have any other volumes in the
//     queue
//   * Have interruptible attempts, so that if we're taking longer
//     timeouts, we can abort them to service new volumes.
//
// These are too complicated to be worth the engineering effort at this
// time.

// always gotta call the WithTimeout cancel

// if unpublishing was successful, then call the callback

// Get returns a volume published path for the provided volume ID.  If the volume doesn't exist, returns empty string.
func (r *volumes) Get(volumeID string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// TODO(dperny): use a structured error

// don't put this line here, it spams like crazy.
// log.L.WithField("method", "(*volumes).Get").Debugf("Path not published for volume:%v", volumeID)

// Add adds one or more volumes to the volume map.
func (r *volumes) Add(volumes ...api.VolumeAssignment) { _ = "STUB: not implemented"; return }

// if we get an Add operation, then we will always restart the retries.

// enqueue the volume so that we process it

// Remove removes one or more volumes from this manager. callback is called
// whenever the removal is successful.
func (r *volumes) Remove(volumes []api.VolumeAssignment, callback func(id string)) {
	_ = "STUB: not implemented"
	return
}

// if we get a Remove call, then we always restart the retries and
// attempt removal.

func (r *volumes) publishVolume(ctx context.Context, assignment *api.VolumeAssignment) error {
	_ = "STUB: not implemented"
	return nil
}

// even though this may have succeeded already, the call to NodeStageVolume
// is idempotent, so we can retry it every time.

func (r *volumes) unpublishVolume(ctx context.Context, assignment *api.VolumeAssignment) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *volumes) Plugins() exec.VolumePluginManager {
	_ = "STUB: not implemented"

	// taskRestrictedVolumesProvider restricts the ids to the task.
	return *new(exec.VolumePluginManager)
}

type taskRestrictedVolumesProvider struct {
	volumes   exec.VolumeGetter
	volumeIDs map[string]struct{}
}

func (sp *taskRestrictedVolumesProvider) Get(volumeID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Restrict provides a getter that only allows access to the volumes
// referenced by the task.
func Restrict(volumes exec.VolumeGetter, t *api.Task) exec.VolumeGetter {
	_ = "STUB: not implemented"
	return *new(exec.VolumeGetter)
}
