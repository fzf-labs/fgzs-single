package repo

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"
	"gorm.io/gorm"
)

var _ iSysRoleRepo = (*SysRoleRepo)(nil)

type SysRoleRepo struct {
	ctx context.Context
	db  *gorm.DB
}
type iSysRoleRepo interface {
	IdToName(ids []int64) (map[int64]string, error)
}

func NewSysRoleRepo(ctx context.Context, DB *gorm.DB) *SysRoleRepo {
	return &SysRoleRepo{ctx: ctx, db: DB}
}

func (s *SysRoleRepo) IdToName(ids []int64) (map[int64]string, error) {
	result := make(map[int64]string)
	sysRoleDao := dao.Use(s.db).SysRole
	sysRoles, err := sysRoleDao.WithContext(s.ctx).Where(sysRoleDao.ID.In(ids...)).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	for _, v := range sysRoles {
		result[v.ID] = v.Name
	}
	return result, nil
}
