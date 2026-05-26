package networkallocator

import (
	"errors"

	"github.com/moby/swarmkit/v2/api"
)

// InertProvider is a network allocator [Provider] which does not allocate networks.
type InertProvider struct{}

var _ Provider = InertProvider{}

// NewAllocator returns an instance of [Inert].
func (InertProvider) NewAllocator(*Config) (NetworkAllocator, error) {
	_ = "STUB: not implemented"
	return *

	// PredefinedNetworks returns a nil slice.
	new(NetworkAllocator), nil
}

func (InertProvider) PredefinedNetworks() []PredefinedNetworkData {
	_ = "STUB: not implemented"

	// SetDefaultVXLANUDPPort is a no-op.
	return nil
}

func (InertProvider) SetDefaultVXLANUDPPort(uint32) error {
	_ = "STUB: not implemented"

	// ValidateIPAMDriver returns an InvalidArgument error unless d is nil.
	return nil
}

func (InertProvider) ValidateIPAMDriver(d *api.Driver) error { _ = "STUB: not implemented"; return nil }

// ValidateIngressNetworkDriver returns an InvalidArgument error unless d is nil.
func (InertProvider) ValidateIngressNetworkDriver(d *api.Driver) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateNetworkDriver returns an InvalidArgument error unless d is nil.
func (InertProvider) ValidateNetworkDriver(d *api.Driver) error {
	_ = "STUB: not implemented"
	return nil
}

// Inert is a [NetworkAllocator] which does not allocate networks.
type Inert struct{}

var _ NetworkAllocator = Inert{}

var errUnavailable = errors.New("network support is unavailable")

// Allocate returns an error unless n.Spec.Ingress is true.
func (Inert) Allocate(n *api.Network) error { _ = "STUB: not implemented"; return nil }

// AllocateAttachment unconditionally returns an error.
func (Inert) AllocateAttachment(_ *api.Node, _ *api.NetworkAttachment) error {
	_ = "STUB: not implemented"
	return nil

	// AllocateService succeeds iff the service specifies no network attachments.
}

func (Inert) AllocateService(s *api.Service) error { _ = "STUB: not implemented"; return nil }

// AllocateTask succeeds iff the task specifies no network attachments.
func (Inert) AllocateTask(t *api.Task) error { _ = "STUB: not implemented"; return nil }

// Deallocate does nothing, successfully.
func (Inert) Deallocate(_ *api.Network) error {
	_ = "STUB: not implemented"

	// DeallocateAttachment does nothing, successfully.
	return nil
}

func (Inert) DeallocateAttachment(_ *api.Node, _ *api.NetworkAttachment) error {
	_ = "STUB: not implemented"

	// DeallocateService does nothing, successfully.
	return nil
}

func (Inert) DeallocateService(_ *api.Service) error {
	_ = "STUB: not implemented"

	// DeallocateTask does nothing, successfully.
	return nil
}

func (Inert) DeallocateTask(_ *api.Task) error {
	_ = "STUB: not implemented"

	// IsAllocated returns true iff [Inert.Allocate] would return nil.
	return nil
}

func (Inert) IsAllocated(n *api.Network) bool { _ = "STUB: not implemented"; return false }

// IsAttachmentAllocated returns false.
func (Inert) IsAttachmentAllocated(_ *api.Node, _ *api.NetworkAttachment) bool {
	_ = "STUB: not implemented"

	// IsServiceAllocated returns true iff [Inert.AllocateService] would return nil.
	return false
}

func (Inert) IsServiceAllocated(s *api.Service, _ ...func(*ServiceAllocationOpts)) bool {
	_ = "STUB: not implemented"
	return false
}

// IsTaskAllocated returns true iff [Inert.AllocateTask] would return nil.
func (Inert) IsTaskAllocated(t *api.Task) bool { _ = "STUB: not implemented"; return false }
