package ca

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"sync"
	"time"

	cfconfig "github.com/cloudflare/cfssl/config"
	events "github.com/docker/go-events"
	"github.com/opencontainers/go-digest"
	"github.com/pkg/errors"
	"google.golang.org/grpc/credentials"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/connectionbroker"
	"github.com/moby/swarmkit/v2/watch"
)

const (
	rootCACertFilename  = "swarm-root-ca.crt"
	rootCAKeyFilename   = "swarm-root-ca.key"
	nodeTLSCertFilename = "swarm-node.crt"
	nodeTLSKeyFilename  = "swarm-node.key"

	// DefaultRootCN represents the root CN that we should create roots CAs with by default
	DefaultRootCN = "swarm-ca"
	// ManagerRole represents the Manager node type, and is used for authorization to endpoints
	ManagerRole = "swarm-manager"
	// WorkerRole represents the Worker node type, and is used for authorization to endpoints
	WorkerRole = "swarm-worker"
	// CARole represents the CA node type, and is used for clients attempting to get new certificates issued
	CARole = "swarm-ca"

	generatedSecretEntropyBytes = 16
	joinTokenBase               = 36
	// ceil(log(2^128-1, 36))
	maxGeneratedSecretLength = 25
	// ceil(log(2^256-1, 36))
	base36DigestLen = 50
)

var (
	// GetCertRetryInterval is how long to wait before retrying a node
	// certificate or root certificate request.
	GetCertRetryInterval = 2 * time.Second

	// errInvalidJoinToken is returned when attempting to parse an invalid join
	// token (e.g. when attempting to get the version, fipsness, or the root ca
	// digest)
	errInvalidJoinToken = errors.New("invalid join token")
)

// strongTLSCiphers defines a secure, modern set of TLS cipher suites
// with known weak algorithms removed.
var strongTLSCiphers = []uint16{
	tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
	tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
	tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
}

// SecurityConfig is used to represent a node's security configuration. It includes information about
// the RootCA and ServerTLSCreds/ClientTLSCreds transport authenticators to be used for MTLS
type SecurityConfig struct {
	// mu protects against concurrent access to fields inside the structure.
	mu sync.Mutex

	// renewalMu makes sure only one certificate renewal attempt happens at
	// a time. It should never be locked after mu is already locked.
	renewalMu sync.Mutex

	rootCA        *RootCA
	keyReadWriter *KeyReadWriter

	certificate *tls.Certificate
	issuerInfo  *IssuerInfo

	ServerTLSCreds *MutableTLSCreds
	ClientTLSCreds *MutableTLSCreds

	// An optional queue for anyone interested in subscribing to SecurityConfig updates
	queue *watch.Queue
}

// CertificateUpdate represents a change in the underlying TLS configuration being returned by
// a certificate renewal event.
type CertificateUpdate struct {
	Role string
	Err  error
}

// ParsedJoinToken is the data from a join token, once parsed
type ParsedJoinToken struct {
	// Version is the version of the join token that is being parsed
	Version int

	// RootDigest is the digest of the root CA certificate of the cluster, which
	// is always part of the join token so that the root CA can be verified
	// upon initial node join
	RootDigest digest.Digest

	// Secret is the randomly-generated secret part of the join token - when
	// rotating a join token, this is the value that is changed unless some other
	// property of the cluster (like the root CA) is changed.
	Secret string

	// FIPS indicates whether the join token specifies that the cluster mandates
	// that all nodes must have FIPS mode enabled.
	FIPS bool
}

