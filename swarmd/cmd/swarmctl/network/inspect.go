package network

import (
	"errors"

	"github.com/moby/swarmkit/swarmd/cmd/swarmctl/common"
	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/cobra"
)

var (
	inspectCmd = &cobra.Command{
		Use:   "inspect <network ID>",
		Short: "Inspect a network",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("network ID missing")
			}

			if len(args) > 1 {
				return errors.New("inspect command takes exactly 1 argument")
			}

			c, err := common.Dial(cmd)
			if err != nil {
				return err
			}
			network, err := GetNetwork(common.Context(cmd), c, args[0])
			if err != nil {
				return err
			}

			printNetworkSummary(network)

			return nil
		},
	}
)

func printNetworkSummary(network *api.Network) { _ = "STUB: not implemented"; return }

// Ignore flushing errors - there's nothing we can do.
