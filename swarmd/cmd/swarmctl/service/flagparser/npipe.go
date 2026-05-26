package flagparser

import (
	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/pflag"
)

// parseNpipe only supports a very simple version of anonymous npipes for
// testing the most basic of data flows. Replace with a --mount flag, similar
// to what we have in docker service.
func parseNpipe(flags *pflag.FlagSet, spec *api.ServiceSpec) error {
	_ = "STUB: not implemented"
	return nil
}
