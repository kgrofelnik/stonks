package stonks_test

import (
	"testing"

	keepertest "github.com/kgrofelnik/stonks/testutil/keeper"
	"github.com/kgrofelnik/stonks/testutil/nullify"
	"github.com/kgrofelnik/stonks/x/stonks/module"
	"github.com/kgrofelnik/stonks/x/stonks/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StonksKeeper(t)
	stonks.InitGenesis(ctx, k, genesisState)
	got := stonks.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
