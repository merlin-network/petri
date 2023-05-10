package feeburner_test

import (
	"testing"

	"github.com/merlin-network/petri/app"

	"github.com/merlin-network/petri/testutil/feeburner/keeper"
	"github.com/merlin-network/petri/testutil/feeburner/nullify"
	"github.com/merlin-network/petri/x/feeburner"
	"github.com/merlin-network/petri/x/feeburner/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	_ = app.GetDefaultConfig()

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keeper.FeeburnerKeeper(t)
	feeburner.InitGenesis(ctx, *k, genesisState)
	got := feeburner.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
