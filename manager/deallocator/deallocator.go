package deallocator

import (
	"context"

	"github.com/docker/go-events"
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

// Deallocator waits for services to fully shutdown (ie no containers left)
// and then proceeds to deallocate service-level resources (e.g. networks),
// and finally services themselves
// in particular, the Deallocator should be the only place where services, or
// service-level resources, are ever deleted!
//
// It’s worth noting that this new component’s role is quite different from
// the task reaper’s: tasks are purely internal to Swarmkit, and their status
// is entirely managed by the system itself. In contrast, the deallocator is
// responsible for safely deleting entities that are directly controlled by the
// user.
//
// NOTE: since networks are the only service-level resources as of now,
// it has been deemed over-engineered to have a generic way to
// handle other types of service-level resources; if we ever start
// having more of those and thus want to reconsider this choice, it
// might be worth having a look at this archived branch, that does
// implement a way of separating the code for the deallocator itself
// from each resource-speficic way of handling it
// https://github.com/docker/swarmkit/compare/a84c01f49091167dd086c26b45dc18b38d52e4d9...wk8:wk8/generic_deallocator#diff-75f4f75eee6a6a7a7268c672203ea0ac
type Deallocator struct {
	store *store.MemoryStore

	// for services that are shutting down, we keep track of how many
	// tasks still exist for them
	services map[string]*serviceWithTaskCounts

	// mainly used for tests, so that we can peek
	// into the DB state in between events
	// the bool notifies whether any DB update was actually performed
	eventChan chan bool

	stopChan chan struct{}
	doneChan chan struct{}
}

// used in our internal state's `services` right above
type serviceWithTaskCounts struct {
	service   *api.Service
	taskCount int
}

// New creates a new deallocator
func New(store *store.MemoryStore) *Deallocator { _ = "STUB: not implemented"; return nil }

// Run starts the deallocator, which then starts cleaning up services
// and their resources when relevant (ie when no tasks still exist
// for a given service)
// This is a blocking function
func (deallocator *Deallocator) Run(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// look for services that are marked for deletion
// there's no index on the `PendingDelete` field in the store,
// so we just iterate over all of them and filter manually
// this is okay since we only do this at leadership change

// now we also need to look at all existing service-level networks
// that may be marked for deletion

// if we have an error here, we can't proceed any further

// eventsChanCancel()

// now let's populate our internal taskCounts

// and deallocate networks that may be marked for deletion and aren't used any more

// now we just need to wait for events

// Stop stops the deallocator's routine
// FIXME (jrouge): see the comment on TaskReaper.Stop() and see when to properly stop this
// plus unit test on this!
func (deallocator *Deallocator) Stop() { _ = "STUB: not implemented"; return }

// always a bno-op, except when running tests tests
// see the comment about `Deallocator`s' `eventChan` field
func (deallocator *Deallocator) notifyEventChan(updated bool) { _ = "STUB: not implemented"; return }

// if a service is marked for deletion, this checks whether it's ready to be
// deleted yet, and does it if relevant
func (deallocator *Deallocator) processService(ctx context.Context, service *api.Service) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// if in doubt, let's proceed to clean up the service anyway
// better to clean up resources that shouldn't be cleaned up yet
// than ending up with a service and some resources lost in limbo forever

// no tasks remaining for this service, we can clean it up

func (deallocator *Deallocator) deallocateService(ctx context.Context, service *api.Service) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// first, let's delete the service

// all errors are just for logging here, we do a best effort at cleaning up everything we can

// then all of its networks, provided no other service uses them

// see https://github.com/docker/swarmkit/blob/e2aafdd3453d2ab103dd97364f79ea6b857f9446/api/specs.proto#L80-L84
// we really should have a helper function on services to do this...

// proceeds to deallocate a network if it's pending deletion and there no
// longer are any services using it
// actually deletes the network if it's marked for deletion and no services are
// using it any more (or the only one using it has ID `ignoreServiceID`, if not
// nil - this comes in handy when there's been an error deleting a service)
// This function can be called either when deallocating a whole service, or
// because there was an `EventUpdateNetwork` event - in the former case, the
// transaction will be that of the service deallocation, in the latter it will be nil
func (deallocator *Deallocator) processNetwork(ctx context.Context, tx store.Tx, network *api.Network, ignoreServiceID *string) (updated bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Processes new events, and dispatches to the right method depending on what
// type of event it is.
// The boolean part of the return tuple indicates whether anything was actually
// removed from the store
func (deallocator *Deallocator) processNewEvent(ctx context.Context, event events.Event) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
