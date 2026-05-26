package main

import (
	"github.com/moby/swarmkit/v2/ca"
	"github.com/moby/swarmkit/v2/manager"
)

func certPaths(swarmdir string) *ca.SecurityConfigPaths { _ = "STUB: not implemented"; return nil }

func getDEKData(krw *ca.KeyReadWriter) (manager.RaftDEKData, error) {
	_ = "STUB: not implemented"
	return *new(manager.RaftDEKData), nil
}

func getKRW(swarmdir, unlockKey string) (*ca.KeyReadWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// loads all the key data into the KRW object

func moveDirAside(dirname string) error { _ = "STUB: not implemented"; return nil }

func decryptRaftData(swarmdir, outdir, unlockKey string) error {
	_ = "STUB: not implemented"
	return nil
}

// always use false for FIPS, since we want to be able to decrypt logs written using
// any algorithm (not just FIPS-compatible ones)

func downgradeKey(swarmdir, unlockKey string) error { _ = "STUB: not implemented"; return nil }
