package contractmanager_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/merlin-network/petri/testutil/contractmanager/keeper"
	"github.com/merlin-network/petri/testutil/contractmanager/nullify"
	"github.com/merlin-network/petri/x/contractmanager"
	"github.com/merlin-network/petri/x/contractmanager/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		FailuresList: []types.Failure{
			{
				Address: "address1",
				Id:      1,
			},
			{
				Address: "address1",
				Id:      2,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ContractManagerKeeper(t, nil)
	contractmanager.InitGenesis(ctx, *k, genesisState)
	got := contractmanager.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.FailuresList, got.FailuresList)
	// this line is used by starport scaffolding # genesis/test/assert
}
