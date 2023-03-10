package config

import (
	"github.com/fzf-labs/fpkg/db"
	"github.com/fzf-labs/fpkg/jwt"
	"github.com/fzf-labs/fpkg/oss"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Gorm  db.GormMysqlConfig
	Redis redis.RedisConf
	Jwt   jwt.Config

	AliOss oss.AliConfig

	Upload Upload
}

type Upload struct {
	Path string
	Host string
}
