package svc

import (
	"fgzs-single/internal/app/admin/internal/config"
	"fgzs-single/internal/app/admin/internal/middleware"
	"fgzs-single/internal/core"
	"fgzs-single/pkg/db"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config           config.Config
	Gorm             *gorm.DB
	Redis            *redis.Redis
	JwtMiddleware    rest.Middleware
	AuthMiddleware   rest.Middleware
	SysLogMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	newGorm := db.NewGorm(&c.Gorm)
	newRedis := core.NewRedis(c.Redis)
	return &ServiceContext{
		Config:           c,
		Gorm:             newGorm,
		Redis:            newRedis,
		JwtMiddleware:    middleware.NewJwtMiddleware(&c.Jwt, newRedis).Handle,
		AuthMiddleware:   middleware.NewAuthMiddleware().Handle,
		SysLogMiddleware: middleware.NewSysLogMiddleware(newGorm).Handle,
	}
}
