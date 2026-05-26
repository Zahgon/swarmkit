package storage

import (
	"github.com/moby/swarmkit/v2/manager/encryption"
	"go.etcd.io/etcd/server/v3/etcdserver/api/snap"
	"go.etcd.io/raft/v3/raftpb"
)

// This package wraps the go.etcd.io/etcd/server/v3/api/snap package, and encrypts
// the bytes of whatever snapshot is passed to it, and decrypts the bytes of
// whatever snapshot it reads.

// Snapshotter is the interface presented by go.etcd.io/etcd/server/v3/api/snap.Snapshotter that we depend upon
type Snapshotter interface {
	SaveSnap(snapshot raftpb.Snapshot) error
	Load() (*raftpb.Snapshot, error)
}

// SnapFactory provides an interface for the different ways to get a Snapshotter object.
// For instance, the etcd/snap package itself provides this
type SnapFactory interface {
	New(dirpath string) Snapshotter
}

var _ Snapshotter = &wrappedSnap{}
var _ Snapshotter = &snap.Snapshotter{}
var _ SnapFactory = snapCryptor{}

// wrappedSnap wraps a go.etcd.io/etcd/server/v3/api/snap.Snapshotter, and handles
// encrypting/decrypting.
type wrappedSnap struct {
	*snap.Snapshotter
	encrypter encryption.Encrypter
	decrypter encryption.Decrypter
}

// SaveSnap encrypts the snapshot data (if an encrypter is exists) before passing it onto the
// wrapped snap.Snapshotter's SaveSnap function.
func (s *wrappedSnap) SaveSnap(snapshot raftpb.Snapshot) error {
	_ = "STUB: not implemented"
	return nil
}

// Load decrypts the snapshot data (if a decrypter is exists) after reading it using the
// wrapped snap.Snapshotter's Load function.
func (s *wrappedSnap) Load() (*raftpb.Snapshot, error) { _ = "STUB: not implemented"; return nil, nil }

// snapCryptor is an object that provides the same functions as `etcd/wal`
// and `etcd/snap` that we need to open a WAL object or Snapshotter object
type snapCryptor struct {
	encrypter encryption.Encrypter
	decrypter encryption.Decrypter
}

// NewSnapFactory returns a new object that can read from and write to encrypted
// snapshots on disk
func NewSnapFactory(encrypter encryption.Encrypter, decrypter encryption.Decrypter) SnapFactory {
	_ = "STUB: not implemented"
	return *new(SnapFactory)
}

// NewSnapshotter returns a new Snapshotter with the given encrypters and decrypters
func (sc snapCryptor) New(dirpath string) Snapshotter {
	_ = "STUB: not implemented"
	return *new(Snapshotter)
}

type originalSnap struct{}

func (o originalSnap) New(dirpath string) Snapshotter {
	_ = "STUB: not implemented"
	return *new(Snapshotter)
}

// OriginalSnap is the original `snap` package as an implementation of the SnapFactory interface
var OriginalSnap SnapFactory = originalSnap{}

// MigrateSnapshot reads the latest existing snapshot from one directory, encoded one way, and writes
// it to a new directory, encoded a different way
func MigrateSnapshot(oldDir, newDir string, oldFactory, newFactory SnapFactory) error {
	_ = "STUB: not implemented"
	// use temporary snapshot directory so initialization appears atomic
	return nil
}

// if there's no snapshot, the migration succeeded

// write the new snapshot to the temporary location

// ListSnapshots lists all the snapshot files in a particular directory and returns
// the snapshot files in reverse lexical order (newest first)
func ListSnapshots(dirpath string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Sort snapshot filenames in reverse lexical order
