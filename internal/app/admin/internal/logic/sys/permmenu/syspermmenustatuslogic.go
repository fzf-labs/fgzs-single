package permmenu

import (
	"context"
	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"
	"fgzs-single/internal/dal/dao"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysPermMenuStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysPermMenuStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysPermMenuStatusLogic {
	return &SysPermMenuStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysPermMenuStatusLogic) SysPermMenuStatus(req *types.SysPermMenuStatusReq) (resp *types.SysPermMenuStatusResp, err error) {
	resp = new(types.SysPermMenuStatusResp)
	sysPermMenuDao := dao.Use(l.svcCtx.Gorm).SysPermMenu
	_, err = sysPermMenuDao.WithContext(l.ctx).Where(sysPermMenuDao.ID.Eq(req.Id)).UpdateSimple(sysPermMenuDao.Status.Value(req.Status))
	if err != nil {
		return nil, err
	}
	return
}
