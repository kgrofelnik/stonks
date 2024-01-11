package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kgrofelnik/stonks/x/stonks/types"
	"github.com/tmc/langchaingo/llms/openai"
)

func (k msgServer) Ask(goCtx context.Context, msg *types.MsgAsk) (*types.MsgAskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	llm, err := openai.New()
	if err != nil {
		fmt.Println(err)
	}
	completion, err := llm.Call(ctx, msg.Question)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(completion)
	return &types.MsgAskResponse{Answer: completion, Id: 1}, nil
}
