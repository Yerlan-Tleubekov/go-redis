package cache

import (
	"encoding/json"
	"errors"

	"github.com/yerlan-tleubekov/go-redis/internal/models"
)

type UserCache interface {
	GetUser(string) (*models.User, error)
	SetUser(string, string) error
}

func (cache *RedisCache) GetUser(login string) (*models.User, error) {
	var user models.User

	client := cache.getClient()

	userJSON, err := client.Get(cache.ctx, login).Result()

	if err != nil {
		return nil, errors.New("User doesnt exist")
	}

	if err = json.Unmarshal([]byte(userJSON), &user); err != nil {
		return nil, errors.New("server error")
	}

	return &user, nil

}

func (cache *RedisCache) SetUser(login, user string) error {

	client := cache.getClient()

	_, err := client.Set(cache.ctx, login, user, cache.expires).Result()

	if err != nil {
		return errors.New("Server error")
	}

	return nil
}
