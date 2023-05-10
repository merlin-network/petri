package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/merlin-network/petri/x/feeburner/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) TotalBurnedPetrisAmount(goCtx context.Context, _ *types.QueryTotalBurnedPetrisAmountRequest) (*types.QueryTotalBurnedPetrisAmountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	totalBurnedPetrisAmount := k.GetTotalBurnedPetrisAmount(ctx)

	return &types.QueryTotalBurnedPetrisAmountResponse{TotalBurnedPetrisAmount: totalBurnedPetrisAmount}, nil
}
