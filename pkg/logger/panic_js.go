//go:build js

package logger

import (
	"io"
)

func initPanicFile(panicFile string) io.WriteCloser {
	return nopWriteCloser{}
}

type nopWriteCloser struct{}

func (nopWriteCloser) Write(p []byte) (int, error) { return len(p), nil }
func (nopWriteCloser) Close() error                { return nil }
