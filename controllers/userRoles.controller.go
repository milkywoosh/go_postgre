package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/repositories"
	"github.com/milkyway/gin_beginer/services"
)

type UserRolesController struct {
	UserRolesRepo    repositories.UserRolesRepo
	UserRolesService services.UserRolesService
}

// NOTE: harusnya function call ke DB dipisah dari controllers !

// constructor
func NewUserRolesController(arg_db *sql.DB) UserRolesController {
	return UserRolesController{
		UserRolesRepo: repositories.UserRolesRepo{
			DB: arg_db,
		}, // mnote : harusnya userroles compose ke service!
		UserRolesService: services.UserRolesService{
			UserRolesRepo: repositories.UserRolesRepo{
				DB: arg_db,
			},
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
	var id_param_int int
	var err error = nil
	id_param, ok := ctx.Params.Get("user_id")
	if !ok {
		ru.badRequestErrorResp("failed get param id", "error", ctx)
		return
	}

	id_param_int, err = strconv.Atoi(id_param)
	if err != nil {
		ru.badRequestErrorResp("failed get param", "id conversion failed", ctx)
		return
	}

	var AllRolesData []repositories.RoleStruct
	AllRolesData, err = ru.UserRolesRepo.FetchRoleByID(id_param_int, ctx)
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
	var err error = nil
	if err = ctx.ShouldBindJSON(&AssignRoleDataReqBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"fail":    "fail DataUserReqBody",
		})
		return
	}

	err = ru.UserRolesRepo.InsertNewUserRole(ctx, AssignRoleDataReqBody.Username, AssignRoleDataReqBody.RoleName)
	if err != nil {
		ru.unprocessableEntityErrorResp("err insert new user_role", err.Error(), ctx)
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"err":     "no_data",
		"message": "success add new role",
	})
}

func (ru UserRolesController) DeleteRoleServiceBeginTx(ctx *gin.Context) {

	// assign role 1 by 1
	// request => username, rolename insert user_roles role_id, user_id values (1=admin, 203=benten, if each not found? handled by rollback?)
	type reqAssignRole struct {
		Username string `json:"username"`
		RoleName string `json:"rolename"`
	}

	var DeleteRoleDataReqBody reqAssignRole
	var err error = nil
	if err = ctx.ShouldBindJSON(&DeleteRoleDataReqBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"fail":    "fail DataUserReqBody",
		})
		return
	}

	err = ru.UserRolesService.DeleteRoleOfUser(ctx, DeleteRoleDataReqBody.Username, DeleteRoleDataReqBody.RoleName)
	if err != nil {
		ru.unprocessableEntityErrorResp("failed delete role transaction", err.Error(), ctx)
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"err":     "no_data",
		"message": "success delete role",
	})
}

func (ru UserRolesController) GetRolesByUsername(ctx *gin.Context) {
	var username_param string
	var ok bool
	var err error = nil

	username_param, ok = ctx.Params.Get("username")
	if !ok {
		ru.badRequestErrorResp("failed get param username", "error", ctx)
		return
	}

	var info_roles []repositories.RoleStruct = []repositories.RoleStruct{}
	info_roles, err = ru.UserRolesService.WhatRoles(ctx, username_param)
	if err != nil {
		ru.unprocessableEntityErrorResp("error check what roles", err.Error(), ctx)
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"data":    info_roles,
		"message": "success get roles",
	})

}
