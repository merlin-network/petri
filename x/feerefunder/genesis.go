package feerefunder

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/petri-labs/petri/x/feerefunder/keeper"
	"github.com/petri-labs/petri/x/feerefunder/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)

	for _, info := range genState.FeeInfos {
		k.StoreFeeInfo(ctx, info)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.FeeInfos = k.GetAllFeeInfos(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
