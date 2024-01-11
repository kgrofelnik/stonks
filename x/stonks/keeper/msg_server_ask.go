package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kgrofelnik/stonks/x/stonks/types"
)

func (k msgServer) Ask(goCtx context.Context, msg *types.MsgAsk) (*types.MsgAskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAskResponse{}, nil
}
