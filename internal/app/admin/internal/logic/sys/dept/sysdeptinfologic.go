package dept

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"

	"github.com/fzf-labs/fpkg/util/timeutil"
	"gorm.io/gorm"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDeptInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDeptInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDeptInfoLogic {
	return &SysDeptInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDeptInfoLogic) SysDeptInfo(req *types.SysDeptInfoReq) (resp *types.SysDeptInfoResp, err error) {
	resp = new(types.SysDeptInfoResp)
	sysDeptDao := dao.Use(l.svcCtx.Gorm).SysDept
	sysDept, err := sysDeptDao.WithContext(l.ctx).Where(sysDeptDao.ID.Eq(req.Id)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	resp.Info = types.SysDeptInfo{
		ID:          sysDept.ID,
		Pid:         sysDept.Pid,
		Name:        sysDept.Name,
		FullName:    sysDept.FullName,
		Responsible: sysDept.Responsible,
		Phone:       sysDept.Phone,
		Email:       sysDept.Email,
		Type:        sysDept.Type,
		Status:      sysDept.Status,
		Sort:        sysDept.Sort,
		CreatedAt:   timeutil.ToDateTimeStringByTime(sysDept.CreatedAt),
		UpdatedAt:   timeutil.ToDateTimeStringByTime(sysDept.UpdatedAt),
		Children:    nil,
	}
	return
}
