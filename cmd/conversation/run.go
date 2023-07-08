package conversation

import (
	"gcli/cmd/conversation/history"
	"gcli/cmd/service"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "conv",
		Short: "Command for working with conversations.",
		Args:  cobra.NoArgs,
	}
	c.AddCommand(cmdUse(), cmdList(), cmdDelete(), history.Cmd())
	return c
}

func cmdList() *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "Display all conversations.",
		Args:  cobra.NoArgs,
		Run:   service.S.ConvList().Run,
	}
}

func cmdUse() *cobra.Command {
	return &cobra.Command{
		Use:   "use [conversation id]",
		Short: "Use a conversation with given id. If the conversation does not exist, it will be created.",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run:   service.S.ConvUse().Run,
	}
}

func cmdDelete() *cobra.Command {
	return &cobra.Command{
		Use:   "del",
		Short: "Delete current conversation.",
		Args:  cobra.NoArgs,
		Run:   service.S.ConvDel().Run,
	}
}
