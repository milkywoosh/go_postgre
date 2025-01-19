package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/milkyway/gin_beginer/controllers"
	"github.com/milkyway/gin_beginer/initializer"
	"github.com/milkyway/gin_beginer/routes"
)

var (
	server *gin.Engine
	// KENAPA HARUS PAKE POINTER TYPE????

	// User Manager
	UsersController      controllers.UsersController
	UsersRouteController routes.UsersRouteController

	UserRolesController      controllers.UserRolesController
	RoleUsersRouteController routes.RoleUsersRouteController

	BooksController      controllers.BooksController
	BooksRouteController routes.BooksRouteController
	// purchase manager

)

// init() function is RUN BEFORE main() function
// it is used to initiate connection to database
func init() {
	// load config by doing checking from very root dir with this => "."
	config, err := initializer.LoadConfig(".") // why 2x init ?
	if err != nil {
		log.Fatal("could not load environment variable", err)
	}

	initializer.StartConnectDB(&config)

	UsersController = controllers.NewUsersController(initializer.DB)
	UsersRouteController = routes.NewRouteUsersController(UsersController)

	UserRolesController = controllers.NewUserRolesController(initializer.DB)
	RoleUsersRouteController = routes.NewRouteUserRolesController(UserRolesController)

	BooksController = controllers.NewBooksController(initializer.DB)
	BooksRouteController = routes.NewRouteBooksController(BooksController)

	server = gin.Default()

}

func main() {
	fmt.Println("test")

	config, err := initializer.LoadConfig(".") // why 2x init ?
	if err != nil {
		log.Fatal("could not load environment variable", err)
	}

	// default
	corsConfig := cors.DefaultConfig()
	// trial
	// corsConfig := cors.Default()

	// port 8000 apa ? 3000 apa ?
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClienOrigin}
	// default
	corsConfig.AllowCredentials = true // kalo false??
	// trial
	// corsConfig.AllowCredentials = false // kalo false??

	// Default
	server.Use(cors.New(corsConfig))
	// limit multipart size upload
	// max 5 mb = 1 << 20
	// max 5 kb = 5 << 10
	server.MaxMultipartMemory = 1 * 1024 * 1024
	// trial
	// server.Use(corsConfig)

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Gin"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	UsersRouteController.UsersRoute(router)
	RoleUsersRouteController.RoleUsersRoute(router)
	BooksRouteController.BooksRoutes(router)

	// Default => deal with firewall
	run_server := fmt.Sprintf("127.0.0.1:%s", config.ServerPort)

	err = server.Run(run_server)

	log.Fatal(err)

	// Trial => byPass firewall checking
	// run_server := fmt.Sprintf("127.0.0.1:%d", 8000)
	// server.Run(run_server)
}
