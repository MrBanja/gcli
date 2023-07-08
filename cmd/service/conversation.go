package service

import (
	"fmt"
	"gcli/internal/history"
	"gcli/internal/util"
	"github.com/MrBanja/openaiAPI"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

/////// CONVERSATION USE
////

type ConvUse struct {
	history *history.History
}

func (s *Service) ConvUse() *ConvUse {
	return &ConvUse{history: s.history}
}

func (c *ConvUse) Run(cmd *cobra.Command, args []string) {
	convIDS := c.history.Keys()
	currentConvID := viper.GetString("current_conversation_id")
	desiredConvID := args[0]

	if strings.ToLower(currentConvID) == strings.ToLower(desiredConvID) {
		fmt.Printf("Conversatino %s already set as current\n", desiredConvID)
		return
	}

	for _, convID := range convIDS {
		if strings.ToLower(convID) == strings.ToLower(desiredConvID) {
			fmt.Printf("Switching to conversation %s\n", convID)
			viper.Set("current_conversation_id", convID)
			viper.WriteConfig()
			return
		}
	}

	fmt.Printf("Creating new conversation %s\n", desiredConvID)
	viper.Set("current_conversation_id", desiredConvID)
	err := viper.WriteConfig()
	util.HandleError(err, "Config write error")
	err = c.history.Set(desiredConvID, []openaiAPI.Message{})
	util.HandleError(err, "History write error")
}

/////// CONVERSATION LIST
////

type ConvList struct {
	history *history.History
}

func (s *Service) ConvList() *ConvList {
	return &ConvList{history: s.history}
}

func (c *ConvList) Run(cmd *cobra.Command, args []string) {
	convIDS := c.history.Keys()
	currentConvID := viper.GetString("current_conversation_id")
	fmt.Println("Current conversation ID:", currentConvID)
	fmt.Println("Conversation IDs:")
	for _, convID := range convIDS {
		if strings.ToLower(convID) == strings.ToLower(currentConvID) {
			fmt.Printf("* %s\n", convID)
		} else {
			fmt.Printf("  %s\n", convID)
		}
	}
}

/////// CONVERSATION DELETE
////

type ConvDel struct {
	history *history.History
}

func (s *Service) ConvDel() *ConvDel {
	return &ConvDel{history: s.history}
}

func (c *ConvDel) Run(cmd *cobra.Command, args []string) {
	currentConvID := viper.GetString("current_conversation_id")
	err := c.history.Delete(currentConvID)
	util.HandleError(err, "History delete error")
	fmt.Printf("Conversation %s deleted\n", currentConvID)

	convIDS := c.history.Keys()
	if len(convIDS) == 0 {
		fmt.Println("No conversations left")
		fmt.Printf("Creating new empty conversation %s\n", currentConvID)
		viper.Set("current_conversation_id", currentConvID)
		err := viper.WriteConfig()
		util.HandleError(err, "Config write error")
		err = c.history.Set(currentConvID, []openaiAPI.Message{})
		util.HandleError(err, "History write error")
		fmt.Println("Created new empty conversation")
		return
	}

	currentConvID = convIDS[0]
	viper.Set("current_conversation_id", currentConvID)
	util.HandleError(viper.WriteConfig(), "Config write error")

	fmt.Println("Switching to conversation ID:", currentConvID)
}
