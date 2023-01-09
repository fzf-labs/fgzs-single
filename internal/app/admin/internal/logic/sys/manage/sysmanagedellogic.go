package manage

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"
	"gorm.io/gorm"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysManageDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysManageDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysManageDelLogic {
	return &SysManageDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysManageDelLogic) SysManageDel(req *types.SysManageDelReq) (resp *types.SysManageDelReq, err error) {
	sysAdminDao := dao.Use(l.svcCtx.Gorm).SysAdmin
	_, err = sysAdminDao.WithContext(l.ctx).Where(sysAdminDao.ID.In(req.Ids...)).Delete()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	return
}
