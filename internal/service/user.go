package service

import (
	"net/http"

	"github.com/yerlan-tleubekov/go-redis/internal/models"
)

type IUser interface {
	GetUser(string) (*models.User, int, error)
}

func (service *Service) GetUser(login string) (*models.User, int, error) {

	user, err := service.repository.GetUser(login)

	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return user, http.StatusOK, nil
}
