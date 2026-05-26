//go:build !windows
// +build !windows

package xnet

import (
	"net"
	"time"
)

// ListenLocal opens a local socket for control communication
func ListenLocal(socket string) (net.Listener, error) {
	_ = "STUB: not implemented"
	// on unix it's just a unix socket
	return *new(net.Listener), nil
}

// DialTimeoutLocal is a DialTimeout function for local sockets
func DialTimeoutLocal(socket string, timeout time.Duration) (net.Conn, error) {
	_ = "STUB: not implemented"
	// on unix, we dial a unix socket
	return *new(net.Conn), nil
}
