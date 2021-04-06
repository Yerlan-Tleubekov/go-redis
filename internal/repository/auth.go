package repository

import "fmt"

func (repo *Repository) SignUp() {

}

func (repo *Repository) SignIn(userID int, token string) error {
	userIDStr := fmt.Sprint(userID)

	if err := repo.memCache.Set(userIDStr, token); err != nil {
		return err
	}

	return nil

}
