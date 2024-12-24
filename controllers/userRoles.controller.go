package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/repositories"
)

type UserRolesController struct {
	DB            *sql.DB // note => refactor this into Repository folder to make code cleaner
	UserRolesRepo repositories.UserRolesRepo
}

// NOTE: harusnya function call ke DB dipisah dari controllers !

// constructor
func NewUserRolesController(arg_db *sql.DB) UserRolesController {
	return UserRolesController{
		UserRolesRepo: repositories.UserRolesRepo{
			DB: arg_db,
		},
	}
}

// coba dipisah sebagai interface ?
func (ru UserRolesController) badRequestErrorResp(message string, err string, ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": message,
		"err":     err,
	})
}

func (ru UserRolesController) unprocessableEntityErrorResp(message string, err string, ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": message,
		"err":     err,
	})
}

func (ru UserRolesController) GetRoleOfUser(ctx *gin.Context) {
	id_user_query_param, ok := ctx.Params.Get("user_id")
	if !ok {
		ru.badRequestErrorResp("failed get param id", "error", ctx)
		return
	}

	var AllRolesData []repositories.RoleStruct
	var err error
	AllRolesData, err = ru.UserRolesRepo.FetchRoleByID(id_user_query_param, ctx)
	if err != nil {
		ru.unprocessableEntityErrorResp("err fetch role by id", err.Error(), ctx)
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"data_user": AllRolesData,
		"message":   "ok",
	})
}

// // check perbedaan
// func (ru UserRolesController) AssignRolesBeginTx(ctx *gin.Context) {

// 	// assign role 1 by 1
// 	// request => username, rolename insert user_roles role_id, user_id values (1=admin, 203=benten, if each not found? handled by rollback?)
// 	type reqAssignRole struct {
// 		Username string `json:"username"`
// 		RoleName string `json:"rolename"`
// 	}

// 	var AssignRoleDataReqBody reqAssignRole
// 	var err error
// 	if err = ctx.ShouldBindJSON(&AssignRoleDataReqBody); err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"message": err.Error(),
// 			"fail":    "fail DataUserReqBody",
// 		})
// 		return
// 	}

// 	tx, err := ru.DB.BeginTx(ctx, nil)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"message": "err1 trx",
// 			"info":    err,
// 		})
// 		return
// 	}

// 	var user_id int
// 	err = tx.QueryRow("select id from users u where u.username = $1", AssignRoleDataReqBody.Username).Scan(&user_id)

// 	switch {
// 	case err == sql.ErrNoRows:
// 		log.Printf("no username %s", AssignRoleDataReqBody.Username)
// 		ru.badRequestErrorResp("err check username trx", err.Error(), ctx)
// 		return
// 	case err != nil:
// 		log.Printf("err %v", err)
// 	case !(err != nil):
// 		log.Printf("err nil check username %v", err)
// 	default:
// 		log.Printf("err check def username %v", err)
// 	}

// 	var role_id int
// 	err = tx.QueryRow("select id from roles r where r.role_name = $1", AssignRoleDataReqBody.RoleName).Scan(&role_id)

// 	switch {
// 	case err == sql.ErrNoRows:
// 		log.Printf("no role_name %s", AssignRoleDataReqBody.RoleName)
// 		if rollbackErr := tx.Rollback(); rollbackErr != nil {
// 			log.Fatalf("unable to rollback: %v", rollbackErr)
// 		}
// 		ru.badRequestErrorResp("err check role_name trx", err.Error(), ctx)
// 		return
// 	case err != nil:
// 		if rollbackErr := tx.Rollback(); rollbackErr != nil {
// 			log.Fatalf("unable to rollback: %v", rollbackErr)
// 		}
// 		log.Fatalf("gett err after rollback %v", err.Error())
// 	case !(err != nil):
// 		log.Printf("err nil check role_name %v", err)
// 	default:
// 		log.Printf("err check def role_name %v", err)
// 	}

// 	insert_user_roles := "INSERT INTO user_roles (role_id, user_id) VALUES ($1, $2)"

// 	// var resultInsertUserRoles sql.Result // note : return <nil>, error kalo value di pake => fmt.Println()

