package config

import (
	"errors"

	"github.com/moby/swarmkit/swarmd/cmd/swarmctl/common"
	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/cobra"
)

func printConfigSummary(config *api.Config) { _ = "STUB: not implemented"; return }

var (
	inspectCmd = &cobra.Command{
		Use:   "inspect <config ID or name>",
		Short: "Inspect a config",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("inspect command takes a single config ID or name")
			}

			client, err := common.Dial(cmd)
			if err != nil {
				return err
			}

			config, err := getConfig(common.Context(cmd), client, args[0])
			if err != nil {
				return err
			}

			printConfigSummary(config)
			return nil
		},
	}
)
