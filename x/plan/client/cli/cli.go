package cli

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCommands(cdc *codec.Codec) []*cobra.Command {
	return flags.GetCommands(
		queryPlan(cdc),
		queryPlans(cdc),
	)
}

func GetTxCommands(cdc *codec.Codec) []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "plan",
		Short: "Plan module sub-commands",
	}

	cmd.AddCommand(flags.PostCommands(
		txAdd(cdc),
		txSetStatus(cdc),
		txAddNode(cdc),
		txRemoveNode(cdc),
	)...)

	return []*cobra.Command{cmd}
}
