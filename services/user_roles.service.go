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
	var err error = nil
	roles_info, err = urs.UserRolesRepo.FetchRolesByUsername(ctx, username)

	if err != nil {
		return roles_info, err
	}
	return roles_info, nil
}

func (urs UserRolesService) DeleteRoleOfUser(ctx context.Context, username, role_name string) error {
	var err error = nil
	err = urs.UserRolesRepo.DeleteUserRole(ctx, username, role_name)
	if err != nil {
		return err
	}
	return nil
}

func (urs UserRolesService) AssignRoleService(ctx context.Context, username, role_name string) (string, error) {
	var err error = nil
	err = urs.UserRolesRepo.AddNewRole(ctx, username, role_name)
	if err != nil {
		return "error", err
	}
	return "no error", nil
}
