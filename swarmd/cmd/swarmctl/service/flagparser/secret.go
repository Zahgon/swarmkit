package flagparser

import (
	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/cobra"
)

// expects secrets in the format SECRET_NAME:TARGET_NAME
func parseSecretString(secretString string) (secretName, presentName string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// ParseAddSecret validates secrets passed on the command line
func ParseAddSecret(cmd *cobra.Command, spec *api.ServiceSpec, flagName string) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(diogo): defaults to File targets, but in the future will take different types

// ParseRemoveSecret removes a set of secrets from the task spec's secret references
func ParseRemoveSecret(cmd *cobra.Command, spec *api.ServiceSpec, flagName string) error {
	_ = "STUB: not implemented"
	return nil
}
