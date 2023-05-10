package types_test

import (
	"testing"

	"github.com/merlin-network/petri/app"

	"github.com/merlin-network/petri/x/cron/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	app.GetDefaultConfig()

	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				Params: types.Params{
					SecurityAddress: "petri17dtl0mjt3t77kpuhg2edqzjpszulwhgzcdvagh",
					Limit:           1,
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "invalid genesis state - params are invalid",
			genState: &types.GenesisState{
				Params: types.Params{
					SecurityAddress: "",
					Limit:           0,
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
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
