package role

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"gorm.io/gorm"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysRoleDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysRoleDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysRoleDelLogic {
	return &SysRoleDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysRoleDelLogic) SysRoleDel(req *types.SysRoleDelReq) (*types.SysRoleDelReq, error) {
	resp := new(types.SysRoleDelReq)
	sysRoleDao := dao.Use(l.svcCtx.Gorm).SysRole
	_, err := sysRoleDao.WithContext(l.ctx).Where(sysRoleDao.ID.In(req.Ids...)).Delete()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return resp, nil
}
