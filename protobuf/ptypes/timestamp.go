package ptypes

import (
	"time"

	gogotypes "github.com/gogo/protobuf/types"
)

// MustTimestampProto converts time.Time to a google.protobuf.Timestamp proto.
// It panics if input timestamp is invalid.
func MustTimestampProto(t time.Time) *gogotypes.Timestamp { _ = "STUB: not implemented"; return nil }
