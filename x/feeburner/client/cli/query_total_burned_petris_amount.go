package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/petri-labs/petri/x/feeburner/types"
	"github.com/spf13/cobra"
)

func CmdQueryTotalBurnedPetrisAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-burned-petris-amount",
		Short: "shows total amount of burned petris",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.TotalBurnedPetrisAmount(context.Background(), &types.QueryTotalBurnedPetrisAmountRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
