package repository

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/yerlan-tleubekov/go-redis/internal/models"
)

type Authorization interface {
	SignIn(int, string) error
	SignUp(*models.User) error
}

func (repo *Repository) SignUp(user *models.User) error {

	userJSON, err := json.Marshal(user)

	if err != nil {
		return errors.New("Server error")
	}

	if err = repo.memCache.SetUser(user.Login, string(userJSON)); err != nil {
		return err
	}

	return nil

}

func (repo *Repository) SignIn(userID int, token string) error {
	userIDStr := fmt.Sprint(userID)

	if err := repo.memCache.SetToken(userIDStr, token); err != nil {
		return err
	}

	return nil

}
