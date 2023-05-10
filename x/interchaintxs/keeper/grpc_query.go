package keeper

import (
	"github.com/petri-labs/petri/x/interchaintxs/types"
)

var _ types.QueryServer = Keeper{}
