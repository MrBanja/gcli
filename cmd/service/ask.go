package service

import (
	"fmt"
	"gcli/internal/controller"
	"gcli/internal/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type Ask struct {
	flowController *controller.FlowController
}

func (s *Service) Ask() *Ask {
	return &Ask{flowController: s.flowController}
}

func (a *Ask) PreRun(cmd *cobra.Command, args []string) {
	if t := viper.GetString("openai.token"); t == "" {
		fmt.Fprintln(os.Stderr, "OpenAI token not found. You can set it by executing `gcli config set openai.token [your token]`")
		os.Exit(1)
	}
}

func (a *Ask) Run(cmd *cobra.Command, args []string) {
	err := a.flowController.Stream(args[0])
	util.HandleError(err, "Controller error")
}
