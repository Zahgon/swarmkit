package remotes

import (
	"fmt"
	"sync"

	"github.com/moby/swarmkit/v2/api"
)

var errRemotesUnavailable = fmt.Errorf("no remote hosts provided")

// DefaultObservationWeight provides a weight to use for positive observations
// that will balance well under repeated observations.
const DefaultObservationWeight = 10

// Remotes keeps track of remote addresses by weight, informed by
// observations.
type Remotes interface {
	// Weight returns the remotes with their current weights.
	Weights() map[api.Peer]int

	// Select a remote from the set of available remotes with optionally
	// excluding ID or address.
	Select(...string) (api.Peer, error)

	// Observe records an experience with a particular remote. A positive weight
	// indicates a good experience and a negative weight a bad experience.
	//
	// The observation will be used to calculate a moving weight, which is
	// implementation dependent. This method will be called such that repeated
	// observations of the same master in each session request are favored.
	Observe(peer api.Peer, weight int)

	// ObserveIfExists records an experience with a particular remote if when a
	// remote exists.
	ObserveIfExists(peer api.Peer, weight int)

	// Remove the remote from the list completely.
	Remove(addrs ...api.Peer)
}

// NewRemotes returns a Remotes instance with the provided set of addresses.
// Entries provided are heavily weighted initially.
func NewRemotes(peers ...api.Peer) Remotes { _ = "STUB: not implemented"; return *new(Remotes) }

type remotesWeightedRandom struct {
	remotes map[api.Peer]int
	mu      sync.Mutex

	// workspace to avoid reallocation. these get lazily allocated when
	// selecting values.
	cdf   []float64
	peers []api.Peer
}

func (mwr *remotesWeightedRandom) Weights() map[api.Peer]int { _ = "STUB: not implemented"; return nil }

func (mwr *remotesWeightedRandom) Select(excludes ...string) (api.Peer, error) {
	_ = "STUB: not implemented"
	return *new(api.Peer), nil
}

// NOTE(stevvooe): We then use a weighted random selection algorithm
// (http://stackoverflow.com/questions/4463561/weighted-random-selection-from-array)
// to choose the master to connect to.
//
// It is possible that this is insufficient. The following may inform a
// better solution:

// https://github.com/LK4D4/sample
//
// The first link applies exponential distribution weight choice reservoir
// sampling. This may be relevant if we view the master selection as a
// distributed reservoir sampling problem.

// bias to zero-weighted remotes have same probability. otherwise, we
// always select first entry when all are zero.

// clear out workspace

// calculate CDF over weights

// if this peer is excluded, ignore it by continuing the loop to label Loop

// treat these as zero, to keep there selection unlikely.

func (mwr *remotesWeightedRandom) Observe(peer api.Peer, weight int) {
	_ = "STUB: not implemented"
	return
}

func (mwr *remotesWeightedRandom) ObserveIfExists(peer api.Peer, weight int) {
	_ = "STUB: not implemented"
	return
}

func (mwr *remotesWeightedRandom) Remove(addrs ...api.Peer) { _ = "STUB: not implemented"; return }

const (
	// remoteWeightSmoothingFactor for exponential smoothing. This adjusts how
	// much of the // observation and old value we are using to calculate the new value.
	// See
	// https://en.wikipedia.org/wiki/Exponential_smoothing#Basic_exponential_smoothing
	// for details.
	remoteWeightSmoothingFactor = 0.5
	remoteWeightMax             = 1 << 8
)

func clip(x float64) float64 {
	_ = "STUB: not implemented"

	// treat garbage as such
	// acts like a no-op for us.
	return 0
}

func (mwr *remotesWeightedRandom) observe(peer api.Peer, weight float64) {
	_ = "STUB: not implemented"

	// While we have a decent, ad-hoc approach here to weight subsequent
	// observations, we may want to look into applying forward decay:
	//
	//  http://dimacs.rutgers.edu/~graham/pubs/papers/fwddecay.pdf
	//
	// We need to get better data from behavior in a cluster.
	return
}

// makes the math easier to read below

// Multiply the new value to current value, and appy smoothing against the old
// value.
