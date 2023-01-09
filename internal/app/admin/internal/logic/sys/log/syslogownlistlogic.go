package log

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/dal/repo"
	"fgzs-single/internal/errorx"
	"fgzs-single/internal/meta"
	"fgzs-single/pkg/page"
	"fgzs-single/pkg/util/timeutil"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogOwnListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysLogOwnListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogOwnListLogic {
	return &SysLogOwnListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysLogOwnListLogic) SysLogOwnList(req *types.SysLogOwnListReq) (resp *types.SysLogOwnListResp, err error) {
	resp = &types.SysLogOwnListResp{
		List:      make([]types.SysLog, 0),
		Paginator: types.Paginator{},
	}
	adminId := meta.GetAdminId(l.ctx)
	sysLogDao := dao.Use(l.svcCtx.Gorm).SysLog
	count, err := sysLogDao.WithContext(l.ctx).Where(sysLogDao.AdminID.Eq(adminId)).Count()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	paginator := page.Paginator(req.Page, req.PageSize, int(count))
	sysLogs, err := sysLogDao.WithContext(l.ctx).Where(sysLogDao.AdminID.Eq(adminId)).Offset(paginator.Offset).Limit(paginator.Limit).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}

	paths := make([]string, 0)
	for _, v := range sysLogs {
		paths = append(paths, v.URI)
	}
	uriToName, err := repo.NewSysApiRepo(l.ctx, l.svcCtx.Gorm).IdToName(paths)
	if err != nil {
		return nil, err
	}
	for _, v := range sysLogs {
		resp.List = append(resp.List, types.SysLog{
			ID:        v.ID,
			AdminID:   v.AdminID,
			Username:  "",
			IP:        v.IP,
			URI:       v.URI,
			UriDesc:   uriToName[v.URI],
			Useragent: v.Useragent,
			HTTPCode:  0,
			Req:       v.Req.String(),
			Resp:      v.Resp.String(),
			CreatedAt: timeutil.ToDateTimeStringByTime(v.CreatedAt),
		})
	}
	err = copier.Copy(&resp.Paginator, paginator)
	if err != nil {
		return nil, err
	}
	return
}
