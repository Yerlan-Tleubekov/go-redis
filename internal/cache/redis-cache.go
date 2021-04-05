package cache

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yerlan-tleubekov/go-redis/internal/models"
)

type RedisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) *RedisCache {
	return &RedisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return &redis.NewClient(&redis.Options{
		Addr: cache.host,
		Password: '',
		DB: cache.db
	})
}
