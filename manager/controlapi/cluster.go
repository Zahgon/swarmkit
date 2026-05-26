package controlapi

import (
	"context"
	"time"

	"github.com/moby/swarmkit/v2/api"
)

const (
	// expiredCertGrace is the amount of time to keep a node in the
	// blacklist beyond its certificate expiration timestamp.
	expiredCertGrace = 24 * time.Hour * 7
	// inbuilt default subnet size
	inbuiltSubnetSize = 24
	// VXLAN default port
	defaultVXLANPort = 4789
)

var (
	// inbuilt default address pool
	inbuiltDefaultAddressPool = []string{"10.0.0.0/8"}
)

func validateClusterSpec(spec *api.ClusterSpec) error { _ = "STUB: not implemented"; return nil }

// Validate that expiry time being provided is valid, and over our minimum

// Validate that AcceptancePolicies only include Secrets that are bcrypted
// TODO(diogo): Add a global list of acceptance algorithms. We only support bcrypt for now.

// Validate that heartbeatPeriod time being provided is valid

// GetCluster returns a Cluster given a ClusterID.
// - Returns `InvalidArgument` if ClusterID is not provided.
// - Returns `NotFound` if the Cluster is not found.
func (s *Server) GetCluster(_ context.Context, request *api.GetClusterRequest) (*api.GetClusterResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WARN: we should never return cluster here. We need to redact the private fields first.

// UpdateCluster updates a Cluster referenced by ClusterID with the given ClusterSpec.
// - Returns `NotFound` if the Cluster is not found.
// - Returns `InvalidArgument` if the ClusterSpec is malformed.
// - Returns `Unimplemented` if the ClusterSpec references unimplemented features.
// - Returns an error if the update fails.
func (s *Server) UpdateCluster(ctx context.Context, request *api.UpdateClusterRequest) (*api.UpdateClusterResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This ensures that we have the current rootCA with which to generate tokens (expiration doesn't matter
// for generating the tokens)

// WARN: we should never return cluster here. We need to redact the private fields first.

func filterClusters(candidates []*api.Cluster, filters ...func(*api.Cluster) bool) []*api.Cluster {
	_ = "STUB: not implemented"
	return nil
}

// ListClusters returns a list of all clusters.
func (s *Server) ListClusters(_ context.Context, request *api.ListClustersRequest) (*api.ListClustersResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WARN: we should never return cluster here. We need to redact the private fields first.

// redactClusters is a method that enforces a whitelist of fields that are ok to be
// returned in the Cluster object. It should filter out all sensitive information.
func redactClusters(clusters []*api.Cluster) []*api.Cluster { _ = "STUB: not implemented"; return nil }

// Only add public fields to the new clusters

// Copy all the mandatory fields
// Do not copy secret keys

// the cert is not a secret, but if API users get the cluster spec and then update,
// then because the cert is included but not the key, the user can get update errors
// or unintended consequences (such as telling swarm to forget about the key so long
// as there is a corresponding external CA)

// This is just for CLI display. Set the inbuilt default pool for
// user reference.

func expireBlacklistedCerts(cluster *api.Cluster) { _ = "STUB: not implemented"; return }
