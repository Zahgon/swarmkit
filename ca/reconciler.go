package ca

import (
	"context"
	"sync"
	"time"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/state/store"
	"github.com/pkg/errors"
)

// IssuanceStateRotateMaxBatchSize is the maximum number of nodes we'll tell to rotate their certificates in any given update
const IssuanceStateRotateMaxBatchSize = 30

func hasIssuer(n *api.Node, info *IssuerInfo) bool { _ = "STUB: not implemented"; return false }

var errRootRotationChanged = errors.New("target root rotation has changed")

// rootRotationReconciler keeps track of all the nodes in the store so that we can determine which ones need reconciliation when nodes are updated
// or the root CA is updated.  This is meant to be used with watches on nodes and the cluster, and provides functions to be called when the
// cluster's RootCA has changed and when a node is added, updated, or removed.
type rootRotationReconciler struct {
	mu                  sync.Mutex
	clusterID           string
	batchUpdateInterval time.Duration
	ctx                 context.Context
	store               *store.MemoryStore

	currentRootCA    *api.RootCA
	currentIssuer    IssuerInfo
	unconvergedNodes map[string]*api.Node

	wg     sync.WaitGroup
	cancel func()
}

// IssuerFromAPIRootCA returns the desired issuer given an API root CA object
func IssuerFromAPIRootCA(rootCA *api.RootCA) (*IssuerInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// assumption:  UpdateRootCA will never be called with a `nil` root CA because the caller will be acting in response to
// a store update event
func (r *rootRotationReconciler) UpdateRootCA(newRootCA *api.RootCA) {
	_ = "STUB: not implemented"
	return
}

// check if the issuer has changed, first

// If the issuer has changed, iterate through all the nodes to figure out which ones need rotation

// from here on out, there will be no more errors that cause us to have to abandon updating the Root CA,
// so we can start making changes to r's fields

// there's already a loop going, so cancel it

// assumption:  UpdateNode will never be called with a `nil` node because the caller will be acting in response to
// a store update event
func (r *rootRotationReconciler) UpdateNode(node *api.Node) { _ = "STUB: not implemented"; return }

// if we're not in the middle of a root rotation ignore the update

// assumption:  DeleteNode will never be called with a `nil` node because the caller will be acting in response to
// a store update event
func (r *rootRotationReconciler) DeleteNode(node *api.Node) { _ = "STUB: not implemented"; return }

func (r *rootRotationReconciler) runReconcilerLoop(ctx context.Context, loopRootCA *api.RootCA) {
	_ = "STUB: not implemented"
	return
}

// if the root rotation has changed, this loop will be cancelled anyway, so may as well abort early

// This function assumes that the expected root CA has root rotation.  This is intended to be used by
// `reconcileNodeRootsAndCerts`, which uses the root CA from the `lastSeenClusterRootCA`, and checks
// that it has a root rotation before calling this function.
func (r *rootRotationReconciler) finishRootRotation(tx store.Tx, expectedRootCA *api.RootCA) error {
	_ = "STUB: not implemented"
	return nil
}

// If the RootCA object has changed (because another root rotation was started or because some other node
// had finished the root rotation), we cannot finish the root rotation that we were working on.

// we don't actually have to parse out the default node expiration from the cluster - we are just using
// the ca.RootCA object to generate new tokens and the digest

func (r *rootRotationReconciler) batchUpdateNodes(toUpdate []*api.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// Directly update the nodes rather than get + update, and ignore version errors.  Since
// `rootRotationReconciler` should be hooked up to all node update/delete/create events, we should have
// close to the latest versions of all the nodes.  If not, the node will updated later and the
// next batch of updates should catch it.
