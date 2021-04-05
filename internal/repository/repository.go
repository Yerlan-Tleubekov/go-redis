package repository

type Repository struct {
	db int
}

func NewRepository(db int, memCache *cache.RedisCache) *Repository {
	return &Repository{db}
}
