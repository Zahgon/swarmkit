package testutils

import (
	"net"
	"net/http"
	"sync"

	"github.com/cloudflare/cfssl/config"
	"github.com/moby/swarmkit/v2/ca"
)

var crossSignPolicy = config.SigningProfile{
	Usage: []string{"cert sign", "crl sign"},
	// we don't want the intermediate to last for very long
	Expiry:       ca.DefaultNodeCertExpiration,
	Backdate:     ca.CertBackdate,
	CAConstraint: config.CAConstraint{IsCA: true},
	ExtensionWhitelist: map[string]bool{
		ca.BasicConstraintsOID.String(): true,
	},
}

// NewExternalSigningServer creates and runs a new ExternalSigningServer which
// uses the given rootCA to sign node certificates. A server key and cert are
// generated and saved into the given basedir and then a TLS listener is
// started on a random available port. On success, an HTTPS server will be
// running in a separate goroutine. The URL of the singing endpoint is
// available in the returned *ExternalSignerServer value. Calling the Close()
// method will stop the server.
func NewExternalSigningServer(rootCA ca.RootCA, basedir string) (*ExternalSigningServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Make a valid server cert for localhost.

// create our own copy of the local signer so we don't mutate the rootCA's signer as we enable and disable CA signing

// Create TLS credentials for the external CA server which we will run.

// ExternalSigningServer runs an HTTPS server with an endpoint at a specified
// URL which signs node certificate requests from a swarm manager client.
type ExternalSigningServer struct {
	listener  net.Listener
	NumIssued uint64
	URL       string
	flaky     uint32
	handler   *signHandler
}

// Stop stops this signing server by closing the underlying TCP/TLS listener.
func (ess *ExternalSigningServer) Stop() error { _ = "STUB: not implemented"; return nil }

// Flake makes the signing server return HTTP 500 errors.
func (ess *ExternalSigningServer) Flake() { _ = "STUB: not implemented"; return }

// Deflake restores normal operation after a call to Flake.
func (ess *ExternalSigningServer) Deflake() { _ = "STUB: not implemented"; return }

// EnableCASigning updates the root CA signer to be able to sign CAs
func (ess *ExternalSigningServer) EnableCASigning() error { _ = "STUB: not implemented"; return nil }

// DisableCASigning prevents the server from being able to sign CA certificates
func (ess *ExternalSigningServer) DisableCASigning() { _ = "STUB: not implemented"; return }

type signHandler struct {
	mu          sync.Mutex
	numIssued   *uint64
	flaky       *uint32
	localSigner *ca.LocalSigner
	origPolicy  *config.Signing
}

func (h *signHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Check client authentication via mutual TLS.

// The client certificate OU should be for a swarm manager.

// The client certificate must have an Org.

// Decode the certificate signing request.

// The signReq should have additional subject info.

// The client's Org should match the Org in the sign request subject.

// Sign the requested certificate.

// Increment the number of certs issued.

// Return a successful JSON response.
