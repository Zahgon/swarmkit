package flagparser

import (
	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// AddServiceFlags add all supported service flags to the flagset.
func AddServiceFlags(flags *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// TODO(stevvooe): Replace these with a more interesting mount flag.

// Merge merges a flagset into a service spec.
func Merge(cmd *cobra.Command, spec *api.ServiceSpec, c api.ControlClient) error {
	_ = "STUB: not implemented"
	return nil
}
