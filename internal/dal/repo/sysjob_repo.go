package repo

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"

	"gorm.io/gorm"
)

var _ iSysJobRepo = (*SysJobRepo)(nil)

type SysJobRepo struct {
	ctx context.Context
	db  *gorm.DB
}
type iSysJobRepo interface {
	IdToName(ids []int64) (map[int64]string, error)
}

func NewSysJobRepo(ctx context.Context, DB *gorm.DB) *SysJobRepo {
	return &SysJobRepo{ctx: ctx, db: DB}
}

func (s *SysJobRepo) IdToName(ids []int64) (map[int64]string, error) {
	result := make(map[int64]string)
	sysJobDao := dao.Use(s.db).SysJob
	sysJobs, err := sysJobDao.WithContext(s.ctx).Where(sysJobDao.ID.In(ids...)).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	for _, v := range sysJobs {
		result[v.ID] = v.Name
	}
	return result, nil
}
