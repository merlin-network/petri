package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/petri-labs/petri/app/params"
	feerefundertypes "github.com/petri-labs/petri/x/feerefunder/types"
	tokenfactorytypes "github.com/petri-labs/petri/x/tokenfactory/types"

	"github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/petri-labs/petri/app"
	"github.com/petri-labs/petri/testutil"
	"github.com/petri-labs/petri/wasmbinding/bindings"
	icqtypes "github.com/petri-labs/petri/x/interchainqueries/types"
	ictxtypes "github.com/petri-labs/petri/x/interchaintxs/types"
)

type CustomQuerierTestSuite struct {
	testutil.IBCConnectionTestSuite
}

func (suite *CustomQuerierTestSuite) TestInterchainQueryResult() {
	var (
		petri = suite.GetPetriZoneApp(suite.ChainA)
		ctx     = suite.ChainA.GetContext()
		owner   = keeper.RandomAccountAddress(suite.T()) // We don't care what this address is
	)

	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(ctx, owner, "../testdata/reflect.wasm")
	contractAddress := suite.InstantiateReflectContract(ctx, owner, codeID)
	suite.Require().NotEmpty(contractAddress)

	// Register and submit query result
	clientKey := host.FullClientStateKey(suite.Path.EndpointB.ClientID)
	lastID := petri.InterchainQueriesKeeper.GetLastRegisteredQueryKey(ctx) + 1
	petri.InterchainQueriesKeeper.SetLastRegisteredQueryKey(ctx, lastID)
	registeredQuery := &icqtypes.RegisteredQuery{
		Id: lastID,
		Keys: []*icqtypes.KVKey{
			{Path: host.StoreKey, Key: clientKey},
		},
		QueryType:    string(icqtypes.InterchainQueryTypeKV),
		UpdatePeriod: 1,
		ConnectionId: suite.Path.EndpointA.ConnectionID,
	}
	petri.InterchainQueriesKeeper.SetLastRegisteredQueryKey(ctx, lastID)
	err := petri.InterchainQueriesKeeper.SaveQuery(ctx, registeredQuery)
	suite.Require().NoError(err)

	chainBResp := suite.ChainB.App.Query(abci.RequestQuery{
		Path:   fmt.Sprintf("store/%s/key", host.StoreKey),
		Height: suite.ChainB.LastHeader.Header.Height - 1,
		Data:   clientKey,
		Prove:  true,
	})

	expectedQueryResult := &icqtypes.QueryResult{
		KvResults: []*icqtypes.StorageValue{{
			Key:           chainBResp.Key,
			Proof:         chainBResp.ProofOps,
			Value:         chainBResp.Value,
			StoragePrefix: host.StoreKey,
		}},
		// we don't have tests to test transactions proofs verification since it's a tendermint layer, and we don't have access to it here
		Block:    nil,
		Height:   uint64(chainBResp.Height),
		Revision: suite.ChainA.LastHeader.GetHeight().GetRevisionNumber(),
	}
	err = petri.InterchainQueriesKeeper.SaveKVQueryResult(ctx, lastID, expectedQueryResult)
	suite.Require().NoError(err)

	// Query interchain query result
	query := bindings.PetriQuery{
		InterchainQueryResult: &bindings.QueryRegisteredQueryResultRequest{
			QueryID: lastID,
		},
	}
	resp := icqtypes.QueryRegisteredQueryResultResponse{}
	err = suite.queryCustom(ctx, contractAddress, query, &resp)
	suite.Require().NoError(err)

	suite.Require().Equal(uint64(chainBResp.Height), resp.Result.Height)
	suite.Require().Equal(suite.ChainA.LastHeader.GetHeight().GetRevisionNumber(), resp.Result.Revision)
	suite.Require().Empty(resp.Result.Block)
	suite.Require().NotEmpty(resp.Result.KvResults)
	suite.Require().Equal([]*icqtypes.StorageValue{{
		Key:           chainBResp.Key,
		Proof:         nil,
		Value:         chainBResp.Value,
		StoragePrefix: host.StoreKey,
	}}, resp.Result.KvResults)
}

