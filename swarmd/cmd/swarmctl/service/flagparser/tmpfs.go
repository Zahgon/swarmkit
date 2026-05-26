package flagparser

import (
	"github.com/moby/swarmkit/v2/api"
	"github.com/spf13/pflag"
)

// parseTmpfs supports a simple tmpfs decl, similar to docker run.
//
// This should go away.
func parseTmpfs(flags *pflag.FlagSet, spec *api.ServiceSpec) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(stevvooe): Nasty inline parsing code, replace with mount syntax.

// repeated colon is illegal

// BUG(stevvooe): Cobra stringslice actually doesn't correctly
// handle comma separated values, so multiple flags aren't
// really supported. We'll have to replace StringSlice with a
// type that doesn't use the csv parser. This is good enough
// for now.

// try to parse this into bytes

// remove suffix and try again

// reparse the meat
