package agent

import (
	"github.com/moby/swarmkit/v2/api"
	bolt "go.etcd.io/bbolt"
)

// Layout:
//
//	bucket(v1.tasks.<id>) ->
//		data (task protobuf)
//		status (task status protobuf)
//		assigned (key present)
var (
	bucketKeyStorageVersion = []byte("v1")
	bucketKeyTasks          = []byte("tasks")
	bucketKeyAssigned       = []byte("assigned")
	bucketKeyData           = []byte("data")
	bucketKeyStatus         = []byte("status")
)

// InitDB prepares a database for writing task data.
//
// Proper buckets will be created if they don't already exist.
func InitDB(db *bolt.DB) error { _ = "STUB: not implemented"; return nil }

// GetTask retrieves the task with id from the datastore.
func GetTask(tx *bolt.Tx, id string) (*api.Task, error) { _ = "STUB: not implemented"; return nil, nil }

// WalkTasks walks all tasks in the datastore.
func WalkTasks(tx *bolt.Tx, fn func(task *api.Task) error) error {
	_ = "STUB: not implemented"
	return nil
}

// TaskAssigned returns true if the task is assigned to the node.
func TaskAssigned(tx *bolt.Tx, id string) bool { _ = "STUB: not implemented"; return false }

// GetTaskStatus returns the current status for the task.
func GetTaskStatus(tx *bolt.Tx, id string) (*api.TaskStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WalkTaskStatus calls fn for the status of each task.
func WalkTaskStatus(tx *bolt.Tx, fn func(id string, status *api.TaskStatus) error) error {
	_ = "STUB: not implemented"
	return nil
}

// PutTask places the task into the database.
func PutTask(tx *bolt.Tx, task *api.Task) error { _ = "STUB: not implemented"; return nil }

// blank out the status.

// PutTaskStatus updates the status for the task with id.
func PutTaskStatus(tx *bolt.Tx, id string, status *api.TaskStatus) error {
	_ = "STUB: not implemented"
	// this used to be withCreateTaskBucketIfNotExists, but that could lead
	// to weird race conditions, and was not necessary.
	return nil
}

// DeleteTask completely removes the task from the database.
func DeleteTask(tx *bolt.Tx, id string) error { _ = "STUB: not implemented"; return nil }

// SetTaskAssignment sets the current assignment state.
func SetTaskAssignment(tx *bolt.Tx, id string, assigned bool) error {
	_ = "STUB: not implemented"
	return nil
}

func createBucketIfNotExists(tx *bolt.Tx, keys ...[]byte) (*bolt.Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func withCreateTaskBucketIfNotExists(tx *bolt.Tx, id string, fn func(bkt *bolt.Bucket) error) error {
	_ = "STUB: not implemented"
	return nil
}

func withTaskBucket(tx *bolt.Tx, id string, fn func(bkt *bolt.Bucket) error) error {
	_ = "STUB: not implemented"
	return nil
}

func getTaskBucket(tx *bolt.Tx, id string) *bolt.Bucket { _ = "STUB: not implemented"; return nil }

func getTasksBucket(tx *bolt.Tx) *bolt.Bucket { _ = "STUB: not implemented"; return nil }

func getBucket(tx *bolt.Tx, keys ...[]byte) *bolt.Bucket { _ = "STUB: not implemented"; return nil }
