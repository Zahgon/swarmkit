package cli

import (
	"github.com/moby/swarmkit/v2/api"
)

// ExternalCAOpt is a Value type for parsing external CA specifications.
type ExternalCAOpt struct {
	values []*api.ExternalCA
}

// Set parses an external CA option.
func (m *ExternalCAOpt) Set(value string) error { _ = "STUB: not implemented"; return nil }

// Type returns the type of this option.
func (m *ExternalCAOpt) Type() string { _ = "STUB: not implemented"; return "" }

// String returns a string repr of this option.
func (m *ExternalCAOpt) String() string { _ = "STUB: not implemented"; return "" }

// Value returns the external CAs
func (m *ExternalCAOpt) Value() []*api.ExternalCA {
	_ = "STUB: not implemented"

	// parseExternalCA parses an external CA specification from the command line,
	// such as protocol=cfssl,url=https://example.com.
	return nil
}

func parseExternalCA(caSpec string) (*api.ExternalCA, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
