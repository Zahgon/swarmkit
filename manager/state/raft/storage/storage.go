package storage

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	"go.etcd.io/raft/v3/raftpb"
)

// ErrNoWAL is returned if there are no WALs on disk
var ErrNoWAL = errors.New("no WAL present")

type walSnapDirs struct {
	wal  string
	snap string
}

// the wal/snap directories in decreasing order of preference/version
var versionedWALSnapDirs = []walSnapDirs{
	{wal: "wal-v3-encrypted", snap: "snap-v3-encrypted"},
	{wal: "wal-v3", snap: "snap-v3"},
	{wal: "wal", snap: "snap"},
}

// EncryptedRaftLogger saves raft data to disk
type EncryptedRaftLogger struct {
	StateDir      string
	EncryptionKey []byte

	// FIPS specifies whether the encryption should be FIPS-compliant
	FIPS bool

	// mutex is locked for writing only when we need to replace the wal object and snapshotter
	// object, not when we're writing snapshots or wals (in which case it's locked for reading)
	encoderMu   sync.RWMutex
	wal         WAL
	snapshotter Snapshotter
}

// BootstrapFromDisk creates a new snapshotter and wal, and also reads the latest snapshot and WALs from disk
func (e *EncryptedRaftLogger) BootstrapFromDisk(ctx context.Context, oldEncryptionKeys ...[]byte) (*raftpb.Snapshot, WALData, error) {
	_ = "STUB: not implemented"
	return nil, *new(WALData), nil
}

// If snapshots created by the etcd-v2 code exist, or by swarmkit development version,
// read the latest snapshot and write it encoded to the new path.  The new path
// prevents etc-v2 creating snapshots that are visible to us, but not encoded and
// out of sync with our WALs, after a downgrade.

// ensure the new directory exists

// Create a snapshotter and load snapshot data

// If wals created by the etcd-v2 wal code exist, read the latest ones based
// on this snapshot and encode them to wals in the new path to avoid adding
// backwards-incompatible entries to those files.

// BootstrapNew creates a new snapshotter and WAL writer, expecting that there is nothing on disk
func (e *EncryptedRaftLogger) BootstrapNew(metadata []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// the wal directory must not already exist upon creation

func (e *EncryptedRaftLogger) walDir() string { _ = "STUB: not implemented"; return "" }

func (e *EncryptedRaftLogger) snapDir() string { _ = "STUB: not implemented"; return "" }

// RotateEncryptionKey swaps out the encoders and decoders used by the wal and snapshotter
func (e *EncryptedRaftLogger) RotateEncryptionKey(newKey []byte) { _ = "STUB: not implemented"; return }

// if the wal exists, the snapshotter exists
// We don't want to have to close the WAL, because we can't open a new one.
// We need to know the previous snapshot, because when you open a WAL you
// have to read out all the entries from a particular snapshot, or you can't
// write.  So just rotate the encoders out from under it.  We already
// have a lock on writing to snapshots and WALs.

// SaveSnapshot actually saves a given snapshot to both the WAL and the snapshot.
func (e *EncryptedRaftLogger) SaveSnapshot(snapshot raftpb.Snapshot) error {
	_ = "STUB: not implemented"
	return nil
}

// GC garbage collects snapshots and wals older than the provided index and term
func (e *EncryptedRaftLogger) GC(index uint64, term uint64, keepOldSnapshots uint64) error {
	_ = "STUB: not implemented"
	// Delete any older snapshots
	return nil
}

// Ignore any snapshots that are older than the current snapshot.
// Delete the others. Rather than doing lexical comparisons, we look
// at what exists before/after the current snapshot in the slice.
// This means that if the current snapshot doesn't appear in the
// directory for some strange reason, we won't delete anything, which
// is the safe behavior.

// Remove any WAL files that only contain data from before the oldest
// remaining snapshot.

// Parse index out of oldest snapshot's filename

// If all WAL files started with indices below the oldest snapshot's
// index, we can delete all but the newest WAL file.

// SaveEntries saves only entries to disk
func (e *EncryptedRaftLogger) SaveEntries(st raftpb.HardState, entries []raftpb.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Close closes the logger - it will have to be bootstrapped again to start writing
func (e *EncryptedRaftLogger) Close(ctx context.Context) { _ = "STUB: not implemented"; return }

// Clear closes the existing WAL and removes the WAL and snapshot.
func (e *EncryptedRaftLogger) Clear(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
