package ca

import (
	"context"
	"sync"
	"time"

	"github.com/docker/go-events"
	"github.com/moby/swarmkit/v2/connectionbroker"
)

// RenewTLSExponentialBackoff sets the exponential backoff when trying to renew TLS certificates that have expired
var RenewTLSExponentialBackoff = events.ExponentialBackoffConfig{
	Base:   time.Second * 5,
	Factor: time.Second * 5,
	Max:    1 * time.Hour,
}

// TLSRenewer handles renewing TLS certificates, either automatically or upon
// request.
type TLSRenewer struct {
	mu           sync.Mutex
	s            *SecurityConfig
	connBroker   *connectionbroker.Broker
	renew        chan struct{}
	expectedRole string
	rootPaths    CertPaths
}

// NewTLSRenewer creates a new TLS renewer. It must be started with Start.
func NewTLSRenewer(s *SecurityConfig, connBroker *connectionbroker.Broker, rootPaths CertPaths) *TLSRenewer {
	_ = "STUB: not implemented"
	return nil
}

// SetExpectedRole sets the expected role. If a renewal is forced, and the role
// doesn't match this expectation, renewal will be retried with exponential
// backoff until it does match.
func (t *TLSRenewer) SetExpectedRole(role string) { _ = "STUB: not implemented"; return }

// Renew causes the TLSRenewer to renew the certificate (nearly) right away,
// instead of waiting for the next automatic renewal.
func (t *TLSRenewer) Renew() { _ = "STUB: not implemented"; return }

// Start will continuously monitor for the necessity of renewing the local certificates, either by
// issuing them locally if key-material is available, or requesting them from a remote CA.
func (t *TLSRenewer) Start(ctx context.Context) <-chan CertificateUpdate {
	_ = "STUB: not implemented"
	return nil
}

// Our starting default will be 5 minutes

// Since the expiration of the certificate is managed remotely we should update our
// retry timer on every iteration of this loop.
// Retrieve the current certificate expiration information.

// We failed to read the expiration, let's stick with the starting default

// If we have an expired certificate, try to renew immediately: the hope that this is a temporary clock skew, or
// we can issue our own TLS certs.

// retry immediately(ish) with exponential backoff

// A forced renewal was requested, but did not succeed yet.
// retry immediately(ish) with exponential backoff

// Random retry time between 50% and 80% of the total time to expiration

// Pause briefly before attempting the renewal,
// to give the CA a chance to reconcile the
// desired role.

// ignore errors - it will just try again later
