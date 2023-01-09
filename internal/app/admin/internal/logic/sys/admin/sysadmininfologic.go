package admin

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"
	"fgzs-single/internal/meta"
	"fgzs-single/pkg/util/jsonutil"
	"gorm.io/gorm"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysAdminInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysAdminInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysAdminInfoLogic {
	return &SysAdminInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysAdminInfoLogic) SysAdminInfo(req *types.SysAdminInfoReq) (*types.SysAdminInfoResp, error) {
	resp := new(types.SysAdminInfoResp)
	adminId := meta.GetAdminId(l.ctx)
	sysAdminDao := dao.Use(l.svcCtx.Gorm).SysAdmin
	sysAdmin, err := sysAdminDao.WithContext(l.ctx).Where(sysAdminDao.ID.Eq(adminId)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	if sysAdmin.Status != 1 {
		return nil, errorx.AccountIsBanned
	}
	roleIds := make([]int64, 0)
	err = jsonutil.Decode(sysAdmin.RoleIds, &roleIds)
	if err != nil {
		return nil, errorx.DataFormattingError.WithDetail(err)
	}
	resp.Info = types.SysAdminInfo{
		ID:       sysAdmin.ID,
		Username: sysAdmin.Username,
		Nickname: sysAdmin.Nickname,
		Avatar:   sysAdmin.Avatar,
		Gender:   sysAdmin.Gender,
		Email:    sysAdmin.Email,
		Mobile:   sysAdmin.Mobile,
		JobID:    sysAdmin.JobID,
		DeptID:   sysAdmin.DeptID,
		RoleIds:  roleIds,
		Motto:    sysAdmin.Motto,
	}
	return resp, nil
}
