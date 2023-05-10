package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	ictxtypes "github.com/petri-labs/petri/x/interchaintxs/types"

	adminkeeper "github.com/cosmos/admin-module/x/adminmodule/keeper"
	admintypes "github.com/cosmos/admin-module/x/adminmodule/types"

	"github.com/petri-labs/petri/app/params"

	"github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	"github.com/stretchr/testify/require"

	"github.com/petri-labs/petri/app"
	"github.com/petri-labs/petri/testutil"
	"github.com/petri-labs/petri/wasmbinding"
	"github.com/petri-labs/petri/wasmbinding/bindings"
	feetypes "github.com/petri-labs/petri/x/feerefunder/types"
	icqkeeper "github.com/petri-labs/petri/x/interchainqueries/keeper"
	icqtypes "github.com/petri-labs/petri/x/interchainqueries/types"
	ictxkeeper "github.com/petri-labs/petri/x/interchaintxs/keeper"

	tokenfactorytypes "github.com/petri-labs/petri/x/tokenfactory/types"
)

const FeeCollectorAddress = "petri1vguuxez2h5ekltfj9gjd62fs5k4rl2zy5hfrncasykzw08rezpfsd2rhm7"

type CustomMessengerTestSuite struct {
	testutil.IBCConnectionTestSuite
	petri         *app.App
	ctx             sdk.Context
	messenger       *wasmbinding.CustomMessenger
	contractOwner   sdk.AccAddress
	contractAddress sdk.AccAddress
}

func (suite *CustomMessengerTestSuite) SetupTest() {
	suite.IBCConnectionTestSuite.SetupTest()
	suite.petri = suite.GetPetriZoneApp(suite.ChainA)
	suite.ctx = suite.ChainA.GetContext()
	suite.messenger = &wasmbinding.CustomMessenger{}
	suite.messenger.Ictxmsgserver = ictxkeeper.NewMsgServerImpl(suite.petri.InterchainTxsKeeper)
	suite.messenger.Keeper = suite.petri.InterchainTxsKeeper
	suite.messenger.Icqmsgserver = icqkeeper.NewMsgServerImpl(suite.petri.InterchainQueriesKeeper)
	suite.messenger.Adminserver = adminkeeper.NewMsgServerImpl(suite.petri.AdminmoduleKeeper)
	suite.messenger.Bank = &suite.petri.BankKeeper
	suite.messenger.TokenFactory = suite.petri.TokenFactoryKeeper
	suite.messenger.CronKeeper = &suite.petri.CronKeeper
	suite.messenger.AdminKeeper = &suite.petri.AdminmoduleKeeper
	suite.contractOwner = keeper.RandomAccountAddress(suite.T())

	suite.messenger.TokenFactory.SetParams(suite.ctx, tokenfactorytypes.NewParams(
		sdk.NewCoins(sdk.NewInt64Coin(tokenfactorytypes.DefaultPetriDenom, 10_000_000)),
		FeeCollectorAddress,
	))
}

func (suite *CustomMessengerTestSuite) TestRegisterInterchainAccount() {
	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(suite.ctx, suite.contractOwner, "../testdata/reflect.wasm")
	suite.contractAddress = suite.InstantiateReflectContract(suite.ctx, suite.contractOwner, codeID)
	suite.Require().NotEmpty(suite.contractAddress)

	// Craft RegisterInterchainAccount message
	msg, err := json.Marshal(bindings.PetriMsg{
		RegisterInterchainAccount: &bindings.RegisterInterchainAccount{
			ConnectionId:        suite.Path.EndpointA.ConnectionID,
			InterchainAccountId: testutil.TestInterchainID,
		},
	})
	suite.NoError(err)

	// Dispatch RegisterInterchainAccount message
	events, data, err := suite.messenger.DispatchMsg(suite.ctx, suite.contractAddress, suite.Path.EndpointA.ChannelConfig.PortID, types.CosmosMsg{
		Custom: msg,
	})
	suite.NoError(err)
	suite.Nil(events)
	suite.Equal([][]byte{[]byte(`{}`)}, data)
}

