//go:build !js

package tools

import (
	"os"

	"github.com/creack/pty"
)

func ptyOpen() (*os.File, *os.File, error) {
	return pty.Open()
}
