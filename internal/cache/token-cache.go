package cache

import (
	"encoding/json"
	"time"

	"github.com/yerlan-tleubekov/go-redis/internal/models"
)

type TokenCache interface {
	Set(key string, value *models.UserToken)
	Get(key string) *models.UserToken
}

func (cache *redisCache) Set(key string, value *models.UserToken) {
	client := cache.getClient()

	json, err := json.Marshal(value)

	if err != nil {
		panic(err)
	}

	client.Set(key, json, cache.expires*time.Second)

}

func (cache *redisCache) Get(key string) *models.UserToken {
	client := cache.getClient()

	val, err := client.Get(key).Result()

	if err != nil {
		return nil
	}

	userToken := models.UserToken{}

	err := json.Unmarshal([]byte(val), &userToken)

	if err != nil {
		return nil
	}

	return &userToken
}
