package version

import (
	"io"
)

// FprintVersion outputs the version string to the writer, in the following
// format, followed by a newline:
//
//	<cmd> <project> <version>
//
// For example, a binary "registry" built from github.com/docker/distribution
// with version "v2.0" would print the following:
//
//	registry github.com/docker/distribution v2.0
func FprintVersion(w io.Writer) { _ = "STUB: not implemented"; return }

// PrintVersion outputs the version information, from Fprint, to stdout.
func PrintVersion() { _ = "STUB: not implemented"; return }
