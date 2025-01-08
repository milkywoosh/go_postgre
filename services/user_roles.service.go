package services

import (
	"context"
	"database/sql"

	"github.com/milkyway/gin_beginer/repositories"
)

// logic business
type UserRolesService struct {
	UserRolesRepo repositories.UserRolesRepo
}

func NewUserRolesService(arg_db *sql.DB) UserRolesService {
	return UserRolesService{
		UserRolesRepo: repositories.UserRolesRepo{
			DB: arg_db,
		},
	}
}

func (urs UserRolesService) WhatRoles(ctx context.Context, username string) ([]repositories.RoleStruct, error) {
	var roles_info []repositories.RoleStruct = []repositories.RoleStruct{}
	var err error
	roles_info, err = urs.UserRolesRepo.FetchRolesByUsername(ctx, username)

	if err != nil {
		return roles_info, err
	}
	return roles_info, nil
}
