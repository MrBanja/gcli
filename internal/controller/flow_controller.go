package controller

import (
	"context"
	"fmt"
	"time"

	"gcli/internal/history"
	"gcli/internal/util"
	"github.com/MrBanja/openaiAPI"
	"github.com/inancgumus/screen"
	"github.com/spf13/viper"
)

var openai *openaiAPI.OpenAI

func InitFlowController() error {
	if openai != nil {
		return fmt.Errorf("flow controller already initialized")
	}
	openai = openaiAPI.New(viper.GetString("openai.token"), openaiAPI.Model4, 60*time.Second)
	return nil
}

func Stream(prompt string) error {
	convID := viper.GetString("current_conversation_id")
	messages, err := history.Get(convID)
	if err != nil {
		return err
	}

	resp := openai.SendWithStream(context.Background(), prompt, messages)
	messageContent := ""

L:
	for {
		select {
		case err := <-resp.Error():
			return fmt.Errorf("OpenAI response error %w", err)
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

	if err := util.GPrint("# RESPONSE"); err != nil {
		return err
	}
	if err := util.GPrint(messageContent); err != nil {
		return err
	}

	messages = append(
		messages,
		openaiAPI.Message{Role: openaiAPI.RoleUser, Content: prompt},
		openaiAPI.Message{Role: openaiAPI.RoleAssistant, Content: messageContent},
	)
	if err := history.Set(convID, messages); err != nil {
		return nil
	}
	return nil
}
