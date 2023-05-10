package keeper_test

import (
	"testing"

	"github.com/merlin-network/petri/app"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/merlin-network/petri/testutil/feeburner/keeper"
	"github.com/merlin-network/petri/x/feeburner/types"
)

func TestParamsQuery(t *testing.T) {
	_ = app.GetDefaultConfig()

	keeper, ctx := testkeeper.FeeburnerKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
