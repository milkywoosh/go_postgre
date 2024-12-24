package repositories

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type UserRolesRepo struct {
	DB *sql.DB
}

type RoleStruct struct {
	Username string `json:"username"`
	RoleName string `json:"rolename"`
}

func (ur UserRolesRepo) FetchRoleByID(user_id string, ctx *gin.Context) ([]RoleStruct, error) {

	get_by_id_query := `
		SELECT r.role_name, u.username from user_roles ur
		INNER JOIN users u on u.id = ur.user_id
		INNER JOIN roles r on r.id = ur.role_id
		where u.id = $1 
	`
	var EachRoleData RoleStruct
	var AllRolesData []RoleStruct

	rows, err := ur.DB.QueryContext(ctx, get_by_id_query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&EachRoleData.RoleName,
			&EachRoleData.Username,
		); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.

			return nil, err
		}

		AllRolesData = append(AllRolesData, EachRoleData)

	}

	rerr := rows.Close()
	if rerr != nil {
		return nil, err
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return AllRolesData, nil

}
