package file

import (
	"context"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileStaticLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileStaticLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileStaticLogic {
	return &FileStaticLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileStaticLogic) FileStatic(req *types.FileStaticReq) (resp *types.FileStaticResp, err error) {
	// todo: add your logic here and delete this line

	return
}
