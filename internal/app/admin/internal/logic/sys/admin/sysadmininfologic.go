package admin

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/define/cachekey"
	"fgzs-single/internal/errorx"
	"fgzs-single/internal/meta"
	"fgzs-single/pkg/util/jsonutil"
	"gorm.io/gorm"
	"strconv"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysAdminInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysAdminInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysAdminInfoLogic {
	return &SysAdminInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysAdminInfoLogic) SysAdminInfo(req *types.SysAdminInfoReq) (*types.SysAdminInfoResp, error) {
	resp := new(types.SysAdminInfoResp)
	adminId := meta.GetAdminId(l.ctx)
	cacheKey := cachekey.SysAdminInfo.BuildCacheKey(strconv.FormatInt(adminId, 10))
	res, err := cacheKey.RocksCache(l.svcCtx.RocksCache, func() (string, error) {
		sysAdminDao := dao.Use(l.svcCtx.Gorm).SysAdmin
		sysAdmin, err := sysAdminDao.WithContext(l.ctx).Where(sysAdminDao.ID.Eq(adminId)).First()
		if err != nil && err != gorm.ErrRecordNotFound {
			return "", errorx.DataSqlErr.WithDetail(err)
		}
		if sysAdmin.Status != 1 {
			return "", errorx.AccountIsBanned
		}
		roleIds := make([]int64, 0)
		err = jsonutil.Decode(sysAdmin.RoleIds, &roleIds)
		if err != nil {
			return "", errorx.DataFormattingError.WithDetail(err)
		}
		res, err := jsonutil.EncodeToString(types.SysAdminInfo{
			ID:       sysAdmin.ID,
			Username: sysAdmin.Username,
			Nickname: sysAdmin.Nickname,
			Avatar:   sysAdmin.Avatar,
			Gender:   sysAdmin.Gender,
			Email:    sysAdmin.Email,
			Mobile:   sysAdmin.Mobile,
			JobID:    sysAdmin.JobID,
			DeptID:   sysAdmin.DeptID,
			RoleIds:  roleIds,
			Motto:    sysAdmin.Motto,
		})
		if err != nil {
			return "", err
		}
		return res, nil
	})
	if err != nil {
		return nil, err
	}
	err = jsonutil.DecodeString(res, resp.Info)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
