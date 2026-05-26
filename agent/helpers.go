package agent

import "context"

// runctx blocks until the function exits, closed is closed, or the context is
// cancelled. Call as part of go statement.
func runctx(ctx context.Context, closed chan struct{}, errs chan error, fn func(ctx context.Context) error) {
	_ = "STUB: not implemented"
	return
}
