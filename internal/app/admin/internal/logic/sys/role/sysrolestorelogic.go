package role

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/dal/model"
	"fgzs-single/internal/errorx"
	"fgzs-single/pkg/util/sliutil"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysRoleStoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysRoleStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysRoleStoreLogic {
	return &SysRoleStoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysRoleStoreLogic) SysRoleStore(req *types.SysRoleStoreReq) (*types.SysRoleStoreResp, error) {
	resp := new(types.SysRoleStoreResp)
	sysRoleDao := dao.Use(l.svcCtx.Gorm).SysRole
	toString := sliutil.JoinInt64ToString(req.PermMenuIds)
	if req.Id > 0 {
		_, err := sysRoleDao.WithContext(l.ctx).Where(sysRoleDao.ID.Eq(req.Id)).Updates(&model.SysRole{
			Pid:         req.Pid,
			Name:        req.Name,
			PermMenuIds: toString,
			Remark:      req.Remark,
			Status:      req.Status,
		})
		if err != nil {
			return nil, errorx.DataSqlErr.WithDetail(err)
		}
	} else {
		err := sysRoleDao.WithContext(l.ctx).Create(&model.SysRole{
			Pid:         req.Pid,
			Name:        req.Name,
			PermMenuIds: toString,
			Remark:      req.Remark,
			Status:      req.Status,
		})
		if err != nil {
			return nil, errorx.DataSqlErr.WithDetail(err)
		}
	}

	return resp, nil
}
