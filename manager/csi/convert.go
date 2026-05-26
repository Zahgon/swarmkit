package csi

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/moby/swarmkit/v2/api"
)

// convert.go contains functions for converting swarm objects into CSI requests
// and back again.

// makeTopology converts a swarmkit topology into a CSI topology.
func makeTopologyRequirement(t *api.TopologyRequirement) *csi.TopologyRequirement {
	_ = "STUB: not implemented"
	return nil
}

// makeTopologies converts a slice of swarmkit topologies into a slice of CSI
// topologies.
func makeTopologies(ts []*api.Topology) []*csi.Topology { _ = "STUB: not implemented"; return nil }

// makeTopology converts a swarmkit topology into a CSI topology. These types
// are essentially homologous, with the swarm type being copied verbatim from
// the CSI type (for build reasons).
func makeTopology(t *api.Topology) *csi.Topology { _ = "STUB: not implemented"; return nil }

// makeCapcityRange converts the swarmkit CapacityRange object to the
// equivalent CSI object
func makeCapacityRange(cr *api.CapacityRange) *csi.CapacityRange {
	_ = "STUB: not implemented"
	return nil
}

// unmakeTopologies transforms a CSI-type topology into the equivalent swarm
// type. it is called "unmakeTopologies" because it performs the inverse of
// "makeTopologies".
func unmakeTopologies(topologies []*csi.Topology) []*api.Topology {
	_ = "STUB: not implemented"
	return nil
}

// unmakeTopology transforms a CSI-type topology into the equivalent swarm
// type.
func unmakeTopology(topology *csi.Topology) *api.Topology { _ = "STUB: not implemented"; return nil }

// makeVolumeInfo converts a csi.Volume object into a swarmkit VolumeInfo
// object.
func makeVolumeInfo(csiVolume *csi.Volume) *api.VolumeInfo { _ = "STUB: not implemented"; return nil }
