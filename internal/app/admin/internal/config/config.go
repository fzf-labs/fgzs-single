package config

import (
	"fgzs-single/internal/core"
	"fgzs-single/pkg/db"
	"fgzs-single/pkg/jwt"
	"fgzs-single/pkg/oss"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Gorm  db.MysqlConfig
	Cache cache.CacheConf
	Redis core.RedisConfig
	Jwt   jwt.Config

	AliOss oss.AliConfig

	Upload Upload
}

type Upload struct {
	Path string
	Host string
}
