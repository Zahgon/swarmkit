package testutils

import (
	"context"
	"crypto"
	"testing"
	"time"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/ca"
	"github.com/moby/swarmkit/v2/connectionbroker"
	"github.com/moby/swarmkit/v2/manager/state/store"
	"google.golang.org/grpc"
)

// TestCA is a structure that encapsulates everything needed to test a CA Server
type TestCA struct {
	RootCA                      ca.RootCA
	ExternalSigningServer       *ExternalSigningServer
	MemoryStore                 *store.MemoryStore
	Addr, TempDir, Organization string
	Paths                       *ca.SecurityConfigPaths
	Server                      *grpc.Server
	ServingSecurityConfig       *ca.SecurityConfig
	CAServer                    *ca.Server
	Context                     context.Context
	NodeCAClients               []api.NodeCAClient
	CAClients                   []api.CAClient
	Conns                       []*grpc.ClientConn
	WorkerToken                 string
	ManagerToken                string
	ConnBroker                  *connectionbroker.Broker
	KeyReadWriter               *ca.KeyReadWriter
	ctxCancel                   func()
	securityConfigCleanups      []func() error
}

// Stop cleans up after TestCA
func (tc *TestCA) Stop() { _ = "STUB: not implemented"; return }

// NewNodeConfig returns security config for a new node, given a role
func (tc *TestCA) NewNodeConfig(role string) (*ca.SecurityConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WriteNewNodeConfig returns security config for a new node, given a role
// saving the generated key and certificates to disk
func (tc *TestCA) WriteNewNodeConfig(role string) (*ca.SecurityConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewNodeConfigOrg returns security config for a new node, given a role and an org
func (tc *TestCA) NewNodeConfigOrg(role, org string) (*ca.SecurityConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// External controls whether or not NewTestCA() will create a TestCA server
// configured to use an external signer or not.
var External bool

// NewTestCA is a helper method that creates a TestCA and a bunch of default
// connections and security configs.
func NewTestCA(t *testing.T, krwGenerators ...func(ca.CertPaths) *ca.KeyReadWriter) *TestCA {
	_ = "STUB: not implemented"
	return nil
}

// NewFIPSTestCA is a helper method that creates a mandatory fips TestCA and a bunch of default
// connections and security configs.
func NewFIPSTestCA(t *testing.T) *TestCA { _ = "STUB: not implemented"; return nil }

// NewTestCAFromAPIRootCA is a helper method that creates a TestCA and a bunch of default
// connections and security configs, given a temp directory and an api.RootCA to use for creating
// a cluster and for signing.
func NewTestCAFromAPIRootCA(t *testing.T, tempBaseDir string, apiRootCA api.RootCA, krwGenerators []func(ca.CertPaths) *ca.KeyReadWriter) *TestCA {
	_ = "STUB: not implemented"
	return nil
}

func newTestCA(t *testing.T, tempBaseDir string, apiRootCA api.RootCA, krwGenerators []func(ca.CertPaths) *ca.KeyReadWriter, fips bool) *TestCA {
	_ = "STUB: not implemented"
	return nil
}

// Write the root certificate to disk, using decent permissions

// Start the CA API server - ensure that the external server doesn't have any intermediates

// remove the key from the API root CA so that once the CA server starts up, it won't have a local signer

// remove the key from the API root CA so that once the CA server starts up, it won't have a local signer

// Wait for caServer to be ready to serve

func createNode(s *store.MemoryStore, nodeID, role string, csr, cert []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func genSecurityConfig(s *store.MemoryStore, rootCA ca.RootCA, krw *ca.KeyReadWriter, role, org, tmpDir string, nonSigningRoot bool) (*ca.SecurityConfig, func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Obtain a signed Certificate

// If we were instructed to persist the files

// Load a valid tls.Certificate from the chain and the key

func createClusterObject(t *testing.T, s *store.MemoryStore, clusterID string, apiRootCA api.RootCA, caRootCA *ca.RootCA, externalCAs ...*api.ExternalCA) *api.Cluster {
	_ = "STUB: not implemented"
	return nil
}

// CreateRootCertAndKey returns a generated certificate and key for a root CA
func CreateRootCertAndKey(rootCN string) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	// Create a simple CSR for the CA using the default CA validator and policy
	return nil, nil, nil
}

// Generate the CA and get the certificate and private key

// ReDateCert takes an existing cert and changes the not before and not after date, to make it easier
// to test expiry
func ReDateCert(t *testing.T, cert, signerCert, signerKey []byte, notBefore, notAfter time.Time) []byte {
	_ = "STUB: not implemented"
	return nil
}

// CreateCertFromSigner creates a Certificate authority for a new Swarm Cluster given an existing key only.
func CreateCertFromSigner(rootCN string, priv crypto.Signer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
