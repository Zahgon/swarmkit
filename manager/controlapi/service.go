package controlapi

import (
	"context"
	"errors"
	"time"

	"github.com/moby/swarmkit/v2/api"
	"github.com/moby/swarmkit/v2/manager/state/store"
)

var (
	errNetworkUpdateNotSupported = errors.New("networks must be migrated to TaskSpec before being changed")
	errRenameNotSupported        = errors.New("renaming services is not supported")
	errModeChangeNotAllowed      = errors.New("service mode change is not allowed")
)

const minimumDuration = 1 * time.Millisecond

func validateResources(r *api.Resources) error { _ = "STUB: not implemented"; return nil }

func validateResourceRequirements(r *api.ResourceRequirements) error {
	_ = "STUB: not implemented"
	return nil
}

func validateRestartPolicy(rp *api.RestartPolicy) error { _ = "STUB: not implemented"; return nil }

func validatePlacement(placement *api.Placement) error { _ = "STUB: not implemented"; return nil }

func validateUpdate(uc *api.UpdateConfig) error { _ = "STUB: not implemented"; return nil }

func validateContainerSpec(taskSpec api.TaskSpec) error {
	_ = "STUB: not implemented"
	// Building a empty/dummy Task to validate the templating and
	// the resulting container spec as well. This is a *best effort*
	// validation.
	return nil
}

// validateImage validates image name in containerSpec
func validateImage(image string) error { _ = "STUB: not implemented"; return nil }

// validateMounts validates if there are duplicate mounts in containerSpec
func validateMounts(mounts []api.Mount) error { _ = "STUB: not implemented"; return nil }

// validateHealthCheck validates configs about container's health check
func validateHealthCheck(hc *api.HealthConfig) error { _ = "STUB: not implemented"; return nil }

func validateGenericRuntimeSpec(taskSpec api.TaskSpec) error { _ = "STUB: not implemented"; return nil }

func validateTaskSpec(taskSpec api.TaskSpec) error { _ = "STUB: not implemented"; return nil }

// Check to see if the secret reference portion of the spec is valid

// Check to see if the config reference portion of the spec is valid

func validateEndpointSpec(epSpec *api.EndpointSpec) error {
	_ = "STUB: not implemented"
	// Endpoint spec is optional
	return nil
}

// Publish mode = "ingress" represents Routing-Mesh and current implementation
// of routing-mesh relies on IPVS based load-balancing with input=published-port.
// But Endpoint-Spec mode of DNSRR relies on multiple A records and cannot be used
// with routing-mesh (PublishMode="ingress") which cannot rely on DNSRR.
// But PublishMode="host" doesn't provide Routing-Mesh and the DNSRR is applicable
// for the backend network and hence we accept that configuration.

// If published port is not specified, it does not conflict
// with any others.

// validateSecretRefsSpec finds if the secrets passed in spec are valid and have no
// conflicting targets.
func validateSecretRefsSpec(spec api.TaskSpec) error { _ = "STUB: not implemented"; return nil }

// Keep a map to track all the targets that will be exposed
// The string returned is only used for logging. It could as well be struct{}{}

// SecretID and SecretName are mandatory, we have invalid references without them

// Every secret reference requires a Target

// If this is a file target, we will ensure filename uniqueness

// If this target is already in use, we have conflicting targets

// validateConfigRefsSpec finds if the configs passed in spec are valid and have no
// conflicting targets.
func validateConfigRefsSpec(spec api.TaskSpec) error { _ = "STUB: not implemented"; return nil }

// check if we're using a config as a CredentialSpec -- if so, we need to
// verify

// if there is no config in the credspec, then this will just be
// assigned to emptystring anyway, so we don't need to check
// existence.

// Keep a map to track all the targets that will be exposed
// The string returned is only used for logging. It could as well be struct{}{}

// ConfigID and ConfigName are mandatory, we have invalid references without them

// Every config reference requires a Target

// If this is a file target, we will ensure filename uniqueness

// Validate the file name

// If this target is already in use, we have conflicting targets

