package flagparser

import (
	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/cobra"
)

// ParseAddCapability validates capabilities passed on the command line
func ParseAddCapability(cmd *cobra.Command, spec *api.ServiceSpec, flagName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Index adds so we don't have to double loop

// Check if any of the adds are in drop so we can remove them from the drop list.

// De-dup the list to be added

// ParseDropCapability validates capabilities passed on the command line
func ParseDropCapability(cmd *cobra.Command, spec *api.ServiceSpec, flagName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Index removals so we don't have to double loop

// Check if any of the adds are in add so we can remove them from the add list.

// De-dup the list to be dropped
