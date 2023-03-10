package config

import (
	"github.com/fzf-labs/fpkg/db"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Gorm       db.GormMysqlConfig
	Redis      redis.RedisConf
	OpenAI     OpenAIConfig
	IpLocation IpLocationConfig
}

type OpenAIConfig struct {
	ChatGPT string
}

type IpLocationConfig struct {
	Path string
}
