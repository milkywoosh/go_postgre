package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/controllers"
	"github.com/milkyway/gin_beginer/utils"
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
	router.GET("/all-person-data", utils.AuthenticateMiddleware, pc.personController.AllPerson)
	router.POST("/insert-new-person", pc.personController.InsertNewPerson)
	// router.POST("/create-new", pc.personController.CreateNewPerson)
	// router.POST("/upload-xl", pc.personController.UploadMultiplePerson)
	// harusnya buat router group "teacher"
	// router.GET("/call-deffunc/:id", pc.personController.PlSqlCallDefinedFuncOne)
}
