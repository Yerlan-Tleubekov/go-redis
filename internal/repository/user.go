package repository

import (
	"encoding/json"

	"github.com/yerlan-tleubekov/go-redis/internal/models"
)

type IUser interface {
	CreateUser(*models.User) error
	GetUser(string) (*models.User, error)
}

func (repo *Repository) CreateUser(user *models.User) error {
	userJSON, err := json.Marshal(user)

	if err != nil {
		return err
	}

	if err = repo.memCache.SetUser(user.Login, string(userJSON)); err != nil {
		return err
	}

	return nil
}

func (repo *Repository) GetUser(login string) (*models.User, error) {

	user, err := repo.memCache.GetUser(login)

	if err != nil {
		return nil, err
	}

	return user, nil
}
