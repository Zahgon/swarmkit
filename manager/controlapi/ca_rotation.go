package controlapi

import (
	"context"
	"crypto/tls"
	"net"

	"github.com/cloudflare/cfssl/helpers"
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/ca"
)

var minRootExpiration = 1 * helpers.OneYear

// determines whether an api.RootCA, api.RootRotation, or api.CAConfig has a signing key (local signer)
func hasSigningKey(a interface{}) bool { _ = "STUB: not implemented"; return false }

// Creates a cross-signed intermediate and new api.RootRotation object.
// This function assumes that the root cert and key and the external CAs have already been validated.
func newRootRotationObject(ctx context.Context, securityConfig *ca.SecurityConfig, apiRootCA *api.RootCA, newCARootCA ca.RootCA, extCAs []*api.ExternalCA, version uint64) (*api.RootCA, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// we have to sign with the original signer, not whatever is in the SecurityConfig's RootCA (which may have an intermediate signer, if
// a root rotation is already in progress)

// the original CA and the new CA both require external CAs

// We need the same credentials but to connect to the original URLs (in case we are in the middle of a root rotation already)

// Checks that a CA URL is connectable using the credentials we have and that its server certificate is signed by the
// root CA that we expect.  This uses a TCP dialer rather than an HTTP client; because we have custom TLS configuration,
// if we wanted to use an HTTP client we'd have to create a new transport for every connection.  The docs specify that
// Transports cache connections for future re-use, which could cause many open connections.
func validateExternalCAURL(dialer *net.Dialer, tlsOpts *tls.Config, caURL string) error {
	_ = "STUB: not implemented"
	return nil
}

// It either has no port or is otherwise invalid (e.g. too many colons).  If it's otherwise invalid the dialer
// will error later, so just assume it's no port and set the port to the default HTTPS port.

// Validates that there is at least 1 reachable, valid external CA for the given CA certificate.  Returns true if there is, false otherwise.
// Requires that the wanted cert is already normalized.
func validateHasAtLeastOneExternalCA(ctx context.Context, externalCAs map[string][]*api.ExternalCA, securityConfig *ca.SecurityConfig,
	wantedCert []byte, desc string) ([]*api.ExternalCA, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validates that the list of external CAs have valid certs associated with them, and produce a mapping of subject/pubkey:external
// for later validation of required external CAs
func getNormalizedExtCAs(caConfig *api.CAConfig, normalizedCurrentRootCACert []byte) (map[string][]*api.ExternalCA, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if no associated cert is provided, assume it's the current root cert

// validateAndUpdateCA validates a cluster's desired CA configuration spec, and returns a RootCA value on success representing
// current RootCA as it should be.  Validation logic and return values are as follows:
//  1. Validates that the contents are complete - e.g. a signing key is not provided without a signing cert, and that external
//     CAs are not removed if they are needed.  Otherwise, returns an error.
//  2. If no desired signing cert or key are provided, then either:
//     - we are happy with the current CA configuration (force rotation value has not changed), and we return the current RootCA
//     object as is
//     - we want to generate a new internal CA cert and key (force rotation value has changed), and we return the updated RootCA
//     object
//  3. Check if the cert is the same key. We cannot rotate to a cert with the same key. As of go 1.19, the logic for certificate
//     trust chain validation changed, and a chain including two certs with the same key will not validate. This case would
//     usually occur when reissuing the same cert with a later expiration date. Because of this validation failure, our root
//     rotation algorithm fails. While it might be possible to adjust the rotation procedure to accommodate such a cert change,
//     it is somewhat of an edge case, and, more importantly, we do not currently possess the cryptographic expertise to safely
//     make such a change. So, as a result, this operation is disallowed. The new root cert must have a new key.
//  4. Signing cert and key have been provided: validate that these match (the cert and key match). Otherwise, return an error.
//  5. Return the updated RootCA object according to the following criteria:
//     - If the desired cert is the same as the current CA cert then abort any outstanding rotations. The current signing key
//     is replaced with the desired signing key (this could lets us switch between external->internal or internal->external
//     without an actual CA rotation, which is not needed because any leaf cert issued with one CA cert can be validated using
//     the second CA certificate).
//     - If the desired cert is the same as the current to-be-rotated-to CA cert then a new root rotation is not needed. The
//     current to-be-rotated-to signing key is replaced with the desired signing key (this could lets us switch between
//     external->internal or internal->external without an actual CA rotation, which is not needed because any leaf cert
//     issued with one CA cert can be validated using the second CA certificate).
//     - Otherwise, start a new root rotation using the desired signing cert and desired signing key as the root rotation
//     signing cert and key.  If a root rotation is already in progress, just replace it and start over.
func validateCAConfig(ctx context.Context, securityConfig *ca.SecurityConfig, cluster *api.Cluster) (*api.RootCA, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ensure this is normalized before we use it

// validate that the list of external CAs is not malformed

// If we are going from external -> internal, but providing the external CA's signing key,
// then we don't need to validate any external CAs.  We can in fact abort any outstanding root
// rotations if we are just adding a key.  Because we have a key, we don't care if there are
// no external CAs matching the certificate.

// validate that the key and cert indeed match - if they don't then just fail now rather
// than go through all the external CA URLs, which is a more expensive operation

// if the desired CA cert and key are not set, then we are happy with the current root CA configuration, unless
// the ForceRotate version has changed

// we also need to make sure that if the current root rotation requires an external CA, those external CAs are
// still valid

// no change, return as is

// A desired cert and maybe key were provided - we need to make sure the cert and key (if provided) match.

// The new certificate's expiry must be at least one year away

// check if we can abort any existing root rotations

// See step 3 in the doc comment. We cannot upgrade a cert with the same
// key.

// check if this is the same desired cert as an existing root rotation

// ok, everything's different; we have to begin a new root rotation which means generating a new cross-signed cert
