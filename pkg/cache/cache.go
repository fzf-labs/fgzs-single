package cache

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/syncx"
	"sync"
	"time"
)

var (
	// can't use one SingleFlight per conn, because multiple conns may share the same cache key.
	singleFlight = syncx.NewSingleFlight()
)

// Key 实际key参数
type Key struct {
	keyPrefix *KeyPrefix
	buildKey  string
	lock      sync.RWMutex
}

// Key 获取构建好的key
func (p *Key) Key() string {
	return p.buildKey
}

// TTL 获取缓存key的过期时间time.Duration
func (p *Key) TTL() time.Duration {
	return p.keyPrefix.ExpirationTime
}

// TTLSecond 获取缓存key的过去时间 Second
func (p *Key) TTLSecond() int {
	return int(p.keyPrefix.ExpirationTime / time.Second)
}

// AutoCache 自动缓存
func (p *Key) AutoCache(rd *redis.Redis, result interface{}, fn func() (string, error)) error {
	val, err := singleFlight.Do(p.Key(), func() (interface{}, error) {
		p.lock.RLock()
		defer p.lock.RUnlock()
		res, err := rd.Get(p.Key())
		if err != nil && err != redis.Nil {
			return nil, err
		}
		if res != "" {
			return res, nil
		}
		res, err = fn()
		if err != nil {
			return nil, err
		}
		err = rd.Setex(p.Key(), res, p.TTLSecond())
		if err != nil {
			return nil, err
		}
		return res, nil
	})
	if err != nil {
		return err
	}
	s := val.(string)
	err = json.Unmarshal([]byte(s), result)
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除数据
func (p *Key) Delete(rd *redis.Redis) error {
	_, err := rd.Del(p.Key())
	if err != nil {
		return err
	}
	return nil
}
