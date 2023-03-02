package tool

import (
	"context"

	"fgzs-single/internal/app/web/internal/svc"
	"fgzs-single/internal/app/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatGPTLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatGPTLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatGPTLogic {
	return &ChatGPTLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatGPTLogic) ChatGPT(req *types.ChatGPTReq) (resp *types.ChatGPTResp, err error) {

	return
}
