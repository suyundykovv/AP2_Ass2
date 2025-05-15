package repository

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type CacheStore struct {
	cache *cache.Cache
}

func NewCacheStore() *CacheStore {
	return &CacheStore{
		cache: cache.New(12*time.Hour, 1*time.Hour), // Default expiration 12h, cleanup every 1h
	}
}

func (c *CacheStore) Set(key string, value interface{}) {
	c.cache.Set(key, value, cache.DefaultExpiration)
}

func (c *CacheStore) Get(key string) (interface{}, bool) {
	return c.cache.Get(key)
}

func (c *CacheStore) Delete(key string) {
	c.cache.Delete(key)
}

func (c *CacheStore) Flush() {
	c.cache.Flush()
}
