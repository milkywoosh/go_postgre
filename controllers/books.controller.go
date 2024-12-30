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
)

type BooksController struct {
	BooksService services.BooksService
}

// constructor
// note : possibly using interface ? look chatGPT
func NewBooksController(db_arg *sql.DB) BooksController {
	return BooksController{
		BooksService: services.BooksService{
			BooksRepo: repositories.BooksRepo{
				DB: db_arg,
			},
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
	var book_info_rows []models.Books
	var err error
	var id_param string
	var id_param_int int
	id_param, ok := ctx.Params.Get("id")

	if !ok {
		bc.badRequestErrorResp("failed get param", "id param not found", ctx)
		return
	}
	id_param_int, err = strconv.Atoi(id_param)
	if err != nil {
		bc.badRequestErrorResp("failed get param", "id conversion failed", ctx)
		return
	}

	book_info_rows, err = bc.BooksService.GetBookInfo(ctx, id_param_int)
	if err != nil {
		bc.unprocessableEntityErrorResp("err get book by ID", err.Error(), ctx)
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"book_info": book_info_rows,
		"message":   "ok",
	})

}

func (bc BooksController) SearchBooksByName(ctx *gin.Context) {
	var reqBody models.Books
	var err error

	err = ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		bc.badRequestErrorResp("failed get req body", err.Error(), ctx)
		return
	}

	var searchResults []models.BooksLike

	searchResults, err = bc.BooksService.SearchLikeName(ctx, reqBody.BookName)
	if err != nil {
		bc.unprocessableEntityErrorResp("err search book like keyword", err.Error(), ctx)
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"result":  searchResults,
		"token":   "no token",
		"message": "ok",
	})

}

func (bc BooksController) SearchBooksByAuthorName(ctx *gin.Context) {
	var reqBody models.Author
	var err error

	err = ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		bc.badRequestErrorResp("failed get req body", err.Error(), ctx)
		return
	}

	var searchResults []models.BooksAuthorLike

	searchResults, err = bc.BooksService.SearchBookByAuthor(ctx, reqBody.AuthorName)
	if err != nil {
		bc.unprocessableEntityErrorResp("err search book like keyword", err.Error(), ctx)
		return
	}

	fmt.Println("ini kondisi setelah proses call to repo")

	// note: ctx.JSON sama dgn json.Marshal(), implement Marshaler interface
	ctx.JSON(http.StatusAccepted, gin.H{
		"result":  searchResults, //searchResults,
		"token":   "no token",
		"message": "ok",
	})

}
