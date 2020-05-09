// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package cache

import (
	"github.com/zerjioang/gotools/lib/concurrentmap"
	"github.com/zerjioang/gotools/lib/logger"
	"github.com/zerjioang/gotools/util/str"
)

var (
	defaultCache *MemoryCache
)

func init() {
	defaultCache = NewMemoryCache()
}

type MemoryCache struct {
	c concurrentmap.ConcurrentMap
}

func (cache MemoryCache) Get(key []byte) (interface{}, bool) {
	logger.Debug("reading value from global memory cache")
	v, ok := cache.c.Get(str.UnsafeString(key))
	return v, ok
}

func (cache *MemoryCache) Set(key []byte, value interface{}) {
	logger.Debug("settings new value on global memory cache")
	cache.c.Set(str.UnsafeString(key), value)
}

func NewMemoryCache() *MemoryCache {
	logger.Debug("creating new memory cache")
	m := new(MemoryCache)
	m.c = concurrentmap.New()
	return m
}

func Instance() *MemoryCache {
	return defaultCache
}
