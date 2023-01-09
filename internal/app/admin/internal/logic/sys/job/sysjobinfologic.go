package job

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

type SysJobInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysJobInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysJobInfoLogic {
	return &SysJobInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysJobInfoLogic) SysJobInfo(req *types.SysJobInfoReq) (resp *types.SysJobInfoResp, err error) {
	resp = new(types.SysJobInfoResp)
	sysJobDao := dao.Use(l.svcCtx.Gorm).SysJob
	sysJob, err := sysJobDao.WithContext(l.ctx).Where(sysJobDao.ID.Eq(req.Id)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	resp.Info = types.SysJobInfo{
		ID:        sysJob.ID,
		Name:      sysJob.Name,
		Code:      sysJob.Code,
		Remark:    sysJob.Remark,
		Status:    sysJob.Status,
		Sort:      sysJob.Sort,
		CreatedAt: timeutil.ToDateTimeStringByTime(sysJob.CreatedAt),
		UpdatedAt: timeutil.ToDateTimeStringByTime(sysJob.UpdatedAt),
	}
	return
}
