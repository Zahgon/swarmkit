package main

import (
	"github.com/moby/swarmkit/v2/manager/state/raft/storage"
	"github.com/moby/swarmkit/v2/manager/state/store"
	"go.etcd.io/raft/v3/raftpb"
)

func loadData(swarmdir, unlockKey string) (*storage.WALData, *raftpb.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Encrypted WAL is present

// always set FIPS=false, because we want to decrypt logs stored using any
// algorithm, not just FIPS-compatible ones

// Try unencrypted WAL

func dumpWAL(swarmdir, unlockKey string, start, end uint64, redact bool) error {
	_ = "STUB: not implemented"
	return nil
}

// redact sensitive information

func dumpSnapshot(swarmdir, unlockKey string, redact bool) error {
	_ = "STUB: not implemented"
	return nil
}

// expunge everything that may have key material

// objSelector provides some criteria to select objects.
type objSelector struct {
	all  bool
	id   string
	name string
}

func bySelection(selector objSelector) store.By { _ = "STUB: not implemented"; return *new(store.By) }

// find nothing

func dumpObject(swarmdir, unlockKey, objType string, selector objSelector) error {
	_ = "STUB: not implemented"
	return nil
}
