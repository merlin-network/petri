package types_test

import (
	"testing"

	"github.com/merlin-network/petri/app"

	"github.com/stretchr/testify/require"

	keepertest "github.com/merlin-network/petri/testutil/feeburner/keeper"
	"github.com/merlin-network/petri/testutil/feeburner/nullify"
	"github.com/merlin-network/petri/x/feeburner"
	"github.com/merlin-network/petri/x/feeburner/types"
)

func TestGenesis(t *testing.T) {
	_ = app.GetDefaultConfig()

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx := keepertest.FeeburnerKeeper(t)
	feeburner.InitGenesis(ctx, *k, genesisState)
	got := feeburner.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "empty petri denom",
			genState: &types.GenesisState{
				Params: types.Params{
					PetriDenom: "",
				},
			},
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
