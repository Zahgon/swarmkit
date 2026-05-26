package store

import (
	memdb "github.com/hashicorp/go-memdb"
	"github.com/moby/swarmkit/v2/api"
)

const tableVolume = "volume"

func init() {
	register(ObjectStoreConfig{
		Table: &memdb.TableSchema{
			Name: tableVolume,
			Indexes: map[string]*memdb.IndexSchema{
				indexID: {
					Name:    indexID,
					Unique:  true,
					Indexer: api.VolumeIndexerByID{},
				},
				indexName: {
					Name:    indexName,
					Unique:  true,
					Indexer: api.VolumeIndexerByName{},
				},
				indexCustom: {
					Name:         indexCustom,
					Indexer:      api.VolumeCustomIndexer{},
					AllowMissing: true,
				},
				indexVolumeGroup: {
					Name:    indexVolumeGroup,
					Indexer: volumeIndexerByGroup{},
				},
				indexDriver: {
					Name:    indexDriver,
					Indexer: volumeIndexerByDriver{},
				},
			},
		},
		Save: func(tx ReadTx, snapshot *api.StoreSnapshot) error {
			var err error
			snapshot.Volumes, err = FindVolumes(tx, All)
			return err
		},
		Restore: func(tx Tx, snapshot *api.StoreSnapshot) error {
			toStoreObj := make([]api.StoreObject, len(snapshot.Volumes))
			for i, x := range snapshot.Volumes {
				toStoreObj[i] = x
			}
			return RestoreTable(tx, tableVolume, toStoreObj)
		},
		ApplyStoreAction: func(tx Tx, sa api.StoreAction) error {
			switch v := sa.Target.(type) {
			case *api.StoreAction_Volume:
				obj := v.Volume
				switch sa.Action {
				case api.StoreActionKindCreate:
					return CreateVolume(tx, obj)
				case api.StoreActionKindUpdate:
					return UpdateVolume(tx, obj)
				case api.StoreActionKindRemove:
					return DeleteVolume(tx, obj.ID)
				}
			}
			return errUnknownStoreAction
		},
	})
}

func CreateVolume(tx Tx, v *api.Volume) error { _ = "STUB: not implemented"; return nil }

func UpdateVolume(tx Tx, v *api.Volume) error {
	_ = "STUB: not implemented"
	// ensure the name is either not in use, or is in use by this volume.
	return nil
}

func DeleteVolume(tx Tx, id string) error { _ = "STUB: not implemented"; return nil }

func GetVolume(tx ReadTx, id string) *api.Volume { _ = "STUB: not implemented"; return nil }

func FindVolumes(tx ReadTx, by By) ([]*api.Volume, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type volumeIndexerByGroup struct{}

func (vi volumeIndexerByGroup) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vi volumeIndexerByGroup) FromObject(obj interface{}) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

type volumeIndexerByDriver struct{}

func (vi volumeIndexerByDriver) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vi volumeIndexerByDriver) FromObject(obj interface{}) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false,

		// this should never happen -- existence of the volume driver is checked
		// at the controlapi level. However, guard against the unforeseen.
		nil, nil
}
