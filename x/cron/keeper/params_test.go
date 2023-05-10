package keeper_test

import (
	"testing"

	"github.com/petri-labs/petri/testutil"

	"github.com/petri-labs/petri/app"

	testkeeper "github.com/petri-labs/petri/testutil/cron/keeper"

	"github.com/petri-labs/petri/x/cron/types"
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
