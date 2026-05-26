package service

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
)

func getService(ctx context.Context, c api.ControlClient, input string) (*api.Service, error) {
	_ = "STUB: not implemented"
	// GetService to match via full ID.
	return nil, nil
}

// If any error (including NotFound), ListServices to match via full name.

func getServiceReplicasTxt(s *api.Service, running int) string {
	_ = "STUB: not implemented"
	return ""
}
