package allocator

import (
	"context"
	"sync"

	"github.com/docker/go-events"
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/allocator/networkallocator"
	"github.com/moby/swarmkit/v2/manager/state"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

// Allocator controls how the allocation stage in the manager is handled.
type Allocator struct {
	// The manager store.
	store *store.MemoryStore

	// the ballot used to synchronize across all allocators to ensure
	// all of them have completed their respective allocations so that the
	// task can be moved to ALLOCATED state.
	taskBallot *taskBallot

	// context for the network allocator that will be needed by
	// network allocator.
	netCtx *networkContext

	// stopChan signals to the allocator to stop running.
	stopChan chan struct{}
	// doneChan is closed when the allocator is finished running.
	doneChan chan struct{}

	nwkAllocator networkallocator.NetworkAllocator
}

// taskBallot controls how the voting for task allocation is
// coordinated b/w different allocators. This the only structure that
// will be written by all allocator goroutines concurrently. Hence the
// mutex.
type taskBallot struct {
	sync.Mutex

	// List of registered voters who have to cast their vote to
	// indicate their allocation complete
	voters []string

	// List of votes collected for every task so far from different voters.
	votes map[string][]string
}

// allocActor controls the various phases in the lifecycle of one kind of allocator.
type allocActor struct {
	// Task voter identity of the allocator.
	taskVoter string

	// Action routine which is called for every event that the
	// allocator received.
	action func(context.Context, events.Event)

	// Init routine which is called during the initialization of
	// the allocator.
	init func(ctx context.Context) error
}

// New returns a new instance of Allocator for use during allocation
// stage of the manager.
func New(store *store.MemoryStore, na networkallocator.NetworkAllocator) *Allocator {
	_ = "STUB: not implemented"
	return nil
}

// Run starts all allocator go-routines and waits for Stop to be called.
func (a *Allocator) Run(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Setup cancel context for all goroutines to use.
	return nil
}

// Assign a pointer for variable capture

// init might return an allocator specific context
// which is a child of the passed in context to hold
// allocator specific state

// Stop stops the allocator
func (a *Allocator) Stop() {
	_ = "STUB: not implemented"

	// Wait for all allocator goroutines to truly exit
	return
}

func (a *Allocator) init(ctx context.Context, aa *allocActor) (<-chan events.Event, func(), error) {
	watch, watchCancel := state.Watch(a.store.WatchQueue(),
		api.EventCreateNetwork{},
		api.EventDeleteNetwork{},
		api.EventCreateService{},
		api.EventUpdateService{},
		api.EventDeleteService{},
		api.EventCreateTask{},
		api.EventUpdateTask{},
		api.EventDeleteTask{},
		api.EventCreateNode{},
		api.EventUpdateNode{},
		api.EventDeleteNode{},
		state.EventCommit{},
	)

	if err := aa.init(ctx); err != nil {
		watchCancel()
		return nil, nil, err
	}

	return watch, watchCancel, nil
}

func (a *Allocator) run(ctx context.Context, aa allocActor, watch <-chan events.Event) {
	_ = "STUB: not implemented"
	return
}

func (a *Allocator) registerToVote(name string) { _ = "STUB: not implemented"; return }

func (a *Allocator) taskAllocateVote(voter string, id string) bool {
	_ = "STUB: not implemented"
	return false
}

// If voter has already voted, return false

// check if voter is in x

// We haven't gotten enough votes yet

// Not every registered voter has registered a vote.
