package keeper_test

import (
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/merlin-network/petri/x/tokenfactory/types"
)

func (suite *KeeperTestSuite) TestGenesis() {
	genesisState := types.GenesisState{
		FactoryDenoms: []types.GenesisDenom{
			{
				Denom: "factory/petri1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2/bitcoin",
				AuthorityMetadata: types.DenomAuthorityMetadata{
					Admin: "petri1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2",
				},
			},
			{
				Denom: "factory/petri1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2/diff-admin",
				AuthorityMetadata: types.DenomAuthorityMetadata{
					Admin: "petri1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2",
				},
			},
			{
				Denom: "factory/petri1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2/litecoin",
				AuthorityMetadata: types.DenomAuthorityMetadata{
					Admin: "petri1m9l358xunhhwds0568za49mzhvuxx9ux8xafx2",
				},
			},
		},
	}
	app := suite.GetPetriZoneApp(suite.ChainA)
	context := app.BaseApp.NewContext(false, tmproto.Header{})
	// Test both with bank denom metadata set, and not set.
	for i, denom := range genesisState.FactoryDenoms {
		// hacky, sets bank metadata to exist if i != 0, to cover both cases.
		if i != 0 {
			app.BankKeeper.SetDenomMetaData(context, banktypes.Metadata{Base: denom.GetDenom()})
		}
	}

	app.TokenFactoryKeeper.SetParams(context, types.Params{})
	app.TokenFactoryKeeper.InitGenesis(context, genesisState)
	exportedGenesis := app.TokenFactoryKeeper.ExportGenesis(context)
	suite.Require().NotNil(exportedGenesis)
	suite.Require().Equal(genesisState, *exportedGenesis)
}
