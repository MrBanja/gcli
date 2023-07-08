package conf

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Cmd = &cobra.Command{
		Use:   "config",
		Short: "Set configuration",
		Args:  cobra.NoArgs,
	}
	CmdSet = &cobra.Command{
		Use:   "set [key] [value]",
		Short: "Set config",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			viper.Set(args[0], args[1])
			fmt.Printf("%s is set\n", args[0])
			viper.WriteConfig()
		},
	}
)

func init() {
	Cmd.AddCommand(CmdSet)
}
