package permmenu

import (
	"context"
	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"

	"github.com/fzf-labs/fpkg/util/timeutil"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysPermMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysPermMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysPermMenuListLogic {
	return &SysPermMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysPermMenuListLogic) SysPermMenuList(req *types.SysPermMenuListReq) (*types.SysPermMenuListResp, error) {
	resp := new(types.SysPermMenuListResp)
	sysPermMenuDao := dao.Use(l.svcCtx.Gorm).SysPermMenu
	sysPermMenus, err := sysPermMenuDao.WithContext(l.ctx).Find()
	if err != nil {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	roles := make([]types.SysPermMenu, 0)
	for _, v := range sysPermMenus {
		roles = append(roles, types.SysPermMenu{
			Id:        v.ID,
			Pid:       v.Pid,
			Type:      v.Type,
			Title:     v.Title,
			Name:      v.Name,
			Path:      v.Path,
			Icon:      v.Icon,
			MenuType:  v.MenuType,
			URL:       v.URL,
			Component: v.Component,
			Extend:    v.Extend,
			Remark:    v.Remark,
			Sort:      v.Sort,
			Status:    v.Status,
			CreatedAt: timeutil.ToDateTimeStringByTime(v.CreatedAt),
			UpdatedAt: timeutil.ToDateTimeStringByTime(v.UpdatedAt),
			Children:  nil,
		})
	}
	resp.List = sysPermMenuGenerateTree(roles)
	return resp, nil
}
func sysPermMenuGenerateTree(list []types.SysPermMenu) []types.SysPermMenu {
	var trees []types.SysPermMenu
	// Define the top-level root and child nodes
	var roots, childs []types.SysPermMenu
	for _, v := range list {
		if v.Pid == 0 {
			// Determine the top-level root node
			roots = append(roots, v)
		}
		childs = append(childs, v)
	}

	for _, v := range roots {
		childTree := &types.SysPermMenu{
			Id:        v.Id,
			Pid:       v.Pid,
			Type:      v.Type,
			Title:     v.Title,
			Name:      v.Name,
			Path:      v.Path,
			Icon:      v.Icon,
			MenuType:  v.MenuType,
			URL:       v.URL,
			Component: v.Component,
			Keepalive: v.Keepalive,
			Extend:    v.Extend,
			Remark:    v.Remark,
			Sort:      v.Sort,
			Status:    v.Status,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Children:  make([]types.SysPermMenu, 0),
		}
		// recursive
		sysPermMenuRecursiveTree(childTree, childs)

		trees = append(trees, *childTree)
	}
	return trees
}
func sysPermMenuRecursiveTree(tree *types.SysPermMenu, allNodes []types.SysPermMenu) {
	for _, v := range allNodes {
		if v.Pid == 0 {
			// If the current node is the top-level root node, skip
			continue
		}
		if tree.Id == v.Pid {
			childTree := &types.SysPermMenu{
				Id:        v.Id,
				Pid:       v.Pid,
				Type:      v.Type,
				Title:     v.Title,
				Name:      v.Name,
				Path:      v.Path,
				Icon:      v.Icon,
				MenuType:  v.MenuType,
				URL:       v.URL,
				Component: v.Component,
				Keepalive: v.Keepalive,
				Extend:    v.Extend,
				Remark:    v.Remark,
				Sort:      v.Sort,
				Status:    v.Status,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
				Children:  make([]types.SysPermMenu, 0),
			}
			sysPermMenuRecursiveTree(childTree, allNodes)
			tree.Children = append(tree.Children, *childTree)
		}
	}
}
