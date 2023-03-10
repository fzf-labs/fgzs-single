package permmenu

import (
	"context"
	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"
	"github.com/fzf-labs/fpkg/util/timeutil"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysPermMenuInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysPermMenuInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysPermMenuInfoLogic {
	return &SysPermMenuInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysPermMenuInfoLogic) SysPermMenuInfo(req *types.SysPermMenuInfoReq) (*types.SysPermMenuInfoResp, error) {
	resp := new(types.SysPermMenuInfoResp)
	sysPermMenuDao := dao.Use(l.svcCtx.Gorm).SysPermMenu
	sysPermMenu, err := sysPermMenuDao.WithContext(l.ctx).Where(sysPermMenuDao.ID.Eq(req.Id)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	resp.Info = types.SysPermMenu{
		Id:        sysPermMenu.ID,
		Pid:       sysPermMenu.Pid,
		Type:      sysPermMenu.Type,
		Title:     sysPermMenu.Title,
		Name:      sysPermMenu.Name,
		Path:      sysPermMenu.Path,
		Icon:      sysPermMenu.Icon,
		MenuType:  sysPermMenu.MenuType,
		URL:       sysPermMenu.URL,
		Component: sysPermMenu.Component,
		Extend:    sysPermMenu.Extend,
		Remark:    sysPermMenu.Remark,
		Sort:      sysPermMenu.Sort,
		Status:    sysPermMenu.Status,
		CreatedAt: timeutil.ToDateTimeStringByTime(sysPermMenu.CreatedAt),
		UpdatedAt: timeutil.ToDateTimeStringByTime(sysPermMenu.UpdatedAt),
		Children:  nil,
	}
	return resp, nil
}
