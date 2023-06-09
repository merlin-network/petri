package feerefunder_test

import (
	"testing"

	"github.com/merlin-network/petri/app/params"
	"github.com/merlin-network/petri/testutil/feerefunder/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/merlin-network/petri/testutil/interchainqueries/nullify"
	"github.com/merlin-network/petri/x/feerefunder"
	"github.com/merlin-network/petri/x/feerefunder/types"
)

const TestContractAddressPetri = "petri14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s5c2epq"

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		FeeInfos: []types.FeeInfo{{
			Payer:    TestContractAddressPetri,
			PacketId: types.NewPacketID("port", "channel-1", 64),
			Fee: types.Fee{
				RecvFee:    sdk.NewCoins(sdk.NewCoin(params.DefaultDenom, sdk.NewInt(0))),
				AckFee:     sdk.NewCoins(sdk.NewCoin(params.DefaultDenom, sdk.NewInt(types.DefaultFees.AckFee.AmountOf(params.DefaultDenom).Int64()+1))),
				TimeoutFee: sdk.NewCoins(sdk.NewCoin(params.DefaultDenom, sdk.NewInt(types.DefaultFees.TimeoutFee.AmountOf(params.DefaultDenom).Int64()+1))),
			},
		}},
	}

	require.EqualValues(t, genesisState.Params, types.DefaultParams())

	k, ctx := keeper.FeeKeeper(t, nil, nil)
	feerefunder.InitGenesis(ctx, *k, genesisState)
	got := feerefunder.ExportGenesis(ctx, *k)

	require.EqualValues(t, got.Params, types.DefaultParams())
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