func (suite *CustomMessengerTestSuite) TestRegisterInterchainAccountLongID() {
	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(suite.ctx, suite.contractOwner, "../testdata/reflect.wasm")
	suite.contractAddress = suite.InstantiateReflectContract(suite.ctx, suite.contractOwner, codeID)
	suite.Require().NotEmpty(suite.contractAddress)

	// Craft RegisterInterchainAccount message
	msg, err := json.Marshal(bindings.PetriMsg{
		RegisterInterchainAccount: &bindings.RegisterInterchainAccount{
			ConnectionId: suite.Path.EndpointA.ConnectionID,
			// the limit is 47, this line is 50 characters long
			InterchainAccountId: "01234567890123456789012345678901234567890123456789",
		},
	})
	suite.NoError(err)

	// Dispatch RegisterInterchainAccount message
	_, _, err = suite.messenger.DispatchMsg(suite.ctx, suite.contractAddress, suite.Path.EndpointA.ChannelConfig.PortID, types.CosmosMsg{
		Custom: msg,
	})
	suite.Error(err)
	suite.ErrorIs(err, ictxtypes.ErrLongInterchainAccountID)
}

func (suite *CustomMessengerTestSuite) TestRegisterInterchainQuery() {
	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(suite.ctx, suite.contractOwner, "../testdata/reflect.wasm")
	suite.contractAddress = suite.InstantiateReflectContract(suite.ctx, suite.contractOwner, codeID)
	suite.Require().NotEmpty(suite.contractAddress)

	err := testutil.SetupICAPath(suite.Path, suite.contractAddress.String())
	suite.Require().NoError(err)

	// Top up contract balance
	senderAddress := suite.ChainA.SenderAccounts[0].SenderAccount.GetAddress()
	coinsAmnt := sdk.NewCoins(sdk.NewCoin(params.DefaultDenom, sdk.NewInt(int64(10_000_000))))
	bankKeeper := suite.petri.BankKeeper
	err = bankKeeper.SendCoins(suite.ctx, senderAddress, suite.contractAddress, coinsAmnt)
	suite.NoError(err)

	// Craft RegisterInterchainQuery message
	msg, err := json.Marshal(bindings.PetriMsg{
		RegisterInterchainQuery: &bindings.RegisterInterchainQuery{
			QueryType: string(icqtypes.InterchainQueryTypeKV),
			Keys: []*icqtypes.KVKey{
				{Path: host.StoreKey, Key: host.FullClientStateKey(suite.Path.EndpointB.ClientID)},
			},
			TransactionsFilter: "{}",
			ConnectionId:       suite.Path.EndpointA.ConnectionID,
			UpdatePeriod:       20,
		},
	})
	suite.NoError(err)

	// Dispatch RegisterInterchainQuery message
	events, data, err := suite.messenger.DispatchMsg(suite.ctx, suite.contractAddress, suite.Path.EndpointA.ChannelConfig.PortID, types.CosmosMsg{
		Custom: msg,
	})
	suite.NoError(err)
	suite.Nil(events)
	suite.Equal([][]byte{[]byte(`{"id":1}`)}, data)
}

func (suite *CustomMessengerTestSuite) TestCreateDenomMsg() {
	codeID := suite.StoreReflectCode(suite.ctx, suite.contractOwner, "../testdata/reflect.wasm")
	suite.contractAddress = suite.InstantiateReflectContract(suite.ctx, suite.contractOwner, codeID)
	suite.Require().NotEmpty(suite.contractAddress)

	senderAddress := suite.ChainA.SenderAccounts[0].SenderAccount.GetAddress()
	coinsAmnt := sdk.NewCoins(sdk.NewCoin(params.DefaultDenom, sdk.NewInt(int64(10_000_000))))
	bankKeeper := suite.petri.BankKeeper
	err := bankKeeper.SendCoins(suite.ctx, senderAddress, suite.contractAddress, coinsAmnt)
	suite.NoError(err)

	fullMsg := bindings.PetriMsg{
		CreateDenom: &bindings.CreateDenom{
			Subdenom: "SUN",
		},
	}

	data, _ := suite.executeCustomMsg(suite.contractAddress, fullMsg)

	suite.Equal([][]byte(nil), data)
}

