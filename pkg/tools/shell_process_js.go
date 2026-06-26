//go:build js

package tools

import "os/exec"

func prepareCommandForTermination(cmd *exec.Cmd) {
}

func terminateProcessTree(cmd *exec.Cmd) error {
	return nil
}
