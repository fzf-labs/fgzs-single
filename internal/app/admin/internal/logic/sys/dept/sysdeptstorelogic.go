package dept

import (
	"context"
	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/dal/model"
	"fgzs-single/internal/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDeptStoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDeptStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDeptStoreLogic {
	return &SysDeptStoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDeptStoreLogic) SysDeptStore(req *types.SysDeptStoreReq) (resp *types.SysDeptStoreResp, err error) {
	resp = new(types.SysDeptStoreResp)
	sysDeptDao := dao.Use(l.svcCtx.Gorm).SysDept
	if req.Id > 0 {
		_, err = sysDeptDao.WithContext(l.ctx).Select(sysDeptDao.Pid, sysDeptDao.Name, sysDeptDao.FullName, sysDeptDao.Responsible, sysDeptDao.Phone, sysDeptDao.Email, sysDeptDao.Type, sysDeptDao.Status, sysDeptDao.Sort).Where(sysDeptDao.ID.Eq(req.Id)).Updates(model.SysDept{
			Pid:         req.Pid,
			Name:        req.Name,
			FullName:    req.FullName,
			Responsible: req.Responsible,
			Phone:       req.Phone,
			Email:       req.Email,
			Type:        req.Type,
			Status:      req.Status,
			Sort:        req.Sort,
		})
		if err != nil {
			return nil, errorx.DataSqlErr.WithDetail(err)
		}

	} else {
		err := sysDeptDao.WithContext(l.ctx).Create(&model.SysDept{
			Pid:         req.Pid,
			Name:        req.Name,
			FullName:    req.FullName,
			Responsible: req.Responsible,
			Phone:       req.Phone,
			Email:       req.Email,
			Type:        req.Type,
			Status:      req.Status,
			Sort:        req.Sort,
		})
		if err != nil {
			return nil, errorx.DataSqlErr.WithDetail(err)
		}
	}
	return
}