func (suite *CustomMessengerTestSuite) TestMintMsg() {
	var (
		petri = suite.GetPetriZoneApp(suite.ChainA)
		ctx     = suite.ChainA.GetContext()
		lucky   = keeper.RandomAccountAddress(suite.T()) // We don't care what this address is
	)

	codeID := suite.StoreReflectCode(suite.ctx, suite.contractOwner, "../testdata/reflect.wasm")
	suite.contractAddress = suite.InstantiateReflectContract(suite.ctx, suite.contractOwner, codeID)
	suite.Require().NotEmpty(suite.contractAddress)

	senderAddress := suite.ChainA.SenderAccounts[0].SenderAccount.GetAddress()
	coinsAmnt := sdk.NewCoins(sdk.NewCoin(params.DefaultDenom, sdk.NewInt(int64(20_000_000))))
	bankKeeper := suite.petri.BankKeeper
	err := bankKeeper.SendCoins(suite.ctx, senderAddress, suite.contractAddress, coinsAmnt)
	suite.NoError(err)

	// lucky was broke
	balances := petri.BankKeeper.GetAllBalances(suite.ctx, lucky)
	require.Empty(suite.T(), balances)

	// Create denom for minting
	fullMsg := bindings.PetriMsg{
		CreateDenom: &bindings.CreateDenom{
			Subdenom: "SUN",
		},
	}

	suite.executeCustomMsg(suite.contractAddress, fullMsg)

	sunDenom := fmt.Sprintf("factory/%s/%s", suite.contractAddress.String(), fullMsg.CreateDenom.Subdenom)

	amount, ok := sdk.NewIntFromString("808010808")
	require.True(suite.T(), ok)

	fullMsg = bindings.PetriMsg{
		MintTokens: &bindings.MintTokens{
			Denom:         sunDenom,
			Amount:        amount,
			MintToAddress: lucky.String(),
		},
	}

	suite.executeCustomMsg(suite.contractAddress, fullMsg)

	balances = petri.BankKeeper.GetAllBalances(ctx, lucky)
	require.Len(suite.T(), balances, 1)
	coin := balances[0]
	require.Equal(suite.T(), amount, coin.Amount)
	require.Contains(suite.T(), coin.Denom, "factory/")
	require.Equal(suite.T(), sunDenom, coin.Denom)

	// mint the same denom again
	suite.executeCustomMsg(suite.contractAddress, fullMsg)

	balances = petri.BankKeeper.GetAllBalances(ctx, lucky)
	require.Len(suite.T(), balances, 1)
	coin = balances[0]
	require.Equal(suite.T(), amount.MulRaw(2), coin.Amount)
	require.Contains(suite.T(), coin.Denom, "factory/")
	require.Equal(suite.T(), sunDenom, coin.Denom)

	// now mint another amount / denom
	// create it first
	fullMsg = bindings.PetriMsg{
		CreateDenom: &bindings.CreateDenom{
			Subdenom: "MOON",
		},
	}
	suite.executeCustomMsg(suite.contractAddress, fullMsg)

	moonDenom := fmt.Sprintf("factory/%s/%s", suite.contractAddress.String(), fullMsg.CreateDenom.Subdenom)

	amount = amount.SubRaw(1)
	fullMsg = bindings.PetriMsg{
		MintTokens: &bindings.MintTokens{
			Denom:         moonDenom,
			Amount:        amount,
			MintToAddress: lucky.String(),
		},
	}

	suite.executeCustomMsg(suite.contractAddress, fullMsg)

	balances = petri.BankKeeper.GetAllBalances(ctx, lucky)
	require.Len(suite.T(), balances, 2)
	coin = balances[0]
	require.Equal(suite.T(), amount, coin.Amount)
	require.Contains(suite.T(), coin.Denom, "factory/")
	require.Equal(suite.T(), moonDenom, coin.Denom)
}

