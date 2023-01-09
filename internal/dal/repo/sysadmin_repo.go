package repo

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"
	"gorm.io/gorm"
)

var _ iSysAdminRepo = (*SysAdminRepo)(nil)

type SysAdminRepo struct {
	ctx context.Context
	db  *gorm.DB
}
type iSysAdminRepo interface {
	IdToName(ids []int64) (map[int64]string, map[int64]string, error)
}

func NewSysAdminRepo(ctx context.Context, DB *gorm.DB) *SysAdminRepo {
	return &SysAdminRepo{ctx: ctx, db: DB}
}

func (s *SysAdminRepo) IdToName(ids []int64) (map[int64]string, map[int64]string, error) {
	idToUserName := make(map[int64]string)
	idToNickName := make(map[int64]string)
	sysAdminDao := dao.Use(s.db).SysAdmin
	sysAdmins, err := sysAdminDao.WithContext(s.ctx).Where(sysAdminDao.ID.In(ids...)).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, nil, errorx.DataSqlErr.WithDetail(err)
	}
	for _, v := range sysAdmins {
		idToUserName[v.ID] = v.Username
		idToNickName[v.ID] = v.Nickname
	}
	return idToUserName, idToNickName, nil
}
