package flagparser

import (
	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/pflag"
)

func parsePorts(flags *pflag.FlagSet, spec *api.ServiceSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// In swarmctl all ports are by default
// PublishModeHost

func parsePortConfig(portConfig string) (string, api.PortConfig_Protocol, uint32, uint32, error) {
	_ = "STUB: not implemented"
	return "", *new(api.PortConfig_Protocol), 0, 0, nil
}

func parsePortSpec(portSpec string) (api.PortConfig_Protocol, uint32, error) {
	_ = "STUB: not implemented"
	return *new(api.PortConfig_Protocol), 0, nil
}
