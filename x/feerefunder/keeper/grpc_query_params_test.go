package keeper_test

import (
	"testing"

	testkeeper "github.com/merlin-network/petri/testutil/feerefunder/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/merlin-network/petri/x/feerefunder/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.FeeKeeper(t, nil, nil)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
