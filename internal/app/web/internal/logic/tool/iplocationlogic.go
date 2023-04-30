package tool

import (
	"context"
	"fgzs-single/internal/app/web/internal/svc"
	"fgzs-single/internal/app/web/internal/types"

	"github.com/fzf-labs/fpkg/iplocation"

	"github.com/zeromicro/go-zero/core/logx"
)

type IpLocationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIpLocationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IpLocationLogic {
	return &IpLocationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IpLocationLogic) IpLocation(req *types.IpLocationReq) (resp *types.IpLocationResp, err error) {
	resp = new(types.IpLocationResp)
	resp.Ip = req.Ip
	ipLocation, err := iplocation.NewIpLocation(l.svcCtx.Config.IpLocation.Path)
	if err != nil {
		return resp, nil
	}
	location, err := ipLocation.SearchLocation(req.Ip)
	if err != nil {
		return nil, err
	}

	resp.Location = location
	return
}
