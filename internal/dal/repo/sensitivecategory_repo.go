package repo

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"
	"gorm.io/gorm"
)

var _ iSensitiveCategoryRepo = (*SensitiveCategoryRepo)(nil)

type SensitiveCategoryRepo struct {
	ctx context.Context
	db  *gorm.DB
}
type iSensitiveCategoryRepo interface {
	IdToName(ids []int64) (map[int64]string, error)
}

func NewSensitiveCategoryRepo(ctx context.Context, DB *gorm.DB) *SensitiveCategoryRepo {
	return &SensitiveCategoryRepo{ctx: ctx, db: DB}
}

func (s *SensitiveCategoryRepo) IdToName(ids []int64) (map[int64]string, error) {
	result := make(map[int64]string)
	sensitiveCategoryDao := dao.Use(s.db).SensitiveCategory
	res, err := sensitiveCategoryDao.WithContext(s.ctx).Where(sensitiveCategoryDao.ID.In(ids...)).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	for _, v := range res {
		result[v.ID] = v.Name
	}
	return result, nil
}
