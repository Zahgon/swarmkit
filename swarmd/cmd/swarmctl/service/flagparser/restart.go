package flagparser

import (
	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/pflag"
)

func parseRestart(flags *pflag.FlagSet, spec *api.ServiceSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// set new service's restart policy as RestartOnAny
