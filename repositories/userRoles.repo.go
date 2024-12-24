package repositories

import "database/sql"

type RoleUsersRepo struct {
	DB *sql.DB
}

func