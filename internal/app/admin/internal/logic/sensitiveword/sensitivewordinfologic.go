package sensitiveword

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
	sensitiveWordCategoryDao := dao.Use(l.svcCtx.Gorm).SensitiveWordCategory
	sensitiveWord, err := sensitiveWordDao.WithContext(l.ctx).Where(sensitiveWordDao.ID.Eq(req.Id)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	sensitiveWordCategory, err := sensitiveWordCategoryDao.WithContext(l.ctx).Where(sensitiveWordCategoryDao.ID.Eq(sensitiveWord.CategoryID)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	resp.Info = types.SensitiveWordInfo{
		Id:           sensitiveWord.ID,
		CategoryID:   sensitiveWord.CategoryID,
		CategoryName: sensitiveWordCategory.Name,
		Word:         sensitiveWord.Word,
		CreatedAt:    timeutil.ToDateTimeStringByTime(sensitiveWord.CreatedAt),
		UpdatedAt:    timeutil.ToDateTimeStringByTime(sensitiveWord.UpdatedAt),
	}
	return
}
