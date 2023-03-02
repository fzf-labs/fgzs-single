package tool

import (
	"context"

	"fgzs-single/internal/app/web/internal/svc"
	"fgzs-single/internal/app/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WebsocketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWebsocketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WebsocketLogic {
	return &WebsocketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WebsocketLogic) Websocket(req *types.WebsocketReq) (resp *types.WebsocketResp, err error) {
	return
}
