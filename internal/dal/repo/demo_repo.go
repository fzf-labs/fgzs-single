package repo

import (
	"context"
	"gorm.io/gorm"
)

var _ iDemoRepo = (*DemoRepo)(nil)

type DemoRepo struct {
	ctx context.Context
	db  *gorm.DB
}
type iDemoRepo interface {
	IdToName(ids []int64) (map[int64]string, error)
}

func NewDemoRepo(ctx context.Context, DB *gorm.DB) *DemoRepo {
	return &DemoRepo{ctx: ctx, db: DB}
}

func (s *DemoRepo) IdToName(ids []int64) (map[int64]string, error) {
	return nil, nil
}
