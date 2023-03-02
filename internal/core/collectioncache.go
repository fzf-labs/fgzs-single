package core

import (
	"github.com/zeromicro/go-zero/core/collection"
	"time"
)

// NewDefaultCollectionCache 默认进程内缓存
func NewDefaultCollectionCache() *collection.Cache {
	return NewCollectionCache("default", time.Minute, 10000)
}

// NewCollectionCache 进程内缓存
func NewCollectionCache(name string, duration time.Duration, limit int) *collection.Cache {
	cache, err := collection.NewCache(duration, collection.WithLimit(limit), collection.WithName(name))
	if err != nil {
		panic("collectionCache err:" + err.Error())
	}
	return cache
}
