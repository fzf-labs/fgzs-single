package permmenu

import (
	"context"
	"fgzs-single/internal/dal/dao"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysPermMenuDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysPermMenuDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysPermMenuDelLogic {
	return &SysPermMenuDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysPermMenuDelLogic) SysPermMenuDel(req *types.SysPermMenuDelReq) (*types.SysPermMenuDelResp, error) {
	resp := new(types.SysPermMenuDelResp)
	sysPermMenuDao := dao.Use(l.svcCtx.Gorm).SysPermMenu
	_, err := sysPermMenuDao.WithContext(l.ctx).Where(sysPermMenuDao.ID.In(req.Ids...)).Delete()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
