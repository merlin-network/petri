package keeper

import (
	"github.com/merlin-network/petri/x/interchaintxs/types"
)

var _ types.QueryServer = Keeper{}
