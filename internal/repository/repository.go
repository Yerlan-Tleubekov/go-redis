package repository

import (
	"github.com/yerlan-tleubekov/go-redis/internal/cache"
)

type Repository struct {
	db       int
	memCache *cache.RedisCache
}

// func NewRepository(db int, memCache *cache.RedisCache) *Repository {
func NewRepository(db int, memCache *cache.RedisCache) *Repository {
	return &Repository{db, memCache}
}
