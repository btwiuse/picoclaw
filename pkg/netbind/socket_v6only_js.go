//go:build js

package netbind

import "syscall"

func applyIPv6OnlyControl(enabled bool) func(string, string, syscall.RawConn) error {
	return nil
}
