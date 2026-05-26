package controlapi

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
)

// GetTask returns a Task given a TaskID.
// - Returns `InvalidArgument` if TaskID is not provided.
// - Returns `NotFound` if the Task is not found.
func (s *Server) GetTask(_ context.Context, request *api.GetTaskRequest) (*api.GetTaskResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoveTask removes a Task referenced by TaskID.
// - Returns `InvalidArgument` if TaskID is not provided.
// - Returns `NotFound` if the Task is not found.
// - Returns an error if the deletion fails.
func (s *Server) RemoveTask(_ context.Context, request *api.RemoveTaskRequest) (*api.RemoveTaskResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func filterTasks(candidates []*api.Task, filters ...func(*api.Task) bool) []*api.Task {
	_ = "STUB: not implemented"
	return nil
}

// ListTasks returns a list of all tasks.
func (s *Server) ListTasks(_ context.Context, request *api.ListTasksRequest) (*api.ListTasksResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
