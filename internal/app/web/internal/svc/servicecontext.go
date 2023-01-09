package svc

import (
	"fgzs-single/internal/app/web/internal/config"
	"fgzs-single/internal/app/web/internal/middleware"
	"fgzs-single/internal/core"
	"fgzs-single/pkg/db"
	goRedis "github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	Gorm     *gorm.DB
	Redis    *redis.Redis
	RedisDB1 *goRedis.Client
	//设备校验中间件
	DeviceCheckMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		Gorm:                  db.NewGorm(&c.Gorm),
		Redis:                 core.NewRedis(c.Redis),
		RedisDB1:              core.NewGoRedis(c.Redis, core.DB1),
		DeviceCheckMiddleware: middleware.NewDeviceCheckMiddleware().Handle,
	}
}
