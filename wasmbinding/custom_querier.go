package wasmbinding

import (
	"encoding/json"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/petri-labs/petri/wasmbinding/bindings"
)

// CustomQuerier returns a function that is an implementation of custom querier mechanism for specific messages
func CustomQuerier(qp *QueryPlugin) func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
	return func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
		var contractQuery bindings.PetriQuery
		if err := json.Unmarshal(request, &contractQuery); err != nil {
			return nil, sdkerrors.Wrapf(err, "failed to unmarshal petri query: %v", err)
		}

		switch {
		case contractQuery.InterchainQueryResult != nil:
			queryID := contractQuery.InterchainQueryResult.QueryID

			response, err := qp.GetInterchainQueryResult(ctx, queryID)
			if err != nil {
				return nil, sdkerrors.Wrapf(err, "failed to get interchain query result: %v", err)
			}

			bz, err := json.Marshal(response)
			if err != nil {
				return nil, sdkerrors.Wrapf(err, "failed to marshal interchain query result: %v", err)
			}

			return bz, nil
		case contractQuery.InterchainAccountAddress != nil:

			interchainAccountAddress, err := qp.GetInterchainAccountAddress(ctx, contractQuery.InterchainAccountAddress)
			if err != nil {
				return nil, sdkerrors.Wrapf(err, "failed to get interchain account address: %v", err)
			}

			bz, err := json.Marshal(interchainAccountAddress)
			if err != nil {
				return nil, sdkerrors.Wrapf(err, "failed to marshal interchain account query response: %v", err)
			}

			return bz, nil
		case contractQuery.RegisteredInterchainQueries != nil:
			registeredQueries, err := qp.GetRegisteredInterchainQueries(ctx, contractQuery.RegisteredInterchainQueries)
			if err != nil {
				return nil, sdkerrors.Wrapf(err, "failed to get registered queries: %v", err)
			}

			bz, err := json.Marshal(registeredQueries)
			if err != nil {
				return nil, sdkerrors.Wrapf(err, "failed to marshal interchain account query response: %v", err)
			}

			return bz, nil
		case contractQuery.RegisteredInterchainQuery != nil:
			registeredQuery, err := qp.GetRegisteredInterchainQuery(ctx, contractQuery.RegisteredInterchainQuery)
			if err != nil {
				return nil, sdkerrors.Wrapf(err, "failed to get registered queries: %v", err)
			}

			bz, err := json.Marshal(registeredQuery)
			if err != nil {
				return nil, sdkerrors.Wrapf(err, "failed to marshal interchain account query response: %v", err)
			}

			return bz, nil
		case contractQuery.TotalBurnedPetrisAmount != nil:
			totalBurnedPetris, err := qp.GetTotalBurnedPetrisAmount(ctx, contractQuery.TotalBurnedPetrisAmount)
			if err != nil {
				return nil, sdkerrors.Wrapf(err, "failed to get total burned petris amount: %v", err)
			}

			bz, err := json.Marshal(totalBurnedPetris)
			if err != nil {
				return nil, sdkerrors.Wrapf(err, "failed to marshal total burned petris amount response: %v", err)
			}

			return bz, nil
		case contractQuery.MinIbcFee != nil:
			minFee, err := qp.GetMinIbcFee(ctx, contractQuery.MinIbcFee)
			if err != nil {
				return nil, sdkerrors.Wrapf(err, "failed to get min fee: %v", err)
			}

			bz, err := json.Marshal(minFee)
			if err != nil {
				return nil, sdkerrors.Wrapf(err, "failed to marshal min fee response: %v", err)
			}

			return bz, nil

		case contractQuery.FullDenom != nil:
			creator := contractQuery.FullDenom.CreatorAddr
			subdenom := contractQuery.FullDenom.Subdenom

			fullDenom, err := GetFullDenom(creator, subdenom)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "unable to get full denom")
			}

			res := bindings.FullDenomResponse{
				Denom: fullDenom,
			}

			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "failed to JSON marshal FullDenomResponse response.")
			}

			return bz, nil

		case contractQuery.DenomAdmin != nil:
			res, err := qp.GetDenomAdmin(ctx, contractQuery.DenomAdmin.Subdenom)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "unable to get denom admin")
			}

			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "failed to JSON marshal DenomAdminResponse response.")
			}

			return bz, nil

		default:
			return nil, wasmvmtypes.UnsupportedRequest{Kind: "unknown petri query type"}
		}
	}
}
