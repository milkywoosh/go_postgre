package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/models"
	"github.com/milkyway/gin_beginer/utils"
)

type UsersController struct {
	DB *sql.DB
}

func NewUsersController(db_arg *sql.DB) UsersController {
	return UsersController{
		DB: db_arg,
	}
}

func (uc UsersController) InsertNewUser(ctx *gin.Context) {

	var tx *sql.Tx
	var err error
	var UsersModel *models.Users

	if err = ctx.ShouldBindJSON(&UsersModel); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tx, err = uc.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}
	insertUsersQry := `insert into users(username, fullname) values($1, $2)`
	// test case duplicate entry username : success
	_, err = tx.ExecContext(ctx, insertUsersQry, &UsersModel.UserName, &UsersModel.FullName)

	if err != nil {
		tx.Rollback()

		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error":   err.Error(),
			"message": "err execContext",
		})
		return
	}

	err = tx.Commit()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{
			"error":   err.Error(),
			"message": "err commit",
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "insert new user success",
	})
	return
}

func (uc UsersController) GetUserByID(ctx *gin.Context) {
	// var Users *models.Users
	var err error

	// bisa pake ini user request PATH params :id
	// id_param := ctx.Param("id")
	// bisa pake ini get path :id
	idx_query_param, ok := ctx.Params.Get("id")

	fmt.Println(idx_query_param)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	get_by_id_query := `select username from users u where u.id_user = $1`

	var username *string
	var rows *sql.Row = uc.DB.QueryRowContext(ctx, get_by_id_query, idx_query_param)
	if err = rows.Scan(&username); err != nil {
		ctx.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{
			"message": "StatusFailedDependency",
			"error":   err.Error(),
		})
		return
	}
	var token string
	fmt.Println("tess: ", *username)
	token, err = utils.CreateToken(*username)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{
			"message": "Token Err",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"data_user": *username,
		"token":     token,
		"message":   "ok",
	})

}

func (uc UsersController) Login(c *gin.Context) {

	type Log struct {
		Username string
		Password string
	}

	var DataUser *Log

	if err := c.BindJSON(&DataUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	fmt.Println("===> ", DataUser.Username, DataUser.Password)

	// Dummy credential check
	if string(DataUser.Username) == "ben1" && string(DataUser.Password) == "password" {
		tokenString, err := utils.CreateToken(DataUser.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Error creating token")
			return
		}

		fmt.Printf("Token created: %s\n", tokenString)
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
