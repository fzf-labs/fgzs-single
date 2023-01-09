package admin

import (
	"context"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysAdminLogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysAdminLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysAdminLogoutLogic {
	return &SysAdminLogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysAdminLogoutLogic) SysAdminLogout(req *types.SysAdminLogoutReq) (resp *types.SysAdminLogoutResp, err error) {

	return
}
