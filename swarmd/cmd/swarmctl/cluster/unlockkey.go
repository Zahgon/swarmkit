package cluster

import (
	"errors"

	"github.com/spf13/cobra"
)

// get the unlock key

func displayUnlockKey(cmd *cobra.Command) error { _ = "STUB: not implemented"; return nil }

var (
	unlockKeyCmd = &cobra.Command{
		Use:   "unlock-key <cluster name>",
		Short: "Get the unlock key for a cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("cluster name missing")
			}

			if len(args) > 1 {
				return errors.New("unlock-key command takes exactly 1 argument")
			}

			return displayUnlockKey(cmd)
		},
	}
)