func (suite *CustomMessengerTestSuite) TestUpdateInterchainQuery() {
	// reuse register interchain query test to get query registered
	suite.TestRegisterInterchainQuery()

	// Craft UpdateInterchainQuery message
	msg, err := json.Marshal(bindings.PetriMsg{
		UpdateInterchainQuery: &bindings.UpdateInterchainQuery{
			QueryId:         1,
			NewKeys:         nil,
			NewUpdatePeriod: 111,
		},
	})
	suite.NoError(err)

	// Dispatch UpdateInterchainQuery message
	events, data, err := suite.messenger.DispatchMsg(suite.ctx, suite.contractAddress, suite.Path.EndpointA.ChannelConfig.PortID, types.CosmosMsg{
		Custom: msg,
	})
	suite.NoError(err)
	suite.Nil(events)
	suite.Equal([][]byte{[]byte(`{}`)}, data)
}

func (suite *CustomMessengerTestSuite) TestUpdateInterchainQueryFailed() {
	// Craft UpdateInterchainQuery message
	msg, err := json.Marshal(bindings.PetriMsg{
		UpdateInterchainQuery: &bindings.UpdateInterchainQuery{
			QueryId:         1,
			NewKeys:         nil,
			NewUpdatePeriod: 1,
		},
	})
	suite.NoError(err)

	// Dispatch UpdateInterchainQuery message
	owner, err := sdk.AccAddressFromBech32(testutil.TestOwnerAddress)
	suite.NoError(err)
	events, data, err := suite.messenger.DispatchMsg(suite.ctx, owner, suite.Path.EndpointA.ChannelConfig.PortID, types.CosmosMsg{
		Custom: msg,
	})
	expectedErrMsg := "failed to update interchain query: failed to update interchain query: failed to get query by query id: there is no query with id: 1"
	suite.Require().ErrorContains(err, expectedErrMsg)
	suite.Nil(events)
	suite.Nil(data)
}

func (suite *CustomMessengerTestSuite) TestRemoveInterchainQuery() {
	// Reuse register interchain query test to get query registered
	suite.TestRegisterInterchainQuery()

	// Craft RemoveInterchainQuery message
	msg, err := json.Marshal(bindings.PetriMsg{
		RemoveInterchainQuery: &bindings.RemoveInterchainQuery{
			QueryId: 1,
		},
	})
	suite.NoError(err)

	// Dispatch RemoveInterchainQuery message
	suite.NoError(err)
	events, data, err := suite.messenger.DispatchMsg(suite.ctx, suite.contractAddress, suite.Path.EndpointA.ChannelConfig.PortID, types.CosmosMsg{
		Custom: msg,
	})
	suite.NoError(err)
	suite.Nil(events)
	suite.Equal([][]byte{[]byte(`{}`)}, data)
}

func (suite *CustomMessengerTestSuite) TestRemoveInterchainQueryFailed() {
	// Craft RemoveInterchainQuery message
	msg, err := json.Marshal(bindings.PetriMsg{
		RemoveInterchainQuery: &bindings.RemoveInterchainQuery{
			QueryId: 1,
		},
	})
	suite.NoError(err)

	// Dispatch RemoveInterchainQuery message
	owner, err := sdk.AccAddressFromBech32(testutil.TestOwnerAddress)
	suite.NoError(err)
	events, data, err := suite.messenger.DispatchMsg(suite.ctx, owner, suite.Path.EndpointA.ChannelConfig.PortID, types.CosmosMsg{
		Custom: msg,
	})
	expectedErrMsg := "failed to remove interchain query: failed to remove interchain query: failed to get query by query id: there is no query with id: 1"
	suite.Require().ErrorContains(err, expectedErrMsg)
	suite.Nil(events)
	suite.Nil(data)
}

