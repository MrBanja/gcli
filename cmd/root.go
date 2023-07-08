package cmd

import (
	"fmt"
	"gcli/cmd/conf"
	"gcli/cmd/conversation"
	"gcli/cmd/service"
	"gcli/internal/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gcli [query]",
		Short: "Ask ChatGPT",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	}
)

func Execute() error {
	rootCmd.Run = service.S.Ask().Run
	rootCmd.PreRun = service.S.Ask().PreRun
	return rootCmd.Execute()
}

func init() {
	initConfig()

	home, err := os.UserHomeDir()
	util.HandleError(err, "HomeDir not found")

	viper.SetDefault("history.path", fmt.Sprintf("%s/.config/gcli/history.json", home))
	viper.SetDefault("current_conversation_id", "init")

	service.S = service.New()
	util.HandleError(service.S.Init(), "Init error")
	util.HandleError(viper.WriteConfig(), "Config writing error")
	rootCmd.AddCommand(conversation.Cmd(), conf.Cmd)
}

func initConfig() {
	home, err := os.UserHomeDir()
	util.HandleError(err, "HomeDir not found")

	pathToConfigDir := fmt.Sprintf("%s/.config/gcli", home)
	err = os.MkdirAll(pathToConfigDir, os.ModePerm)
	util.HandleError(err, "Can not create conf path")

	pathToConfig := fmt.Sprintf("%s/config.yaml", pathToConfigDir)

	_, err = os.Stat(pathToConfig)
	if err != nil {
		f, err := os.Create(pathToConfig)
		f.Close()
		util.HandleError(err, "Config creation error")
		fmt.Println("Config file created:", pathToConfig)
	}

	viper.SetConfigFile(pathToConfig)
	if err := viper.ReadInConfig(); err != nil {
		util.HandleError(err, "Config parsing error")
	}
}