func (s *Server) validateNetworks(networks []*api.NetworkAttachmentConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMode(s *api.ServiceSpec) error { _ = "STUB: not implemented"; return nil }

// this check shouldn't be required as the point of uint64 is to
// constrain the possible values to positive numbers, but it almost
// certainly is required because there are almost certainly blind casts
// from int64 to uint64, and uint64(-1) is almost certain to crash the
// cluster because of how large it is.

func validateJob(spec *api.ServiceSpec) error { _ = "STUB: not implemented"; return nil }

func validateServiceSpec(spec *api.ServiceSpec) error { _ = "STUB: not implemented"; return nil }

// job-mode services are validated differently. most notably, they do not
// have UpdateConfigs, which is why this case statement skips update
// validation.

func isJobSpec(spec *api.ServiceSpec) bool { _ = "STUB: not implemented"; return false }

// checkPortConflicts does a best effort to find if the passed in spec has port
// conflicts with existing services.
// `serviceID string` is the service ID of the spec in service update. If
// `serviceID` is not "", then conflicts check will be skipped against this
// service (the service being updated).
func (s *Server) checkPortConflicts(spec *api.ServiceSpec, serviceID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Multiple services with same port in host publish mode can
// coexist - this is handled by the scheduler.

// If service ID is the same (and not "") then this is an update

// checkSecretExistence finds if the secret exists
func (s *Server) checkSecretExistence(tx store.Tx, spec *api.ServiceSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// Check to see if the secret exists and secretRef.SecretName matches the actual secretName

// checkConfigExistence finds if the config exists
func (s *Server) checkConfigExistence(tx store.Tx, spec *api.ServiceSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// Check to see if the config exists and configRef.ConfigName matches the actual configName

// CreateService creates and returns a Service based on the provided ServiceSpec.
// - Returns `InvalidArgument` if the ServiceSpec is malformed.
// - Returns `Unimplemented` if the ServiceSpec references unimplemented features.
// - Returns `AlreadyExists` if the ServiceID conflicts.
// - Returns an error if the creation fails.
func (s *Server) CreateService(_ context.Context, request *api.CreateServiceRequest) (*api.CreateServiceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(aluzzardi): Consider using `Name` as a primary key to handle
// duplicate creations. See #65

// Check to see if all the secrets being added exist as objects
// in our datastore

// Enhance the name-confict error to include the service name. The original
// `ErrNameConflict` error-message is included for backward-compatibility
// with older consumers of the API performing string-matching.

// GetService returns a Service given a ServiceID.
// - Returns `InvalidArgument` if ServiceID is not provided.
// - Returns `NotFound` if the Service is not found.
func (s *Server) GetService(_ context.Context, request *api.GetServiceRequest) (*api.GetServiceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateService updates a Service referenced by ServiceID with the given ServiceSpec.
// - Returns `NotFound` if the Service is not found.
// - Returns `InvalidArgument` if the ServiceSpec is malformed.
// - Returns `Unimplemented` if the ServiceSpec references unimplemented features.
// - Returns an error if the update fails.
func (s *Server) UpdateService(_ context.Context, request *api.UpdateServiceRequest) (*api.UpdateServiceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// It's not okay to update Service.Spec.Networks on its own.
// However, if Service.Spec.Task.Networks is also being
// updated, that's okay (for example when migrating from the
// deprecated Spec.Networks field to Spec.Task.Networks).

// Check to see if all the secrets being added exist as objects
// in our datastore

// orchestrator is designed to be stateless, so it should not deal
// with service mode change (comparing current config with previous config).
// proper way to change service mode is to delete and re-add.

// if the service has a JobStatus, that means it must be a Job, and we
// should increment the JobIteration

// Set spec version. Note that this will not match the
// service's Meta.Version after the store update. The
// versions for the spec and the service itself are not
// meant to be directly comparable.

// Reset update status

// RemoveService removes a Service referenced by ServiceID.
// - Returns `InvalidArgument` if ServiceID is not provided.
// - Returns `NotFound` if the Service is not found.
// - Returns an error if the deletion fails.
func (s *Server) RemoveService(_ context.Context, request *api.RemoveServiceRequest) (*api.RemoveServiceResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func filterServices(candidates []*api.Service, filters ...func(*api.Service) bool) []*api.Service {
	_ = "STUB: not implemented"
	return nil
}

// ListServices returns a list of all services.
func (s *Server) ListServices(_ context.Context, request *api.ListServicesRequest) (*api.ListServicesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListServiceStatuses returns a `ListServiceStatusesResponse` with the status
// of the requested services, formed by computing the number of running vs
// desired tasks. It is provided as a shortcut or helper method, which allows a
// client to avoid having to calculate this value by listing all Tasks.  If any
// service requested does not exist, it will be returned but with empty status
// values.
func (s *Server) ListServiceStatuses(_ context.Context, req *api.ListServiceStatusesRequest) (*api.ListServiceStatusesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// no matter what, add this status to the list.

// if there is another kind of error here (not sure what it
// could be) then still return 0/0 for this service.

// use a boolean to see global vs replicated. this avoids us having to
// iterate the task list twice.

// jobIteration is the iteration that the Job is currently
// operating on, to distinguish Tasks in old executions from tasks
// in the current one. if nil, service is not a Job

// a service might be deleted, but it may still have tasks. in that
// case, we will be using 0 as the desired task count.

// figure out how many tasks the service requires. for replicated
// services, this is easy: we can just check the replicas field. for
// global services, this is a bit harder and we'll need to do some
// numbercrunchin

// global applies to both GlobalJob and regular Global

// now, figure out how many tasks are running. Pretty easy, and
// universal across both global and replicated services

// if the service is a Job, jobIteration will be non-nil. This
// means we should check if the task belongs to the current job
// iteration. If not, skip accounting the task.

// additionally, since we've verified that the service is a
// job and the task belongs to this iteration, we should
// increment CompletedTasks

// if the service is global, a shortcut for figuring out the
// number of tasks desired is to look at all tasks, and take a
// count of the ones whose desired state is not Shutdown.

// for jobs, this is any task with desired state Completed
// which is not actually in that state.
