package admin

import (
	"context"
	"github.com/mojocn/base64Captcha"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysAdminLoginCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysAdminLoginCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysAdminLoginCaptchaLogic {
	return &SysAdminLoginCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// SysAdminLoginCaptcha 登录验证码
func (l *SysAdminLoginCaptchaLogic) SysAdminLoginCaptcha(req *types.SysAdminLoginCaptchaReq) (resp *types.SysAdminLoginCaptchaResp, err error) {
	resp = new(types.SysAdminLoginCaptchaResp)
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	captchaId, picPath, err := cp.Generate()
	if err != nil {
		return nil, err
	}
	resp.CaptchaId = captchaId
	resp.CaptchaImg = picPath
	return
}
