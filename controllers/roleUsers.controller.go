package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleUsersController struct {
	DB *sql.DB // note => refactor this into Repository folder to make code cleaner
}

// NOTE: harusnya function call ke DB dipisah dari controllers !

// constructor
func NewRoleUsersController(arg_db *sql.DB) RoleUsersController {
	return RoleUsersController{
		DB: arg_db,
	}
}

func (ru RoleUsersController) GetRoleOfUser(ctx *gin.Context) {
	id_user_query_param, ok := ctx.Params.Get("user_id")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "failed get param id",
		})
		return
	}

	get_by_id_query := `
		SELECT r.role_name, u.username from user_roles ur
		INNER JOIN users u on u.id = ur.user_id
		INNER JOIN roles r on r.id = ur.role_id
		where u.id = $1 
	`

	type RoleStruct struct {
		Username string `json:"username"`
		RoleName string `json:"rolename"`
	}

	var EachRoleData RoleStruct
	var AllRolesData []RoleStruct

	rows, err := ru.DB.QueryContext(ctx, get_by_id_query, id_user_query_param)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "err1",
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&EachRoleData.RoleName,
			&EachRoleData.Username,
		); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.

			log.Fatal(err)
		}

		AllRolesData = append(AllRolesData, EachRoleData)

	}

	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(rerr)
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"data_user": AllRolesData,
		"message":   "ok",
	})
}

func (ru RoleUsersController) AssignRoles(ctx *gin.Context) {

	ctx.JSON(http.StatusAccepted, gin.H{
		"data_user": "success",
		"message":   "ok",
	})
}