func (suite *CustomMessengerTestSuite) TestSubmitTx() {
	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(suite.ctx, suite.contractOwner, "../testdata/reflect.wasm")
	suite.contractAddress = suite.InstantiateReflectContract(suite.ctx, suite.contractOwner, codeID)
	suite.Require().NotEmpty(suite.contractAddress)

	senderAddress := suite.ChainA.SenderAccounts[0].SenderAccount.GetAddress()
	coinsAmnt := sdk.NewCoins(sdk.NewCoin(params.DefaultDenom, sdk.NewInt(int64(10_000_000))))
	bankKeeper := suite.petri.BankKeeper
	err := bankKeeper.SendCoins(suite.ctx, senderAddress, suite.contractAddress, coinsAmnt)
	suite.NoError(err)

	err = testutil.SetupICAPath(suite.Path, suite.contractAddress.String())
	suite.Require().NoError(err)

	events, data, err := suite.messenger.DispatchMsg(
		suite.ctx,
		suite.contractAddress,
		suite.Path.EndpointA.ChannelConfig.PortID,
		types.CosmosMsg{
			Custom: suite.craftMarshaledMsgSubmitTxWithNumMsgs(1),
		},
	)
	suite.NoError(err)

	var response bindings.SubmitTxResponse
	err = json.Unmarshal(data[0], &response)
	suite.NoError(err)
	suite.Nil(events)
	suite.Equal(uint64(1), response.SequenceId)
	suite.Equal("channel-2", response.Channel)
}

func (suite *CustomMessengerTestSuite) TestSubmitTxTooMuchTxs() {
	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(suite.ctx, suite.contractOwner, "../testdata/reflect.wasm")
	suite.contractAddress = suite.InstantiateReflectContract(suite.ctx, suite.contractOwner, codeID)
	suite.Require().NotEmpty(suite.contractAddress)

	err := testutil.SetupICAPath(suite.Path, suite.contractAddress.String())
	suite.Require().NoError(err)

	_, _, err = suite.messenger.DispatchMsg(
		suite.ctx,
		suite.contractAddress,
		suite.Path.EndpointA.ChannelConfig.PortID,
		types.CosmosMsg{
			Custom: suite.craftMarshaledMsgSubmitTxWithNumMsgs(20),
		},
	)
	suite.ErrorContains(err, "MsgSubmitTx contains more messages than allowed")
}

func (suite *CustomMessengerTestSuite) TestSoftwareUpgradeProposal() {
	// Set admin so that we can execute this proposal without permission error
	suite.petri.AdminmoduleKeeper.SetAdmin(suite.ctx, suite.contractAddress.String())

	// Craft SubmitAdminProposal message
	msg, err := json.Marshal(bindings.PetriMsg{
		SubmitAdminProposal: &bindings.SubmitAdminProposal{
			AdminProposal: bindings.AdminProposal{
				SoftwareUpgradeProposal: &bindings.SoftwareUpgradeProposal{
					Title:       "Test",
					Description: "Test",
					Plan: bindings.Plan{
						Name:   "TestPlan",
						Height: 150,
						Info:   "TestInfo",
					},
				},
			},
		},
	})
	suite.NoError(err)

	// Dispatch SubmitAdminProposal message
	events, data, err := suite.messenger.DispatchMsg(suite.ctx, suite.contractAddress, suite.Path.EndpointA.ChannelConfig.PortID, types.CosmosMsg{
		Custom: msg,
	})
	suite.NoError(err)
	suite.Nil(events)

	expected, err := json.Marshal(&admintypes.MsgSubmitProposalResponse{
		ProposalId: 1,
	})
	suite.NoError(err)
	suite.Equal([][]uint8{expected}, data)

	// Test with other proposer that is not admin should return failure
	otherAddress, err := sdk.AccAddressFromBech32("petri13jrwrtsyjjuynlug65r76r2zvfw5xjcq6532h2")
	suite.NoError(err)
	_, _, err = suite.messenger.DispatchMsg(suite.ctx, otherAddress, suite.Path.EndpointA.ChannelConfig.PortID, types.CosmosMsg{
		Custom: msg,
	})
	suite.Error(err)

	// Check CancelSubmitAdminProposal

	// Craft CancelSubmitAdminProposal message
	msg, err = json.Marshal(bindings.PetriMsg{
		SubmitAdminProposal: &bindings.SubmitAdminProposal{
			AdminProposal: bindings.AdminProposal{
				CancelSoftwareUpgradeProposal: &bindings.CancelSoftwareUpgradeProposal{
					Title:       "Test",
					Description: "Test",
				},
			},
		},
	})
	suite.NoError(err)

	// Dispatch SubmitAdminProposal message
	events, data, err = suite.messenger.DispatchMsg(suite.ctx, suite.contractAddress, suite.Path.EndpointA.ChannelConfig.PortID, types.CosmosMsg{
		Custom: msg,
	})
	suite.NoError(err)
	suite.Nil(events)
	expected, err = json.Marshal(&admintypes.MsgSubmitProposalResponse{
		ProposalId: 2,
	})
	suite.NoError(err)
	suite.Equal([][]uint8{expected}, data)
}

