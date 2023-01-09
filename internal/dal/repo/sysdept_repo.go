package repo

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"
	"gorm.io/gorm"
)

var _ iSysDeptRepo = (*SysDeptRepo)(nil)

type SysDeptRepo struct {
	ctx context.Context
	db  *gorm.DB
}

func NewSysDeptRepo(ctx context.Context, DB *gorm.DB) *SysDeptRepo {
	return &SysDeptRepo{ctx: ctx, db: DB}
}

type iSysDeptRepo interface {
	IdToName(ids []int64) (map[int64]string, error)
}

func (s *SysDeptRepo) IdToName(ids []int64) (map[int64]string, error) {
	result := make(map[int64]string)
	sysDeptDao := dao.Use(s.db).SysDept
	sysDepts, err := sysDeptDao.WithContext(s.ctx).Where(sysDeptDao.ID.In(ids...)).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	for _, v := range sysDepts {
		result[v.ID] = v.Name
	}
	return result, nil
}
