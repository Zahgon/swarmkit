package capability

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/moby/swarmkit/v2/api"
)

func CheckArguments(req *api.VolumeAssignment) error { _ = "STUB: not implemented"; return nil }

func MakeCapability(am *api.VolumeAccessMode) *csi.VolumeCapability {
	_ = "STUB: not implemented"
	return nil
}

// Block type is empty.
