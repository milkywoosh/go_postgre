package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/controllers"
)

// create inherited struct type from PersonController
type PersonRouteController struct {
	personController controllers.PersonController
}

func NewRoutePersonController(personController controllers.PersonController) PersonRouteController {
	return PersonRouteController{personController}
}

func (pc *PersonRouteController) PersonRoute(routerGroup *gin.RouterGroup) {
	router := routerGroup.Group("person")
	// middleware Deserialize LATER IMPLEMENT
	router.GET("/subject-info", pc.personController.PersonSubjectInfo)
	router.POST("/create-new", pc.personController.CreateNewPerson)
}
