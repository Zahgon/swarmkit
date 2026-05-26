package ca

import (
	"context"
	"crypto/tls"
	"crypto/x509/pkix"

	"github.com/moby/swarmkit/v2/api"
)

type localRequestKeyType struct{}

// LocalRequestKey is a context key to mark a request that originating on the
// local node. The associated value is a RemoteNodeInfo structure describing the
// local node.
var LocalRequestKey = localRequestKeyType{}

// LogTLSState logs information about the TLS connection and remote peers
func LogTLSState(ctx context.Context, tlsState *tls.ConnectionState) {
	_ = "STUB: not implemented"
	return
}

// "peer.verifiedChain": verifiedChain},

// getCertificateSubject extracts the subject from a verified client certificate
func getCertificateSubject(tlsState *tls.ConnectionState) (pkix.Name, error) {
	_ = "STUB: not implemented"
	return *new(pkix.Name), nil
}

func tlsConnStateFromContext(ctx context.Context) (*tls.ConnectionState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// certSubjectFromContext extracts pkix.Name from context.
func certSubjectFromContext(ctx context.Context) (pkix.Name, error) {
	_ = "STUB: not implemented"
	return *new(pkix.Name), nil
}

// AuthorizeOrgAndRole takes in a context and a list of roles, and returns
// the Node ID of the node.
func AuthorizeOrgAndRole(ctx context.Context, org string, blacklistedCerts map[string]*api.BlacklistedCertificate, ou ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Check if the current certificate has an OU that authorizes
// access to this method

// authorizeOrg takes in a certificate subject and an organization, and returns
// the Node ID of the node.
func authorizeOrg(certSubj pkix.Name, org string, blacklistedCerts map[string]*api.BlacklistedCertificate) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// AuthorizeForwardedRoleAndOrg checks for proper roles and organization of caller. The RPC may have
// been proxied by a manager, in which case the manager is authenticated and
// so is the certificate information that it forwarded. It returns the node ID
// of the original client.
func AuthorizeForwardedRoleAndOrg(ctx context.Context, authorizedRoles, forwarderRoles []string, org string, blacklistedCerts map[string]*api.BlacklistedCertificate) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// This was a forwarded request. Authorize the forwarder, and
// check if the forwarded role matches one of the authorized
// roles.

// There wasn't any node being forwarded, check if this is a direct call by the expected role

// intersectArrays returns true when there is at least one element in common
// between the two arrays
func intersectArrays(orig, tgt []string) bool { _ = "STUB: not implemented"; return false }

// RemoteNodeInfo describes a node sending an RPC request.
type RemoteNodeInfo struct {
	// Roles is a list of roles contained in the node's certificate
	// (or forwarded by a trusted node).
	Roles []string

	// Organization is the organization contained in the node's certificate
	// (or forwarded by a trusted node).
	Organization string

	// NodeID is the node's ID, from the CN field in its certificate
	// (or forwarded by a trusted node).
	NodeID string

	// ForwardedBy contains information for the node that forwarded this
	// request. It is set to nil if the request was received directly.
	ForwardedBy *RemoteNodeInfo

	// RemoteAddr is the address that this node is connecting to the cluster
	// from.
	RemoteAddr string
}

// RemoteNode returns the node ID and role from the client's TLS certificate.
// If the RPC was forwarded, the original client's ID and role is returned, as
// well as the forwarder's ID. This function does not do authorization checks -
// it only looks up the node ID.
func RemoteNode(ctx context.Context) (RemoteNodeInfo, error) {
	_ = "STUB: not implemented"
	// If we have a value on the context that marks this as a local
	// request, we return the node info from the context.
	return *new(RemoteNodeInfo), nil
}
