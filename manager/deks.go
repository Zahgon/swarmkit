package manager

import (
	"fmt"

	"github.com/moby/swarmkit/v2/ca"
	"github.com/moby/swarmkit/v2/manager/state/raft"
)

// This module contains the data structures and control flow to manage rotating the raft
// DEK and also for interacting with KeyReadWriter to maintain the raft DEK information in
// the PEM headers fo the TLS key for the node.

const (
	// the raft DEK (data encryption key) is stored in the TLS key as a header
	// these are the header values
	pemHeaderRaftDEK              = "raft-dek"
	pemHeaderRaftPendingDEK       = "raft-dek-pending"
	pemHeaderRaftDEKNeedsRotation = "raft-dek-needs-rotation"
)

// RaftDEKData contains all the data stored in TLS pem headers.
type RaftDEKData struct {

	// EncryptionKeys contain the current and pending raft DEKs
	raft.EncryptionKeys

	// NeedsRotation indicates whether another rotation needs to be happen after
	// the current one.
	NeedsRotation bool

	// The FIPS boolean is not serialized, but is internal state which indicates how
	// the raft DEK headers should be encrypted (e.g. using FIPS compliant algorithms)
	FIPS bool
}

// RaftDEKData should implement the PEMKeyHeaders interface
var _ ca.PEMKeyHeaders = RaftDEKData{}

// UnmarshalHeaders loads the current state of the DEKs into a new RaftDEKData object (which is returned) given the
// current TLS headers and the current KEK.
func (r RaftDEKData) UnmarshalHeaders(headers map[string]string, kekData ca.KEKData) (ca.PEMKeyHeaders, error) {
	_ = "STUB: not implemented"
	return *new(ca.PEMKeyHeaders), nil
}

// MarshalHeaders returns new PEM headers given the current KEK - it uses the current KEK to
// serialize/encrypt the current DEK state that is maintained in the current RaftDEKData object.
func (r RaftDEKData) MarshalHeaders(kekData ca.KEKData) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// return a function that updates the dek data on write success

// UpdateKEK sets NeedRotation to true if we go from unlocked to locked.
func (r RaftDEKData) UpdateKEK(oldKEK, candidateKEK ca.KEKData) ca.PEMKeyHeaders {
	_ = "STUB: not implemented"
	return *new(ca.PEMKeyHeaders)
}

// Returns whether the old KEK should be replaced with the new KEK, whether we went from
// unlocked to locked, and whether there was an error (the versions are the same, but the
// keks are different)
func compareKEKs(oldKEK, candidateKEK ca.KEKData) (bool, bool, error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

// RaftDEKManager manages the raft DEK keys by interacting with KeyReadWriter, calling the necessary functions
// to update the TLS headers when the raft DEK needs to change, or to re-encrypt everything when the KEK changes.
type RaftDEKManager struct {
	kw         ca.KeyWriter
	rotationCh chan struct{}
	FIPS       bool
}

var errNoUpdateNeeded = fmt.Errorf("don't need to rotate or update")

// this error is returned if the KeyReadWriter's PEMKeyHeaders object is no longer a RaftDEKData object -
// this can happen if the node is no longer a manager, for example
var errNotUsingRaftDEKData = fmt.Errorf("RaftDEKManager can no longer store and manage TLS key headers")

// NewRaftDEKManager returns a RaftDEKManager that uses the current key writer
// and header manager
func NewRaftDEKManager(kw ca.KeyWriter, fips bool) (*RaftDEKManager, error) {
	_ = "STUB: not implemented"
	// If there is no current DEK, generate one and write it to disk
	return nil, nil
}

// it wasn't a raft DEK manager before - just replace it

// NeedsRotation returns a boolean about whether we should do a rotation
func (r *RaftDEKManager) NeedsRotation() bool { _ = "STUB: not implemented"; return false }

// GetKeys returns the current set of DEKs.  If NeedsRotation is true, and there
// is no existing PendingDEK, it will try to create one.  If it successfully creates
// and writes a PendingDEK, it sets NeedRotation to false.  If there are any errors
// doing so, just return the original set of keys.
func (r *RaftDEKManager) GetKeys() raft.EncryptionKeys {
	_ = "STUB: not implemented"
	return *new(raft.EncryptionKeys)
}

// RotationNotify the channel used to notify subscribers as to whether there
// should be a rotation done
func (r *RaftDEKManager) RotationNotify() chan struct{} { _ = "STUB: not implemented"; return nil }

// UpdateKeys will set the updated encryption keys in the headers.  This finishes
// a rotation, and is expected to set the CurrentDEK to the previous PendingDEK.
func (r *RaftDEKManager) UpdateKeys(newKeys raft.EncryptionKeys) error {
	_ = "STUB: not implemented"
	return nil
}

// If there is no current DEK, we are basically wiping out all DEKs (no header object)

// MaybeUpdateKEK does a KEK rotation if one is required.  Returns whether
// the kek was updated, whether it went from unlocked to locked, and any errors.
func (r *RaftDEKManager) MaybeUpdateKEK(candidateKEK ca.KEKData) (bool, bool, error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

// if we don't need to rotate the KEK, don't bother updating

func decodePEMHeaderValue(headerValue string, kek []byte, fips bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodePEMHeaderValue(headerValue []byte, kek []byte, fips bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
