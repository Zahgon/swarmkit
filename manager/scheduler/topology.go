package scheduler

import (
	"github.com/moby/swarmkit/v2/api"
)

// IsInTopology takes a Topology `top` (which is reported by a Node) and a list
// of Topologies `accessible` (which comes from a created volume, in the form
// of the AccessibleTopology) and returns true if `top` lies within
// `accessible` (meaning a node with that Topology can access a volume with
// that AccessibleTopology).
//
// In order for `top` to lie within `accessible`, there must exist a topology
// in `accessible` such that for every subdomain/segment pair in that topology,
// there exists an equivalent subdomain/segment pair in `top`.
//
// For examples, see the test for this function.
//
// NOTE(dperny): It is unclear whether a topology can be partial. For example,
// can an accessible topology contain only a "region" subdomain, without a
// "zone" subdomain? This function assumes yes.
func IsInTopology(top *api.Topology, accessible []*api.Topology) bool {
	_ = "STUB: not implemented"
	// if any part of the topology equation is missing, then this does fit.
	return false
}

// go through each accessible topology

// and for each topology, go through every segment

// if the segment for this subdomain is different in the `top`,
// then, `top` does not lie within this topology.

// go to the next topology in the list

// if we get through all of the segments specified in this topology,
// and they have all matched, then `top` lies within `accessible`.

// if we have iterated through all topologies, and never once finished
// iterating through all topological segments, then `top` does not lie
// within `accessible`.