func (suite *CustomQuerierTestSuite) TestInterchainQueryResultNotFound() {
	var (
		ctx   = suite.ChainA.GetContext()
		owner = keeper.RandomAccountAddress(suite.T()) // We don't care what this address is
	)

	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(ctx, owner, "../testdata/reflect.wasm")
	contractAddress := suite.InstantiateReflectContract(ctx, owner, codeID)
	suite.Require().NotEmpty(contractAddress)

	// Query interchain query result
	query := bindings.PetriQuery{
		InterchainQueryResult: &bindings.QueryRegisteredQueryResultRequest{
			QueryID: 1,
		},
	}
	resp := icqtypes.QueryRegisteredQueryResultResponse{}
	err := suite.queryCustom(ctx, contractAddress, query, &resp)
	expectedErrMsg := fmt.Sprintf("Generic error: Querier contract error: codespace: interchainqueries, code: %d: query wasm contract failed", icqtypes.ErrNoQueryResult.ABCICode())
	suite.Require().ErrorContains(err, expectedErrMsg)
}

func (suite *CustomQuerierTestSuite) TestInterchainAccountAddress() {
	var (
		ctx   = suite.ChainA.GetContext()
		owner = keeper.RandomAccountAddress(suite.T()) // We don't care what this address is
	)

	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(ctx, owner, "../testdata/reflect.wasm")
	contractAddress := suite.InstantiateReflectContract(ctx, owner, codeID)
	suite.Require().NotEmpty(contractAddress)

	err := testutil.SetupICAPath(suite.Path, contractAddress.String())
	suite.Require().NoError(err)

	query := bindings.PetriQuery{
		InterchainAccountAddress: &bindings.QueryInterchainAccountAddressRequest{
			OwnerAddress:        contractAddress.String(),
			InterchainAccountID: testutil.TestInterchainID,
			ConnectionID:        suite.Path.EndpointA.ConnectionID,
		},
	}
	resp := ictxtypes.QueryInterchainAccountAddressResponse{}
	err = suite.queryCustom(ctx, contractAddress, query, &resp)
	suite.Require().NoError(err)

	hostPetriApp, ok := suite.ChainB.App.(*app.App)
	suite.Require().True(ok)

	expected := hostPetriApp.ICAHostKeeper.GetAllInterchainAccounts(suite.ChainB.GetContext())[0].AccountAddress // we expect only one registered ICA
	suite.Require().Equal(expected, resp.InterchainAccountAddress)
}

func (suite *CustomQuerierTestSuite) TestUnknownInterchainAcc() {
	var (
		ctx   = suite.ChainA.GetContext()
		owner = keeper.RandomAccountAddress(suite.T()) // We don't care what this address is
	)

	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(ctx, owner, "../testdata/reflect.wasm")
	contractAddress := suite.InstantiateReflectContract(ctx, owner, codeID)
	suite.Require().NotEmpty(contractAddress)

	err := testutil.SetupICAPath(suite.Path, contractAddress.String())
	suite.Require().NoError(err)

	query := bindings.PetriQuery{
		InterchainAccountAddress: &bindings.QueryInterchainAccountAddressRequest{
			OwnerAddress:        testutil.TestOwnerAddress,
			InterchainAccountID: "wrong_account_id",
			ConnectionID:        suite.Path.EndpointA.ConnectionID,
		},
	}
	resp := ictxtypes.QueryInterchainAccountAddressResponse{}
	expectedErrorMsg := "Generic error: Querier contract error: codespace: interchaintxs, code: 1102: query wasm contract failed"

	err = suite.queryCustom(ctx, contractAddress, query, &resp)
	suite.Require().ErrorContains(err, expectedErrorMsg)
}

func (suite *CustomQuerierTestSuite) TestMinIbcFee() {
	var (
		ctx   = suite.ChainA.GetContext()
		owner = keeper.RandomAccountAddress(suite.T()) // We don't care what this address is
	)

	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(ctx, owner, "../testdata/reflect.wasm")
	contractAddress := suite.InstantiateReflectContract(ctx, owner, codeID)
	suite.Require().NotEmpty(contractAddress)

	query := bindings.PetriQuery{
		MinIbcFee: &bindings.QueryMinIbcFeeRequest{},
	}
	resp := bindings.QueryMinIbcFeeResponse{}

	err := suite.queryCustom(ctx, contractAddress, query, &resp)
	suite.Require().NoError(err)
	suite.Require().Equal(
		feerefundertypes.Fee{
			RecvFee: sdk.Coins{},
			AckFee: sdk.Coins{
				sdk.Coin{Denom: "untrn", Amount: sdk.NewInt(1000)},
			},
			TimeoutFee: sdk.Coins{
				sdk.Coin{Denom: "untrn", Amount: sdk.NewInt(1000)},
			},
		},
		resp.MinFee,
	)
}

