package cache

import (
	"fmt"
	"strings"
	"time"
)

var KeyPrefixes = map[string]*KeyPrefix{}

// KeyPrefix 缓存key前缀管理
type KeyPrefix struct {
	PrefixName     string
	Remark         string
	ExpirationTime time.Duration
}

// Key 实际key参数
type Key struct {
	keyPrefix *KeyPrefix
	buildKey  string
}

func NewCacheKey(prefixName string, expirationTime time.Duration, remark string) *KeyPrefix {
	if _, ok := KeyPrefixes[prefixName]; ok {
		panic(fmt.Sprintf("cache key %s is exsit, please change one", prefixName))
	}
	key := &KeyPrefix{PrefixName: prefixName, ExpirationTime: expirationTime, Remark: remark}
	KeyPrefixes[prefixName] = key
	return key
}

// BuildCacheKey 构建一个带有前缀的缓存key 使用 ":" 分隔
func (p *KeyPrefix) BuildCacheKey(keys ...string) Key {
	cacheKey := Key{
		keyPrefix: p,
	}
	if len(keys) == 0 {
		cacheKey.buildKey = p.PrefixName
	} else {
		cacheKey.buildKey = strings.Join(append([]string{p.PrefixName}, keys...), ":")
	}
	return cacheKey
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
