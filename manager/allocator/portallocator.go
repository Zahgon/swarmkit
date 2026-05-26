package allocator

import (
	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/internal/idm"
)

const (
	// Start of the dynamic port range from which node ports will
	// be allocated when the user did not specify a port.
	dynamicPortStart = 30000

	// End of the dynamic port range from which node ports will be
	// allocated when the user did not specify a port.
	dynamicPortEnd = 32767

	// The start of master port range which will hold all the
	// allocation state of ports allocated so far regardless of
	// whether it was user defined or not.
	masterPortStart = 1

	// The end of master port range which will hold all the
	// allocation state of ports allocated so far regardless of
	// whether it was user defined or not.
	masterPortEnd = 65535
)

type portAllocator struct {
	// portspace definition per protocol
	portSpaces map[api.PortConfig_Protocol]*portSpace
}

type portSpace struct {
	protocol         api.PortConfig_Protocol
	masterPortSpace  *idm.IDM
	dynamicPortSpace *idm.IDM
}

type allocatedPorts map[api.PortConfig]map[uint32]*api.PortConfig

// addState add the state of an allocated port to the collection.
// `allocatedPorts` is a map of portKey:publishedPort:portState.
// In case the value of the portKey is missing, the map
// publishedPort:portState is created automatically
func (ps allocatedPorts) addState(p *api.PortConfig) { _ = "STUB: not implemented"; return }

// delState delete the state of an allocated port from the collection.
// `allocatedPorts` is a map of portKey:publishedPort:portState.
//
// If publishedPort is non-zero, then it is user defined. We will try to
// remove the portState from `allocatedPorts` directly and return
// the portState (or nil if no portState exists)
//
// If publishedPort is zero, then it is dynamically allocated. We will try
// to remove the portState from `allocatedPorts`, as long as there is
// a portState associated with a non-zero publishedPort.
// Note multiple dynamically allocated ports might exists. In this case,
// we will remove only at a time so both allocated ports are tracked.
//
// Note because of the potential co-existence of user-defined and dynamically
// allocated ports, delState has to be called for user-defined port first.
// dynamically allocated ports should be removed later.
func (ps allocatedPorts) delState(p *api.PortConfig) *api.PortConfig {
	_ = "STUB: not implemented"
	return nil
}

// If name, port, protocol values don't match then we
// are not allocated.

// If SwarmPort was user defined but the port state
// SwarmPort doesn't match we are not allocated.

// Delete state from allocatedPorts

// If PublishedPort == 0 and we don't have non-zero port
// then we are not allocated

// Delete state from allocatedPorts

func newPortAllocator() *portAllocator { _ = "STUB: not implemented"; return nil }

func newPortSpace(protocol api.PortConfig_Protocol) *portSpace {
	_ = "STUB: not implemented"
	return nil
}

// getPortConfigKey returns a map key for doing set operations with
// ports. The key consists of name, protocol and target port which
// uniquely identifies a port within a single Endpoint.
func getPortConfigKey(p *api.PortConfig) api.PortConfig {
	_ = "STUB: not implemented"
	return *new(api.PortConfig)
}

func reconcilePortConfigs(s *api.Service) []*api.PortConfig {
	_ = "STUB: not implemented"
	// If runtime state hasn't been created or if port config has
	// changed from port state return the port config from Spec.
	return nil
}

// Process the portConfig with portConfig.PublishMode != api.PublishModeIngress
// and PublishedPort != 0 (high priority)

// If the PublishMode is not Ingress simply pick up the port config.

// Otherwise we only process PublishedPort != 0 in this round

// Remove record from portState

// For PublishedPort != 0 prefer the portConfig

// Iterate portConfigs with PublishedPort == 0 (low priority)

// Ignore ports which are not PublishModeIngress (already processed)
// And we only process PublishedPort == 0 in this round
// So the following:
//  `portConfig.PublishMode == api.PublishModeIngress && portConfig.PublishedPort == 0`

// If the portConfig is exactly the same as portState
// except if SwarmPort is not user-define then prefer
// portState to ensure sticky allocation of the same
// port that was allocated before.

// Remove record from portState

// For all other cases prefer the portConfig

func (pa *portAllocator) serviceAllocatePorts(s *api.Service) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// We might have previous allocations which we want to stick
// to if possible. So instead of strictly going by port
// configs in the Spec reconcile the list of port configs from
// both the Spec and runtime state.

// Port configuration might have changed. Cleanup all old allocations first.

// Free all the ports allocated so far which
// should be present in s.Endpoints.ExposedPorts

// Make a copy of port config to create runtime state

// Do an actual allocation only if the PublishMode is Ingress

func (pa *portAllocator) serviceDeallocatePorts(s *api.Service) { _ = "STUB: not implemented"; return }

// Do an actual free only if the PublishMode is
// Ingress

func (pa *portAllocator) hostPublishPortsNeedUpdate(s *api.Service) bool {
	_ = "STUB: not implemented"
	return false
}

func (pa *portAllocator) isPortsAllocatedOnInit(s *api.Service, onInit bool) bool {
	_ = "STUB: not implemented"
	// If service has no user-defined endpoint and allocated endpoint,
	// we assume it is allocated and return true.
	return false
}

// If service has allocated endpoint while has no user-defined endpoint,
// we assume allocated endpoints are redundant, and they need deallocated.
// If service has no allocated endpoint while has user-defined endpoint,
// we assume it is not allocated.

// If we don't have same number of port states as port configs
// we assume it is not allocated.

// build a map of host mode ports we've seen. if in the spec we get
// a host port that's not in the service, then we need to do
// allocation. if we get the same target port but something else
// has changed, then HostPublishPortsNeedUpdate will cover that
// case. see docker/swarmkit#2376

// Iterate portConfigs with PublishedPort != 0 (high priority)

// Ignore ports which are not PublishModeIngress

// Iterate portConfigs with PublishedPort == 0 (low priority)

// Ignore ports which are not PublishModeIngress

// If SwarmPort was not defined by user and the func
// is called during allocator initialization state then
// we are not allocated.

// check if the target port is already in the port config. if it
// isn't, then it's our problem.

// NOTE(dperny) there could be a further case where we check if
// there are host ports in the config that aren't in the spec, but
// that's only possible if there's a mismatch in the number of
// ports, which is handled by a length check earlier in the code

func (ps *portSpace) allocate(p *api.PortConfig) (err error) { _ = "STUB: not implemented"; return nil }

// If it falls in the dynamic port range check out
// from dynamic port space first.

// Check out an arbitrary port from dynamic port space.

// Make sure we allocate the same port from the master space.

func (ps *portSpace) free(p *api.PortConfig) { _ = "STUB: not implemented"; return }
