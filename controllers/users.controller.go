package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/models"
	"github.com/milkyway/gin_beginer/repositories"
	"github.com/milkyway/gin_beginer/utils"
	"golang.org/x/crypto/bcrypt"
)

type UsersController struct {
	// DB        *sql.DB // note => refactor this into Repository folder to make code cleaner
	UsersRepo repositories.UsersRepo
}

// NOTE: harusnya function call ke DB dipisah dari controllers !

// constructor
func NewUsersController(db_arg *sql.DB) UsersController {
	return UsersController{
		UsersRepo: repositories.UsersRepo{
			DB: db_arg,
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

// reference: https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
func (uc UsersController) HashPasswordUser(password string) (string, error) {
	hashed_pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return string(""), nil
	}
	return string(hashed_pass), nil
}

func (uc UsersController) DecryptPasswordUser(hashed_password string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))
	if err != nil {
		return false, err
	}
	// if error nil
	return true, nil
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

	hash_pass, err := uc.HashPasswordUser(UsersModel.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"message": err.Error(),
			"info":    "fail2",
		})
		return
	}

	_, err = uc.DecryptPasswordUser(hash_pass, UsersModel.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"info":    "fail3",
		})
		return
	}

	err = uc.UsersRepo.InsertNewUser(UsersModel, hash_pass, ctx)
	if err != nil {
		uc.unprocessableEntityErrorResp("err insert new user", err.Error(), ctx)
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "insert new user success",
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

	username, err = uc.UsersRepo.FetchUserByID(ctx, id_param_int)
	if err != nil {
		uc.unprocessableEntityErrorResp("err fetch user", err.Error(), ctx)
		return
	}

	var token string
	fmt.Println("tess: ", username)
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

	var DataUserReqBody *models.Users
	// var username string
	var hash_password string

	var err error
	if err = c.ShouldBindJSON(&DataUserReqBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"fail":    "fail DataUserReqBody",
		})
		return
	}

	log.Println("DataUserReqBody.Username", DataUserReqBody.Username)
	_, hash_password, err = uc.UsersRepo.FetchUsernamePassword(c, DataUserReqBody.Username)
	if err != nil {
		log.Println("err fetch??")
		uc.badRequestErrorResp("err FetchUsernamePassword", err.Error(), c)
		return
	}

	var decryptSuccess bool
	decryptSuccess, err = uc.DecryptPasswordUser(hash_password, DataUserReqBody.Password)
	dataFailed := struct {
		Username string
		Email    string
	}{
		// note: struct must me started CAPITAL
		Username: DataUserReqBody.Username,
		Email:    DataUserReqBody.Email,
	}

	fmt.Println(dataFailed)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "password is wrong",
			"data":    dataFailed,
		})
		return
	}

	// SEND TOKEN TO COOKIE ==> validasi role atau username

	if decryptSuccess {
		tokenString, err := utils.CreateToken(DataUserReqBody.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Error creating token")
			return
		}

		// fmt.Printf("Token created: %s\n", tokenString)
		// check Cookie in header postman
		c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusAccepted, gin.H{
			"message": "login succeed",
			"token":   tokenString,
		})
		return
	} else {
		c.AbortWithStatusJSON(http.StatusNonAuthoritativeInfo, gin.H{
			"message": "wrong password",
		})
		return
	}
}
