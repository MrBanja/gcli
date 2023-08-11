package cmd

import (
	"fmt"

	"gcli/cmd/conf"
	"gcli/cmd/conversation"
	"gcli/cmd/version"
	"gcli/internal/controller"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:      "gcli [query]",
	Short:    "Ask ChatGPT",
	Args:     cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE:     runE,
	PostRunE: preRunE,
}

func runE(cmd *cobra.Command, args []string) error {
	return controller.Stream(args[0])
}

func preRunE(cmd *cobra.Command, args []string) error {
	if t := viper.GetString("openai.token"); t == "" {
		return fmt.Errorf("OpenAI token not found. You can set it by executing `gcli config set openai.token [your token]`")
	}
	return nil
}

func Execute() error {
	rootCmd.AddCommand(conversation.Cmd(), conf.Cmd(), version.Cmd())
	return rootCmd.Execute()
}
