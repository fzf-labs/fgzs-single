package admin

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/dal/model"
	"fgzs-single/internal/define/cachekey"
	"fgzs-single/internal/define/constant"
	"fgzs-single/internal/errorx"
	"fgzs-single/internal/meta"
	"fgzs-single/pkg/util/jsonutil"
	"fgzs-single/pkg/util/sliutil"
	"gorm.io/gorm"
	"strconv"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysAdminPermMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysAdminPermMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysAdminPermMenuLogic {
	return &SysAdminPermMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// SysAdminPermMenu 获取权限菜单
func (l *SysAdminPermMenuLogic) SysAdminPermMenu(req *types.SysAdminPermMenuReq) (*types.SysAdminPermMenuResp, error) {
	resp := new(types.SysAdminPermMenuResp)
	adminId := meta.GetAdminId(l.ctx)
	cacheKey := cachekey.SysAdminPermmenu.BuildCacheKey(strconv.FormatInt(adminId, 10))
	res, err := cacheKey.RocksCache(l.svcCtx.RocksCache, func() (string, error) {
		sysAdminDao := dao.Use(l.svcCtx.Gorm).SysAdmin
		sysRoleDao := dao.Use(l.svcCtx.Gorm).SysRole
		sysPermMenuDao := dao.Use(l.svcCtx.Gorm).SysPermMenu
		//获取用户角色->获取用户权限->树状结构
		sysAdmin, err := sysAdminDao.WithContext(l.ctx).Where(sysAdminDao.ID.Eq(adminId)).First()
		if err != nil && err != gorm.ErrRecordNotFound {
			return "", errorx.DataSqlErr.WithDetail(err)
		}
		if sysAdmin.RoleIds == nil {
			return "", errorx.UserNotBoundRole
		}
		roleIds := make([]int64, 0)
		err = jsonutil.Decode(sysAdmin.RoleIds, &roleIds)
		if err != nil {
			return "", err
		}
		sysRoles, err := sysRoleDao.WithContext(l.ctx).Where(sysRoleDao.ID.In(roleIds...)).Find()
		if err != nil && err != gorm.ErrRecordNotFound {
			return "", errorx.DataSqlErr.WithDetail(err)
		}
		if len(sysRoles) == 0 {
			return "", errorx.UserNotBoundRole
		}
		var super bool
		permMenuIds := make([]int64, 0)
		for _, role := range sysRoles {
			if role.PermMenuIds == "" {
				continue
			}
			if role.PermMenuIds == "*" {
				super = true
				break
			}
			int64s, err := sliutil.SplitStringToInt64(role.PermMenuIds)
			if err != nil {
				return "", err
			}
			permMenuIds = append(permMenuIds, int64s...)
		}
		var sysPermMenus []*model.SysPermMenu
		if super {
			sysPermMenus, err = sysPermMenuDao.WithContext(l.ctx).Where(sysPermMenuDao.Status.Eq(constant.StatusEnable)).Find()
			if err != nil {
				return "", err
			}
		} else {
			permMenuIds = sliutil.Unique(permMenuIds)
			sysPermMenus, err = sysPermMenuDao.WithContext(l.ctx).Where(sysPermMenuDao.ID.In(permMenuIds...), sysPermMenuDao.Status.Eq(constant.StatusEnable)).Find()
			if err != nil {
				return "", err
			}
		}
		if len(sysPermMenus) == 0 {
			return "", errorx.UserNotBoundRole
		}
		menus := make([]types.SysAdminMenu, 0)
		for _, v := range sysPermMenus {
			menus = append(menus, types.SysAdminMenu{
				ID:        v.ID,
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
			})
		}
		menus = sysAdminMenuGenerateTree(menus)
		toString, err := jsonutil.EncodeToString(menus)
		if err != nil {
			return "", err
		}
		return toString, nil
	})
	if err != nil {
		return nil, err
	}
	err = jsonutil.DecodeString(res, resp.Menus)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func sysAdminMenuGenerateTree(list []types.SysAdminMenu) []types.SysAdminMenu {
	var trees []types.SysAdminMenu
	// Define the top-level root and child nodes
	var roots, childs []types.SysAdminMenu
	for _, v := range list {
		if v.Pid == 0 {
			// Determine the top-level root node
			roots = append(roots, v)
		}
		childs = append(childs, v)
	}

	for _, v := range roots {
		childTree := &types.SysAdminMenu{
			ID:        v.ID,
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
			Children:  make([]types.SysAdminMenu, 0),
		}
		// recursive
		sysAdminMenuRecursiveTree(childTree, childs)

		trees = append(trees, *childTree)
	}
	return trees
}
func sysAdminMenuRecursiveTree(tree *types.SysAdminMenu, allNodes []types.SysAdminMenu) {
	for _, v := range allNodes {
		if v.Pid == 0 {
			// If the current node is the top-level root node, skip
			continue
		}
		if tree.ID == v.Pid {
			childTree := &types.SysAdminMenu{
				ID:        v.ID,
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
				Children:  make([]types.SysAdminMenu, 0),
			}
			sysAdminMenuRecursiveTree(childTree, allNodes)
			tree.Children = append(tree.Children, *childTree)
		}
	}
}
