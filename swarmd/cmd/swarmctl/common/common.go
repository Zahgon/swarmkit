package common

import (
	"context"

	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
)

// Dial establishes a connection and creates a client.
// It infers connection parameters from CLI options.
func Dial(cmd *cobra.Command) (api.ControlClient, error) {
	_ = "STUB: not implemented"
	return *new(api.ControlClient), nil
}

// DialConn establishes a connection to SwarmKit.
func DialConn(cmd *cobra.Command) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Context returns a request context based on CLI arguments.
func Context(_ *cobra.Command) context.Context {
	_ = "STUB: not implemented"
	// TODO(aluzzardi): Actually create a context.
	return *new(context.Context)
}

// ParseLogDriverFlags parses a silly string format for log driver and options.
// Fully baked log driver config should be returned.
//
// If no log driver is available, nil, nil will be returned.
func ParseLogDriverFlags(flags *pflag.FlagSet) (*api.Driver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
