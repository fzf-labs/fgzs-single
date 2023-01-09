package middleware

import (
	"fgzs-single/internal/errorx"
	"fgzs-single/internal/meta"
	"fgzs-single/internal/response"
	"fgzs-single/pkg/conv"
	"fgzs-single/pkg/jwt"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"time"
)

type JwtMiddleware struct {
	Redis     *redis.Redis
	JwtConfig *jwt.Config
}

func NewJwtMiddleware(config *jwt.Config, redis *redis.Redis) *JwtMiddleware {
	return &JwtMiddleware{
		Redis:     redis,
		JwtConfig: config,
	}
}

func (m *JwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//获取header头中的 Authorization
		authorization := r.Header.Get("Authorization")
		//不存在则报错
		if len(authorization) == 0 {
			response.Err(r, w, errorx.TokenNotRequest)
			return
		}
		//token截取
		var token string
		_, err := fmt.Sscanf(authorization, "Bearer %s", &token)
		if err != nil {
			response.Err(r, w, errorx.TokenFormatErr)
			return
		}
		//鉴权处理
		//token 解析
		j := jwt.NewJwt(m.Redis, m.JwtConfig)
		//后续判断是否过期等操作,这里只检查格式正确即可
		claims, err := j.ParseToken(token)
		if err != nil {
			response.Err(r, w, errorx.TokenInvalidErr)
			return
		}
		now := time.Now().Unix()
		//当前时间超过过期时间 则直接失效
		if now > conv.Int64(claims[jwt.JwtExpired]) {
			response.Err(r, w, errorx.TokenExpired)
			return
		}
		//校验是否单一登录
		err = j.JwtTokenCheck(claims)
		if err != nil {
			jwtBlackTokenCheck, _ := j.JwtBlackTokenCheck(claims)
			if jwtBlackTokenCheck {
				//将sys_uid参数写进context中
				r = r.WithContext(meta.SetAdminId(r.Context(), conv.String(claims[jwt.JwtUID])))
				next(w, r)
			}
			response.Err(r, w, errorx.TokenVerificationFailed)
			return
		}
		//将sys_uid参数写进context中
		r = r.WithContext(meta.SetAdminId(r.Context(), conv.String(claims[jwt.JwtUID])))
		next(w, r)
	}
}