func (suite *CustomMessengerTestSuite) TestTooMuchProposals() {
	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(suite.ctx, suite.contractOwner, "../testdata/reflect.wasm")
	suite.contractAddress = suite.InstantiateReflectContract(suite.ctx, suite.contractOwner, codeID)
	suite.Require().NotEmpty(suite.contractAddress)

	err := testutil.SetupICAPath(suite.Path, suite.contractAddress.String())
	suite.Require().NoError(err)

	// Craft  message with 2 proposals
	msg, err := json.Marshal(bindings.PetriMsg{
		SubmitAdminProposal: &bindings.SubmitAdminProposal{
			AdminProposal: bindings.AdminProposal{
				CancelSoftwareUpgradeProposal: &bindings.CancelSoftwareUpgradeProposal{
					Title:       "Test",
					Description: "Test",
				},
				ClearAdminProposal: &bindings.ClearAdminProposal{
					Title:       "Test",
					Description: "Test",
					Contract:    "Test",
				},
			},
		},
	})
	suite.NoError(err)

	cosmosMsg := types.CosmosMsg{Custom: msg}

	// Dispatch SubmitAdminProposal message
	_, _, err = suite.messenger.DispatchMsg(suite.ctx, suite.contractAddress, suite.Path.EndpointA.ChannelConfig.PortID, cosmosMsg)

	suite.ErrorContains(err, "more than one admin proposal type is present in message")
}

func (suite *CustomMessengerTestSuite) TestNoProposals() {
	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(suite.ctx, suite.contractOwner, "../testdata/reflect.wasm")
	suite.contractAddress = suite.InstantiateReflectContract(suite.ctx, suite.contractOwner, codeID)
	suite.Require().NotEmpty(suite.contractAddress)

	err := testutil.SetupICAPath(suite.Path, suite.contractAddress.String())
	suite.Require().NoError(err)

	// Craft  message with 0 proposals
	msg, err := json.Marshal(bindings.PetriMsg{
		SubmitAdminProposal: &bindings.SubmitAdminProposal{
			AdminProposal: bindings.AdminProposal{},
		},
	})
	suite.NoError(err)

	cosmosMsg := types.CosmosMsg{Custom: msg}

	// Dispatch SubmitAdminProposal message
	_, _, err = suite.messenger.DispatchMsg(suite.ctx, suite.contractAddress, suite.Path.EndpointA.ChannelConfig.PortID, cosmosMsg)

	suite.ErrorContains(err, "no admin proposal type is present in message")
}

