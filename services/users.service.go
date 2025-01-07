package services

// note handle sisi business logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/milkyway/gin_beginer/models"
	"github.com/milkyway/gin_beginer/repositories"
	"github.com/milkyway/gin_beginer/utils"
)

type UsersService struct {
	UsersRepo repositories.UsersRepo
}

func (us UsersService) GetInfoUsernamePassword(ctx context.Context, username string) (string, string, error) {

	var username_val string
	var hash_password string
	var err error

	username_val, hash_password, err = us.UsersRepo.FetchPasswordByUsername(ctx, username)
	if err != nil {
		return "", "", err
	}
	return username_val, hash_password, nil

}

func (us UsersService) AuthLoginProcess(ctx context.Context, username string, password string) (token_string, success_user string, err error) {

	var hash_password string
	success_user, hash_password, err = us.UsersRepo.FetchPasswordByUsername(ctx, username)
	if err != nil {
		return "", "", err
	}

	var decryptSuccess bool
	decryptSuccess, err = utils.DecryptPasswordUser(hash_password, password)

	if err != nil {
		return "", "", err
	}

	if decryptSuccess {
		tokenString, err := utils.CreateToken(username)
		if err != nil {
			return "", "", err
		}

		return tokenString, success_user, nil

	} else {
		return "", "", errors.New("password is wrong")
	}

}

func (us UsersService) RegisterNewUser(ctx context.Context, UsersModel models.Users) (string, error) {
	var err error
	var username string = UsersModel.Username
	var hash_password string
	hash_password, err = utils.HashPasswordUser(UsersModel.Password)

	err = us.UsersRepo.InsertNewUser(ctx, UsersModel, hash_password)
	if err != nil {
		return "failed register", err
	}
	var info string = fmt.Sprintf("%s success register", username)
	return info, nil
}
