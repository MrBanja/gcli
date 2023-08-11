package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string

var cmd = &cobra.Command{
	Use:   "version",
	Short: "Display current version",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Version: %s\n", Version)
		return nil
	},
}

func Cmd() *cobra.Command {
	return cmd
}
