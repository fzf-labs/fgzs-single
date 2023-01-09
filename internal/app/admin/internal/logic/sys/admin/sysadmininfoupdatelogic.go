package admin

import (
	"context"
	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/dal/model"
	"fgzs-single/internal/errorx"
	"fgzs-single/internal/meta"
	"fgzs-single/pkg/crypt"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysAdminInfoUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysAdminInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysAdminInfoUpdateLogic {
	return &SysAdminInfoUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysAdminInfoUpdateLogic) SysAdminInfoUpdate(req *types.SysAdminInfoUpdateReq) (resp *types.SysAdminInfoUpdateResp, err error) {
	resp = new(types.SysAdminInfoUpdateResp)
	adminId := meta.GetAdminId(l.ctx)
	sysAdminDao := dao.Use(l.svcCtx.Gorm).SysAdmin
	sysAdmin, err := sysAdminDao.WithContext(l.ctx).Where(sysAdminDao.ID.Eq(adminId)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	var pwd string
	if req.Password != "" {
		pwd, err = crypt.Encrypt(req.Password + sysAdmin.Salt)
		if err != nil {
			return nil, err
		}
	}
	_, err = sysAdminDao.WithContext(l.ctx).Where(sysAdminDao.ID.Eq(adminId)).Updates(model.SysAdmin{
		Password: pwd,
		Nickname: req.Nickname,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Motto:    req.Motto,
	})
	if err != nil {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	return
}
