package cluster

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
)

func getCluster(ctx context.Context, c api.ControlClient, input string) (*api.Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
