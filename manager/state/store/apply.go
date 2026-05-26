package store

import (
	"github.com/docker/go-events"
)

// Apply takes an item from the event stream of one Store and applies it to
// a second Store.
func Apply(store *MemoryStore, item events.Event) (err error) {
	_ = "STUB: not implemented"
	return nil
}
