package store

import (
	memdb "github.com/hashicorp/go-memdb"
	"github.com/moby/swarmkit/v2/api"
)

// ObjectStoreConfig provides the necessary methods to store a particular object
// type inside MemoryStore.
type ObjectStoreConfig struct {
	Table            *memdb.TableSchema
	Save             func(ReadTx, *api.StoreSnapshot) error
	Restore          func(Tx, *api.StoreSnapshot) error
	ApplyStoreAction func(Tx, api.StoreAction) error
}

// RestoreTable takes a list of new objects of a particular type (e.g. clusters,
// nodes, etc., which conform to the StoreObject interface) and replaces the
// existing objects in the store of that type with the new objects.
func RestoreTable(tx Tx, table string, newObjects []api.StoreObject) error {
	_ = "STUB: not implemented"
	return nil
}
