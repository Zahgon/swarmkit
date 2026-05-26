package ca

import (
	"context"
)

const (
	certForwardedKey = "forwarded_cert"
	certCNKey        = "forwarded_cert_cn"
	certOUKey        = "forwarded_cert_ou"
	certOrgKey       = "forwarded_cert_org"
	remoteAddrKey    = "remote_addr"
)

// forwardedTLSInfoFromContext obtains forwarded TLS CN/OU from the grpc.MD
// object in ctx.
func forwardedTLSInfoFromContext(ctx context.Context) (remoteAddr string, cn string, org string, ous []string) {
	_ = "STUB: not implemented"
	return "", "", "", nil
}

func isForwardedRequest(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// WithMetadataForwardTLSInfo reads certificate from context and returns context where
// ForwardCert is set based on original certificate.
func WithMetadataForwardTLSInfo(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// If there's no TLS cert, forward with blank TLS metadata.
// Note that the presence of this blank metadata is extremely
// important. Without it, it would look like manager is making
// the request directly.
