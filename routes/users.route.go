package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/controllers"
)

type UsersRouteController struct {
	usersController controllers.UsersController
}

func NewRouteUsersController(usersController controllers.UsersController) UsersRouteController {
	return UsersRouteController{usersController}
}

func (uc *UsersRouteController) UsersRoute(routerGroup *gin.RouterGroup) {
	router := routerGroup.Group("users")
	router.POST("/create-new-user", uc.usersController.InsertNewUser)
}
