package category

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SensitiveCategoryDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSensitiveCategoryDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SensitiveCategoryDelLogic {
	return &SensitiveCategoryDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SensitiveCategoryDelLogic) SensitiveCategoryDel(req *types.SensitiveCategoryDelReq) (resp *types.SensitiveCategoryDelReq, err error) {
	sensitiveCategoryDao := dao.Use(l.svcCtx.Gorm).SensitiveCategory
	_, err = sensitiveCategoryDao.WithContext(l.ctx).Where(sensitiveCategoryDao.ID.In(req.Ids...)).Delete()
	if err != nil {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	return
}
