package keeper

import (
	"github.com/petri-labs/petri/x/cron/types"
)

var _ types.QueryServer = Keeper{}