func (suite *CustomQuerierTestSuite) TestFullDenom() {
	var (
		ctx   = suite.ChainA.GetContext()
		owner = keeper.RandomAccountAddress(suite.T()) // We don't care what this address is
	)

	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(ctx, owner, "../testdata/reflect.wasm")
	contractAddress := suite.InstantiateReflectContract(ctx, owner, codeID)
	suite.Require().NotEmpty(contractAddress)

	query := bindings.PetriQuery{
		FullDenom: &bindings.FullDenom{
			CreatorAddr: contractAddress.String(),
			Subdenom:    "test",
		},
	}
	resp := bindings.FullDenomResponse{}
	err := suite.queryCustom(ctx, contractAddress, query, &resp)
	suite.Require().NoError(err)

	expected := fmt.Sprintf("factory/%s/test", contractAddress.String())
	suite.Require().Equal(expected, resp.Denom)
}

func (suite *CustomQuerierTestSuite) TestDenomAdmin() {
	var (
		petri = suite.GetPetriZoneApp(suite.ChainA)
		ctx     = suite.ChainA.GetContext()
		owner   = keeper.RandomAccountAddress(suite.T()) // We don't care what this address is
	)

	petri.TokenFactoryKeeper.SetParams(ctx, tokenfactorytypes.NewParams(
		sdk.NewCoins(sdk.NewInt64Coin(tokenfactorytypes.DefaultPetriDenom, 10_000_000)),
		FeeCollectorAddress,
	))

	// Store code and instantiate reflect contract
	codeID := suite.StoreReflectCode(ctx, owner, "../testdata/reflect.wasm")
	contractAddress := suite.InstantiateReflectContract(ctx, owner, codeID)
	suite.Require().NotEmpty(contractAddress)

	senderAddress := suite.ChainA.SenderAccounts[0].SenderAccount.GetAddress()
	coinsAmnt := sdk.NewCoins(sdk.NewCoin(params.DefaultDenom, sdk.NewInt(int64(10_000_000))))
	bankKeeper := petri.BankKeeper
	err := bankKeeper.SendCoins(ctx, senderAddress, contractAddress, coinsAmnt)
	suite.NoError(err)

	denom, _ := petri.TokenFactoryKeeper.CreateDenom(ctx, contractAddress.String(), "test")

	query := bindings.PetriQuery{
		DenomAdmin: &bindings.DenomAdmin{
			Subdenom: denom,
		},
	}
	resp := bindings.DenomAdminResponse{}
	err = suite.queryCustom(ctx, contractAddress, query, &resp)
	suite.Require().NoError(err)

	suite.Require().Equal(contractAddress.String(), resp.Admin)
}

type ChainRequest struct {
	Reflect wasmvmtypes.QueryRequest `json:"reflect"`
}

type ChainResponse struct {
	Data []byte `json:"data"`
}

func (suite *CustomQuerierTestSuite) queryCustom(ctx sdk.Context, contract sdk.AccAddress, request interface{}, response interface{}) error {
	msgBz, err := json.Marshal(request)
	suite.Require().NoError(err)

	query := ChainRequest{
		Reflect: wasmvmtypes.QueryRequest{Custom: msgBz},
	}

	queryBz, err := json.Marshal(query)
	if err != nil {
		return err
	}

	resBz, err := suite.GetPetriZoneApp(suite.ChainA).WasmKeeper.QuerySmart(ctx, contract, queryBz)
	if err != nil {
		return err
	}

	var resp ChainResponse
	err = json.Unmarshal(resBz, &resp)
	if err != nil {
		return err
	}

	return json.Unmarshal(resp.Data, response)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(CustomQuerierTestSuite))
}
