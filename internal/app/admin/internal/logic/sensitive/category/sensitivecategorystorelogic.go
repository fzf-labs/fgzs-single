package category

import (
	"context"
	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/dal/model"
	"fgzs-single/internal/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type SensitiveCategoryStoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSensitiveCategoryStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SensitiveCategoryStoreLogic {
	return &SensitiveCategoryStoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SensitiveCategoryStoreLogic) SensitiveCategoryStore(req *types.SensitiveCategoryStoreReq) (resp *types.SensitiveCategoryStoreResp, err error) {
	resp = new(types.SensitiveCategoryStoreResp)
	sensitiveCategoryDao := dao.Use(l.svcCtx.Gorm).SensitiveCategory
	if req.Id > 0 {
		_, err = sensitiveCategoryDao.WithContext(l.ctx).Select(sensitiveCategoryDao.ID, sensitiveCategoryDao.Name, sensitiveCategoryDao.Label).Where(sensitiveCategoryDao.ID.Eq(req.Id)).Updates(model.SensitiveCategory{
			ID:    req.Id,
			Label: req.Label,
			Name:  req.Name,
		})
		if err != nil {
			return nil, errorx.DataSqlErr.WithDetail(err)
		}
	} else {
		err := sensitiveCategoryDao.WithContext(l.ctx).Create(&model.SensitiveCategory{
			ID:    req.Id,
			Label: req.Label,
			Name:  req.Name,
		})
		if err != nil {
			return nil, errorx.DataSqlErr.WithDetail(err)
		}
	}
	return
}
