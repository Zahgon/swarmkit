package state

import (
	"github.com/docker/go-events"
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/watch"
)

// EventCommit delineates a transaction boundary.
type EventCommit struct {
	Version *api.Version
}

// Matches returns true if this event is a commit event.
func (e EventCommit) Matches(watchEvent events.Event) bool { _ = "STUB: not implemented"; return false }

// TaskCheckStateGreaterThan is a TaskCheckFunc for checking task state.
func TaskCheckStateGreaterThan(t1, t2 *api.Task) bool { _ = "STUB: not implemented"; return false }

// NodeCheckState is a NodeCheckFunc for matching node state.
func NodeCheckState(n1, n2 *api.Node) bool { _ = "STUB: not implemented"; return false }

// Watch takes a variable number of events to match against. The subscriber
// will receive events that match any of the arguments passed to Watch.
//
// Examples:
//
//	// subscribe to all events
//	Watch(q)
//
//	// subscribe to all UpdateTask events
//	Watch(q, EventUpdateTask{})
//
//	// subscribe to all task-related events
//	Watch(q, EventUpdateTask{}, EventCreateTask{}, EventDeleteTask{})
//
//	// subscribe to UpdateTask for node 123
//	Watch(q, EventUpdateTask{
//		Task:   &api.Task{NodeID: 123},
//		Checks: []TaskCheckFunc{TaskCheckNodeID},
//	})
//
//	// subscribe to UpdateTask for node 123, as well as CreateTask
//	// for node 123 that also has ServiceID set to "abc"
//	Watch(q, EventUpdateTask{
//		Task:   &api.Task{NodeID: 123},
//		Checks: []TaskCheckFunc{TaskCheckNodeID}},
//		EventCreateTask{
//			Task:   &api.Task{NodeID: 123, ServiceID: "abc"},
//			Checks: []TaskCheckFunc{TaskCheckNodeID, func(t1, t2 *api.Task) bool {
//				return t1.ServiceID == t2.ServiceID
//			},
//		},
//	})
func Watch(queue *watch.Queue, specifiers ...api.Event) (eventq chan events.Event, cancel func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Matcher returns an events.Matcher that Matches the specifiers with OR logic.
func Matcher(specifiers ...api.Event) events.MatcherFunc {
	_ = "STUB: not implemented"
	return *new(events.MatcherFunc)
}
