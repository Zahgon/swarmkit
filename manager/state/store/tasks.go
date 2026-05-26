package store

import (
	memdb "github.com/hashicorp/go-memdb"
	"github.com/moby/swarmkit/v2/api"
)

const tableTask = "task"

func init() {
	register(ObjectStoreConfig{
		Table: &memdb.TableSchema{
			Name: tableTask,
			Indexes: map[string]*memdb.IndexSchema{
				indexID: {
					Name:    indexID,
					Unique:  true,
					Indexer: api.TaskIndexerByID{},
				},
				indexName: {
					Name:         indexName,
					AllowMissing: true,
					Indexer:      taskIndexerByName{},
				},
				indexRuntime: {
					Name:         indexRuntime,
					AllowMissing: true,
					Indexer:      taskIndexerByRuntime{},
				},
				indexServiceID: {
					Name:         indexServiceID,
					AllowMissing: true,
					Indexer:      taskIndexerByServiceID{},
				},
				indexNodeID: {
					Name:         indexNodeID,
					AllowMissing: true,
					Indexer:      taskIndexerByNodeID{},
				},
				indexSlot: {
					Name:         indexSlot,
					AllowMissing: true,
					Indexer:      taskIndexerBySlot{},
				},
				indexDesiredState: {
					Name:    indexDesiredState,
					Indexer: taskIndexerByDesiredState{},
				},
				indexTaskState: {
					Name:    indexTaskState,
					Indexer: taskIndexerByTaskState{},
				},
				indexNetwork: {
					Name:         indexNetwork,
					AllowMissing: true,
					Indexer:      taskIndexerByNetwork{},
				},
				indexSecret: {
					Name:         indexSecret,
					AllowMissing: true,
					Indexer:      taskIndexerBySecret{},
				},
				indexConfig: {
					Name:         indexConfig,
					AllowMissing: true,
					Indexer:      taskIndexerByConfig{},
				},
				indexVolumeAttachment: {
					Name:         indexVolumeAttachment,
					AllowMissing: true,
					Indexer:      taskIndexerByVolumeAttachment{},
				},
				indexCustom: {
					Name:         indexCustom,
					Indexer:      api.TaskCustomIndexer{},
					AllowMissing: true,
				},
			},
		},
		Save: func(tx ReadTx, snapshot *api.StoreSnapshot) error {
			var err error
			snapshot.Tasks, err = FindTasks(tx, All)
			return err
		},
		Restore: func(tx Tx, snapshot *api.StoreSnapshot) error {
			toStoreObj := make([]api.StoreObject, len(snapshot.Tasks))
			for i, x := range snapshot.Tasks {
				toStoreObj[i] = x
			}
			return RestoreTable(tx, tableTask, toStoreObj)
		},
		ApplyStoreAction: func(tx Tx, sa api.StoreAction) error {
			switch v := sa.Target.(type) {
			case *api.StoreAction_Task:
				obj := v.Task
				switch sa.Action {
				case api.StoreActionKindCreate:
					return CreateTask(tx, obj)
				case api.StoreActionKindUpdate:
					return UpdateTask(tx, obj)
				case api.StoreActionKindRemove:
					return DeleteTask(tx, obj.ID)
				}
			}
			return errUnknownStoreAction
		},
	})
}

// CreateTask adds a new task to the store.
// Returns ErrExist if the ID is already taken.
func CreateTask(tx Tx, t *api.Task) error { _ = "STUB: not implemented"; return nil }

// UpdateTask updates an existing task in the store.
// Returns ErrNotExist if the node doesn't exist.
func UpdateTask(tx Tx, t *api.Task) error { _ = "STUB: not implemented"; return nil }

// DeleteTask removes a task from the store.
// Returns ErrNotExist if the task doesn't exist.
func DeleteTask(tx Tx, id string) error { _ = "STUB: not implemented"; return nil }

// GetTask looks up a task by ID.
// Returns nil if the task doesn't exist.
func GetTask(tx ReadTx, id string) *api.Task { _ = "STUB: not implemented"; return nil }

// FindTasks selects a set of tasks and returns them.
func FindTasks(tx ReadTx, by By) ([]*api.Task, error) { _ = "STUB: not implemented"; return nil, nil }

type taskIndexerByName struct{}

func (ti taskIndexerByName) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ti taskIndexerByName) FromObject(obj interface{}) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// Add the null character as a terminator

func (ti taskIndexerByName) PrefixFromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type taskIndexerByRuntime struct{}

func (ti taskIndexerByRuntime) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ti taskIndexerByRuntime) FromObject(obj interface{}) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (ti taskIndexerByRuntime) PrefixFromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type taskIndexerByServiceID struct{}

func (ti taskIndexerByServiceID) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ti taskIndexerByServiceID) FromObject(obj interface{}) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false,

		// Add the null character as a terminator
		nil, nil
}

type taskIndexerByNodeID struct{}

func (ti taskIndexerByNodeID) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ti taskIndexerByNodeID) FromObject(obj interface{}) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false,

		// Add the null character as a terminator
		nil, nil
}

type taskIndexerBySlot struct{}

func (ti taskIndexerBySlot) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ti taskIndexerBySlot) FromObject(obj interface{}) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false,

		// Add the null character as a terminator
		nil, nil
}

type taskIndexerByDesiredState struct{}

func (ti taskIndexerByDesiredState) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ti taskIndexerByDesiredState) FromObject(obj interface{}) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false,

		// Add the null character as a terminator
		nil, nil
}

type taskIndexerByNetwork struct{}

func (ti taskIndexerByNetwork) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ti taskIndexerByNetwork) FromObject(obj interface{}) (bool, [][]byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// Add the null character as a terminator

type taskIndexerBySecret struct{}

func (ti taskIndexerBySecret) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ti taskIndexerBySecret) FromObject(obj interface{}) (bool, [][]byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// Add the null character as a terminator

type taskIndexerByConfig struct{}

func (ti taskIndexerByConfig) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ti taskIndexerByConfig) FromObject(obj interface{}) (bool, [][]byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// Add the null character as a terminator

type taskIndexerByVolumeAttachment struct{}

func (ti taskIndexerByVolumeAttachment) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ti taskIndexerByVolumeAttachment) FromObject(obj interface{}) (bool, [][]byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

type taskIndexerByTaskState struct{}

func (ts taskIndexerByTaskState) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ts taskIndexerByTaskState) FromObject(obj interface{}) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false,

		// Add the null character as a terminator
		nil, nil
}
