package sensitiveword

import (
	"context"
	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/dal/model"
	"fgzs-single/internal/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type SensitiveWordStoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSensitiveWordStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SensitiveWordStoreLogic {
	return &SensitiveWordStoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SensitiveWordStoreLogic) SensitiveWordStore(req *types.SensitiveWordStoreReq) (resp *types.SensitiveWordStoreResp, err error) {
	resp = new(types.SensitiveWordStoreResp)
	sensitiveWordDao := dao.Use(l.svcCtx.Gorm).SensitiveWord
	if req.Id > 0 {
		_, err = sensitiveWordDao.WithContext(l.ctx).Select(sensitiveWordDao.ID, sensitiveWordDao.CategoryID, sensitiveWordDao.Word).Where(sensitiveWordDao.ID.Eq(req.Id)).Updates(model.SensitiveWord{
			ID:         req.Id,
			CategoryID: req.CategoryID,
			Word:       req.Word,
		})
		if err != nil {
			return nil, errorx.DataSqlErr.WithDetail(err)
		}
	} else {
		err := sensitiveWordDao.WithContext(l.ctx).Create(&model.SensitiveWord{
			ID:         req.Id,
			CategoryID: req.CategoryID,
			Word:       req.Word,
		})
		if err != nil {
			return nil, errorx.DataSqlErr.WithDetail(err)
		}
	}
	return
}
