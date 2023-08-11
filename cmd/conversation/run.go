package conversation

import (
	"fmt"
	"strings"

	"gcli/cmd/conversation/history"
	hstr "gcli/internal/history"
	"github.com/MrBanja/openaiAPI"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cmd = &cobra.Command{
		Use:   "conv",
		Short: "Command for working with conversations.",
		Args:  cobra.NoArgs,
	}

	cmdList = &cobra.Command{
		Use:   "ls",
		Short: "Display all conversations.",
		Args:  cobra.NoArgs,
		RunE:  listRunE,
	}

	cmdUse = &cobra.Command{
		Use:   "use [conversation id]",
		Short: "Use a conversation with given id. If the conversation does not exist, it will be created.",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE:  useRunE,
	}

	cmdDelete = &cobra.Command{
		Use:   "del",
		Short: "Delete current conversation.",
		Args:  cobra.NoArgs,
		RunE:  deleteRunE,
	}
)

func Cmd() *cobra.Command {
	cmd.AddCommand(cmdUse, cmdList, cmdDelete, history.Cmd())
	return cmd
}

func listRunE(cmd *cobra.Command, args []string) error {
	convIDS := hstr.Keys()
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
	return nil
}

func useRunE(cmd *cobra.Command, args []string) error {
	convIDS := hstr.Keys()
	currentConvID := viper.GetString("current_conversation_id")
	desiredConvID := args[0]

	if strings.ToLower(currentConvID) == strings.ToLower(desiredConvID) {
		fmt.Printf("Conversatino %s already set as current\n", desiredConvID)
		return nil
	}

	for _, convID := range convIDS {
		if strings.ToLower(convID) == strings.ToLower(desiredConvID) {
			fmt.Printf("Switching to conversation %s\n", convID)
			viper.Set("current_conversation_id", convID)
			return viper.WriteConfig()
		}
	}

	fmt.Printf("Creating new conversation %s\n", desiredConvID)
	viper.Set("current_conversation_id", desiredConvID)
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("config write error %w", err)
	}
	if err := hstr.Set(desiredConvID, []openaiAPI.Message{}); err != nil {
		return fmt.Errorf("history write error %w", err)
	}
	return nil
}

func deleteRunE(cmd *cobra.Command, args []string) error {
	currentConvID := viper.GetString("current_conversation_id")
	if err := hstr.Delete(currentConvID); err != nil {
		return fmt.Errorf("history delete error %w", err)
	}

	fmt.Printf("Conversation %s deleted\n", currentConvID)

	convIDS := hstr.Keys()
	if len(convIDS) == 0 {
		fmt.Println("No conversations left")
		fmt.Printf("Creating new empty conversation %s\n", currentConvID)

		viper.Set("current_conversation_id", currentConvID)

		if err := viper.WriteConfig(); err != nil {
			return fmt.Errorf("config write error %w", err)
		}
		if err := hstr.Set(currentConvID, []openaiAPI.Message{}); err != nil {
			return fmt.Errorf("history write error %w", err)
		}

		fmt.Println("Created new empty conversation")
		return nil
	}

	currentConvID = convIDS[0]
	viper.Set("current_conversation_id", currentConvID)
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("config write error %w", err)
	}

	fmt.Println("Switching to conversation ID:", currentConvID)
	return nil
}
