package controlapi

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
)

func (s *Server) CreateVolume(_ context.Context, request *api.CreateVolumeRequest) (*api.CreateVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validate the volume spec

// check all secrets, so that we can return an error indicating ALL
// missing secrets, instead of just the first one.

func (s *Server) UpdateVolume(_ context.Context, request *api.UpdateVolumeRequest) (*api.UpdateVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// compare specs, to see if any invalid fields have changed

// to further guard against changing fields we're not allowed to, don't
// replace the entire spec. just replace the fields we are allowed to
// change

// read the volume back out, so it has the correct meta version
// TODO(dperny): this behavior, while likely more correct, may not be
// consistent with the rest of swarmkit...

func (s *Server) ListVolumes(_ context.Context, request *api.ListVolumesRequest) (*api.ListVolumesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// so the way we do this is with two filtering passes. first, we do a store
// request, filtering on one of the parameters. then, from the result of
// the store request, we filter on the remaining filters. This is necessary
// because the store filters do not expose an AND function.

// short circuit to avoid nil pointer deref

// Names

// NamePrefixes

// IDPrefixes

// Labels

// Groups

// Drivers

func filterVolumes(candidates []*api.Volume, filters ...func(*api.Volume) bool) []*api.Volume {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) GetVolume(_ context.Context, request *api.GetVolumeRequest) (*api.GetVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoveVolume marks a Volume for removal. For a Volume to be removed, it must
// have Availability set to Drain. RemoveVolume does not immediately delete the
// volume, because some clean-up must occur before it can be removed. However,
// calling RemoveVolume is an irrevocable action, and once it occurs, the
// Volume can no longer be used in any way.
func (s *Server) RemoveVolume(_ context.Context, request *api.RemoveVolumeRequest) (*api.RemoveVolumeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If this is a force delete, we force the delete. No survivors. This
// is a last resort to resolve otherwise intractable problems with
// volumes. Using this has the potential to break other things in the
// cluster, because testing every case where we force-remove a volume
// is difficult at best.
