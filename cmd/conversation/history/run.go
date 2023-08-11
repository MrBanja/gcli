package history

import (
	"fmt"

	"gcli/internal/history"

	"gcli/internal/util"
	"github.com/MrBanja/openaiAPI"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cmd = &cobra.Command{
		Use:   "history",
		Short: "Command for working with current conversation history.",
		Args:  cobra.NoArgs,
	}

	cmdClear = &cobra.Command{
		Use:   "clear",
		Short: "Clear current conversation history.",
		Args:  cobra.NoArgs,
		RunE:  clearRunE,
	}

	cmdLast = &cobra.Command{
		Use:   "last",
		Short: "Shows last message of a current conversation.",
		Args:  cobra.NoArgs,
		RunE:  lastRunE,
	}

	cmdAll = &cobra.Command{
		Use:   "all",
		Short: "Shows All messages in conversation.",
		Args:  cobra.NoArgs,
		RunE:  allRunE,
	}
)

func Cmd() *cobra.Command {
	cmd.AddCommand(cmdClear, cmdLast, cmdAll)
	return cmd
}

func clearRunE(cmd *cobra.Command, args []string) error {
	currentConvID := viper.GetString("current_conversation_id")
	if err := history.Set(currentConvID, []openaiAPI.Message{}); err != nil {
		return fmt.Errorf("history error: %w", err)
	}
	fmt.Println("History cleared")
	return nil
}

func lastRunE(cmd *cobra.Command, args []string) error {
	currentConvID := viper.GetString("current_conversation_id")
	msgs, err := history.Get(currentConvID)
	if err != nil {
		return fmt.Errorf("history get error: %w", err)
	}
	if len(msgs) == 0 {
		fmt.Println("No messages in conversation")
		return nil
	}

	lastMsg := msgs[len(msgs)-1]
	fmt.Println("Current conversation ID:", currentConvID)
	fmt.Printf("Last message in conversation from %s\n====\n:", lastMsg.Role)
	if err := util.GPrint(lastMsg.Content); err != nil {
		return err
	}
	return nil
}

func allRunE(cmd *cobra.Command, args []string) error {
	currentConvID := viper.GetString("current_conversation_id")
	msgs, err := history.Get(currentConvID)
	if err != nil {
		return fmt.Errorf("history get error: %w", err)
	}
	if len(msgs) == 0 {
		fmt.Println("No messages in conversation")
		return nil
	}

	fmt.Println("Current conversation ID:", currentConvID)

	for _, msg := range msgs {
		if err := util.GPrint(fmt.Sprintf("#### %s\n%s\n", msg.Role, msg.Content)); err != nil {
			return err
		}
	}
	return nil
}
