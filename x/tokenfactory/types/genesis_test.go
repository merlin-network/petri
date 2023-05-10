package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/merlin-network/petri/app"
	"github.com/merlin-network/petri/x/tokenfactory/types"
)

func TestGenesisState_Validate(t *testing.T) {
	app.GetDefaultConfig()

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
			desc: "valid genesis state",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/petri1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "petri1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2",
						},
					},
				},
			},
			valid: true,
		},
		{
			desc: "different admin from creator",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/petri1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "petri1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2",
						},
					},
				},
			},
			valid: true,
		},
		{
			desc: "empty admin",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/petri1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
				},
			},
			valid: true,
		},
		{
			desc: "no admin",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/petri1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2/bitcoin",
					},
				},
			},
			valid: true,
		},
		{
			desc: "invalid admin",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/petri1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "moose",
						},
					},
				},
			},
			valid: false,
		},
		{
			desc: "multiple denoms",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/petri1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
					{
						Denom: "factory/petri1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2/litecoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
				},
			},
			valid: true,
		},
		{
			desc: "duplicate denoms",
			genState: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{
						Denom: "factory/petri1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
					{
						Denom: "factory/petri1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2/bitcoin",
						AuthorityMetadata: types.DenomAuthorityMetadata{
							Admin: "",
						},
					},
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
