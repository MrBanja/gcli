package history

import (
	"gcli/cmd/service"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "history",
		Short: "Command for working with current conversation history.",
		Args:  cobra.NoArgs,
	}
	c.AddCommand(cmdClear(), cmdLast(), cmdAll())
	return c
}

func cmdClear() *cobra.Command {
	return &cobra.Command{
		Use:   "clear",
		Short: "Clear current conversation history.",
		Args:  cobra.NoArgs,
		Run:   service.S.HistoryClear().Run,
	}
}

func cmdLast() *cobra.Command {
	return &cobra.Command{
		Use:   "last",
		Short: "Shows last message of a current conversation.",
		Args:  cobra.NoArgs,
		Run:   service.S.HistoryLast().Run,
	}
}

func cmdAll() *cobra.Command {
	return &cobra.Command{
		Use:   "all",
		Short: "Shows All messages in conversation.",
		Args:  cobra.NoArgs,
		Run:   service.S.HistoryAll().Run,
	}
}
