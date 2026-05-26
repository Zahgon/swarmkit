package testutils

import (
	"google.golang.org/grpc/codes"
)

// ErrorDesc returns the error description of err if it was produced by the rpc system.
// Otherwise, it returns err.Error() or empty string when err is nil.
func ErrorDesc(err error) string { _ = "STUB: not implemented"; return "" }

// ErrorCode returns the error code for err if it was produced by the rpc system.
// Otherwise, it returns codes.Unknown.
func ErrorCode(err error) codes.Code { _ = "STUB: not implemented"; return *new(codes.Code) }
