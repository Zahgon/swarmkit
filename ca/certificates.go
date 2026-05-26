package ca

import (
	"context"
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"encoding/asn1"
	"time"

	cflog "github.com/cloudflare/cfssl/log"
	cfsigner "github.com/cloudflare/cfssl/signer"
	"github.com/moby/swarmkit/v2/connectionbroker"
	"github.com/opencontainers/go-digest"
	"github.com/pkg/errors"
	"google.golang.org/grpc/credentials"
)

const (
	// Security Strength Equivalence
	//
	// | ECC      | DH/DSA/RSA    |
	// |----------|---------------|
	// | 256      | 3072          |
	// | 384      | 7680          |

	// RootKeySize is the default size of the root CA key
	// It would be ideal for the root key to use P-384, but in P-384 is not optimized in go yet :(
	RootKeySize = 256
	// RootKeyAlgo defines the default algorithm for the root CA Key
	RootKeyAlgo = "ecdsa"
	// RootCAExpiration represents the default expiration for the root CA in seconds (20 years)
	RootCAExpiration = "630720000s"
	// DefaultNodeCertExpiration represents the default expiration for node certificates (3 months)
	DefaultNodeCertExpiration = 2160 * time.Hour
	// CertBackdate represents the amount of time each certificate is backdated to try to avoid
	// clock drift issues.
	CertBackdate = 1 * time.Hour
	// CertLowerRotationRange represents the minimum fraction of time that we will wait when randomly
	// choosing our next certificate rotation
	CertLowerRotationRange = 0.5
	// CertUpperRotationRange represents the maximum fraction of time that we will wait when randomly
	// choosing our next certificate rotation
	CertUpperRotationRange = 0.8
	// MinNodeCertExpiration represents the minimum expiration for node certificates
	MinNodeCertExpiration = 1 * time.Hour
)

// BasicConstraintsOID is the ASN1 Object ID indicating a basic constraints extension
var BasicConstraintsOID = asn1.ObjectIdentifier{2, 5, 29, 19}

// A recoverableErr is a non-fatal error encountered signing a certificate,
// which means that the certificate issuance may be retried at a later time.
type recoverableErr struct {
	err error
}

func (r recoverableErr) Error() string { _ = "STUB: not implemented"; return "" }

// ErrNoLocalRootCA is an error type used to indicate that the local root CA
// certificate file does not exist.
var ErrNoLocalRootCA = errors.New("local root CA certificate does not exist")

// ErrNoValidSigner is an error type used to indicate that our RootCA doesn't have the ability to
// sign certificates.
var ErrNoValidSigner = recoverableErr{err: errors.New("no valid signer found")}

func init() {
	cflog.Level = 5
}

// CertPaths is a helper struct that keeps track of the paths of a
// Cert and corresponding Key
type CertPaths struct {
	Cert, Key string
}

// IssuerInfo contains the subject and public key of the issuer of a certificate
type IssuerInfo struct {
	Subject   []byte
	PublicKey []byte
}

// LocalSigner is a signer that can sign CSRs
type LocalSigner struct {
	cfsigner.Signer

	// Key will only be used by the original manager to put the private
	// key-material in raft, no signing operations depend on it.
	Key []byte

	// Cert is one PEM encoded Certificate used as the signing CA.  It must correspond to the key.
	Cert []byte

	// just cached parsed values for validation, etc.
	parsedCert   *x509.Certificate
	cryptoSigner crypto.Signer
}

type x509UnknownAuthError struct {
	error
	failedLeafCert *x509.Certificate
}

