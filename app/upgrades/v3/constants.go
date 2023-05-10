package v3

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	ccvprovider "github.com/cosmos/interchain-security/x/ccv/provider/types"

	"github.com/petri-labs/petri/app/upgrades"
)

const (
	// UpgradeName defines the on-chain upgrades name.
	UpgradeName = "v0.3.0"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			ccvprovider.ModuleName,
		},
	},
}
