package keeper

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	consumertypes "github.com/cosmos/interchain-security/x/ccv/consumer/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/merlin-network/petri/x/feeburner/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

var KeyBurnedFees = []byte("BurnedFees")

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

// RecordBurnedFees adds `amount` to the total amount of burned FURY tokens
func (k Keeper) RecordBurnedFees(ctx sdk.Context, amount sdk.Coin) {
	store := ctx.KVStore(k.storeKey)

	totalBurnedPetrisAmount := k.GetTotalBurnedPetrisAmount(ctx)
	totalBurnedPetrisAmount.Coin = totalBurnedPetrisAmount.Coin.Add(amount)

	store.Set(KeyBurnedFees, k.cdc.MustMarshal(&totalBurnedPetrisAmount))
}

// GetTotalBurnedPetrisAmount gets the total burned amount of FURY tokens
func (k Keeper) GetTotalBurnedPetrisAmount(ctx sdk.Context) types.TotalBurnedPetrisAmount {
	store := ctx.KVStore(k.storeKey)

	var totalBurnedPetrisAmount types.TotalBurnedPetrisAmount
	bzTotalBurnedPetrisAmount := store.Get(KeyBurnedFees)
	if bzTotalBurnedPetrisAmount != nil {
		k.cdc.MustUnmarshal(bzTotalBurnedPetrisAmount, &totalBurnedPetrisAmount)
	}

	if totalBurnedPetrisAmount.Coin.Denom == "" {
		totalBurnedPetrisAmount.Coin = sdk.NewCoin(k.GetParams(ctx).PetriDenom, sdk.NewInt(0))
	}

	return totalBurnedPetrisAmount
}

// BurnAndDistribute is an important part of tokenomics. It does few things:
// 1. Burns FURY fee coins distributed to consumertypes.ConsumerRedistributeName in ICS (https://github.com/cosmos/interchain-security/blob/v0.2.0/x/ccv/consumer/keeper/distribution.go#L17)
// 2. Updates total amount of burned FURY coins
// 3. Sends non-FURY fee tokens to reserve contract address
// Panics if no `consumertypes.ConsumerRedistributeName` module found OR could not burn FURY tokens
func (k Keeper) BurnAndDistribute(ctx sdk.Context) {
	moduleAddr := k.accountKeeper.GetModuleAddress(consumertypes.ConsumerRedistributeName)
	if moduleAddr == nil {
		panic("ConsumerRedistributeName must have module address")
	}

	params := k.GetParams(ctx)
	balances := k.bankKeeper.GetAllBalances(ctx, moduleAddr)
	fundsForReserve := make(sdk.Coins, 0, len(balances))

	for _, balance := range balances {
		if !balance.IsZero() {
			if balance.Denom == params.PetriDenom {
				err := k.bankKeeper.BurnCoins(ctx, consumertypes.ConsumerRedistributeName, sdk.Coins{balance})
				if err != nil {
					panic(sdkerrors.Wrapf(err, "failed to burn FURY tokens during fee processing"))
				}

				k.RecordBurnedFees(ctx, balance)
			} else {
				fundsForReserve = append(fundsForReserve, balance)
			}
		}
	}

	if len(fundsForReserve) > 0 {
		addr, err := sdk.AccAddressFromBech32(params.TreasuryAddress)
		if err != nil {
			// there's no way we face this kind of situation in production, since it means the chain is misconfigured
			// still, in test environments it might be the case when the chain is started without Reserve
			// in such case we just burn the tokens
			err := k.bankKeeper.BurnCoins(ctx, consumertypes.ConsumerRedistributeName, fundsForReserve)
			if err != nil {
				panic(sdkerrors.Wrapf(err, "failed to burn tokens during fee processing"))
			}
		} else {
			err = k.bankKeeper.SendCoins(
				ctx,
				moduleAddr, addr,
				fundsForReserve,
			)
			if err != nil {
				panic(sdkerrors.Wrapf(err, "failed sending funds to Reserve"))
			}
		}
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// FundCommunityPool is method to satisfy DistributionKeeper interface for packet-forward-middleware Keeper.
// The original method sends coins to a community pool of a chain.
// The current method sends coins to a Fee Collector module which collects fee on consumer chain.
func (k Keeper) FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error {
	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, authtypes.FeeCollectorName, amount)
}
