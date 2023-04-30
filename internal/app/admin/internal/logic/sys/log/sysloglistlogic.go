package log

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/dal/repo"
	"fgzs-single/internal/errorx"

	"github.com/fzf-labs/fpkg/page"
	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogListLogic {
	return &SysLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysLogListLogic) SysLogList(req *types.SysLogListReq) (resp *types.SysLogListResp, err error) {
	resp = &types.SysLogListResp{
		List:      make([]types.SysLog, 0),
		Paginator: types.Paginator{},
	}
	sysLogDao := dao.Use(l.svcCtx.Gorm).SysLog
	count, err := sysLogDao.WithContext(l.ctx).Count()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	paginator := page.Paginator(req.Page, req.PageSize, int(count))
	sysLogs, err := sysLogDao.WithContext(l.ctx).Offset(paginator.Offset).Limit(paginator.Limit).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	adminIds := make([]int64, 0)
	paths := make([]string, 0)
	for _, v := range sysLogs {
		adminIds = append(adminIds, v.AdminID)
		paths = append(paths, v.URI)
	}
	idToName, _, err := repo.NewSysAdminRepo(l.ctx, l.svcCtx.Gorm).IdToName(adminIds)
	if err != nil {
		return nil, err
	}
	uriToName, err := repo.NewSysApiRepo(l.ctx, l.svcCtx.Gorm).IdToName(paths)
	if err != nil {
		return nil, err
	}
	for _, v := range sysLogs {
		resp.List = append(resp.List, types.SysLog{
			ID:        v.ID,
			AdminID:   v.AdminID,
			Username:  idToName[v.AdminID],
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
