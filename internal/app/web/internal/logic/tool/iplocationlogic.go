package tool

import (
	"context"
	"fgzs-single/pkg/iplocation"

	"fgzs-single/internal/app/web/internal/svc"
	"fgzs-single/internal/app/web/internal/types"

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
	ipLocation, err := iplocation.NewIpLocation("/etc/ip2region.xdb")
	if err != nil {
		return nil, err
	}
	location, err := ipLocation.SearchLocation(req.Ip)
	if err != nil {
		return nil, err
	}
	resp.Ip = req.Ip
	resp.Location = location
	return
}
