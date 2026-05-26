package defaults

import (
	"time"

	gogotypes "github.com/gogo/protobuf/types"
	"github.com/moby/swarmkit/v2/api"
)

// Service is a ServiceSpec object with all fields filled in using default
// values.
var Service = api.ServiceSpec{
	Task: api.TaskSpec{
		Runtime: &api.TaskSpec_Container{
			Container: &api.ContainerSpec{
				StopGracePeriod: gogotypes.DurationProto(10 * time.Second),
				PullOptions:     &api.ContainerSpec_PullOptions{},
				DNSConfig:       &api.ContainerSpec_DNSConfig{},
			},
		},
		Resources: &api.ResourceRequirements{},
		Restart: &api.RestartPolicy{
			Condition: api.RestartOnAny,
			Delay:     gogotypes.DurationProto(5 * time.Second),
		},
		Placement: &api.Placement{},
	},
	Update: &api.UpdateConfig{
		FailureAction: api.UpdateConfig_PAUSE,
		Monitor:       gogotypes.DurationProto(5 * time.Second),
		Parallelism:   1,
		Order:         api.UpdateConfig_STOP_FIRST,
	},
	Rollback: &api.UpdateConfig{
		FailureAction: api.UpdateConfig_PAUSE,
		Monitor:       gogotypes.DurationProto(5 * time.Second),
		Parallelism:   1,
		Order:         api.UpdateConfig_STOP_FIRST,
	},
}

// InterpolateService returns a ServiceSpec based on the provided spec, which
// has all unspecified values filled in with default values.
func InterpolateService(origSpec *api.ServiceSpec) *api.ServiceSpec {
	_ = "STUB: not implemented"
	return nil
}
