package core

import (
	"context"
	goRedis "github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type RedisConfig struct {
	Host string
	Pass string
	Type string
}

// NewRedis 实例化redis
func NewRedis(conf RedisConfig) *redis.Redis {
	return redis.New(conf.Host, func(r *redis.Redis) {
		r.Type = conf.Type
		r.Pass = conf.Pass
	})
}

type RedisDB int

const (
	DB0 RedisDB = 0
	DB1 RedisDB = 1
)

// NewGoRedis 初始化go-redis客户端
func NewGoRedis(cfg RedisConfig, db RedisDB) *goRedis.Client {
	Client := goRedis.NewClient(&goRedis.Options{
		Addr:     cfg.Host,
		Password: cfg.Pass,
		DB:       int(db),
	})
	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		logx.Errorf("[redis] redis ping err: %+v", err)
		return nil
	}
	return Client
}
