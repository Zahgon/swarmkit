package store

import (
	memdb "github.com/hashicorp/go-memdb"
	"github.com/moby/swarmkit/v2/api"
)

const tableNetwork = "network"

func init() {
	register(ObjectStoreConfig{
		Table: &memdb.TableSchema{
			Name: tableNetwork,
			Indexes: map[string]*memdb.IndexSchema{
				indexID: {
					Name:    indexID,
					Unique:  true,
					Indexer: api.NetworkIndexerByID{},
				},
				indexName: {
					Name:    indexName,
					Unique:  true,
					Indexer: api.NetworkIndexerByName{},
				},
				indexCustom: {
					Name:         indexCustom,
					Indexer:      api.NetworkCustomIndexer{},
					AllowMissing: true,
				},
			},
		},
		Save: func(tx ReadTx, snapshot *api.StoreSnapshot) error {
			var err error
			snapshot.Networks, err = FindNetworks(tx, All)
			return err
		},
		Restore: func(tx Tx, snapshot *api.StoreSnapshot) error {
			toStoreObj := make([]api.StoreObject, len(snapshot.Networks))
			for i, x := range snapshot.Networks {
				toStoreObj[i] = x
			}
			return RestoreTable(tx, tableNetwork, toStoreObj)
		},
		ApplyStoreAction: func(tx Tx, sa api.StoreAction) error {
			switch v := sa.Target.(type) {
			case *api.StoreAction_Network:
				obj := v.Network
				switch sa.Action {
				case api.StoreActionKindCreate:
					return CreateNetwork(tx, obj)
				case api.StoreActionKindUpdate:
					return UpdateNetwork(tx, obj)
				case api.StoreActionKindRemove:
					return DeleteNetwork(tx, obj.ID)
				}
			}
			return errUnknownStoreAction
		},
	})
}

// CreateNetwork adds a new network to the store.
// Returns ErrExist if the ID is already taken.
func CreateNetwork(tx Tx, n *api.Network) error {
	_ = "STUB: not implemented"
	// Ensure the name is not already in use.
	return nil
}

// UpdateNetwork updates an existing network in the store.
// Returns ErrNotExist if the network doesn't exist.
func UpdateNetwork(tx Tx, n *api.Network) error {
	_ = "STUB: not implemented"
	// Ensure the name is either not in use or already used by this same Network.
	return nil
}

// DeleteNetwork removes a network from the store.
// Returns ErrNotExist if the network doesn't exist.
func DeleteNetwork(tx Tx, id string) error { _ = "STUB: not implemented"; return nil }

// GetNetwork looks up a network by ID.
// Returns nil if the network doesn't exist.
func GetNetwork(tx ReadTx, id string) *api.Network { _ = "STUB: not implemented"; return nil }

// FindNetworks selects a set of networks and returns them.
func FindNetworks(tx ReadTx, by By) ([]*api.Network, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