func (suite *CustomMessengerTestSuite) TestAddRemoveSchedule() {
	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(suite.ctx, suite.contractOwner, "../testdata/reflect.wasm")
	suite.contractAddress = suite.InstantiateReflectContract(suite.ctx, suite.contractOwner, codeID)
	suite.Require().NotEmpty(suite.contractAddress)

	// Set admin so that we can execute this proposal without permission error
	suite.petri.AdminmoduleKeeper.SetAdmin(suite.ctx, suite.contractAddress.String())

	// Craft AddSchedule message
	msg, err := json.Marshal(bindings.PetriMsg{
		AddSchedule: &bindings.AddSchedule{
			Name:   "schedule1",
			Period: 5,
			Msgs: []bindings.MsgExecuteContract{
				{
					Contract: suite.contractAddress.String(),
					Msg:      "{\"send\": { \"to\": \"asdf\", \"amount\": 1000 }}",
				},
			},
		},
	})
	suite.NoError(err)

	// Dispatch AddSchedule message
	events, data, err := suite.messenger.DispatchMsg(suite.ctx, suite.contractAddress, suite.Path.EndpointA.ChannelConfig.PortID, types.CosmosMsg{
		Custom: msg,
	})
	suite.NoError(err)
	suite.Nil(events)
	expected, err := json.Marshal(&bindings.AddScheduleResponse{})
	suite.NoError(err)
	suite.Equal([][]uint8{expected}, data)

	// Craft RemoveSchedule message
	msg, err = json.Marshal(bindings.PetriMsg{
		RemoveSchedule: &bindings.RemoveSchedule{
			Name: "schedule1",
		},
	})
	suite.NoError(err)

	// Dispatch AddSchedule message
	events, data, err = suite.messenger.DispatchMsg(suite.ctx, suite.contractAddress, suite.Path.EndpointA.ChannelConfig.PortID, types.CosmosMsg{
		Custom: msg,
	})
	suite.NoError(err)
	suite.Nil(events)
	expected, err = json.Marshal(&bindings.RemoveScheduleResponse{})
	suite.NoError(err)
	suite.Equal([][]uint8{expected}, data)
}

func (suite *CustomMessengerTestSuite) executeCustomMsg(owner sdk.AccAddress, fullMsg bindings.PetriMsg) (result [][]byte, msg []byte) {
	msg, err := json.Marshal(fullMsg)
	suite.NoError(err)

	events, result, err := suite.messenger.DispatchMsg(suite.ctx, owner, suite.Path.EndpointA.ChannelConfig.PortID, types.CosmosMsg{
		Custom: msg,
	})
	suite.NoError(err)
	suite.Nil(events)

	return
}

func (suite *CustomMessengerTestSuite) craftMarshaledMsgSubmitTxWithNumMsgs(numMsgs int) (result []byte) {
	msg := bindings.ProtobufAny{
		TypeURL: "/cosmos.staking.v1beta1.MsgDelegate",
		Value:   []byte{26, 10, 10, 5, 115, 116, 97, 107, 101, 18, 1, 48},
	}
	msgs := make([]bindings.ProtobufAny, 0, numMsgs)
	for i := 0; i < numMsgs; i++ {
		msgs = append(msgs, msg)
	}
	result, err := json.Marshal(struct {
		SubmitTx bindings.SubmitTx `json:"submit_tx"`
	}{
		SubmitTx: bindings.SubmitTx{
			ConnectionId:        suite.Path.EndpointA.ConnectionID,
			InterchainAccountId: testutil.TestInterchainID,
			Msgs:                msgs,
			Memo:                "Jimmy",
			Timeout:             2000,
			Fee: feetypes.Fee{
				RecvFee:    sdk.NewCoins(),
				AckFee:     sdk.NewCoins(sdk.NewCoin(params.DefaultDenom, sdk.NewInt(1000))),
				TimeoutFee: sdk.NewCoins(sdk.NewCoin(params.DefaultDenom, sdk.NewInt(1000))),
			},
		},
	})
	suite.NoError(err)
	return
}

func TestMessengerTestSuite(t *testing.T) {
	suite.Run(t, new(CustomMessengerTestSuite))
}
