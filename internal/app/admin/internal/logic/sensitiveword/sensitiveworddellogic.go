package sensitiveword

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SensitiveWordDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSensitiveWordDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SensitiveWordDelLogic {
	return &SensitiveWordDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SensitiveWordDelLogic) SensitiveWordDel(req *types.SensitiveWordDelReq) (resp *types.SensitiveWordDelReq, err error) {
	sensitiveWordDao := dao.Use(l.svcCtx.Gorm).SensitiveWord
	_, err = sensitiveWordDao.WithContext(l.ctx).Where(sensitiveWordDao.ID.In(req.Ids...)).Delete()
	if err != nil {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	return
}
