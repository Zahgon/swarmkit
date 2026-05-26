package store

import (
	memdb "github.com/hashicorp/go-memdb"
	"github.com/moby/swarmkit/v2/api"
)

const (
	tableCluster = "cluster"

	// DefaultClusterName is the default name to use for the cluster
	// object.
	DefaultClusterName = "default"
)

func init() {
	register(ObjectStoreConfig{
		Table: &memdb.TableSchema{
			Name: tableCluster,
			Indexes: map[string]*memdb.IndexSchema{
				indexID: {
					Name:    indexID,
					Unique:  true,
					Indexer: api.ClusterIndexerByID{},
				},
				indexName: {
					Name:    indexName,
					Unique:  true,
					Indexer: api.ClusterIndexerByName{},
				},
				indexCustom: {
					Name:         indexCustom,
					Indexer:      api.ClusterCustomIndexer{},
					AllowMissing: true,
				},
			},
		},
		Save: func(tx ReadTx, snapshot *api.StoreSnapshot) error {
			var err error
			snapshot.Clusters, err = FindClusters(tx, All)
			return err
		},
		Restore: func(tx Tx, snapshot *api.StoreSnapshot) error {
			toStoreObj := make([]api.StoreObject, len(snapshot.Clusters))
			for i, x := range snapshot.Clusters {
				toStoreObj[i] = x
			}
			return RestoreTable(tx, tableCluster, toStoreObj)
		},
		ApplyStoreAction: func(tx Tx, sa api.StoreAction) error {
			switch v := sa.Target.(type) {
			case *api.StoreAction_Cluster:
				obj := v.Cluster
				switch sa.Action {
				case api.StoreActionKindCreate:
					return CreateCluster(tx, obj)
				case api.StoreActionKindUpdate:
					return UpdateCluster(tx, obj)
				case api.StoreActionKindRemove:
					return DeleteCluster(tx, obj.ID)
				}
			}
			return errUnknownStoreAction
		},
	})
}

// CreateCluster adds a new cluster to the store.
// Returns ErrExist if the ID is already taken.
func CreateCluster(tx Tx, c *api.Cluster) error {
	_ = "STUB: not implemented"
	// Ensure the name is not already in use.
	return nil
}

// UpdateCluster updates an existing cluster in the store.
// Returns ErrNotExist if the cluster doesn't exist.
func UpdateCluster(tx Tx, c *api.Cluster) error {
	_ = "STUB: not implemented"
	// Ensure the name is either not in use or already used by this same Cluster.
	return nil
}

// DeleteCluster removes a cluster from the store.
// Returns ErrNotExist if the cluster doesn't exist.
func DeleteCluster(tx Tx, id string) error { _ = "STUB: not implemented"; return nil }

// GetCluster looks up a cluster by ID.
// Returns nil if the cluster doesn't exist.
func GetCluster(tx ReadTx, id string) *api.Cluster { _ = "STUB: not implemented"; return nil }

// FindClusters selects a set of clusters and returns them.
func FindClusters(tx ReadTx, by By) ([]*api.Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
