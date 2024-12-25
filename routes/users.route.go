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
	// implement middelware
	// reference: https://permify.co/post/jwt-authentication-go/
	router.POST("/login", uc.usersController.Login)
	router.POST("/regist-new-user", uc.usersController.RegistrationNewUser)
	router.GET("/:id", uc.usersController.GetUserByID)
}
