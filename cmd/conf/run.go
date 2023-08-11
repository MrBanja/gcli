package conf

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cmd = &cobra.Command{
		Use:   "config",
		Short: "Set configuration",
		Args:  cobra.NoArgs,
	}
	cmdSet = &cobra.Command{
		Use:   "set [key] [value]",
		Short: "Set config",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			viper.Set(args[0], args[1])
			fmt.Printf("%s is set\n", args[0])
			return viper.WriteConfig()
		},
	}
)

func Cmd() *cobra.Command {
	cmd.AddCommand(cmdSet)
	return cmd
}
