package core

import (
	"github.com/zeromicro/go-zero/core/collection"
	"time"
)

// NewDefaultCollectionCache 默认进程内缓存
func NewDefaultCollectionCache() *collection.Cache {
	cache, err := NewCollectionCache("default", time.Minute, 10000)
	if err != nil {
		panic("collectionCache err:" + err.Error())
	}
	return cache
}

// NewCollectionCache 进程内缓存
func NewCollectionCache(name string, duration time.Duration, limit int) (*collection.Cache, error) {
	cache, err := collection.NewCache(duration, collection.WithLimit(limit), collection.WithName(name))
	if err != nil {
		return nil, err
	}
	return cache, nil
}