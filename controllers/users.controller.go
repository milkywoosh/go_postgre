package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/models"
	"github.com/milkyway/gin_beginer/repositories"
	"github.com/milkyway/gin_beginer/services"
	"github.com/milkyway/gin_beginer/utils"
)

type UsersController struct {
	// DB        *sql.DB // note => refactor this into Repository folder to make code cleaner
	// UsersRepo    repositories.UsersRepo
	UsersService services.UsersService
}

// NOTE: harusnya function call ke DB dipisah dari controllers !

// constructor
func NewUsersController(db_arg *sql.DB) UsersController {
	return UsersController{
		UsersService: services.UsersService{
			UsersRepo: repositories.UsersRepo{
				DB: db_arg,
			},
		},
	}
}

func (uc UsersController) badRequestErrorResp(message string, err string, ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": message,
		"err":     err,
	})
}

func (uc UsersController) unprocessableEntityErrorResp(message string, err string, ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": message,
		"err":     err,
	})
}

func (uc UsersController) RegistrationNewUser(ctx *gin.Context) {

	var err error
	var UsersModel models.Users

	if err = ctx.ShouldBindJSON(&UsersModel); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"info":    "fail1",
		})
		return
	}

	var info string
	info, err = uc.UsersService.RegisterNewUser(ctx, UsersModel)
	if err != nil {
		uc.unprocessableEntityErrorResp("err insert new user", err.Error(), ctx)
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": fmt.Sprintf("insert new user, %s", info),
	})
}

func (uc UsersController) GetUserByID(ctx *gin.Context) {
	// var Users *models.Users
	var username string
	var err error
	var id_param_int int

	// bisa pake ini user request PATH params :id
	// id_param := ctx.Param("id")
	// bisa pake ini get path :id
	idx_query_param, ok := ctx.Params.Get("id")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "failed get param",
		})
		return
	}

	id_param_int, err = strconv.Atoi(idx_query_param)
	if err != nil {
		uc.badRequestErrorResp("failed get param", "id conversion failed", ctx)
		return
	}

	username, err = uc.UsersService.UsersRepo.FetchUserByID(ctx, id_param_int)
	if err != nil {
		uc.unprocessableEntityErrorResp("err fetch user", err.Error(), ctx)
		return
	}

	var token string
	token, err = utils.CreateToken(username)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{
			"message": "Token Err",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"data_user": username,
		"token":     token,
		"message":   "ok",
	})

}

func (uc UsersController) Login(c *gin.Context) {

	var DataUserReqBody models.Users

	var err error
	if err = c.ShouldBindJSON(&DataUserReqBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"fail":    "fail DataUserReqBody",
		})
		return
	}

	var token string
	var info_username string

	token, info_username, err = uc.UsersService.AuthLoginProcess(c, DataUserReqBody.Username, DataUserReqBody.Password)
	fmt.Println(DataUserReqBody.Username, DataUserReqBody.Password)
	if err != nil {
		uc.unprocessableEntityErrorResp("Err auth login process", err.Error(), c)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data_user": info_username,
		"token":     token,
		"message":   "ok",
	})
}
