package store

import (
	memdb "github.com/hashicorp/go-memdb"
	"github.com/moby/swarmkit/v2/api"
	"github.com/pkg/errors"
)

const tableResource = "resource"

var (
	// ErrNoKind is returned by resource create operations if the provided Kind
	// of the resource does not exist
	ErrNoKind = errors.New("object kind is unregistered")
)

func init() {
	register(ObjectStoreConfig{
		Table: &memdb.TableSchema{
			Name: tableResource,
			Indexes: map[string]*memdb.IndexSchema{
				indexID: {
					Name:    indexID,
					Unique:  true,
					Indexer: resourceIndexerByID{},
				},
				indexName: {
					Name:    indexName,
					Unique:  true,
					Indexer: resourceIndexerByName{},
				},
				indexKind: {
					Name:    indexKind,
					Indexer: resourceIndexerByKind{},
				},
				indexCustom: {
					Name:         indexCustom,
					Indexer:      resourceCustomIndexer{},
					AllowMissing: true,
				},
			},
		},
		Save: func(tx ReadTx, snapshot *api.StoreSnapshot) error {
			var err error
			snapshot.Resources, err = FindResources(tx, All)
			return err
		},
		Restore: func(tx Tx, snapshot *api.StoreSnapshot) error {
			toStoreObj := make([]api.StoreObject, len(snapshot.Resources))
			for i, x := range snapshot.Resources {
				toStoreObj[i] = resourceEntry{x}
			}
			return RestoreTable(tx, tableResource, toStoreObj)
		},
		ApplyStoreAction: func(tx Tx, sa api.StoreAction) error {
			switch v := sa.Target.(type) {
			case *api.StoreAction_Resource:
				obj := v.Resource
				switch sa.Action {
				case api.StoreActionKindCreate:
					return CreateResource(tx, obj)
				case api.StoreActionKindUpdate:
					return UpdateResource(tx, obj)
				case api.StoreActionKindRemove:
					return DeleteResource(tx, obj.ID)
				}
			}
			return errUnknownStoreAction
		},
	})
}

type resourceEntry struct {
	*api.Resource
}

func (r resourceEntry) CopyStoreObject() api.StoreObject {
	_ = "STUB: not implemented"
	return *new(api.StoreObject)
}

// ensure that when update events are emitted, we unwrap resourceEntry
func (r resourceEntry) EventUpdate(oldObject api.StoreObject) api.Event {
	_ = "STUB: not implemented"
	return *new(api.Event)
}

func confirmExtension(tx Tx, r *api.Resource) error {
	_ = "STUB: not implemented"
	// There must be an extension corresponding to the Kind field.
	return nil
}

// CreateResource adds a new resource object to the store.
// Returns ErrExist if the ID is already taken.
// Returns ErrNameConflict if a Resource with this Name already exists
// Returns ErrNoKind if the specified Kind does not exist
func CreateResource(tx Tx, r *api.Resource) error { _ = "STUB: not implemented"; return nil }

// TODO(dperny): currently the "name" index is unique, which means only one
// Resource of _any_ Kind can exist with that name. This isn't a problem
// right now, but the ideal case would be for names to be namespaced to the
// kind.

// UpdateResource updates an existing resource object in the store.
// Returns ErrNotExist if the object doesn't exist.
func UpdateResource(tx Tx, r *api.Resource) error { _ = "STUB: not implemented"; return nil }

// DeleteResource removes a resource object from the store.
// Returns ErrNotExist if the object doesn't exist.
func DeleteResource(tx Tx, id string) error { _ = "STUB: not implemented"; return nil }

// GetResource looks up a resource object by ID.
// Returns nil if the object doesn't exist.
func GetResource(tx ReadTx, id string) *api.Resource { _ = "STUB: not implemented"; return nil }

// FindResources selects a set of resource objects and returns them.
func FindResources(tx ReadTx, by By) ([]*api.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type resourceIndexerByKind struct{}

func (ri resourceIndexerByKind) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ri resourceIndexerByKind) FromObject(obj interface{}) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false,

		// Add the null character as a terminator
		nil, nil
}

type resourceIndexerByID struct{}

func (indexer resourceIndexerByID) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (indexer resourceIndexerByID) PrefixFromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (indexer resourceIndexerByID) FromObject(obj interface{}) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

type resourceIndexerByName struct{}

func (indexer resourceIndexerByName) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (indexer resourceIndexerByName) PrefixFromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (indexer resourceIndexerByName) FromObject(obj interface{}) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

type resourceCustomIndexer struct{}

func (indexer resourceCustomIndexer) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (indexer resourceCustomIndexer) PrefixFromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (indexer resourceCustomIndexer) FromObject(obj interface{}) (bool, [][]byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}
