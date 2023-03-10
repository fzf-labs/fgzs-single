package manage

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"
	"github.com/fzf-labs/fpkg/conv"
	"github.com/fzf-labs/fpkg/page"
	"github.com/fzf-labs/fpkg/util/jsonutil"
	"github.com/fzf-labs/fpkg/util/sliutil"
	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"strings"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysManageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysManageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysManageListLogic {
	return &SysManageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysManageListLogic) SysManageList(req *types.SysManageListReq) (resp *types.SysManageListResp, err error) {
	resp = new(types.SysManageListResp)
	roleIds := make([]int64, 0)
	jobIds := make([]int64, 0)
	deptIds := make([]int64, 0)
	roleIdToName := make(map[int64]string)
	jobIdToName := make(map[int64]string)
	deptIdToName := make(map[int64]string)
	list := make([]types.SysManageInfo, 0)
	sysAdminDao := dao.Use(l.svcCtx.Gorm).SysAdmin
	sysRoleDao := dao.Use(l.svcCtx.Gorm).SysRole
	sysJobDao := dao.Use(l.svcCtx.Gorm).SysJob
	sysDeptDao := dao.Use(l.svcCtx.Gorm).SysDept

	query := sysAdminDao.WithContext(l.ctx)
	if req.QuickSearch != "" {
		query = query.Where(sysAdminDao.Nickname.Like(req.QuickSearch))
	} else {
		for _, search := range req.Search {
			if search.Field == "id" {
				query = query.Where(sysAdminDao.ID.Eq(conv.Int64(search.Val)))
			}
			if search.Field == "username" {
				query = query.Where(sysAdminDao.Username.Eq(search.Val))
			}
			if search.Field == "nickname" {
				query = query.Where(sysAdminDao.Nickname.Eq(search.Val))
			}
			if search.Field == "mobile" {
				query = query.Where(sysAdminDao.Mobile.Eq(search.Val))
			}
			if search.Field == "email" {
				query = query.Where(sysAdminDao.Email.Eq(search.Val))
			}
			if search.Field == "status" {
				query = query.Where(sysAdminDao.Status.Eq(conv.Int32(search.Val)))
			}
			if search.Field == "createdAt" {
				ss := strings.Split(search.Val, ",")
				query = query.Where(sysAdminDao.CreatedAt.Gte(timeutil.Carbon().Parse(ss[0]).Carbon2Time()), sysAdminDao.CreatedAt.Lte(timeutil.Carbon().Parse(ss[1]).Carbon2Time()))
			}
		}
	}

	queryCount := query
	total, err := queryCount.Count()
	if err != nil {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	paginator := page.Paginator(req.Page, req.PageSize, int(total))
	sysAdmins, err := query.Offset(paginator.Offset).Limit(paginator.Limit).Find()
	if err != nil {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	for _, v := range sysAdmins {
		tmpRoleIds := make([]int64, 0)
		err := jsonutil.Decode(v.RoleIds, &tmpRoleIds)
		if err != nil {
			return nil, err
		}
		list = append(list, types.SysManageInfo{
			Id:        v.ID,
			Username:  v.Username,
			Nickname:  v.Nickname,
			Avatar:    v.Avatar,
			Gender:    v.Gender,
			Email:     v.Email,
			Mobile:    v.Mobile,
			JobID:     v.JobID,
			DeptID:    v.DeptID,
			RoleIds:   tmpRoleIds,
			JobName:   "",
			DeptName:  "",
			RoleNames: make([]string, 0),
			Motto:     v.Motto,
			Status:    v.Status,
			CreatedAt: timeutil.ToDateTimeStringByTime(v.CreatedAt),
			UpdatedAt: timeutil.ToDateTimeStringByTime(v.UpdatedAt),
		})
		roleIds = append(roleIds, tmpRoleIds...)
		jobIds = append(jobIds, v.JobID)
		deptIds = append(deptIds, v.DeptID)
	}
	roleIds = sliutil.Unique(roleIds)
	jobIds = sliutil.Unique(jobIds)
	deptIds = sliutil.Unique(deptIds)

	sysRoles, err := sysRoleDao.WithContext(l.ctx).Where(sysRoleDao.ID.In(roleIds...)).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	for _, v := range sysRoles {
		roleIdToName[v.ID] = v.Name
	}
	sysJobs, err := sysJobDao.WithContext(l.ctx).Where(sysJobDao.ID.In(jobIds...)).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	for _, v := range sysJobs {
		jobIdToName[v.ID] = v.Name
	}
	sysDepts, err := sysDeptDao.WithContext(l.ctx).Where(sysDeptDao.ID.In(deptIds...)).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	for _, v := range sysDepts {
		deptIdToName[v.ID] = v.Name
	}
	for k := range list {
		list[k].JobName = jobIdToName[list[k].JobID]
		list[k].DeptName = deptIdToName[list[k].DeptID]
		if len(list[k].RoleIds) > 0 {
			for _, id := range list[k].RoleIds {
				list[k].RoleNames = append(list[k].RoleNames, roleIdToName[id])
			}
		}
	}
	resp.List = list
	err = copier.Copy(&resp.Paginator, paginator)
	if err != nil {
		return nil, err
	}
	return
}
