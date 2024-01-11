package keeper

import (
	"github.com/kgrofelnik/stonks/x/stonks/types"
)

var _ types.QueryServer = Keeper{}
