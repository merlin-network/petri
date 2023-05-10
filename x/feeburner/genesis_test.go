package feeburner_test

import (
	"testing"

	"github.com/petri-labs/petri/app"

	"github.com/petri-labs/petri/testutil/feeburner/keeper"
	"github.com/petri-labs/petri/testutil/feeburner/nullify"
	"github.com/petri-labs/petri/x/feeburner"
	"github.com/petri-labs/petri/x/feeburner/types"
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
