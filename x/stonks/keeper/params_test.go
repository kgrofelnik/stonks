package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/kgrofelnik/stonks/testutil/keeper"
	"github.com/kgrofelnik/stonks/x/stonks/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.StonksKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