// RootCA is the representation of everything we need to sign certificates and/or to verify certificates
//
// RootCA.Cert:          [CA cert1][CA cert2]
// RootCA.Intermediates: [intermediate CA1][intermediate CA2][intermediate CA3]
// RootCA.signer.Cert:   [signing CA cert]
// RootCA.signer.Key:    [signing CA key]
//
// Requirements:
//
//   - [signing CA key] must be the private key for [signing CA cert], and either both or none must be provided
//   - [intermediate CA1] must have the same public key and subject as [signing CA cert], because otherwise when
//     appended to a leaf certificate, the intermediates will not form a chain (because [intermediate CA1] won't because
//     the signer of the leaf certificate)
//   - [intermediate CA1] must be signed by [intermediate CA2], which must be signed by [intermediate CA3]
//   - When we issue a certificate, the intermediates will be appended so that the certificate looks like:
//     [leaf signed by signing CA cert][intermediate CA1][intermediate CA2][intermediate CA3]
//   - [leaf signed by signing CA cert][intermediate CA1][intermediate CA2][intermediate CA3] is guaranteed to form a
//     valid chain from [leaf signed by signing CA cert] to one of the root certs ([signing CA cert], [CA cert1], [CA cert2])
//     using zero or more of the intermediate certs ([intermediate CA1][intermediate CA2][intermediate CA3]) as intermediates
//
// Example 1:  Simple root rotation
//
// - Initial state:
//   - RootCA.Cert:          [Root CA1 self-signed]
//   - RootCA.Intermediates: []
//   - RootCA.signer.Cert:   [Root CA1 self-signed]
//   - Issued TLS cert:      [leaf signed by Root CA1]
//
// - Intermediate state (during root rotation):
//   - RootCA.Cert:          [Root CA1 self-signed]
//   - RootCA.Intermediates: [Root CA2 signed by Root CA1]
//   - RootCA.signer.Cert:   [Root CA2 signed by Root CA1]
//   - Issued TLS cert:      [leaf signed by Root CA2][Root CA2 signed by Root CA1]
//
// - Final state:
//   - RootCA.Cert:          [Root CA2 self-signed]
//   - RootCA.Intermediates: []
//   - RootCA.signer.Cert:   [Root CA2 self-signed]
//   - Issued TLS cert:      [leaf signed by Root CA2]
type RootCA struct {
	// Certs contains a bundle of self-signed, PEM encoded certificates for the Root CA to be used
	// as the root of trust.
	Certs []byte

	// Intermediates contains a bundle of PEM encoded intermediate CA certificates to append to any
	// issued TLS (leaf) certificates. The first one must have the same public key and subject as the
	// signing root certificate, and the rest must form a chain, each one certifying the one above it,
	// as per RFC5246 section 7.4.2.
	Intermediates []byte

	// Pool is the root pool used to validate TLS certificates
	Pool *x509.CertPool

	// Digest of the serialized bytes of the certificate(s)
	Digest digest.Digest

	// This signer will be nil if the node doesn't have the appropriate key material
	signer *LocalSigner
}

// Signer is an accessor for the local signer that returns an error if this root cannot sign.
func (rca *RootCA) Signer() (*LocalSigner, error) { _ = "STUB: not implemented"; return nil, nil }

