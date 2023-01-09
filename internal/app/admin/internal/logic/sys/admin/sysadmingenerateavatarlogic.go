package admin

import (
	"context"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysAdminGenerateAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysAdminGenerateAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysAdminGenerateAvatarLogic {
	return &SysAdminGenerateAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysAdminGenerateAvatarLogic) SysAdminGenerateAvatar(req *types.SysAdminGenerateAvatarReq) (resp *types.SysAdminGenerateAvatarResp, err error) {
	// todo: add your logic here and delete this line

	return
}
