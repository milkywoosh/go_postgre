package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/models"
	"github.com/milkyway/gin_beginer/utils"
	"golang.org/x/crypto/bcrypt"
)

type UsersController struct {
	DB *sql.DB
}

func NewUsersController(db_arg *sql.DB) UsersController {
	return UsersController{
		DB: db_arg,
	}
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

	var tx *sql.Tx
	var err error
	var UsersModel *models.Users

	if err = ctx.ShouldBindJSON(&UsersModel); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"fail":  "fail1",
		})
		return
	}

	hash_pass, err := uc.HashPasswordUser(*UsersModel.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error": err.Error(),
			"fail":  "fail2",
		})
		return
	}

	_, err = uc.DecryptPasswordUser(hash_pass, *UsersModel.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"fail":  "fail3",
		})
		return
	}

	// ctx.JSON(http.StatusAccepted, gin.H{
	// 	"hash_pass": hash_pass,
	// })

	tx, err = uc.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error": err.Error(),
			"fail":  "fail4",
		})
		return
	}
	insertUsersQry := `insert into users(username, fullname, password) values($1, $2, $3)`
	// test case duplicate entry username : success
	_, err = tx.ExecContext(ctx, insertUsersQry, &UsersModel.UserName, &UsersModel.FullName, hash_pass)

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

	var DataUserReqBody *models.Users
	var err error
	if err = c.ShouldBindJSON(&DataUserReqBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"fail":    "fail DataUserReqBody",
		})
		return
	}

	// var UserReq string = DataUserReqBody.Username
	// var PassReq string = DataUserReqBody.Password

	// c.JSON(200, gin.H{
	// 	"user": UserReq,
	// 	"pass": PassReq,
	// })
	// return

	// tampungan hash password fetch from DB
	var hash_password *string
	query_get_user := `select u.username, u.password, u.fullname from users u WHERE u.username = $1 limit 1`
	var row *sql.Row = uc.DB.QueryRowContext(c, query_get_user, DataUserReqBody.UserName)
	// scan: tampungan data fetch from DB
	err = row.Scan(&DataUserReqBody.UserName, &hash_password, &DataUserReqBody.FullName)

	if err != nil {
		log.Fatal("==>> ", err)
	}
	var decryptSuccess bool
	decryptSuccess, err = uc.DecryptPasswordUser(*hash_password, *DataUserReqBody.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "password is wrong",
			"data":    DataUserReqBody,
		})
		return
	}

	// SEND TOKEN TO COOKIE ==> validasi role atau username

	if decryptSuccess {
		tokenString, err := utils.CreateToken(*DataUserReqBody.UserName)
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
