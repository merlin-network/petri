package wasmbinding

import (
	feeburnerkeeper "github.com/petri-labs/petri/x/feeburner/keeper"
	feerefunderkeeper "github.com/petri-labs/petri/x/feerefunder/keeper"
	icqkeeper "github.com/petri-labs/petri/x/interchainqueries/keeper"
	icacontrollerkeeper "github.com/petri-labs/petri/x/interchaintxs/keeper"

	tokenfactorykeeper "github.com/petri-labs/petri/x/tokenfactory/keeper"
)

type QueryPlugin struct {
	icaControllerKeeper *icacontrollerkeeper.Keeper
	icqKeeper           *icqkeeper.Keeper
	feeBurnerKeeper     *feeburnerkeeper.Keeper
	feeRefunderKeeper   *feerefunderkeeper.Keeper
	tokenFactoryKeeper  *tokenfactorykeeper.Keeper
}

// NewQueryPlugin returns a reference to a new QueryPlugin.
func NewQueryPlugin(icaControllerKeeper *icacontrollerkeeper.Keeper, icqKeeper *icqkeeper.Keeper, feeBurnerKeeper *feeburnerkeeper.Keeper, feeRefunderKeeper *feerefunderkeeper.Keeper, tfk *tokenfactorykeeper.Keeper) *QueryPlugin {
	return &QueryPlugin{
		icaControllerKeeper: icaControllerKeeper,
		icqKeeper:           icqKeeper,
		feeBurnerKeeper:     feeBurnerKeeper,
		feeRefunderKeeper:   feeRefunderKeeper,
		tokenFactoryKeeper:  tfk,
	}
}
