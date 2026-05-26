//go:build windows
// +build windows

package xnet

import (
	"net"
	"time"
)

// ListenLocal opens a local socket for control communication
func ListenLocal(socket string) (net.Listener, error) {
	_ = "STUB: not implemented"
	// set up ACL for the named pipe
	// allow Administrators and SYSTEM
	return *new(net.Listener), nil
}

// Use message mode so that CloseWrite() is supported
// Use 64KB buffers to improve performance

// on windows, our socket is actually a named pipe

// DialTimeoutLocal is a DialTimeout function for local sockets
func DialTimeoutLocal(socket string, timeout time.Duration) (net.Conn, error) {
	_ = "STUB: not implemented"
	// On windows, we dial a named pipe
	return *new(net.Conn), nil
}
