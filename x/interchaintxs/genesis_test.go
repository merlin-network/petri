package interchaintxs_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/petri-labs/petri/testutil/interchaintxs/keeper"
	"github.com/petri-labs/petri/testutil/interchaintxs/nullify"
	"github.com/petri-labs/petri/x/interchaintxs"
	"github.com/petri-labs/petri/x/interchaintxs/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx := keepertest.InterchainTxsKeeper(t, nil, nil, nil, nil, nil)
	interchaintxs.InitGenesis(ctx, *k, genesisState)
	got := interchaintxs.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
