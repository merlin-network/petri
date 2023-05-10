package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	feekeeperutil "github.com/petri-labs/petri/testutil/feeburner/keeper"
	"github.com/petri-labs/petri/x/feeburner/types"
)

func TestGrpcQuery_TotalBurnedPetrisAmount(t *testing.T) {
	feeKeeper, ctx := feekeeperutil.FeeburnerKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)

	feeKeeper.RecordBurnedFees(ctx, sdk.NewCoin(types.DefaultPetriDenom, sdk.NewInt(100)))

	request := types.QueryTotalBurnedPetrisAmountRequest{}
	response, err := feeKeeper.TotalBurnedPetrisAmount(wctx, &request)
	require.NoError(t, err)
	require.Equal(t, &types.QueryTotalBurnedPetrisAmountResponse{TotalBurnedPetrisAmount: types.TotalBurnedPetrisAmount{Coin: sdk.Coin{Denom: types.DefaultPetriDenom, Amount: sdk.NewInt(100)}}}, response)
}
