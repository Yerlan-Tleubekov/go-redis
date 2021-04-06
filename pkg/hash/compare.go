package hash

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type ICompare interface {
	CompareHashAndUserPassword(string, string)
}

func CompareHashAndUserPassword(hashedPassword, userPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword)); err != nil {
		return errors.New("Incorrect password")
	}

	return nil
}
