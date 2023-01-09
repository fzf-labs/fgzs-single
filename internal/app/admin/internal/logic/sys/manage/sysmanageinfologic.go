package manage

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/dal/repo"
	"fgzs-single/internal/errorx"
	"fgzs-single/pkg/util/jsonutil"
	"fgzs-single/pkg/util/timeutil"
	"gorm.io/gorm"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysManageInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysManageInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysManageInfoLogic {
	return &SysManageInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysManageInfoLogic) SysManageInfo(req *types.SysManageInfoReq) (*types.SysManageInfoResp, error) {
	resp := new(types.SysManageInfoResp)
	sysAdminDao := dao.Use(l.svcCtx.Gorm).SysAdmin
	sysAdmin, err := sysAdminDao.WithContext(l.ctx).Where(sysAdminDao.ID.Eq(req.Id)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	jobIdToName, err := repo.NewSysJobRepo(l.ctx, l.svcCtx.Gorm).IdToName([]int64{sysAdmin.JobID})
	if err != nil {
		return nil, err
	}
	roleIdToName, err := repo.NewSysRoleRepo(l.ctx, l.svcCtx.Gorm).IdToName([]int64{sysAdmin.JobID})
	if err != nil {
		return nil, err
	}
	deptToName, err := repo.NewSysDeptRepo(l.ctx, l.svcCtx.Gorm).IdToName([]int64{sysAdmin.JobID})
	if err != nil {
		return nil, err
	}
	roleIds := make([]int64, 0)
	roleNames := make([]string, 0)
	err = jsonutil.Decode(sysAdmin.RoleIds, &roleIds)
	if err != nil {
		return nil, err
	}
	for _, v := range roleIds {
		roleNames = append(roleNames, roleIdToName[v])
	}
	resp.Info = types.SysManageInfo{
		Id:        sysAdmin.ID,
		Username:  sysAdmin.Username,
		Nickname:  sysAdmin.Nickname,
		Avatar:    sysAdmin.Avatar,
		Gender:    sysAdmin.Gender,
		Email:     sysAdmin.Email,
		Mobile:    sysAdmin.Mobile,
		JobID:     sysAdmin.JobID,
		DeptID:    sysAdmin.DeptID,
		RoleIds:   roleIds,
		JobName:   jobIdToName[sysAdmin.JobID],
		DeptName:  deptToName[sysAdmin.JobID],
		RoleNames: roleNames,
		Motto:     sysAdmin.Motto,
		Status:    sysAdmin.Status,
		CreatedAt: timeutil.ToDateTimeStringByTime(sysAdmin.CreatedAt),
		UpdatedAt: timeutil.ToDateTimeStringByTime(sysAdmin.UpdatedAt),
	}
	return resp, nil
}
