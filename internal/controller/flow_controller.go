package controller

import (
	"context"
	"fmt"
	"gcli/internal/history"
	"gcli/internal/util"
	"github.com/MrBanja/openaiAPI"
	"github.com/inancgumus/screen"
	"github.com/spf13/viper"
)

type FlowController struct {
	openai  *openaiAPI.OpenAI
	history *history.History
}

func New(openai *openaiAPI.OpenAI, history *history.History) *FlowController {
	return &FlowController{
		openai:  openai,
		history: history,
	}
}

func (f *FlowController) Stream(prompt string) error {
	convID := viper.GetString("current_conversation_id")
	messages, err := f.history.Get(convID)
	if err != nil {
		return err
	}

	resp := f.openai.SendWithStream(context.Background(), prompt, messages)
	messageContent := ""

L:
	for {
		select {
		case err := <-resp.Error():
			util.HandleError(err, "OpenAI response error")
		case msg, ok := <-resp.Data():
			if !ok {
				break L
			}
			fmt.Print(msg)
			messageContent += msg
		}
	}

	screen.Clear()
	screen.MoveTopLeft()
	util.PrintOrExit("# RESPONSE")
	util.PrintOrExit(messageContent)

	messages = append(
		messages,
		openaiAPI.Message{Role: openaiAPI.RoleUser, Content: prompt},
		openaiAPI.Message{Role: openaiAPI.RoleAssistant, Content: messageContent},
	)
	if err := f.history.Set(convID, messages); err != nil {
		return nil
	}
	return nil
}
