package v044

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	feeburnertypes "github.com/merlin-network/petri/x/feeburner/types"
	tokenfactorytypes "github.com/merlin-network/petri/x/tokenfactory/types"

	"github.com/merlin-network/petri/app/upgrades"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	keepers *upgrades.UpgradeKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info("Starting module migrations...")
		vm, err := mm.RunMigrations(ctx, configurator, vm)
		if err != nil {
			return vm, err
		}

		ctx.Logger().Info("Migrating SlashingKeeper Params...")
		oldSlashingParams := keepers.SlashingKeeper.GetParams(ctx)
		oldSlashingParams.SignedBlocksWindow = int64(36000)

		keepers.SlashingKeeper.SetParams(ctx, oldSlashingParams)

		ctx.Logger().Info("Migrating FeeBurner Params...")
		s, ok := keepers.ParamsKeeper.GetSubspace(feeburnertypes.ModuleName)
		if !ok {
			return nil, errors.New("global fee burner params subspace not found")
		}
		var reserveAddress string
		s.Get(ctx, feeburnertypes.KeyReserveAddress, &reserveAddress)

		var petriDenom string
		s.Get(ctx, feeburnertypes.KeyPetriDenom, &petriDenom)

		feeburnerDefaultParams := feeburnertypes.DefaultParams()
		feeburnerDefaultParams.TreasuryAddress = reserveAddress
		feeburnerDefaultParams.PetriDenom = petriDenom
		keepers.FeeBurnerKeeper.SetParams(ctx, feeburnerDefaultParams)

		ctx.Logger().Info("Migrating TokenFactory Params...")
		tokenfactoryDefaultParams := tokenfactorytypes.DefaultParams()
		tokenfactoryDefaultParams.FeeCollectorAddress = reserveAddress
		keepers.TokenFactoryKeeper.SetParams(ctx, tokenfactoryDefaultParams)

		ctx.Logger().Info("Upgrade complete")
		return vm, err
	}
}
