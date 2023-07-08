package service

import (
	"gcli/internal/controller"
	"gcli/internal/history"
	"gcli/internal/util"
	"github.com/MrBanja/openaiAPI"
	"github.com/spf13/viper"
	"time"
)

type Service struct {
	client         *openaiAPI.OpenAI
	history        *history.History
	flowController *controller.FlowController
}

var S *Service

func New() *Service {
	h, err := history.New(viper.GetString("history.path"))
	util.HandleError(err, "History init error")
	client := openaiAPI.New(viper.GetString("openai.token"), openaiAPI.Model4, 60*time.Second)
	f := controller.New(client, h)
	return &Service{client: client, history: h, flowController: f}
}

func (s *Service) Init() error {
	currentConvID := viper.GetString("current_conversation_id")
	convIDS := s.history.Keys()
	if len(convIDS) == 0 {
		err := s.history.Set(currentConvID, []openaiAPI.Message{})
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}
