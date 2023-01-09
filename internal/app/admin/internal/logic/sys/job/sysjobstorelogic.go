package job

import (
	"context"
	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/dal/model"
	"fgzs-single/internal/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysJobStoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysJobStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysJobStoreLogic {
	return &SysJobStoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysJobStoreLogic) SysJobStore(req *types.SysJobStoreReq) (resp *types.SysJobStoreResp, err error) {
	resp = new(types.SysJobStoreResp)
	sysJobDao := dao.Use(l.svcCtx.Gorm).SysJob
	if req.Id > 0 {
		_, err = sysJobDao.WithContext(l.ctx).Select(sysJobDao.Name, sysJobDao.Code, sysJobDao.Remark, sysJobDao.Sort, sysJobDao.Status).Where(sysJobDao.ID.Eq(req.Id)).Updates(model.SysJob{
			Name:   req.Name,
			Code:   req.Code,
			Remark: req.Remark,
			Sort:   req.Sort,
			Status: req.Status,
		})
		if err != nil {
			return nil, errorx.DataSqlErr.WithDetail(err)
		}
	} else {
		err := sysJobDao.WithContext(l.ctx).Create(&model.SysJob{
			Name:   req.Name,
			Code:   req.Code,
			Remark: req.Remark,
			Sort:   req.Sort,
			Status: req.Status,
		})
		if err != nil {
			return nil, errorx.DataSqlErr.WithDetail(err)
		}
	}
	return
}
