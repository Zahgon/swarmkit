package store

import (
	memdb "github.com/hashicorp/go-memdb"
	"github.com/moby/swarmkit/v2/api"
)

const tableService = "service"

func init() {
	register(ObjectStoreConfig{
		Table: &memdb.TableSchema{
			Name: tableService,
			Indexes: map[string]*memdb.IndexSchema{
				indexID: {
					Name:    indexID,
					Unique:  true,
					Indexer: api.ServiceIndexerByID{},
				},
				indexName: {
					Name:    indexName,
					Unique:  true,
					Indexer: api.ServiceIndexerByName{},
				},
				indexRuntime: {
					Name:         indexRuntime,
					AllowMissing: true,
					Indexer:      serviceIndexerByRuntime{},
				},
				indexNetwork: {
					Name:         indexNetwork,
					AllowMissing: true,
					Indexer:      serviceIndexerByNetwork{},
				},
				indexSecret: {
					Name:         indexSecret,
					AllowMissing: true,
					Indexer:      serviceIndexerBySecret{},
				},
				indexConfig: {
					Name:         indexConfig,
					AllowMissing: true,
					Indexer:      serviceIndexerByConfig{},
				},
				indexCustom: {
					Name:         indexCustom,
					Indexer:      api.ServiceCustomIndexer{},
					AllowMissing: true,
				},
			},
		},
		Save: func(tx ReadTx, snapshot *api.StoreSnapshot) error {
			var err error
			snapshot.Services, err = FindServices(tx, All)
			return err
		},
		Restore: func(tx Tx, snapshot *api.StoreSnapshot) error {
			toStoreObj := make([]api.StoreObject, len(snapshot.Services))
			for i, x := range snapshot.Services {
				toStoreObj[i] = x
			}
			return RestoreTable(tx, tableService, toStoreObj)
		},
		ApplyStoreAction: func(tx Tx, sa api.StoreAction) error {
			switch v := sa.Target.(type) {
			case *api.StoreAction_Service:
				obj := v.Service
				switch sa.Action {
				case api.StoreActionKindCreate:
					return CreateService(tx, obj)
				case api.StoreActionKindUpdate:
					return UpdateService(tx, obj)
				case api.StoreActionKindRemove:
					return DeleteService(tx, obj.ID)
				}
			}
			return errUnknownStoreAction
		},
	})
}

// CreateService adds a new service to the store.
// Returns ErrExist if the ID is already taken.
func CreateService(tx Tx, s *api.Service) error {
	_ = "STUB: not implemented"
	// Ensure the name is not already in use.
	return nil
}

// UpdateService updates an existing service in the store.
// Returns ErrNotExist if the service doesn't exist.
func UpdateService(tx Tx, s *api.Service) error {
	_ = "STUB: not implemented"
	// Ensure the name is either not in use or already used by this same Service.
	return nil
}

// DeleteService removes a service from the store.
// Returns ErrNotExist if the service doesn't exist.
func DeleteService(tx Tx, id string) error { _ = "STUB: not implemented"; return nil }

// GetService looks up a service by ID.
// Returns nil if the service doesn't exist.
func GetService(tx ReadTx, id string) *api.Service { _ = "STUB: not implemented"; return nil }

// FindServices selects a set of services and returns them.
func FindServices(tx ReadTx, by By) ([]*api.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type serviceIndexerByRuntime struct{}

func (si serviceIndexerByRuntime) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (si serviceIndexerByRuntime) FromObject(obj interface{}) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (si serviceIndexerByRuntime) PrefixFromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type serviceIndexerByNetwork struct{}

func (si serviceIndexerByNetwork) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (si serviceIndexerByNetwork) FromObject(obj interface{}) (bool, [][]byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// Add the null character as a terminator

type serviceIndexerBySecret struct{}

func (si serviceIndexerBySecret) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (si serviceIndexerBySecret) FromObject(obj interface{}) (bool, [][]byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// Add the null character as a terminator

type serviceIndexerByConfig struct{}

func (si serviceIndexerByConfig) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (si serviceIndexerByConfig) FromObject(obj interface{}) (bool, [][]byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// Add the null character as a terminator
