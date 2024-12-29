package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/controllers"
)

type BooksRouteController struct {
	BooksController controllers.BooksController
}

func NewRouteBooksController(BooksController controllers.BooksController) BooksRouteController {
	return BooksRouteController{BooksController}
}

func (brc BooksRouteController) BooksRoutes(routerGroup *gin.RouterGroup) {
	router := routerGroup.Group("books")

	router.GET("/:id", brc.BooksController.GetBookByID)
	router.GET("/search", brc.BooksController.SearchBooksByName)
	router.GET("/search-by-author", brc.BooksController.SearchBooksByAuthorName)
}
