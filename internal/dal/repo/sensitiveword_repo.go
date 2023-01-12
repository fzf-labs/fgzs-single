package repo

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"
	"gorm.io/gorm"
)

var _ iSensitiveWordRepo = (*SensitiveWordRepo)(nil)

type SensitiveWordRepo struct {
	ctx context.Context
	db  *gorm.DB
}
type iSensitiveWordRepo interface {
	IdToName(ids []int64) (map[int64]string, error)
}

func NewSensitiveWordRepo(ctx context.Context, DB *gorm.DB) *SensitiveWordRepo {
	return &SensitiveWordRepo{ctx: ctx, db: DB}
}

func (s *SensitiveWordRepo) IdToName(ids []int64) (map[int64]string, error) {
	result := make(map[int64]string)
	sensitiveWordCategoryDao := dao.Use(s.db).SensitiveWordCategory
	res, err := sensitiveWordCategoryDao.WithContext(s.ctx).Where(sensitiveWordCategoryDao.ID.In(ids...)).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	for _, v := range res {
		result[v.ID] = v.Name
	}
	return nil, nil
}
