package flagparser

import (
	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/pflag"
)

func parseResourceCPU(flags *pflag.FlagSet, resources *api.Resources, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func parseResourceMemory(flags *pflag.FlagSet, resources *api.Resources, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func parseResource(flags *pflag.FlagSet, spec *api.ServiceSpec) error {
	_ = "STUB: not implemented"
	return nil
}
