package flagparser

import (
	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/pflag"
)

// parseBind only supports a very simple version of bind for testing the most
// basic of data flows. Replace with a --mount flag, similar to what we have in
// docker service.
func parseBind(flags *pflag.FlagSet, spec *api.ServiceSpec) error {
	_ = "STUB: not implemented"
	return nil
}
