package file

import (
	"context"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileInfoLogic {
	return &FileInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileInfoLogic) FileInfo(req *types.FileInfoReq) (resp *types.FileInfoResp, err error) {

	return
}
