package api

import (
	"errors"

	"github.com/docker/go-events"
)

var (
	errUnknownStoreAction = errors.New("unrecognized action type")
	errConflictingFilters = errors.New("conflicting filters specified")
	errNoKindSpecified    = errors.New("no kind of object specified")
	errUnrecognizedAction = errors.New("unrecognized action")
)

// StoreObject is an abstract object that can be handled by the store.
type StoreObject interface {
	GetID() string                           // Get ID
	GetMeta() Meta                           // Retrieve metadata
	SetMeta(Meta)                            // Set metadata
	CopyStoreObject() StoreObject            // Return a copy of this object
	EventCreate() Event                      // Return a creation event
	EventUpdate(oldObject StoreObject) Event // Return an update event
	EventDelete() Event                      // Return a deletion event
}

// Event is the type used for events passed over watcher channels, and also
// the type used to specify filtering in calls to Watch.
type Event interface {
	// TODO(stevvooe): Consider whether it makes sense to squish both the
	// matcher type and the primary type into the same type. It might be better
	// to build a matcher from an event prototype.

	// Matches checks if this item in a watch queue Matches the event
	// description.
	Matches(events.Event) bool
}

// EventCreate is an interface implemented by every creation event type
type EventCreate interface {
	IsEventCreate() bool
}

// EventUpdate is an interface implemented by every update event type
type EventUpdate interface {
	IsEventUpdate() bool
}

// EventDelete is an interface implemented by every delete event type
type EventDelete interface {
	IsEventDelete()
}

func customIndexer(kind string, annotations *Annotations) (bool, [][]byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// Add the null character as a terminator

func fromArgs(args ...interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Add the null character as a terminator

func prefixFromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Strip the null terminator, the rest is a prefix

func checkCustom(a1, a2 Annotations) bool { _ = "STUB: not implemented"; return false }

func checkCustomPrefix(a1, a2 Annotations) bool { _ = "STUB: not implemented"; return false }
