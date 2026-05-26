package store

import (
	memdb "github.com/hashicorp/go-memdb"
	"github.com/moby/swarmkit/v2/api"
)

const tableNode = "node"

func init() {
	register(ObjectStoreConfig{
		Table: &memdb.TableSchema{
			Name: tableNode,
			Indexes: map[string]*memdb.IndexSchema{
				indexID: {
					Name:    indexID,
					Unique:  true,
					Indexer: api.NodeIndexerByID{},
				},
				// TODO(aluzzardi): Use `indexHostname` instead.
				indexName: {
					Name:         indexName,
					AllowMissing: true,
					Indexer:      nodeIndexerByHostname{},
				},
				indexRole: {
					Name:    indexRole,
					Indexer: nodeIndexerByRole{},
				},
				indexMembership: {
					Name:    indexMembership,
					Indexer: nodeIndexerByMembership{},
				},
				indexCustom: {
					Name:         indexCustom,
					Indexer:      api.NodeCustomIndexer{},
					AllowMissing: true,
				},
			},
		},
		Save: func(tx ReadTx, snapshot *api.StoreSnapshot) error {
			var err error
			snapshot.Nodes, err = FindNodes(tx, All)
			return err
		},
		Restore: func(tx Tx, snapshot *api.StoreSnapshot) error {
			toStoreObj := make([]api.StoreObject, len(snapshot.Nodes))
			for i, x := range snapshot.Nodes {
				toStoreObj[i] = x
			}
			return RestoreTable(tx, tableNode, toStoreObj)
		},
		ApplyStoreAction: func(tx Tx, sa api.StoreAction) error {
			switch v := sa.Target.(type) {
			case *api.StoreAction_Node:
				obj := v.Node
				switch sa.Action {
				case api.StoreActionKindCreate:
					return CreateNode(tx, obj)
				case api.StoreActionKindUpdate:
					return UpdateNode(tx, obj)
				case api.StoreActionKindRemove:
					return DeleteNode(tx, obj.ID)
				}
			}
			return errUnknownStoreAction
		},
	})
}

// CreateNode adds a new node to the store.
// Returns ErrExist if the ID is already taken.
func CreateNode(tx Tx, n *api.Node) error { _ = "STUB: not implemented"; return nil }

// UpdateNode updates an existing node in the store.
// Returns ErrNotExist if the node doesn't exist.
func UpdateNode(tx Tx, n *api.Node) error { _ = "STUB: not implemented"; return nil }

// DeleteNode removes a node from the store.
// Returns ErrNotExist if the node doesn't exist.
func DeleteNode(tx Tx, id string) error { _ = "STUB: not implemented"; return nil }

// GetNode looks up a node by ID.
// Returns nil if the node doesn't exist.
func GetNode(tx ReadTx, id string) *api.Node { _ = "STUB: not implemented"; return nil }

// FindNodes selects a set of nodes and returns them.
func FindNodes(tx ReadTx, by By) ([]*api.Node, error) { _ = "STUB: not implemented"; return nil, nil }

type nodeIndexerByHostname struct{}

func (ni nodeIndexerByHostname) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ni nodeIndexerByHostname) FromObject(obj interface{}) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// Add the null character as a terminator

func (ni nodeIndexerByHostname) PrefixFromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type nodeIndexerByRole struct{}

func (ni nodeIndexerByRole) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ni nodeIndexerByRole) FromObject(obj interface{}) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false,

		// Add the null character as a terminator
		nil, nil
}

type nodeIndexerByMembership struct{}

func (ni nodeIndexerByMembership) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ni nodeIndexerByMembership) FromObject(obj interface{}) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false,

		// Add the null character as a terminator
		nil, nil
}
