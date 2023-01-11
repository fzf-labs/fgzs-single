package file

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDelLogic {
	return &FileDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileDelLogic) FileDel(req *types.FileDelReq) (resp *types.FileDelResp, err error) {
	fileUploadDao := dao.Use(l.svcCtx.Gorm).FileUpload
	_, err = fileUploadDao.WithContext(l.ctx).Where(fileUploadDao.ID.In(req.Ids...)).Delete()
	if err != nil {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	return
}
