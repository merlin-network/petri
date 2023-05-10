package keeper_test

import (
	"testing"

	"github.com/merlin-network/petri/testutil"

	"github.com/merlin-network/petri/app"

	testkeeper "github.com/merlin-network/petri/testutil/cron/keeper"

	"github.com/merlin-network/petri/x/cron/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	_ = app.GetDefaultConfig()

	k, ctx := testkeeper.CronKeeper(t, nil, nil)
	params := types.Params{
		SecurityAddress: testutil.TestOwnerAddress,
		Limit:           5,
	}

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
