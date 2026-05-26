package genericresource

import (
	"github.com/moby/swarmkit/v2/api"
)

func newParseError(format string, args ...interface{}) error { _ = "STUB: not implemented"; return nil }

// discreteResourceVal returns an int64 if the string is a discreteResource
// and an error if it isn't
func discreteResourceVal(res string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// allNamedResources returns true if the array of resources are all namedResources
// e.g: res = [red, orange, green]
func allNamedResources(res []string) bool { _ = "STUB: not implemented"; return false }

// ParseCmd parses the Generic Resource command line argument
// and returns a list of *api.GenericResource
func ParseCmd(cmd string) ([]*api.GenericResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse parses a table of GenericResource resources
func Parse(cmds []string) ([]*api.GenericResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// isDiscreteResource returns true if the array of resources is a
// Discrete Resource.
// e.g: res = [1]
func isDiscreteResource(values []string) (int64, bool) { _ = "STUB: not implemented"; return 0, false }
