package log

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/dal/repo"
	"fgzs-single/internal/errorx"

	"github.com/fzf-labs/fpkg/util/timeutil"
	"gorm.io/gorm"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysLogInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogInfoLogic {
	return &SysLogInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysLogInfoLogic) SysLogInfo(req *types.SysLogInfoReq) (resp *types.SysLogInfoResp, err error) {
	resp = new(types.SysLogInfoResp)
	sysLogDao := dao.Use(l.svcCtx.Gorm).SysLog
	log, err := sysLogDao.WithContext(l.ctx).Where(sysLogDao.ID.Eq(req.Id)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	uriToName, err := repo.NewSysApiRepo(l.ctx, l.svcCtx.Gorm).IdToName([]string{log.URI})
	if err != nil {
		return nil, err
	}
	resp.Info = types.SysLog{
		ID:        log.ID,
		AdminID:   log.AdminID,
		Username:  "",
		IP:        log.IP,
		URI:       log.URI,
		UriDesc:   uriToName[log.URI],
		Useragent: log.Useragent,
		Req:       log.Req.String(),
		Resp:      log.Resp.String(),
		CreatedAt: timeutil.ToDateTimeStringByTime(log.CreatedAt),
	}
	return
}
