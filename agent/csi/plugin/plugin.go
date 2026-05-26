package plugin

import (
	"context"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/log"
	"github.com/moby/swarmkit/v2/node/plugin"
)

// SecretGetter is a reimplementation of the exec.SecretGetter interface in the
// scope of the plugin package. This avoids the needing to import exec into the
// plugin package.
type SecretGetter interface {
	Get(secretID string) (*api.Secret, error)
}

type NodePlugin interface {
	GetPublishedPath(volumeID string) string
	NodeGetInfo(ctx context.Context) (*api.NodeCSIInfo, error)
	NodeStageVolume(ctx context.Context, req *api.VolumeAssignment) error
	NodeUnstageVolume(ctx context.Context, req *api.VolumeAssignment) error
	NodePublishVolume(ctx context.Context, req *api.VolumeAssignment) error
	NodeUnpublishVolume(ctx context.Context, req *api.VolumeAssignment) error
}

type volumePublishStatus struct {
	// stagingPath is staging path of volume
	stagingPath string

	// isPublished keeps track if the volume is published.
	isPublished bool

	// publishedPath is published path of volume
	publishedPath string
}

type nodePlugin struct {
	// name is the name of the plugin, which is used in the Driver.Name field.
	name string

	// socket is the path of the unix socket to connect to this plugin at
	socket string

	// scopePath gets the provided path relative to the plugin directory.
	scopePath func(s string) string

	// secrets is the SecretGetter to get volume secret data
	secrets SecretGetter

	// volumeMap is the map from volume ID to Volume. Will place a volume once it is staged,
	// remove it from the map for unstage.
	// TODO: Make this map persistent if the swarm node goes down
	volumeMap map[string]*volumePublishStatus

	// mu for volumeMap
	mu sync.RWMutex

	// staging indicates that the plugin has staging capabilities.
	staging bool

	// cc is the gRPC client connection
	cc *grpc.ClientConn

	// idClient is the CSI Identity Service client
	idClient csi.IdentityClient

	// nodeClient is the CSI Node Service client
	nodeClient csi.NodeClient
}

const (
	// TargetStagePath is the path within the plugin's scope that the volume is
	// to be staged. This does not need to be accessible or propagated outside
	// of the plugin rootfs.
	TargetStagePath string = "/data/staged"
	// TargetPublishPath is the path within the plugin's scope that the volume
	// is to be published. This needs to be the plugin's PropagatedMount.
	TargetPublishPath string = "/data/published"
)

func NewNodePlugin(name string, p plugin.AddrPlugin, secrets SecretGetter) NodePlugin {
	_ = "STUB: not implemented"
	return *new(NodePlugin)
}

// newNodePlugin returns a raw nodePlugin object, not behind an interface. this
// is useful for testing.
func newNodePlugin(name string, p plugin.AddrPlugin, secrets SecretGetter) *nodePlugin {
	_ = "STUB: not implemented"
	return nil
}

// connect is a private method that sets up the identity client and node
// client from a grpc client. it exists separately so that testing code can
// substitute in fake clients without a grpc connection
func (np *nodePlugin) connect(ctx context.Context) error {
	_ = "STUB: not implemented"
	// even though this is a unix socket, we must set WithInsecure or the
	// connection will not be allowed.
	return nil
}

// first, probe the plugin, to ensure that it exists and is ready to go

func (np *nodePlugin) Client(ctx context.Context) (csi.NodeClient, error) {
	_ = "STUB: not implemented"
	return *new(csi.NodeClient), nil
}

func (np *nodePlugin) init(ctx context.Context) error {
	probe, err := np.idClient.Probe(ctx, &csi.ProbeRequest{})
	if err != nil {
		return err
	}
	if probe.Ready != nil && !probe.Ready.Value {
		return status.Error(codes.FailedPrecondition, "Plugin is not Ready")
	}

	c, err := np.Client(ctx)
	if err != nil {
		return err
	}

	resp, err := c.NodeGetCapabilities(ctx, &csi.NodeGetCapabilitiesRequest{})
	if err != nil {
		// TODO(ameyag): handle
		return err
	}
	if resp == nil {
		return nil
	}
	log.G(ctx).Debugf("plugin advertises %d capabilities", len(resp.Capabilities))
	for _, c := range resp.Capabilities {
		if rpc := c.GetRpc(); rpc != nil {
			log.G(ctx).Debugf("plugin has capability %s", rpc)
			switch rpc.Type {
			case csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME:
				np.staging = true
			}
		}
	}

	return nil
}

// GetPublishedPath returns the path at which the provided volume ID is
// published. This path is provided in terms of absolute location on the host,
// not the location in the plugins' scope.
//
// Returns an empty string if the volume does not exist.
func (np *nodePlugin) GetPublishedPath(volumeID string) string {
	_ = "STUB: not implemented"
	return ""
}

func (np *nodePlugin) NodeGetInfo(ctx context.Context) (*api.NodeCSIInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (np *nodePlugin) NodeStageVolume(ctx context.Context, req *api.VolumeAssignment) error {
	_ = "STUB: not implemented"
	return nil
}

func (np *nodePlugin) NodeUnstageVolume(ctx context.Context, req *api.VolumeAssignment) error {
	_ = "STUB: not implemented"
	return nil
}

// Check arguments

// we must unpublish before we unstage. verify here that the volume is not
// published.

// if the volume doesn't exist in the volumeMap, deleting has no effect.

func (np *nodePlugin) NodePublishVolume(ctx context.Context, req *api.VolumeAssignment) error {
	_ = "STUB: not implemented"
	return nil
}

// Some volumes plugins require staging; we track this with a boolean, which
// also implies a staging path in the path map. If the plugin is marked as
// requiring staging but does not have a staging path in the map, that is an
// error.

func (np *nodePlugin) NodeUnpublishVolume(ctx context.Context, req *api.VolumeAssignment) error {
	_ = "STUB: not implemented"
	// Check arguments
	return nil
}

func (np *nodePlugin) makeSecrets(v *api.VolumeAssignment) map[string]string {
	_ = "STUB: not implemented"
	// this should never happen, but program defensively.
	return nil
}

// TODO(dperny): handle error from Get

// makeNodeInfo converts a csi.NodeGetInfoResponse object into a swarmkit NodeCSIInfo
// object.
func makeNodeInfo(csiNodeInfo *csi.NodeGetInfoResponse) *api.NodeCSIInfo {
	_ = "STUB: not implemented"
	return nil
}

// stagePath returns the staging path for a given volume assignment
func stagePath(v *api.VolumeAssignment) string {
	_ = "STUB: not implemented"
	// this really just exists so we use the same trick to determine staging
	// path across multiple methods and can't forget to change it in one place
	// but not another
	return ""
}

// publishPath returns the publishing path for a given volume assignment
func publishPath(v *api.VolumeAssignment) string {
	_ = "STUB: not implemented"
	// ditto as stagePath
	return ""
}
