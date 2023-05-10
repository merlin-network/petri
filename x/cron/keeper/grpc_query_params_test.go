package keeper_test

import (
	"testing"

	testkeeper "github.com/petri-labs/petri/testutil/cron/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/petri-labs/petri/x/cron/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.CronKeeper(t, nil, nil)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
