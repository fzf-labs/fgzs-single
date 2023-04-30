package service

import (
	"context"

	"gorm.io/gorm"
)

var _ iDemoService = (*DemoService)(nil)

type DemoService struct {
	ctx context.Context
	db  *gorm.DB
}
type iDemoService interface {
	IdToName(ids []int64) (map[int64]string, error)
}

func NewDemoService(ctx context.Context, DB *gorm.DB) *DemoService {
	return &DemoService{ctx: ctx, db: DB}
}

func (s *DemoService) IdToName(ids []int64) (map[int64]string, error) {
	return nil, nil
}
