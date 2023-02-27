package svc

import (
	"fgzs-single/internal/app/admin/internal/config"
	"fgzs-single/internal/app/admin/internal/middleware"
	"fgzs-single/internal/core"
	"fgzs-single/pkg/db"
	"github.com/dtm-labs/rockscache"
	goRedis "github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/collection"
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
	CollectionCache *collection.Cache
	//数据一致性缓存
	RocksCache *rockscache.Client

	JwtMiddleware    rest.Middleware
	AuthMiddleware   rest.Middleware
	SysLogMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	newGorm := db.NewGorm(&c.Gorm)
	newRedis := core.NewRedis(c.Redis)
	newGoRedis := core.NewGoRedis(c.Redis, 0)
	return &ServiceContext{
		Config:           c,
		Gorm:             newGorm,
		Redis:            newRedis,
		GoRedis:          newGoRedis,
		CollectionCache:  core.NewDefaultCollectionCache(),
		RocksCache:       core.NewRocksCache(newGoRedis),
		JwtMiddleware:    middleware.NewJwtMiddleware(&c.Jwt, newRedis).Handle,
		AuthMiddleware:   middleware.NewAuthMiddleware().Handle,
		SysLogMiddleware: middleware.NewSysLogMiddleware(newGorm).Handle,
	}
}
