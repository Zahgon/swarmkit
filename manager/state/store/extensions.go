package store

import (
	memdb "github.com/hashicorp/go-memdb"
	"github.com/moby/swarmkit/v2/api"
)

const tableExtension = "extension"

func init() {
	register(ObjectStoreConfig{
		Table: &memdb.TableSchema{
			Name: tableExtension,
			Indexes: map[string]*memdb.IndexSchema{
				indexID: {
					Name:    indexID,
					Unique:  true,
					Indexer: extensionIndexerByID{},
				},
				indexName: {
					Name:    indexName,
					Unique:  true,
					Indexer: extensionIndexerByName{},
				},
				indexCustom: {
					Name:         indexCustom,
					Indexer:      extensionCustomIndexer{},
					AllowMissing: true,
				},
			},
		},
		Save: func(tx ReadTx, snapshot *api.StoreSnapshot) error {
			var err error
			snapshot.Extensions, err = FindExtensions(tx, All)
			return err
		},
		Restore: func(tx Tx, snapshot *api.StoreSnapshot) error {
			toStoreObj := make([]api.StoreObject, len(snapshot.Extensions))
			for i, x := range snapshot.Extensions {
				toStoreObj[i] = extensionEntry{x}
			}
			return RestoreTable(tx, tableExtension, toStoreObj)
		},
		ApplyStoreAction: func(tx Tx, sa api.StoreAction) error {
			switch v := sa.Target.(type) {
			case *api.StoreAction_Extension:
				obj := v.Extension
				switch sa.Action {
				case api.StoreActionKindCreate:
					return CreateExtension(tx, obj)
				case api.StoreActionKindUpdate:
					return UpdateExtension(tx, obj)
				case api.StoreActionKindRemove:
					return DeleteExtension(tx, obj.ID)
				}
			}
			return errUnknownStoreAction
		},
	})
}

type extensionEntry struct {
	*api.Extension
}

func (e extensionEntry) CopyStoreObject() api.StoreObject {
	_ = "STUB: not implemented"
	return *new(api.StoreObject)
}

// ensure that when update events are emitted, we unwrap extensionEntry
func (e extensionEntry) EventUpdate(oldObject api.StoreObject) api.Event {
	_ = "STUB: not implemented"
	return *new(api.Event)
}

// CreateExtension adds a new extension to the store.
// Returns ErrExist if the ID is already taken.
func CreateExtension(tx Tx, e *api.Extension) error {
	_ = "STUB: not implemented"
	// Ensure the name is not already in use.
	return nil
}

// It can't conflict with built-in kinds either.

// UpdateExtension updates an existing extension in the store.
// Returns ErrNotExist if the object doesn't exist.
func UpdateExtension(_ Tx, _ *api.Extension) error {
	_ = "STUB: not implemented"
	// TODO(aaronl): For the moment, extensions are immutable
	return nil
}

// DeleteExtension removes an extension from the store.
// Returns ErrNotExist if the object doesn't exist.
func DeleteExtension(tx Tx, id string) error { _ = "STUB: not implemented"; return nil }

// GetExtension looks up an extension by ID.
// Returns nil if the object doesn't exist.
func GetExtension(tx ReadTx, id string) *api.Extension { _ = "STUB: not implemented"; return nil }

// FindExtensions selects a set of extensions and returns them.
func FindExtensions(tx ReadTx, by By) ([]*api.Extension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type extensionIndexerByID struct{}

func (indexer extensionIndexerByID) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (indexer extensionIndexerByID) PrefixFromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (indexer extensionIndexerByID) FromObject(obj interface{}) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

type extensionIndexerByName struct{}

func (indexer extensionIndexerByName) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (indexer extensionIndexerByName) PrefixFromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (indexer extensionIndexerByName) FromObject(obj interface{}) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

type extensionCustomIndexer struct{}

func (indexer extensionCustomIndexer) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (indexer extensionCustomIndexer) PrefixFromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (indexer extensionCustomIndexer) FromObject(obj interface{}) (bool, [][]byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}
