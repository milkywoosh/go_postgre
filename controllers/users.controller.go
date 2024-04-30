package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/models"
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
