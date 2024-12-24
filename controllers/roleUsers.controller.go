package controllers

import (
	"database/sql"
	"fmt"
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

func (ru RoleUsersController) badRequestErrorResp(message string, err string, ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": message,
		"err":     err,
	})
	return
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

// check perbedaan
func (ru RoleUsersController) AssignRolesBeginTx(ctx *gin.Context) {

	// assign role 1 by 1
	// request => username, rolename insert user_roles role_id, user_id values (1=admin, 203=benten, if each not found? handled by rollback?)
	type reqAssignRole struct {
		Username string `json:"username"`
		RoleName string `json:"rolename"`
	}

	var AssignRoleDataReqBody reqAssignRole
	var err error
	if err = ctx.ShouldBindJSON(&AssignRoleDataReqBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"fail":    "fail DataUserReqBody",
		})
		return
	}

	fmt.Println("AssignRoleDataReqBody", AssignRoleDataReqBody)

	tx, err := ru.DB.BeginTx(ctx, nil)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "err1 trx",
			"info":    err,
		})
		return
	}

	var user_id int
	err = tx.QueryRow("select id from users u where u.username = $1", AssignRoleDataReqBody.Username).Scan(&user_id)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("no username %s", AssignRoleDataReqBody.Username)
		ru.badRequestErrorResp("err check username trx", err.Error(), ctx)
		return

	case err != nil:
		log.Printf("err %v", err)
	case !(err != nil):
		log.Printf("err nil check username %v", err)
	default:
		log.Printf("err check def username %v", err)
	}

	var role_id int
	err = tx.QueryRow("select id from roles r where r.role_name = $1", AssignRoleDataReqBody.RoleName).Scan(&role_id)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("no role_name %s", AssignRoleDataReqBody.RoleName)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("unable to rollback: %v", rollbackErr)
		}
		ru.badRequestErrorResp("err check role_name trx", err.Error(), ctx)
		return
	case err != nil:
		// log.Printf("err => %v", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("unable to rollback: %v", rollbackErr)
		}
		log.Fatalf("gett err after rollback %v", err.Error())
	case !(err != nil):
		log.Printf("err nil check role_name %v", err)
	default:
		log.Printf("err check def role_name %v", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"data_user": "success",
		"message":   "ok",
	})
}

// check perbedaan
func (ru RoleUsersController) AssignRolesExecContext(ctx *gin.Context) {

	// assign role 1 by 1
	// request => username, rolename
	// ru.DB.ExecContext()

	ctx.JSON(http.StatusAccepted, gin.H{
		"data_user": "success",
		"message":   "ok",
	})
}
