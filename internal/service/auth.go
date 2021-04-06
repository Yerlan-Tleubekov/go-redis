package service

import (
	"errors"
	"net/http"

	"github.com/yerlan-tleubekov/go-redis/internal/models"
	"github.com/yerlan-tleubekov/go-redis/pkg/jwt"
)

type Authenticator interface {
	SignUp(user *models.User)
	SignIn(userID int)
}

func (service *Service) SignUp(user *models.User) {

}

func (service *Service) SignIn(userID int) (string, error, int) {

	token, err := jwt.CreateToken(uint64(userID))
	if err != nil {
		return "", errors.New("Server error"), http.StatusInternalServerError
	}

	if err = service.repository.SignIn(userID, token); err != nil {
		return "", err, http.StatusInternalServerError
	}

	return token, nil, http.StatusOK
}
