package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/controllers"
)

type OrdersRouteController struct {
	OrdersController *controllers.OrdersController
}

func NewRouteOrdersController(OrdersController *controllers.OrdersController) *OrdersRouteController {
	return &OrdersRouteController{
		OrdersController,
	}
}

func (orc *OrdersRouteController) OrdersRoutes(routerGroup *gin.RouterGroup) {
	router := routerGroup.Group("orders")

	router.GET("/test", orc.OrdersController.OrdersCheck)
}
