package ca

import (
	"context"
	"crypto/tls"
	"crypto/x509/pkix"
	"net"
	"sync"

	"google.golang.org/grpc/credentials"
)

var (
	// alpnProtoStr is the specified application level protocols for gRPC.
	alpnProtoStr = []string{"h2"}
)

// MutableTLSCreds is the credentials required for authenticating a connection using TLS.
type MutableTLSCreds struct {
	// Mutex for the tls config
	sync.Mutex
	// TLS configuration
	config *tls.Config
	// TLS Credentials
	tlsCreds credentials.TransportCredentials
	// store the subject for easy access
	subject pkix.Name
}

// Info implements the credentials.TransportCredentials interface
func (c *MutableTLSCreds) Info() credentials.ProtocolInfo {
	_ = "STUB: not implemented"
	return *new(credentials.ProtocolInfo)
}

// Clone returns new MutableTLSCreds created from underlying *tls.Config.
// It panics if validation of underlying config fails.
func (c *MutableTLSCreds) Clone() credentials.TransportCredentials {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials)
}

// OverrideServerName overrides *tls.Config.ServerName.
func (c *MutableTLSCreds) OverrideServerName(name string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetRequestMetadata implements the credentials.TransportCredentials interface
func (c *MutableTLSCreds) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	_ = "STUB: not implemented"

	// RequireTransportSecurity implements the credentials.TransportCredentials interface
	return nil, nil
}

func (c *MutableTLSCreds) RequireTransportSecurity() bool {
	_ = "STUB: not implemented"

	// ClientHandshake implements the credentials.TransportCredentials interface
	return false
}

func (c *MutableTLSCreds) ClientHandshake(ctx context.Context, addr string, rawConn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	_ = "STUB: not implemented"
	// borrow all the code from the original TLS credentials
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

// Need to allow conn.Handshake to have access to config,
// would create a deadlock otherwise

// ServerHandshake implements the credentials.TransportCredentials interface
func (c *MutableTLSCreds) ServerHandshake(rawConn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), *new(credentials.AuthInfo), nil
}

// loadNewTLSConfig replaces the currently loaded TLS config with a new one
func (c *MutableTLSCreds) loadNewTLSConfig(newConfig *tls.Config) error {
	_ = "STUB: not implemented"
	return nil
}

// Config returns the current underlying TLS config.
func (c *MutableTLSCreds) Config() *tls.Config { _ = "STUB: not implemented"; return nil }

// Role returns the OU for the certificate encapsulated in this TransportCredentials
func (c *MutableTLSCreds) Role() string { _ = "STUB: not implemented"; return "" }

// Organization returns the O for the certificate encapsulated in this TransportCredentials
func (c *MutableTLSCreds) Organization() string { _ = "STUB: not implemented"; return "" }

// NodeID returns the CN for the certificate encapsulated in this TransportCredentials
func (c *MutableTLSCreds) NodeID() string { _ = "STUB: not implemented"; return "" }

// NewMutableTLS uses c to construct a mutable TransportCredentials based on TLS.
func NewMutableTLS(c *tls.Config) (*MutableTLSCreds, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAndValidateCertificateSubject is a helper method to retrieve and validate the subject
// from the x509 certificate underlying a tls.Certificate
func GetAndValidateCertificateSubject(certs []tls.Certificate) (pkix.Name, error) {
	_ = "STUB: not implemented"
	return *new(pkix.Name), nil
}
