//go:build js

package tools

import (
	"errors"
	"os"
)

func ptyOpen() (*os.File, *os.File, error) {
	return nil, nil, errors.New("PTY is not supported on js/wasm")
}
