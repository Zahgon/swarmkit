package ca

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"sync"
	"time"

	"github.com/cloudflare/cfssl/signer"
	"github.com/pkg/errors"
)

const (
	// ExternalCrossSignProfile is the profile that we will be sending cross-signing CSR sign requests with
	ExternalCrossSignProfile = "CA"

	// CertificateMaxSize is the maximum expected size of a certificate.
	// While there is no specced upper limit to the size of an x509 certificate in PEM format,
	// one with a ridiculous RSA key size (16384) and 26 256-character DNS SAN fields is about 14k.
	// While there is no upper limit on the length of certificate chains, long chains are impractical.
	// To be conservative, and to also account for external CA certificate responses in JSON format
	// from CFSSL, we'll set the max to be 256KiB.
	CertificateMaxSize int64 = 256 << 10
)

// ErrNoExternalCAURLs is an error used it indicate that an ExternalCA is
// configured with no URLs to which it can proxy certificate signing requests.
var ErrNoExternalCAURLs = errors.New("no external CA URLs")

// ExternalCA is able to make certificate signing requests to one of a list
// remote CFSSL API endpoints.
type ExternalCA struct {
	ExternalRequestTimeout time.Duration

	mu            sync.Mutex
	intermediates []byte
	urls          []string
	client        *http.Client
}

// NewExternalCATLSConfig takes a TLS certificate and root pool and returns a TLS config that can be updated
// without killing existing connections
func NewExternalCATLSConfig(certs []tls.Certificate, rootPool *x509.CertPool) *tls.Config {
	_ = "STUB: not implemented"
	return nil
}

// NewExternalCA creates a new ExternalCA which uses the given tlsConfig to
// authenticate to any of the given URLS of CFSSL API endpoints.
func NewExternalCA(intermediates []byte, tlsConfig *tls.Config, urls ...string) *ExternalCA {
	_ = "STUB: not implemented"
	return nil
}

// UpdateTLSConfig updates the HTTP Client for this ExternalCA by creating
// a new client which uses the given tlsConfig.
func (eca *ExternalCA) UpdateTLSConfig(tlsConfig *tls.Config) { _ = "STUB: not implemented"; return }

// UpdateURLs updates the list of CSR API endpoints by setting it to the given urls.
func (eca *ExternalCA) UpdateURLs(urls ...string) { _ = "STUB: not implemented"; return }

// Sign signs a new certificate by proxying the given certificate signing
// request to an external CFSSL API server.
func (eca *ExternalCA) Sign(ctx context.Context, req signer.SignRequest) (cert []byte, err error) {
	_ = "STUB: not implemented"
	// Get the current HTTP client and list of URLs in a small critical
	// section. We will use these to make certificate signing requests.
	return nil, nil
}

// Try each configured proxy URL. Return after the first success. If
// all fail then the last error will be returned.

// CrossSignRootCA takes a RootCA object, generates a CA CSR, sends a signing request with the CA CSR to the external
// CFSSL API server in order to obtain a cross-signed root
func (eca *ExternalCA) CrossSignRootCA(ctx context.Context, rca RootCA) ([]byte, error) {
	_ = "STUB: not implemented"
	// ExtractCertificateRequest generates a new key request, and we want to continue to use the old
	// key.  However, ExtractCertificateRequest will also convert the pkix.Name to csr.Name, which we
	// need in order to generate a signing request
	return nil, nil
}

// cfssl actually ignores non subject alt name extensions in the CSR, so we have to add the CA extension in the signing
// request as well

func makeExternalSignRequest(ctx context.Context, client *http.Client, url string, csrJSON []byte) (cert []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
