package svc

import (
	"fgzs-single/internal/app/web/internal/config"
	"fgzs-single/internal/app/web/internal/middleware"
	"fgzs-single/internal/core"
	"fgzs-single/pkg/db"
	"github.com/dtm-labs/rockscache"
	goRedis "github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
	"time"
)

type ServiceContext struct {
	Config config.Config
	Gorm   *gorm.DB
	Redis  *redis.Redis
	//go-redis
	GoRedis *goRedis.Client
	//进程内缓存
	CollectionCache *collection.Cache
	//ChatGpt进程内缓存
	CollectionCacheChatGpt *collection.Cache
	//数据一致性缓存
	RocksCache *rockscache.Client

	//设备校验中间件
	DeviceCheckMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	newGoRedis := core.NewGoRedis(c.Redis, 0)
	return &ServiceContext{
		Config:                 c,
		Gorm:                   db.NewGorm(&c.Gorm),
		Redis:                  core.NewRedis(c.Redis),
		GoRedis:                newGoRedis,
		CollectionCache:        core.NewDefaultCollectionCache(),
		CollectionCacheChatGpt: core.NewCollectionCache("chatGPT", time.Minute*30, 1000),
		RocksCache:             core.NewRocksCache(newGoRedis),
		DeviceCheckMiddleware:  middleware.NewDeviceCheckMiddleware().Handle,
	}
}
