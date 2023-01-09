package repo

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"
	"gorm.io/gorm"
)

var _ iSysApiRepo = (*SysApiRepo)(nil)

type SysApiRepo struct {
	ctx context.Context
	db  *gorm.DB
}
type iSysApiRepo interface {
	IdToName(paths []string) (map[string]string, error)
}

func NewSysApiRepo(ctx context.Context, DB *gorm.DB) *SysApiRepo {
	return &SysApiRepo{ctx: ctx, db: DB}
}

func (s *SysApiRepo) IdToName(paths []string) (map[string]string, error) {
	result := make(map[string]string)
	sysAPIDao := dao.Use(s.db).SysAPI
	sysAPIS, err := sysAPIDao.WithContext(s.ctx).Where(sysAPIDao.Path.In(paths...)).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	for _, api := range sysAPIS {
		result[api.Path] = api.Desc
	}
	return result, nil
}
