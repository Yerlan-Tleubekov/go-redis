package cache

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/yerlan-tleubekov/go-redis/internal/models"
)

type TokenCache interface {
	Set(key string, value *models.UserToken)
	Get(key string) *models.UserToken
}

func (cache *RedisCache) Set(key string, value string) error {
	client := cache.getClient()

	// json, err := json.Marshal(value)

	status := client.Set(cache.ctx, key, value, cache.expires*time.Second)
	_, err := status.Result()
	if err != nil {
		return errors.New("Cant write user from db")
	}
	return nil

}

func (cache *RedisCache) Get(key string) *models.UserToken {
	client := cache.getClient()

	val, err := client.Get(cache.ctx, key).Result()

	if err != nil {
		return nil
	}

	userToken := models.UserToken{}

	err = json.Unmarshal([]byte(val), &userToken)

	if err != nil {
		return nil
	}

	return &userToken
}
