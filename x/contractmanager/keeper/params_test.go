package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/petri-labs/petri/testutil/contractmanager/keeper"
	"github.com/petri-labs/petri/x/contractmanager/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ContractManagerKeeper(t, nil)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
