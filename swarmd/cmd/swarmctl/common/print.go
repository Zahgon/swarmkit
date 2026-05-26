package common

import (
	"io"

	gogotypes "github.com/gogo/protobuf/types"
)

// PrintHeader prints a nice little header.
func PrintHeader(w io.Writer, columns ...string) { _ = "STUB: not implemented"; return }

// FprintfIfNotEmpty prints only if `s` is not empty.
//
// NOTE(stevvooe): Not even remotely a printf function.. doesn't take args.
func FprintfIfNotEmpty(w io.Writer, format string, v interface{}) {
	_ = "STUB: not implemented"
	return
}

// TimestampAgo returns a relative time string from a timestamp (e.g. "12 seconds ago").
func TimestampAgo(ts *gogotypes.Timestamp) string { _ = "STUB: not implemented"; return "" }
