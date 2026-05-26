package storage

import (
	"context"

	"github.com/moby/swarmkit/v2/manager/encryption"
	"go.etcd.io/etcd/server/v3/storage/wal"
	"go.etcd.io/etcd/server/v3/storage/wal/walpb"
	"go.etcd.io/raft/v3/raftpb"
)

// This package wraps the go.etcd.io/etcd/server/v3/storage/wal package, and encrypts
// the bytes of whatever entry is passed to it, and decrypts the bytes of
// whatever entry it reads.

// WAL is the interface presented by go.etcd.io/etcd/server/v3/storage/wal.WAL that we depend upon
type WAL interface {
	ReadAll() ([]byte, raftpb.HardState, []raftpb.Entry, error)
	ReleaseLockTo(index uint64) error
	Close() error
	Save(st raftpb.HardState, ents []raftpb.Entry) error
	SaveSnapshot(e walpb.Snapshot) error
}

// WALFactory provides an interface for the different ways to get a WAL object.
// For instance, the etcd/wal package itself provides this
type WALFactory interface {
	Create(dirpath string, metadata []byte) (WAL, error)
	Open(dirpath string, walsnap walpb.Snapshot) (WAL, error)
}

var _ WAL = &wrappedWAL{}
var _ WAL = &wal.WAL{}
var _ WALFactory = walCryptor{}

// wrappedWAL wraps a go.etcd.io/etcd/server/v3/storage/wal.WAL, and handles encrypting/decrypting
type wrappedWAL struct {
	*wal.WAL
	encrypter encryption.Encrypter
	decrypter encryption.Decrypter
}

// ReadAll wraps the wal.WAL.ReadAll() function, but it first checks to see if the
// metadata indicates that the entries are encryptd, and if so, decrypts them.
func (w *wrappedWAL) ReadAll() ([]byte, raftpb.HardState, []raftpb.Entry, error) {
	_ = "STUB: not implemented"
	return nil, *new(raftpb.HardState), nil, nil
}

// Save encrypts the entry data (if an encrypter is exists) before passing it onto the
// wrapped wal.WAL's Save function.
func (w *wrappedWAL) Save(st raftpb.HardState, ents []raftpb.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// walCryptor is an object that provides the same functions as `etcd/wal`
// and `etcd/snap` that we need to open a WAL object or Snapshotter object
type walCryptor struct {
	encrypter encryption.Encrypter
	decrypter encryption.Decrypter
}

// NewWALFactory returns an object that can be used to produce objects that
// will read from and write to encrypted WALs on disk.
func NewWALFactory(encrypter encryption.Encrypter, decrypter encryption.Decrypter) WALFactory {
	_ = "STUB: not implemented"
	return *new(WALFactory)
}

// Create returns a new WAL object with the given encrypters and decrypters.
func (wc walCryptor) Create(dirpath string, metadata []byte) (WAL, error) {
	_ = "STUB: not implemented"
	return *new(WAL), nil
}

// Open returns a new WAL object with the given encrypters and decrypters.
func (wc walCryptor) Open(dirpath string, snap walpb.Snapshot) (WAL, error) {
	_ = "STUB: not implemented"
	return *new(WAL), nil
}

type originalWAL struct{}

func (o originalWAL) Create(dirpath string, metadata []byte) (WAL, error) {
	_ = "STUB: not implemented"
	return *new(WAL), nil
}

func (o originalWAL) Open(dirpath string, walsnap walpb.Snapshot) (WAL, error) {
	_ = "STUB: not implemented"
	return *new(WAL), nil
}

// OriginalWAL is the original `wal` package as an implementation of the WALFactory interface
var OriginalWAL WALFactory = originalWAL{}

// WALData contains all the data returned by a WAL's ReadAll() function
// (metadata, hardwate, and entries)
type WALData struct {
	Metadata  []byte
	HardState raftpb.HardState
	Entries   []raftpb.Entry
}

// ReadRepairWAL opens a WAL for reading, and attempts to read it.  If we can't read it, attempts to repair
// and read again.
func ReadRepairWAL(
	ctx context.Context,
	walDir string,
	walsnap walpb.Snapshot,
	factory WALFactory,
) (WAL, WALData, error) {
	_ = "STUB: not implemented"
	return *new(WAL), *new(WALData), nil
}

// we can only repair ErrUnexpectedEOF and we never repair twice.

// TODO(thaJeztah): should ReadRepairWAL be updated to handle cases where
// some (last) of the files cannot be recovered? ("best effort" recovery?)
// Or should an informative error be produced to help the user (which could
// mean: remove the last file?). See TestReadRepairWAL for more details.

// MigrateWALs reads existing WALs (from a particular snapshot and beyond) from one directory, encoded one way,
// and writes them to a new directory, encoded a different way
func MigrateWALs(ctx context.Context, oldDir, newDir string, oldFactory, newFactory WALFactory, snapshot walpb.Snapshot) error {
	_ = "STUB: not implemented"
	return nil
}

// keep temporary wal directory so WAL initialization appears atomic

// ListWALs lists all the wals in a directory and returns the list in lexical
// order (oldest first)
func ListWALs(dirpath string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Sort WAL filenames in lexical order
