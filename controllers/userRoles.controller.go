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
func (ru UserRolesController) AssignRolesBeginTx(ctx *gin.Context) {

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

	err = ru.UserRolesRepo.InsertNewUserRole(AssignRoleDataReqBody.Username, AssignRoleDataReqBody.RoleName, ctx)
	if err != nil {
		ru.unprocessableEntityErrorResp("err insert new user_role", err.Error(), ctx)
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"err":     "no_data",
		"message": "success add new role",
	})
}

func (ru UserRolesController) DeleteRoleBeginTx(ctx *gin.Context) {

	// assign role 1 by 1
	// request => username, rolename insert user_roles role_id, user_id values (1=admin, 203=benten, if each not found? handled by rollback?)
	type reqAssignRole struct {
		Username string `json:"username"`
		RoleName string `json:"rolename"`
	}

	var DeleteRoleDataReqBody reqAssignRole
	var err error
	if err = ctx.ShouldBindJSON(&DeleteRoleDataReqBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"fail":    "fail DataUserReqBody",
		})
		return
	}

	err = ru.UserRolesRepo.DeleteUserRole(DeleteRoleDataReqBody.Username, DeleteRoleDataReqBody.RoleName, ctx)
	if err != nil {
		ru.unprocessableEntityErrorResp("failed delete role transaction", err.Error(), ctx)
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"err":     "no_data",
		"message": "success delete role",
	})
}
