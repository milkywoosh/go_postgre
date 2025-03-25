package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/controllers"
)

type CustomersRouteController struct {
	CustomersController controllers.CustomersController
}

func NewRouteCustomersController(CustomersController controllers.CustomersController) CustomersRouteController {
	return CustomersRouteController{CustomersController}
}

func (crc CustomersRouteController) CustomersRoutes(routerGroup *gin.RouterGroup) {
	router := routerGroup.Group("customers")

	router.GET("/:registry-cust", crc.CustomersController.GetBookByRegistryNumber)
	router.POST("/order-books", crc.CustomersController.OrderBooks)
}
