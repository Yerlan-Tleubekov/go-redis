package jwt

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId uint64) (string, error) {
	key := os.Getenv("ACCESS_SECRET") //this should be in an env file
	atClaims := jwt.MapClaims{}

	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(key))

	if err != nil {
		return "", err
	}

	return token, nil
}
