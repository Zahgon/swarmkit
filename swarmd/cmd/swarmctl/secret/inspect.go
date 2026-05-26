package secret

import (
	"errors"

	"github.com/moby/swarmkit/swarmd/cmd/swarmctl/common"
	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/cobra"
)

func printSecretSummary(secret *api.Secret) { _ = "STUB: not implemented"; return }

var (
	inspectCmd = &cobra.Command{
		Use:   "inspect <secret ID or name>",
		Short: "Inspect a secret",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("inspect command takes a single secret ID or name")
			}

			client, err := common.Dial(cmd)
			if err != nil {
				return err
			}

			secret, err := getSecret(common.Context(cmd), client, args[0])
			if err != nil {
				return err
			}

			printSecretSummary(secret)
			return nil
		},
	}
)
