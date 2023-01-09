package config

import (
	"fgzs-single/internal/core"
	"fgzs-single/pkg/db"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Gorm  db.MysqlConfig
	Cache cache.CacheConf
	Redis core.RedisConfig
}
