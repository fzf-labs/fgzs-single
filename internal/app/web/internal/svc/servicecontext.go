package svc

import (
	"fgzs-single/internal/app/web/internal/config"
	"fgzs-single/internal/app/web/internal/middleware"
	"fgzs-single/internal/core"
	"time"

	"github.com/dtm-labs/rockscache"
	"github.com/fzf-labs/fpkg/cache"
	"github.com/fzf-labs/fpkg/cache/collectioncache"
	"github.com/fzf-labs/fpkg/db"
	goRedis "github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Gorm   *gorm.DB
	Redis  *redis.Redis
	//go-redis
	GoRedis *goRedis.Client
	//进程内缓存
	CollectionCache *collectioncache.Cache
	//ChatGpt进程内缓存
	CollectionCacheChatGpt *collectioncache.Cache
	//数据一致性缓存
	RocksCache *rockscache.Client

	//设备校验中间件
	DeviceCheckMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	newGormMysql, err := db.NewGormMysql(&c.Gorm)
	if err != nil {
		panic("NewGormMysql err")
	}
	newRedis, err := core.NewRedis(c.Redis)
	if err != nil {
		panic("NewRedis err")
	}
	newGoRedis, err := cache.NewGoRedis(c.Redis.Host, c.Redis.Pass, 0)
	if err != nil {
		panic("NewGoRedis err")
	}
	return &ServiceContext{
		Config:                 c,
		Gorm:                   newGormMysql,
		Redis:                  newRedis,
		GoRedis:                newGoRedis,
		CollectionCache:        cache.NewDefaultCollectionCache(),
		CollectionCacheChatGpt: cache.NewCollectionCache("chatGPT", time.Minute*30, 1000),
		RocksCache:             cache.NewRocksCache(newGoRedis),
		DeviceCheckMiddleware:  middleware.NewDeviceCheckMiddleware().Handle,
	}
}
