package flagparser

import (
	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/cobra"
)

// expects configs in the format CONFIG_NAME:TARGET_NAME
func parseConfigString(configString string) (configName, presentName string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// ParseAddConfig validates configs passed on the command line
func ParseAddConfig(cmd *cobra.Command, spec *api.ServiceSpec, flagName string) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(diogo): defaults to File targets, but in the future will take different types

// ParseRemoveConfig removes a set of configs from the task spec's config references
func ParseRemoveConfig(cmd *cobra.Command, spec *api.ServiceSpec, flagName string) error {
	_ = "STUB: not implemented"
	return nil
}
