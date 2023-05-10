package v3

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/petri-labs/petri/app/upgrades"

	crontypes "github.com/petri-labs/petri/x/cron/types"
	icqtypes "github.com/petri-labs/petri/x/interchainqueries/types"
	tokenfactorytypes "github.com/petri-labs/petri/x/tokenfactory/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	keepers *upgrades.UpgradeKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info("Starting module migrations...")

		// todo: FIXME
		keepers.IcqKeeper.SetParams(ctx, icqtypes.DefaultParams())
		keepers.CronKeeper.SetParams(ctx, crontypes.DefaultParams())
		keepers.TokenFactoryKeeper.SetParams(ctx, tokenfactorytypes.DefaultParams())

		vm, err := mm.RunMigrations(ctx, configurator, vm)
		if err != nil {
			return vm, err
		}

		ctx.Logger().Info("Upgrade complete")
		return vm, err
	}
}
