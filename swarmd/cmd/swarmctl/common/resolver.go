package common

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/cobra"
)

// Resolver provides ID to Name resolution.
type Resolver struct {
	cmd   *cobra.Command
	c     api.ControlClient
	ctx   context.Context
	cache map[string]string
}

// NewResolver creates a new Resolver.
func NewResolver(cmd *cobra.Command, c api.ControlClient) *Resolver {
	_ = "STUB: not implemented"
	return nil
}

func (r *Resolver) get(t interface{}, id string) string { _ = "STUB: not implemented"; return "" }

// Resolve will attempt to resolve an ID to a Name by querying the manager.
// Results are stored into a cache.
// If the `-n` flag is used in the command-line, resolution is disabled.
func (r *Resolver) Resolve(t interface{}, id string) string { _ = "STUB: not implemented"; return "" }
