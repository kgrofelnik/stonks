package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kgrofelnik/stonks/x/stonks/types"
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/tools"
)

func (k msgServer) Ask(goCtx context.Context, msg *types.MsgAsk) (*types.MsgAskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	answer, err := askAgent(goCtx, msg.Question)
	if err != nil {
		fmt.Println(err)
	}
	return &types.MsgAskResponse{Answer: answer, Id: 1}, nil
}

func askAgent(goCtx context.Context, question string) (string, error) {
	llm, err := openai.New()
	if err != nil {
		return "", err
	}
	agentTools := []tools.Tool{
		tools.Calculator{},
	}
	executor, err := agents.Initialize(
		llm,
		agentTools,
		agents.ZeroShotReactDescription,
		agents.WithMaxIterations(3),
	)
	if err != nil {
		return "", err
	}
	answer, err := chains.Run(goCtx, executor, question)
	fmt.Println(answer)
	return answer, err
}
