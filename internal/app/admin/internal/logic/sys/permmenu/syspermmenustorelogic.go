package permmenu

import (
	"context"
	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/dal/model"
	"fgzs-single/internal/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type SysPermMenuStoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysPermMenuStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysPermMenuStoreLogic {
	return &SysPermMenuStoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysPermMenuStoreLogic) SysPermMenuStore(req *types.SysPermMenuStoreReq) (*types.SysPermMenuStoreResp, error) {
	resp := new(types.SysPermMenuStoreResp)
	sysPermMenuDao := dao.Use(l.svcCtx.Gorm).SysPermMenu
	if req.Id > 0 {
		_, err := sysPermMenuDao.WithContext(l.ctx).Where(sysPermMenuDao.ID.Eq(req.Id)).Select(
			sysPermMenuDao.Pid,
			sysPermMenuDao.Type,
			sysPermMenuDao.Pid,
			sysPermMenuDao.Title,
			sysPermMenuDao.Name,
			sysPermMenuDao.Path,
			sysPermMenuDao.Icon,
			sysPermMenuDao.MenuType,
			sysPermMenuDao.URL,
			sysPermMenuDao.Component,
			sysPermMenuDao.Extend,
			sysPermMenuDao.Remark,
			sysPermMenuDao.Sort,
			sysPermMenuDao.Status,
		).Updates(&model.SysPermMenu{
			Pid:       req.Pid,
			Type:      req.Type,
			Title:     req.Title,
			Name:      req.Name,
			Path:      req.Path,
			Icon:      req.Icon,
			MenuType:  req.MenuType,
			URL:       req.URL,
			Component: req.Component,
			Extend:    req.Extend,
			Remark:    req.Remark,
			Sort:      req.Sort,
			Status:    req.Status,
		})
		if err != nil {
			return nil, errorx.DataSqlErr.WithDetail(err)
		}
	} else {
		err := sysPermMenuDao.WithContext(l.ctx).Create(&model.SysPermMenu{
			Pid:       req.Pid,
			Type:      req.Type,
			Title:     req.Title,
			Name:      req.Name,
			Path:      req.Path,
			Icon:      req.Icon,
			MenuType:  req.MenuType,
			URL:       req.URL,
			Component: req.Component,
			Extend:    req.Extend,
			Remark:    req.Remark,
			Sort:      req.Sort,
			Status:    req.Status,
		})
		if err != nil {
			return nil, errorx.DataSqlErr.WithDetail(err)
		}
	}
	return resp, nil
}
