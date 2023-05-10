package keeper

import (
	"github.com/merlin-network/petri/x/contractmanager/types"
)

var _ types.QueryServer = Keeper{}
