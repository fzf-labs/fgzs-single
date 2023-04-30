package dept

import (
	"context"
	"fgzs-single/internal/dal/dao"

	"gorm.io/gorm"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDeptDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDeptDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDeptDelLogic {
	return &SysDeptDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDeptDelLogic) SysDeptDel(req *types.SysDeptDelReq) (resp *types.SysDeptDelReq, err error) {
	resp = new(types.SysDeptDelReq)
	sysDeptDao := dao.Use(l.svcCtx.Gorm).SysDept
	_, err = sysDeptDao.WithContext(l.ctx).Where(sysDeptDao.ID.In(req.Ids...)).Delete()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}
