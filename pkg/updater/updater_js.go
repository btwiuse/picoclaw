//go:build js

package updater

import (
	"errors"

	"github.com/spf13/cobra"
)

func NewUpdateCommand(binaryName string) *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Check and apply updates from GitHub releases",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("update is not supported on js/wasm")
		},
	}
}
