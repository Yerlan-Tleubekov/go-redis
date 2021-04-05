package service

import (
	"fmt"

	"github.com/yerlan-tleubekov/go-redis/internal/models"
	"github.com/yerlan-tleubekov/go-redis/pkg/jwt"
)

type Authenticator interface {
	SignUp(user *models.User)
	SignIn(userID int)
}

func (service *Service) SignUp(user *models.User) {

}

func (service *Service) SignIn(userID int) {
	token, err := jwt.CreateToken(uint64(userID))
	if err != nil {
		return
	}

	fmt.Println(token)
}
