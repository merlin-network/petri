package interchaintxs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/petri-labs/petri/x/interchaintxs/keeper"
	"github.com/petri-labs/petri/x/interchaintxs/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	return genesis
}
