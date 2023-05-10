package keeper

import (
	"github.com/petri-labs/petri/x/contractmanager/types"
)

var _ types.QueryServer = Keeper{}
