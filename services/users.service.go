package services

// note handle sisi business logic

import (
	"context"

	"github.com/milkyway/gin_beginer/repositories"
)

type UsersService struct {
	UsersRepo repositories.UsersRepo
}

func (us UsersService) GetInfoUsernamePassword(ctx context.Context, username string) (string, string, error) {

	var username_val string
	var hash_password string
	var err error

	username_val, hash_password, err = us.UsersRepo.FetchUsernamePassword(ctx, username)
	if err != nil {
		return "", "", err
	}
	return username_val, hash_password, nil

}
