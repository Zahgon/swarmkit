package plugin

import (
	google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

// DeepcopyEnabled returns true if deepcopy is enabled for the descriptor.
func DeepcopyEnabled(options *google_protobuf.MessageOptions) bool {
	_ = "STUB: not implemented"
	return false
}
