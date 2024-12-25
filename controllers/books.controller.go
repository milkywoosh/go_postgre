package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/models"
	"github.com/milkyway/gin_beginer/repositories"
)

type BooksController struct {
	BooksRepo repositories.BooksRepo
}

// constructor
func NewBooksController(db_arg *sql.DB) BooksController {
	return BooksController{
		BooksRepo: repositories.BooksRepo{
			DB: db_arg,
		},
	}
}

func (bc BooksController) badRequestErrorResp(message string, err string, ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": message,
		"err":     err,
	})
}

func (bc BooksController) unprocessableEntityErrorResp(message string, err string, ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": message,
		"err":     err,
	})
}

func (bc BooksController) GetBookByID(ctx *gin.Context) {
	var BooksModel models.Books
	var err error
	var id_param string
	id_param, ok := ctx.Params.Get("id")

	fmt.Println("id", id_param)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "failed get param",
		})
		return
	}

	BooksModel.BookName, err = bc.BooksRepo.FetchBookByID(ctx, id_param)
	if err != nil {
		bc.unprocessableEntityErrorResp("err get book by ID", err.Error(), ctx)
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"book_name": BooksModel.BookName,
		"message":   "ok",
	})

}
