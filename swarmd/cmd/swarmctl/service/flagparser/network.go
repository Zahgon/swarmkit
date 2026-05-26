package flagparser

import (
	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/cobra"
)

func parseNetworks(cmd *cobra.Command, spec *api.ServiceSpec, c api.ControlClient) error {
	_ = "STUB: not implemented"
	return nil
}
