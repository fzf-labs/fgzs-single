package category

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"
	"fgzs-single/pkg/util/timeutil"
	"gorm.io/gorm"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SensitiveCategoryInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSensitiveCategoryInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SensitiveCategoryInfoLogic {
	return &SensitiveCategoryInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SensitiveCategoryInfoLogic) SensitiveCategoryInfo(req *types.SensitiveCategoryInfoReq) (resp *types.SensitiveCategoryInfoResp, err error) {
	resp = new(types.SensitiveCategoryInfoResp)
	sensitiveCategoryDao := dao.Use(l.svcCtx.Gorm).SensitiveCategory
	sensitiveCategory, err := sensitiveCategoryDao.WithContext(l.ctx).Where(sensitiveCategoryDao.ID.Eq(req.Id)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	resp.Info = types.SensitiveCategoryInfo{
		Id:        sensitiveCategory.ID,
		Label:     sensitiveCategory.Label,
		Name:      sensitiveCategory.Name,
		CreatedAt: timeutil.ToDateTimeStringByTime(sensitiveCategory.CreatedAt),
		UpdatedAt: timeutil.ToDateTimeStringByTime(sensitiveCategory.UpdatedAt),
	}
	return
}
