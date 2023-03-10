package core

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
)

// NewRedis 实例化redis
func NewRedis(conf redis.RedisConf) (*redis.Redis, error) {
	newRedis, err := redis.NewRedis(conf, func(r *redis.Redis) {
		r.Type = conf.Type
		r.Pass = conf.Pass
	})
	if err != nil {
		return nil, err
	}
	return newRedis, nil
}
