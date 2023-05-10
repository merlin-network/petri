package keeper

import (
	"github.com/merlin-network/petri/x/cron/types"
)

var _ types.QueryServer = Keeper{}
