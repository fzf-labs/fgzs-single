package dept

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"
	"fgzs-single/pkg/util/timeutil"
	"gorm.io/gorm"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDeptListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDeptListLogic {
	return &SysDeptListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDeptListLogic) SysDeptList(req *types.SysDeptListReq) (*types.SysDeptListResp, error) {
	resp := new(types.SysDeptListResp)
	sysDeptDao := dao.Use(l.svcCtx.Gorm).SysDept
	sysDepts, err := sysDeptDao.WithContext(l.ctx).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	depts := make([]types.SysDeptInfo, 0)
	for _, role := range sysDepts {
		depts = append(depts, types.SysDeptInfo{
			ID:          role.ID,
			Pid:         role.Pid,
			Name:        role.Name,
			FullName:    role.FullName,
			Responsible: role.Responsible,
			Phone:       role.Phone,
			Email:       role.Email,
			Type:        role.Type,
			Status:      role.Status,
			Sort:        role.Sort,
			CreatedAt:   role.CreatedAt.Format(timeutil.TimeLayout),
			UpdatedAt:   role.UpdatedAt.Format(timeutil.TimeLayout),
			Children:    nil,
		})
	}
	resp.List = sysDeptGenerateTree(depts)
	return resp, nil
}
func sysDeptGenerateTree(list []types.SysDeptInfo) []types.SysDeptInfo {
	var trees []types.SysDeptInfo
	// Define the top-level root and child nodes
	var roots, childs []types.SysDeptInfo
	for _, v := range list {
		if v.Pid == 0 {
			// Determine the top-level root node
			roots = append(roots, v)
		}
		childs = append(childs, v)
	}

	for _, v := range roots {
		childTree := &types.SysDeptInfo{
			ID:          v.ID,
			Pid:         v.Pid,
			Name:        v.Name,
			FullName:    v.FullName,
			Responsible: v.Responsible,
			Phone:       v.Phone,
			Email:       v.Email,
			Type:        v.Type,
			Status:      v.Status,
			Sort:        v.Sort,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			Children:    make([]types.SysDeptInfo, 0),
		}
		// recursive
		sysDeptRecursiveTree(childTree, childs)

		trees = append(trees, *childTree)
	}
	return trees
}
func sysDeptRecursiveTree(tree *types.SysDeptInfo, allNodes []types.SysDeptInfo) {
	for _, v := range allNodes {
		if v.Pid == 0 {
			// If the current node is the top-level root node, skip
			continue
		}
		if tree.ID == v.Pid {
			childTree := &types.SysDeptInfo{
				ID:          v.ID,
				Pid:         v.Pid,
				Name:        v.Name,
				FullName:    v.FullName,
				Responsible: v.Responsible,
				Phone:       v.Phone,
				Email:       v.Email,
				Type:        v.Type,
				Status:      v.Status,
				Sort:        v.Sort,
				CreatedAt:   v.CreatedAt,
				UpdatedAt:   v.UpdatedAt,
				Children:    make([]types.SysDeptInfo, 0),
			}
			sysDeptRecursiveTree(childTree, allNodes)
			tree.Children = append(tree.Children, *childTree)
		}
	}
}