// IssueAndSaveNewCertificates generates a new key-pair, signs it with the local root-ca, and returns a
// TLS certificate and the issuer information for the certificate.
func (rca *RootCA) IssueAndSaveNewCertificates(kw KeyWriter, cn, ou, org string) (*tls.Certificate, *IssuerInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Obtain a signed Certificate

// should never happen, since if ParseValidateAndSignCSR did not fail this root CA must have a signer

// Create a valid TLSKeyPair out of the PEM encoded private key and certificate

// RequestAndSaveNewCertificates gets new certificates issued, either by signing them locally if a signer is
// available, or by requesting them from the remote server at remoteAddr.  This function returns the TLS
// certificate and the issuer information for the certificate.
func (rca *RootCA) RequestAndSaveNewCertificates(ctx context.Context, kw KeyWriter, config CertificateRequestConfig) (*tls.Certificate, *IssuerInfo, error) {
	_ = "STUB: not implemented"
	// Create a new key/pair and CSR
	return nil, nil, nil
}

// Get the remote manager to issue a CA signed certificate for this node
// Retry up to 5 times in case the manager we first try to contact isn't
// responding properly (for example, it may have just been demoted).

// If the first attempt fails, we should try a remote
// connection. The local node may be a manager that was
// demoted, so the local connection (which is preferred) may
// not work. If we are successful in renewing the certificate,
// the local connection will not be returned by the connection
// broker anymore.

// Wait a moment, in case a leader election was taking place.

// Доверяй, но проверяй.
// Before we overwrite our local key + certificate, let's make sure the server gave us one that is valid
// Create an X509Cert so we can .Verify()
// Check to see if this certificate was signed by our CA, and isn't expired

// TODO(cyli): - right now we need the invalid certificate in order to determine whether or not we should
// download a new root, because we only want to do that in the case of workers.  When we have a single
// codepath for updating the root CAs for both managers and workers, this snippet can go.

// ValidateChain, if successful, will always return at least 1 parsed cert and at least 1 chain containing
// at least 2 certificates:  the leaf and the root.

// Create a valid TLSKeyPair out of the PEM encoded private key and certificate

// ValidateCertChain will always return at least 1 cert, so indexing at 0 is safe

// Wait a moment, in case a leader election was taking place.

func (rca *RootCA) getKEKUpdate(ctx context.Context, leafCert *x509.Certificate, keypair tls.Certificate, config CertificateRequestConfig) (*KEKData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if the server does not support keks, return as if no encryption key was specified

// If this is a worker, set to never encrypt. We always want to set to the lock key to nil,
// in case this was a manager that was demoted to a worker.

// PrepareCSR creates a CFSSL Sign Request based on the given raw CSR and
// overrides the Subject and Hosts with the given extra args.
func PrepareCSR(csrBytes []byte, cn, ou, org string) cfsigner.SignRequest {
	_ = "STUB: not implemented"
	// All managers get added the subject-alt-name of CA, so they can be
	// used for cert issuance.
	return *new(cfsigner.SignRequest)
}

// OU is used for Authentication of the node type. The CN has the random
// node ID.

// Adding ou as DNS alt name, so clients can connect to ManagerRole and CARole

// ParseValidateAndSignCSR returns a signed certificate from a particular rootCA and a CSR.
func (rca *RootCA) ParseValidateAndSignCSR(csrBytes []byte, cn, ou, org string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CrossSignCACertificate takes a CA root certificate and generates an intermediate CA from it signed with the current root signer
func (rca *RootCA) CrossSignCACertificate(otherCAPEM []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create a new cert with exactly the same parameters, including the public key and exact NotBefore and NotAfter

// make sure we can sign with the signer key

func validateSignatureAlgorithm(cert *x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

// NewRootCA creates a new RootCA object from unparsed PEM cert bundle and key byte
// slices. key may be nil, and in this case NewRootCA will return a RootCA
// without a signer.
func NewRootCA(rootCertBytes, signCertBytes, signKeyBytes []byte, certExpiry time.Duration, intermediates []byte) (RootCA, error) {
	_ = "STUB: not implemented"
	// Parse all the certificates in the cert bundle
	return *new(RootCA), nil
}

// Check to see if we have at least one valid cert

// Create a Pool with all of the certificates found

// Check to see if all of the certificates are valid, self-signed root CA certs

// Calculate the digest for our Root CA bundle

// The intermediates supplied must be able to chain up to the root certificates, so that when they are appended to
// a leaf certificate, the leaf certificate can be validated through the intermediates to the root certificates.

// If a signer is provided and there are intermediates, then either the first intermediate would be the signer CA
// certificate (in which case it'd have the same subject and public key), or it would be a cross-signed
// intermediate with the same subject and public key as our signing CA certificate (which could be either an
// intermediate cert or a self-signed root cert).

// ValidateCertChain checks checks that the certificates provided chain up to the root pool provided.  In addition
// it also enforces that every cert in the bundle certificates form a chain, each one certifying the one above,
// as per RFC5246 section 7.4.2, and that every certificate (whether or not it is necessary to form a chain to the root
// pool) is currently valid and not yet expired (unless allowExpiry is set to true).
// This is additional validation not required by go's Certificate.Verify (which allows invalid certs in the
// intermediate pool), because this function is intended to be used when reading certs from untrusted locations such as
// from disk or over a network when a CSR is signed, so it is extra pedantic.
// This function always returns all the parsed certificates in the bundle in order, which means there will always be
// at least 1 certificate if there is no error, and the valid chains found by Certificate.Verify
func ValidateCertChain(rootPool *x509.CertPool, certs []byte, allowExpired bool) ([]*x509.Certificate, [][]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	// Parse all the certificates in the cert bundle
	return nil, nil, nil
}

// ensure that they form a chain, each one being signed by the one after it

// Manual expiry validation because we want more information on which certificate in the chain is expired, and
// because this is an easier way to allow expired certs.

// check that the previous cert was signed by this cert

// If we accept expired certs, try to build a valid cert chain using some subset of the certs.  We start off using the
// first certificate's NotAfter as the current time, thus ensuring that the first cert is not expired. If the chain
// still fails to validate due to expiry issues, continue iterating over the rest of the certs.
// If any of the other certs has an earlier NotAfter time, use that time as the current time instead. This insures that
// particular cert, and any that came before it, are not expired.  Note that the root that the certs chain up to
// should also not be expired at that "current" time.

// newLocalSigner validates the signing cert and signing key to create a local signer, which accepts a crypto signer and a cert
func newLocalSigner(keyBytes, certBytes []byte, certExpiry time.Duration, rootPool, intermediatePool *x509.CertPool) (*LocalSigner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The key should not be encrypted, but it could be in PKCS8 format rather than PKCS1

// We will always use the first certificate inside of the root bundle as the active one

func ensureCertKeyMatch(cert *x509.Certificate, key crypto.PublicKey) error {
	_ = "STUB: not implemented"
	return nil
}

// GetLocalRootCA validates if the contents of the file are a valid self-signed
// CA certificate, and returns the PEM-encoded Certificate if so
func GetLocalRootCA(paths CertPaths) (RootCA, error) {
	_ = "STUB: not implemented"
	// Check if we have a Certificate file
	return *new(RootCA), nil
}

// There may not be a local key. It's okay to pass in a nil
// key. We'll get a root CA without a signer.

func getGRPCConnection(creds credentials.TransportCredentials, connBroker *connectionbroker.Broker, forceRemote bool) (*connectionbroker.Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetRemoteCA returns the remote endpoint's CA certificate bundle
func GetRemoteCA(ctx context.Context, d digest.Digest, connBroker *connectionbroker.Broker) (RootCA, error) {
	_ = "STUB: not implemented"
	// This TLS Config is intentionally using InsecureSkipVerify. We use the
	// digest instead to check the integrity of the CA certificate.
	return *new(RootCA), nil
}

// If a bundle of certificates are provided, the digest covers the entire bundle and not just
// one of the certificates in the bundle.  Otherwise, a node can be MITMed while joining if
// the MITM CA provides a single certificate which matches the digest, and providing arbitrary
// other non-verified root certs that the manager certificate actually chains up to.

// NewRootCA will validate that the certificates are otherwise valid and create a RootCA object.
// Since there is no key, the certificate expiry does not matter and will not be used.

// CreateRootCA creates a Certificate authority for a new Swarm Cluster, potentially
// overwriting any existing CAs.
func CreateRootCA(rootCN string) (RootCA, error) {
	_ = "STUB: not implemented"
	// Create a simple CSR for the CA using the default CA validator and policy
	return *new(RootCA), nil
}

// Generate the CA and get the certificate and private key

// GetRemoteSignedCertificate submits a CSR to a remote CA server address,
// and that is part of a CA identified by a specific certificate pool.
func GetRemoteSignedCertificate(ctx context.Context, csr []byte, rootCAPool *x509.CertPool, config CertificateRequestConfig) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is our only non-MTLS request, and it happens when we are boostraping our TLS certs
// We're using CARole as server name, so an external CA doesn't also have to have ManagerRole in the cert SANs

// Create a CAClient to retrieve a new Certificate

// Send the Request and retrieve the request token

// Exponential backoff with Max of 30 seconds to wait for a new retry

// Send the Request and retrieve the certificate

// Because IssueNodeCertificate succeeded, if this call failed likely it is due to an issue with this
// particular connection, so we need to get another.  We should try a remote connection - the local node
// may be a manager that was demoted, so the local connection (which is preferred) may not work.

// If there was no deadline exceeded error, and the certificate was issued, return

// The certificate in the response must match the CSR
// we submitted. If we are getting a response for a
// certificate that was previously issued, we need to
// retry until the certificate gets updated per our
// current request.

// If NodeCertificateStatus timed out, we're still pending, the issuance failed, or
// the state is unknown let's continue trying after an exponential backoff

// readCertValidity returns the certificate issue and expiration time
func readCertValidity(kr KeyReader) (time.Time, time.Time, error) {
	_ = "STUB: not implemented"
	return *

	// Read the Cert
	new(time.Time), *new(time.Time), nil
}

// Create an x509 certificate out of the contents on disk

// SaveRootCA saves a RootCA object to disk
func SaveRootCA(rootCA RootCA, paths CertPaths) error {
	_ = "STUB: not implemented"
	// Make sure the necessary dirs exist and they are writable
	return nil
}

// If the root certificate got returned successfully, save the rootCA to disk.

// GenerateNewCSR returns a newly generated key and CSR signed with said key
func GenerateNewCSR() ([]byte, []byte, error) { _ = "STUB: not implemented"; return nil, nil, nil }

// NormalizePEMs takes a bundle of PEM-encoded certificates in a certificate bundle,
// decodes them, removes headers, and re-encodes them to make sure that they have
// consistent whitespace.  Note that this is intended to normalize x509 certificates
// in PEM format, hence the stripping out of headers.
func NormalizePEMs(certs []byte) []byte { _ = "STUB: not implemented"; return nil }
