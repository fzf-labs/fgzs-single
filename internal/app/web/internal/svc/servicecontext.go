package svc

import (
	"fgzs-single/internal/app/web/internal/config"
	"fgzs-single/internal/app/web/internal/middleware"
	"fgzs-single/internal/core"
	"fgzs-single/pkg/db"
	goRedis "github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
	"time"
)

type ServiceContext struct {
	Config   config.Config
	Gorm     *gorm.DB
	Redis    *redis.Redis
	RedisDB1 *goRedis.Client
	//设备校验中间件
	DeviceCheckMiddleware rest.Middleware
	//短连接进程内缓存
	TinyUrlCollectionCache *collection.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	tinyUrlCache, err := tinyUrlCollectionCache()
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:                 c,
		Gorm:                   db.NewGorm(&c.Gorm),
		Redis:                  core.NewRedis(c.Redis),
		RedisDB1:               core.NewGoRedis(c.Redis, core.DB1),
		DeviceCheckMiddleware:  middleware.NewDeviceCheckMiddleware().Handle,
		TinyUrlCollectionCache: tinyUrlCache,
	}
}

// tinyUrlCollectionCache 短连接进程内缓存
func tinyUrlCollectionCache() (*collection.Cache, error) {
	cache, err := collection.NewCache(time.Minute*5, collection.WithLimit(1000))
	if err != nil {
		return nil, err
	}
	return cache, nil
}
