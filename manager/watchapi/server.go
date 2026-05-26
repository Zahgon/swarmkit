package watchapi

import (
	"context"
	"errors"
	"sync"

	"github.com/moby/swarmkit/v2/manager/state/store"
)

var (
	errAlreadyRunning = errors.New("broker is already running")
	errNotRunning     = errors.New("broker is not running")
)

// Server is the store API gRPC server.
type Server struct {
	store     *store.MemoryStore
	mu        sync.Mutex
	pctx      context.Context
	cancelAll func()
}

// NewServer creates a store API server.
func NewServer(store *store.MemoryStore) *Server { _ = "STUB: not implemented"; return nil }

// Start starts the watch server.
func (s *Server) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Stop stops the watch server.
func (s *Server) Stop() error { _ = "STUB: not implemented"; return nil }