// ParseJoinToken parses a join token.  Current format is v2, but this is currently used only if the cluster requires
// mandatory FIPS, in order to facilitate mixed version clusters.
// v1: SWMTKN-1-<SHA256 digest of root CA cert in base 36, 0-left-padded to 50 chars>-<16-byte secret in base 36 0-left-padded to 25 chars>
// v2: SWMTKN-2-<0/1 whether its FIPS or not>-<same rest of data as v1>
func ParseJoinToken(token string) (*ParsedJoinToken, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// v1 has 4, v2 has 5

func validateRootCAAndTLSCert(rootCA *RootCA, tlsKeyPair *tls.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

// NewSecurityConfig initializes and returns a new SecurityConfig.
func NewSecurityConfig(rootCA *RootCA, krw *KeyReadWriter, tlsKeyPair *tls.Certificate, issuerInfo *IssuerInfo) (*SecurityConfig, func() error, error) {
	_ = "STUB: not implemented"
	// Create the Server TLS Credentials for this node. These will not be used by workers.
	return nil, nil, nil
}

// Create a TLSConfig to be used when this node connects as a client to another remote node.
// We're using ManagerRole as remote serverName for TLS host verification because both workers
// and managers always connect to remote managers.

// RootCA returns the root CA.
func (s *SecurityConfig) RootCA() *RootCA { _ = "STUB: not implemented"; return nil }

// KeyWriter returns the object that can write keys to disk
func (s *SecurityConfig) KeyWriter() KeyWriter {
	_ = "STUB: not implemented"
	return *

	// KeyReader returns the object that can read keys from disk
	new(KeyWriter)
}

func (s *SecurityConfig) KeyReader() KeyReader {
	_ = "STUB: not implemented"
	return *

	// UpdateRootCA replaces the root CA with a new root CA
	new(KeyReader)
}

func (s *SecurityConfig) UpdateRootCA(rootCA *RootCA) error { _ = "STUB: not implemented"; return nil }

// refuse to update the root CA if the current TLS credentials do not validate against it

// Watch allows you to set a watch on the security config, in order to be notified of any changes
func (s *SecurityConfig) Watch() (chan events.Event, func()) {
	_ = "STUB: not implemented"
	return nil,

		// IssuerInfo returns the issuer subject and issuer public key
		nil
}

func (s *SecurityConfig) IssuerInfo() *IssuerInfo { _ = "STUB: not implemented"; return nil }

// This function expects something else to have taken out a lock on the SecurityConfig.
func (s *SecurityConfig) updateTLSCredentials(certificate *tls.Certificate, issuerInfo *IssuerInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateTLSCredentials updates the security config with an updated TLS certificate and issuer info
func (s *SecurityConfig) UpdateTLSCredentials(certificate *tls.Certificate, issuerInfo *IssuerInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// SigningPolicy creates a policy used by the signer to ensure that the only fields
// from the remote CSRs we trust are: PublicKey, PublicKeyAlgorithm and SignatureAlgorithm.
// It receives the duration a certificate will be valid for
func SigningPolicy(certExpiry time.Duration) *cfconfig.Signing {
	_ = "STUB: not implemented"
	// Force the minimum Certificate expiration to be fifteen minutes
	return nil
}

// Add the backdate

// Only trust the key components from the CSR. Everything else should
// come directly from API call params.

// SecurityConfigPaths is used as a helper to hold all the paths of security relevant files
type SecurityConfigPaths struct {
	Node, RootCA CertPaths
}

// NewConfigPaths returns the absolute paths to all of the different types of files
func NewConfigPaths(baseCertDir string) *SecurityConfigPaths { _ = "STUB: not implemented"; return nil }

// GenerateJoinToken creates a new join token. Current format is v2, but this is
// currently used only if the cluster requires mandatory FIPS, in order to
// facilitate mixed version clusters (the `fips` parameter is set to true).
// Otherwise, v1 is used so as to maintain compatibility in mixed version
// non-FIPS clusters.
// v1: SWMTKN-1-<SHA256 digest of root CA cert in base 36, 0-left-padded to 50 chars>-<16-byte secret in base 36 0-left-padded to 25 chars>
// v2: SWMTKN-2-<0/1 whether its FIPS or not>-<same rest of data as v1>
func GenerateJoinToken(rootCA *RootCA, fips bool) string { _ = "STUB: not implemented"; return "" }

// DownloadRootCA tries to retrieve a remote root CA and matches the digest against the provided token.
func DownloadRootCA(ctx context.Context, paths CertPaths, token string, connBroker *connectionbroker.Broker) (RootCA, error) {
	_ = "STUB: not implemented"

	// Get a digest for the optional CA hash string that we've been provided
	// If we were provided a non-empty string, and it is an invalid hash, return
	// otherwise, allow the invalid digest through.
	return *new(RootCA), nil
}

// Get the remote CA certificate, verify integrity with the
// hash provided. Retry up to 5 times, in case the manager we
// first try to contact is not responding properly (it may have
// just been demoted, for example).

// Save root CA certificate to disk

// LoadSecurityConfig loads TLS credentials from disk, or returns an error if
// these credentials do not exist or are unusable.
func LoadSecurityConfig(ctx context.Context, rootCA RootCA, krw *KeyReadWriter, allowExpired bool) (*SecurityConfig, func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil,

		// At this point we've successfully loaded the CA details from disk, or
		// successfully downloaded them remotely. The next step is to try to
		// load our certificates.
		nil
}

// Read both the Cert and Key from disk

// Check to see if this certificate was signed by our CA, and isn't expired

// ValidateChain, if successful, will always return at least 1 chain containing
// at least 2 certificates:  the leaf and the root.

// Now that we know this certificate is valid, create a TLS Certificate for our
// credentials

// CertificateRequestConfig contains the information needed to request a
// certificate from a remote CA.
type CertificateRequestConfig struct {
	// Token is the join token that authenticates us with the CA.
	Token string
	// Availability allows a user to control the current scheduling status of a node
	Availability api.NodeSpec_Availability
	// ConnBroker provides connections to CAs.
	ConnBroker *connectionbroker.Broker
	// Credentials provides transport credentials for communicating with the
	// remote server.
	Credentials credentials.TransportCredentials
	// ForceRemote specifies that only a remote (TCP) connection should
	// be used to request the certificate. This may be necessary in cases
	// where the local node is running a manager, but is in the process of
	// being demoted.
	ForceRemote bool
	// NodeCertificateStatusRequestTimeout determines how long to wait for a node
	// status RPC result.  If not provided (zero value), will default to 5 seconds.
	NodeCertificateStatusRequestTimeout time.Duration
	// RetryInterval specifies how long to delay between retries, if non-zero.
	RetryInterval time.Duration
	// Organization is the organization to use for a TLS certificate when creating
	// a security config from scratch.  If not provided, a random ID is generated.
	// For swarm certificates, the organization is the cluster ID.
	Organization string
}

// CreateSecurityConfig creates a new key and cert for this node, either locally
// or via a remote CA.
func (rootCA RootCA) CreateSecurityConfig(ctx context.Context, krw *KeyReadWriter, config CertificateRequestConfig) (*SecurityConfig, func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil,

		// Create a new random ID for this certificate
		nil
}

// Request certificate issuance from a remote CA.
// Last argument is nil because at this point we don't have any valid TLS creds

// TODO(cyli): currently we have to only update if it's a worker role - if we have a single root CA update path for
// both managers and workers, we won't need to check any more.
func updateRootThenUpdateCert(ctx context.Context, s *SecurityConfig, connBroker *connectionbroker.Broker, rootPaths CertPaths, failedCert *x509.Certificate) (*tls.Certificate, *IssuerInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// try downloading a new root CA if it's an unknown authority issue, in case there was a root rotation completion
// and we just didn't get the new root

// validate against the existing security config creds

// RenewTLSConfigNow gets a new TLS cert and key, and updates the security config if provided.  This is similar to
// RenewTLSConfig, except while that monitors for expiry, and periodically renews, this renews once and is blocking
func RenewTLSConfigNow(ctx context.Context, s *SecurityConfig, connBroker *connectionbroker.Broker, rootPaths CertPaths) error {
	_ = "STUB: not implemented"
	return nil
}

// Let's request new certs. Renewals don't require a token.

// calculateRandomExpiry returns a random duration between 50% and 80% of the
// original validity period
func calculateRandomExpiry(validFrom, validUntil time.Time) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Our lower bound of renewal will be half of the total expiration time

// Our upper bound of renewal will be 80% of the total expiration time

// Let's select a random number of minutes between min and max, and set our retry for that
// Using randomly selected rotation allows us to avoid certificate thundering herds.

// NewServerTLSConfig returns a tls.Config configured for a TLS Server, given a tls.Certificate
// and the PEM-encoded root CA Certificate
func NewServerTLSConfig(certs []tls.Certificate, rootCAPool *x509.CertPool) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Since we're using the same CA server to issue Certificates to new nodes, we can't
// use tls.RequireAndVerifyClientCert

// NewClientTLSConfig returns a tls.Config configured for a TLS Client, given a tls.Certificate
// the PEM-encoded root CA Certificate, and the name of the remote server the client wants to connect to.
func NewClientTLSConfig(certs []tls.Certificate, rootCAPool *x509.CertPool, serverName string) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewClientTLSCredentials returns GRPC credentials for a TLS GRPC client, given a tls.Certificate
// a PEM-Encoded root CA Certificate, and the name of the remote server the client wants to connect to.
func (rootCA *RootCA) NewClientTLSCredentials(cert *tls.Certificate, serverName string) (*MutableTLSCreds, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewServerTLSCredentials returns GRPC credentials for a TLS GRPC client, given a tls.Certificate
// a PEM-Encoded root CA Certificate, and the name of the remote server the client wants to connect to.
func (rootCA *RootCA) NewServerTLSCredentials(cert *tls.Certificate) (*MutableTLSCreds, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseRole parses an apiRole into an internal role string
func ParseRole(apiRole api.NodeRole) (string, error) { _ = "STUB: not implemented"; return "", nil }

// FormatRole parses an internal role string into an apiRole
func FormatRole(role string) (api.NodeRole, error) {
	_ = "STUB: not implemented"
	return *new(api.NodeRole), nil
}
