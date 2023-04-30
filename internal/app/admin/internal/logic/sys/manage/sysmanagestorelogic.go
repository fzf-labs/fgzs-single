package manage

import (
	"context"
	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/dal/model"
	"fgzs-single/internal/errorx"

	"github.com/fzf-labs/fpkg/crypt"
	avatar2 "github.com/fzf-labs/fpkg/third_api/avatar"
	"github.com/fzf-labs/fpkg/util/jsonutil"
	"github.com/fzf-labs/fpkg/util/strutil"

	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysManageStoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysManageStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysManageStoreLogic {
	return &SysManageStoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysManageStoreLogic) SysManageStore(req *types.SysManageStoreReq) (resp *types.SysManageStoreResp, err error) {
	sysAdminDao := dao.Use(l.svcCtx.Gorm).SysAdmin
	avatar := req.Avatar
	if avatar == "" {
		avatar = avatar2.Url()
	}
	roleIds, err := jsonutil.Encode(req.RoleIds)
	if err != nil {
		return nil, err
	}
	if req.Id > 0 {
		sysAdmin, err := sysAdminDao.WithContext(l.ctx).Where(sysAdminDao.ID.Eq(req.Id)).First()
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, errorx.DataSqlErr.WithDetail(err)
		}
		if sysAdmin == nil {
			return nil, errorx.AccountNotExist

		}
		update := map[string]interface{}{
			"username": req.Username,
			"nickname": req.Nickname,
			"avatar":   avatar,
			"gender":   req.Gender,
			"email":    req.Email,
			"mobile":   req.Mobile,
			"job_id":   req.JobID,
			"dept_id":  req.DeptID,
			"role_ids": roleIds,
			"status":   req.Status,
			"motto":    req.Motto,
		}
		if req.Password != "" {
			pwd, err := crypt.Encrypt(req.Password + sysAdmin.Salt)
			if err != nil {
				return nil, err
			}
			update["Password"] = pwd
		}
		_, err = sysAdminDao.WithContext(l.ctx).Where(sysAdminDao.ID.Eq(req.Id)).Updates(update)
		if err != nil {
			return nil, errorx.DataSqlErr.WithDetail(err)
		}
	} else {
		salt := strutil.Random(16)
		pwd, err := crypt.Encrypt(req.Password + salt)
		if err != nil {
			return nil, err
		}
		sysAdmin := model.SysAdmin{
			Username: req.Username,
			Password: pwd,
			Nickname: req.Nickname,
			Avatar:   avatar,
			Gender:   req.Gender,
			Email:    req.Email,
			Mobile:   req.Mobile,
			JobID:    req.JobID,
			DeptID:   req.DeptID,
			RoleIds:  roleIds,
			Salt:     salt,
			Status:   req.Status,
			Motto:    req.Motto,
		}
		err = sysAdminDao.WithContext(l.ctx).Create(&sysAdmin)
		if err != nil {
			return nil, errorx.DataSqlErr.WithDetail(err)
		}
	}
	return
}
