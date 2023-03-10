package admin

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/define/constant"
	"fgzs-single/internal/errorx"
	"github.com/fzf-labs/fpkg/crypt"
	"github.com/fzf-labs/fpkg/jwt"
	"github.com/mojocn/base64Captcha"
	"gorm.io/gorm"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysAdminLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysAdminLoginLogic {
	return &SysAdminLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysAdminLoginLogic) SysAdminLogin(req *types.SysAdminLoginReq) (*types.SysAdminLoginResp, error) {
	resp := new(types.SysAdminLoginResp)
	//验证码
	verify := base64Captcha.DefaultMemStore.Verify(req.CaptchaId, req.VerifyCode, true)
	if !verify {
		return nil, errorx.VerificationCodeError
	}
	//用户校验
	sysAdminDao := dao.Use(l.svcCtx.Gorm).SysAdmin
	sysAdmin, err := sysAdminDao.WithContext(l.ctx).Where(sysAdminDao.Username.Eq(req.Username)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	if sysAdmin == nil {
		return nil, errorx.AccountNotExist

	}
	if crypt.Compare(sysAdmin.Password, req.Password+sysAdmin.Salt) != nil {
		return nil, errorx.WrongPassword
	}
	if sysAdmin.Status != constant.StatusEnable {
		return nil, errorx.NoAccess
	}
	//颁发token
	kv := make(map[string]interface{})
	kv["uid"] = sysAdmin.ID
	jwtClient := jwt.NewJwt(l.svcCtx.Redis, &l.svcCtx.Config.Jwt)
	token, claims, err := jwtClient.GenerateToken(kv)
	if err != nil {
		return nil, err
	}
	err = jwtClient.JwtTokenStore(claims)
	if err != nil {
		return nil, errorx.TokenStorageFailed
	}
	resp.Token = token.Token
	resp.RefreshAt = token.RefreshAt
	resp.ExpiredAt = token.ExpiredAt
	return resp, nil
}
