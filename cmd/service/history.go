package service

import (
	"fmt"
	"gcli/internal/history"
	"gcli/internal/util"
	"github.com/MrBanja/openaiAPI"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/////// HISTORY CLEAR
////

type HistoryClear struct {
	history *history.History
}

func (s *Service) HistoryClear() *HistoryClear {
	return &HistoryClear{history: s.history}
}

func (h *HistoryClear) Run(cmd *cobra.Command, args []string) {
	currentConvID := viper.GetString("current_conversation_id")
	err := h.history.Set(currentConvID, []openaiAPI.Message{})
	util.HandleError(err, "History error")
	fmt.Println("History cleared")
}

/////// HISTORY LAST
////

type HistoryLast struct {
	history *history.History
}

func (s *Service) HistoryLast() *HistoryLast {
	return &HistoryLast{history: s.history}
}

func (c *HistoryLast) Run(cmd *cobra.Command, args []string) {
	currentConvID := viper.GetString("current_conversation_id")
	msgs, err := c.history.Get(currentConvID)
	util.HandleError(err, "History get error")
	if len(msgs) == 0 {
		fmt.Println("No messages in conversation")
		return
	}

	lastMsg := msgs[len(msgs)-1]
	fmt.Println("Current conversation ID:", currentConvID)
	fmt.Printf("Last message in conversation from %s\n====\n:", lastMsg.Role)
	util.PrintOrExit(lastMsg.Content)
}

/////// HISTORY ALL
////

type HistoryAll struct {
	history *history.History
}

func (s *Service) HistoryAll() *HistoryAll {
	return &HistoryAll{history: s.history}
}

func (c *HistoryAll) Run(cmd *cobra.Command, args []string) {
	currentConvID := viper.GetString("current_conversation_id")
	msgs, err := c.history.Get(currentConvID)
	util.HandleError(err, "History get error")
	if len(msgs) == 0 {
		fmt.Println("No messages in conversation")
		return
	}

	fmt.Println("Current conversation ID:", currentConvID)

	for _, msg := range msgs {
		util.PrintOrExit(fmt.Sprintf("#### %s\n%s\n", msg.Role, msg.Content))
	}
}
