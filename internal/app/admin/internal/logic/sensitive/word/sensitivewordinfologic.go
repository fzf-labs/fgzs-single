package word

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"

	"github.com/fzf-labs/fpkg/util/timeutil"
	"gorm.io/gorm"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SensitiveWordInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSensitiveWordInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SensitiveWordInfoLogic {
	return &SensitiveWordInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SensitiveWordInfoLogic) SensitiveWordInfo(req *types.SensitiveWordInfoReq) (resp *types.SensitiveWordInfoResp, err error) {
	resp = new(types.SensitiveWordInfoResp)
	sensitiveWordDao := dao.Use(l.svcCtx.Gorm).SensitiveWord
	sensitiveCategoryDao := dao.Use(l.svcCtx.Gorm).SensitiveCategory
	sensitiveWord, err := sensitiveWordDao.WithContext(l.ctx).Where(sensitiveWordDao.ID.Eq(req.Id)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	sensitiveWordCategory, err := sensitiveCategoryDao.WithContext(l.ctx).Where(sensitiveCategoryDao.ID.Eq(sensitiveWord.CategoryID)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	resp.Info = types.SensitiveWordInfo{
		Id:           sensitiveWord.ID,
		CategoryID:   sensitiveWord.CategoryID,
		CategoryName: sensitiveWordCategory.Name,
		Text:         sensitiveWord.Text,
		CreatedAt:    timeutil.ToDateTimeStringByTime(sensitiveWord.CreatedAt),
		UpdatedAt:    timeutil.ToDateTimeStringByTime(sensitiveWord.UpdatedAt),
	}
	return
}
