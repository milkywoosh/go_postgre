package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/controllers"
)

type RoleUsersRouteController struct {
	roleUsersController controllers.RoleUsersController
}

func NewRouteRoleUsersController(roleUsersController controllers.RoleUsersController) RoleUsersRouteController {
	return RoleUsersRouteController{roleUsersController}
}

func (ru *RoleUsersRouteController) RoleUsersRoute(routerGroup *gin.RouterGroup) {
	router := routerGroup.Group("roleusers")

	router.GET("/:user_id", ru.roleUsersController.GetRoleOfUser)
}
