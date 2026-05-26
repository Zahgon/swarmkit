package genericresource

import (
	"github.com/moby/swarmkit/v2/api"
)

func discreteToString(d *api.GenericResource_DiscreteResourceSpec) string {
	_ = "STUB: not implemented"
	return ""
}

// Kind returns the kind key as a string
func Kind(res *api.GenericResource) string { _ = "STUB: not implemented"; return "" }

// Value returns the value key as a string
func Value(res *api.GenericResource) string { _ = "STUB: not implemented"; return "" }

// EnvFormat returns the environment string version of the resource
func EnvFormat(res []*api.GenericResource, prefix string) []string {
	_ = "STUB: not implemented"
	return nil
}
