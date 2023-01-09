package job

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"gorm.io/gorm"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysJobDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysJobDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysJobDelLogic {
	return &SysJobDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysJobDelLogic) SysJobDel(req *types.SysJobDelReq) (resp *types.SysJobDelReq, err error) {
	resp = new(types.SysJobDelReq)
	sysJobDao := dao.Use(l.svcCtx.Gorm).SysJob
	_, err = sysJobDao.WithContext(l.ctx).Where(sysJobDao.ID.In(req.Ids...)).Delete()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}
