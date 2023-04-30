package system

import (
	"context"
	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/fzf-labs/fpkg/util/osutil"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type StatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatLogic {
	return &StatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatLogic) Stat(req *types.StatReq) (resp *types.StatResp, err error) {
	resp = new(types.StatResp)
	cpuInfo, err := osutil.GetCpuInfo()
	if err != nil {
		return nil, err
	}
	memInfo, err := osutil.GetMemInfo()
	if err != nil {
		return nil, err
	}
	diskInfo, err := osutil.GetDiskInfo()
	if err != nil {
		return nil, err
	}
	sysInfo, err := osutil.GetSysInfo()
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp.Cpu, cpuInfo)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp.Memory, memInfo)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp.Disk, diskInfo)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp.Sys, sysInfo)
	if err != nil {
		return nil, err
	}
	return
}
