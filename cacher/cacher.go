package cacher

import (
	"sync"
	"time"
)

type Cacher interface {
	Get(string) (string, bool)
	Set(string, string)
}

type Cache struct {
	ttl     time.Duration
	storage map[string]CacheItem
	mu      sync.Mutex
}

type CacheItem struct {
	value     string
	createdAt time.Time
}

func (r *Cache) Get(key string) (string, bool) {
	defer r.mu.Unlock()
	r.mu.Lock()
	val, ok := r.storage[key]

	if !ok {
		return "", false
	}

	if time.Since(val.createdAt) > r.ttl {
		delete(r.storage, key)
		return "", false
	}

	return val.value, ok
}

func (r *Cache) Set(key, v string) {
	defer r.mu.Unlock()
	r.mu.Lock()
	r.storage[key] = CacheItem{
		createdAt: time.Now(),
		value:     v,
	}
}

func NewCache(ttl time.Duration) *Cache {
	m := make(map[string]CacheItem)
	return &Cache{
		ttl:     ttl,
		storage: m,
		mu:      sync.Mutex{},
	}
}
