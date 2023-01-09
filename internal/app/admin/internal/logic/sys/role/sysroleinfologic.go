package role

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"
	"fgzs-single/pkg/util/sliutil"
	"fgzs-single/pkg/util/timeutil"
	"gorm.io/gorm"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysRoleInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysRoleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysRoleInfoLogic {
	return &SysRoleInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysRoleInfoLogic) SysRoleInfo(req *types.SysRoleInfoReq) (*types.SysRoleInfoResp, error) {
	resp := new(types.SysRoleInfoResp)
	sysRoleDao := dao.Use(l.svcCtx.Gorm).SysRole
	sysRole, err := sysRoleDao.WithContext(l.ctx).Where(sysRoleDao.ID.Eq(req.Id)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	permMenuIds := make([]int64, 0)
	if sysRole.PermMenuIds == "*" {
		sysPermMenu := dao.Use(l.svcCtx.Gorm).SysPermMenu
		err := sysPermMenu.WithContext(l.ctx).Pluck(sysPermMenu.ID, &permMenuIds)
		if err != nil {
			return nil, errorx.DataSqlErr.WithDetail(err)
		}
	} else {
		permMenuIds, err = sliutil.SplitStringToInt64(sysRole.PermMenuIds)
		if err != nil {
			return nil, err
		}
	}
	resp.Info = types.SysRole{
		Id:          sysRole.ID,
		Pid:         sysRole.Pid,
		Name:        sysRole.Name,
		Remark:      sysRole.Remark,
		Status:      sysRole.Status,
		Sort:        sysRole.Sort,
		PermMenuIds: permMenuIds,
		CreatedAt:   timeutil.ToDateTimeStringByTime(sysRole.CreatedAt),
		UpdatedAt:   timeutil.ToDateTimeStringByTime(sysRole.UpdatedAt),
		Children:    nil,
	}
	return resp, nil
}
