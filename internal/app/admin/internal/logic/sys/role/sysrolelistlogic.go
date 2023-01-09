package role

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"
	"fgzs-single/pkg/util/timeutil"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysRoleListLogic {
	return &SysRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysRoleListLogic) SysRoleList(req *types.SysRoleListReq) (*types.SysRoleListResp, error) {
	resp := new(types.SysRoleListResp)
	sysRoleDao := dao.Use(l.svcCtx.Gorm).SysRole
	sysRoles, err := sysRoleDao.WithContext(l.ctx).Find()
	if err != nil {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	roles := make([]types.SysRole, 0)
	for _, role := range sysRoles {
		roles = append(roles, types.SysRole{
			Id:        role.ID,
			Pid:       role.Pid,
			Name:      role.Name,
			Remark:    role.Remark,
			Status:    role.Status,
			Sort:      role.Sort,
			CreatedAt: role.CreatedAt.Format(timeutil.TimeLayout),
			UpdatedAt: role.UpdatedAt.Format(timeutil.TimeLayout),
			Children:  nil,
		})
	}
	resp.List = sysRoleGenerateTree(roles)
	return resp, nil
}
func sysRoleGenerateTree(list []types.SysRole) []types.SysRole {
	var trees []types.SysRole
	// Define the top-level root and child nodes
	var roots, childs []types.SysRole
	for _, v := range list {
		if v.Pid == 0 {
			// Determine the top-level root node
			roots = append(roots, v)
		}
		childs = append(childs, v)
	}

	for _, v := range roots {
		childTree := &types.SysRole{
			Id:        v.Id,
			Pid:       v.Pid,
			Name:      v.Name,
			Remark:    v.Remark,
			Status:    v.Status,
			Sort:      v.Sort,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Children:  make([]types.SysRole, 0),
		}
		// recursive
		sysRoleRecursiveTree(childTree, childs)

		trees = append(trees, *childTree)
	}
	return trees
}
func sysRoleRecursiveTree(tree *types.SysRole, allNodes []types.SysRole) {
	for _, v := range allNodes {
		if v.Pid == 0 {
			// If the current node is the top-level root node, skip
			continue
		}
		if tree.Id == v.Pid {
			childTree := &types.SysRole{
				Id:        v.Id,
				Pid:       v.Pid,
				Name:      v.Name,
				Remark:    v.Remark,
				Status:    v.Status,
				Sort:      v.Sort,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
				Children:  make([]types.SysRole, 0),
			}
			sysRoleRecursiveTree(childTree, allNodes)
			tree.Children = append(tree.Children, *childTree)
		}
	}
}
