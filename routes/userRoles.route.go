package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/controllers"
)

type RoleUsersRouteController struct {
	UserRolesController controllers.UserRolesController
}

func NewRouteUserRolesController(UserRolesController controllers.UserRolesController) RoleUsersRouteController {
	return RoleUsersRouteController{UserRolesController}
}

func (ru *RoleUsersRouteController) RoleUsersRoute(routerGroup *gin.RouterGroup) {
	router := routerGroup.Group("roleusers")

	router.GET("/:user_id", ru.UserRolesController.GetRoleOfUser)
	router.GET("/get_roles/:username", ru.UserRolesController.GetRolesByUsername)
	router.POST("/assign_role", ru.UserRolesController.AssignRolesBeginTx)
	router.DELETE("/delete_role", ru.UserRolesController.DeleteRoleBeginTx)
}
