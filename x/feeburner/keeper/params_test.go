package keeper_test

import (
	"testing"

	"github.com/merlin-network/petri/app"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/merlin-network/petri/testutil/feeburner/keeper"
	"github.com/merlin-network/petri/x/feeburner/types"
)

func TestGetParams(t *testing.T) {
	_ = app.GetDefaultConfig()

	k, ctx := testkeeper.FeeburnerKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
