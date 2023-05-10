package interchainqueries_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/petri-labs/petri/testutil/interchainqueries/keeper"
	"github.com/petri-labs/petri/testutil/interchainqueries/nullify"
	"github.com/petri-labs/petri/x/interchainqueries"
	"github.com/petri-labs/petri/x/interchainqueries/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		RegisteredQueries: []*types.RegisteredQuery{
			{
				Id: 4,
			},
			{
				Id: 3,
			},
			{
				Id: 2,
			},
			{
				Id: 1,
			},
		},
	}

	require.EqualValues(t, genesisState.Params, types.DefaultParams())

	k, ctx := keepertest.InterchainQueriesKeeper(t, nil, nil, nil, nil)
	interchainqueries.InitGenesis(ctx, *k, genesisState)
	got := interchainqueries.ExportGenesis(ctx, *k)
	lastQueryID := k.GetLastRegisteredQueryKey(ctx)

	require.EqualValues(t, got.Params, types.DefaultParams())
	require.NotNil(t, got)
	require.EqualValues(t, 4, lastQueryID)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RegisteredQueries, got.RegisteredQueries)
}

func TestGenesisNullQueries(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx := keepertest.InterchainQueriesKeeper(t, nil, nil, nil, nil)
	interchainqueries.InitGenesis(ctx, *k, genesisState)
	got := interchainqueries.ExportGenesis(ctx, *k)

	require.ElementsMatch(t, genesisState.RegisteredQueries, got.RegisteredQueries)
}
