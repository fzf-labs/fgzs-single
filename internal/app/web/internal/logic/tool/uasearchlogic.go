package tool

import (
	"context"

	"github.com/fzf-labs/fpkg/browser"
	"github.com/jinzhu/copier"

	"fgzs-single/internal/app/web/internal/svc"
	"fgzs-single/internal/app/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UaSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUaSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UaSearchLogic {
	return &UaSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UaSearchLogic) UaSearch(req *types.UaSearchReq) (resp *types.UaSearchResp, err error) {
	resp = new(types.UaSearchResp)
	location := browser.UaParse(req.Ua)
	resp.Ua = req.Ua
	err = copier.Copy(&resp.Location, location)
	if err != nil {
		return nil, err
	}
	return
}
