package node

import (
	"context"
	"errors"

	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/cobra"
)

var (
	errNoChange = errors.New("node attribute was already set to the requested value")
	flagLabel   = "label"
)

func changeNodeAvailability(cmd *cobra.Command, args []string, availability api.NodeSpec_Availability) error {
	_ = "STUB: not implemented"
	return nil
}

func changeNodeRole(cmd *cobra.Command, args []string, role api.NodeRole) error {
	_ = "STUB: not implemented"
	return nil
}

func getNode(ctx context.Context, c api.ControlClient, input string) (*api.Node, error) {
	_ = "STUB: not implemented"
	// GetNode to match via full ID.
	return nil, nil
}

// If any error (including NotFound), ListServices to match via full name.

func updateNode(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// overwrite existing labels
