package config

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
)

func getConfig(ctx context.Context, c api.ControlClient, input string) (*api.Config, error) {
	_ = "STUB: not implemented"
	// not sure what it is, match by name or id prefix
	return nil, nil
}

// ok, multiple matches.  Prefer exact ID over exact name.  If no exact matches, return an error