// 	_, err = tx.ExecContext(ctx, insert_user_roles, role_id, user_id)
// 	if err != nil {
// 		fmt.Println("errrr here!")
// 		if rollbackErr := tx.Rollback(); rollbackErr != nil {
// 			log.Fatalf("unable to rollback: %v", rollbackErr)
// 		}
// 		log.Println("rollbackErr: ", err.Error())
// 		ru.unprocessableEntityErrorResp("err insert assign new role", err.Error(), ctx)
// 		return
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		// note : err commit terjadi ketika sudah rollback tapi belum di return ??
// 		ru.unprocessableEntityErrorResp("err commit trx assign new role", err.Error(), ctx)
// 		return
// 	}

// 	ctx.JSON(http.StatusAccepted, gin.H{
// 		"err":     "no_data",
// 		"message": "success add new role",
// 	})
// }

// func (ru UserRolesController) DeleteRoleBeginTx(ctx *gin.Context) {

// 	// assign role 1 by 1
// 	// request => username, rolename insert user_roles role_id, user_id values (1=admin, 203=benten, if each not found? handled by rollback?)
// 	type reqAssignRole struct {
// 		Username string `json:"username"`
// 		RoleName string `json:"rolename"`
// 	}

// 	var DeleteRoleDataReqBody reqAssignRole
// 	var err error
// 	if err = ctx.ShouldBindJSON(&DeleteRoleDataReqBody); err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"message": err.Error(),
// 			"fail":    "fail DataUserReqBody",
// 		})
// 		return
// 	}

// 	tx, err := ru.DB.BeginTx(ctx, nil)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"message": "err1 trx",
// 			"info":    err,
// 		})
// 		return
// 	}

// 	var user_id int
// 	err = tx.QueryRow("select id from users u where u.username = $1", DeleteRoleDataReqBody.Username).Scan(&user_id)

// 	switch {
// 	case err == sql.ErrNoRows:
// 		log.Printf("no username %s", DeleteRoleDataReqBody.Username)
// 		ru.badRequestErrorResp("err check username trx", err.Error(), ctx)
// 		return
// 	case err != nil:
// 		log.Printf("err %v", err)
// 	case !(err != nil):
// 		log.Printf("err nil check username %v", err)
// 	default:
// 		log.Printf("err check def username %v", err)
// 	}

// 	var role_id int
// 	err = tx.QueryRow("select id from roles r where r.role_name = $1", DeleteRoleDataReqBody.RoleName).Scan(&role_id)

// 	switch {
// 	case err == sql.ErrNoRows:
// 		log.Printf("no role_name %s", DeleteRoleDataReqBody.RoleName)
// 		if rollbackErr := tx.Rollback(); rollbackErr != nil {
// 			log.Fatalf("unable to rollback: %v", rollbackErr)
// 		}
// 		ru.badRequestErrorResp("err check role_name trx", err.Error(), ctx)
// 		return
// 	case err != nil:
// 		if rollbackErr := tx.Rollback(); rollbackErr != nil {
// 			log.Fatalf("unable to rollback: %v", rollbackErr)
// 		}
// 		log.Fatalf("gett err after rollback %v", err.Error())
// 	case !(err != nil):
// 		log.Printf("err nil check role_name %v", err)
// 	default:
// 		log.Printf("err check def role_name %v", err)
// 	}

// 	delete_user_roles := "DELETE FROM user_roles WHERE role_id=$1 AND user_id=$2"

// 	_, err = tx.ExecContext(ctx, delete_user_roles, role_id, user_id)
// 	if err != nil {
// 		fmt.Println("errrr here!")
// 		if rollbackErr := tx.Rollback(); rollbackErr != nil {
// 			log.Fatalf("unable to rollback: %v", rollbackErr)
// 		}
// 		log.Println("rollbackErr: ", err.Error())
// 		ru.unprocessableEntityErrorResp("err delete role user", err.Error(), ctx)
// 		return
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		// note : err commit terjadi ketika sudah rollback tapi belum di return ??
// 		ru.unprocessableEntityErrorResp("err commit trx delete user role", err.Error(), ctx)
// 		return
// 	}

// 	ctx.JSON(http.StatusAccepted, gin.H{
// 		"err":     "no_data",
// 		"message": "success delete role",
// 	})
// }
