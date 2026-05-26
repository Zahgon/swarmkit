package watchapi

import (
	"github.com/moby/swarmkit/v2/api"
)

// Watch starts a stream that returns any changes to objects that match
// the specified selectors. When the stream begins, it immediately sends
// an empty message back to the client. It is important to wait for
// this message before taking any actions that depend on an established
// stream of changes for consistency.
func (s *Server) Watch(request *api.WatchRequest, stream api.Watch_WatchServer) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(aaronl): Send current version in this WatchMessage?
