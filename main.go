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
	"github.com/milkyway/gin_beginer/repositories"
	"github.com/milkyway/gin_beginer/routes"
	"github.com/milkyway/gin_beginer/services"
)

var (
	server *gin.Engine
	// KENAPA HARUS PAKE POINTER TYPE????

	// repos
	BooksRepo     *repositories.BooksRepo
	CustomersRepo *repositories.CustomersRepo

	// service
	OrdersService *services.OrdersService

	// controller
	UsersController      *controllers.UsersController
	UsersRouteController *routes.UsersRouteController

	UserRolesController      *controllers.UserRolesController
	RoleUsersRouteController *routes.RoleUsersRouteController

	BooksController      *controllers.BooksController
	BooksRouteController *routes.BooksRouteController
	// purchase manager

	CustomersController      *controllers.CustomersController
	CustomersRouteController *routes.CustomersRouteController

	OrdersController      *controllers.OrdersController
	OrdersRouteController *routes.OrdersRouteController
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

	// TEST : pindah dari init() ke main()

}

func main() {

	// repo ============
	BooksRepo = repositories.NewBooksRepo(initializer.DB)
	CustomersRepo = repositories.NewCustomersRepo(initializer.DB)
	// ============
	OrdersService = services.NewOrdersService(BooksRepo, CustomersRepo)
	OrdersController = controllers.NewOrdersController(OrdersService)
	OrdersRouteController = routes.NewRouteOrdersController(OrdersController)

	UsersController = controllers.NewUsersController(initializer.DB)
	UsersRouteController = routes.NewRouteUsersController(UsersController)

	UserRolesController = controllers.NewUserRolesController(initializer.DB)
	RoleUsersRouteController = routes.NewRouteUserRolesController(UserRolesController)

	BooksController = controllers.NewBooksController(initializer.DB)
	BooksRouteController = routes.NewRouteBooksController(BooksController)

	CustomersController = controllers.NewCustomersController(initializer.DB)
	CustomersRouteController = routes.NewRouteCustomersController(CustomersController)

	server = gin.Default()

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

	router.GET("/check", func(ctx *gin.Context) {
		books_repo := repositories.NewBooksRepo(initializer.DB)

		book_id, err := books_repo.BookIsExistedByID(ctx, 1300)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": err.Error(),
				"data":    book_id,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "message",
			"data":    book_id,
		})

	})

	UsersRouteController.UsersRoute(router)
	RoleUsersRouteController.RoleUsersRoute(router)
	BooksRouteController.BooksRoutes(router)
	CustomersRouteController.CustomersRoutes(router)
	OrdersRouteController.OrdersRoutes(router)

	// Default => deal with firewall
	run_server := fmt.Sprintf("127.0.0.1:%s", config.ServerPort)

	err = server.Run(run_server)

	log.Fatal(err)

	// Trial => byPass firewall checking
	// run_server := fmt.Sprintf("127.0.0.1:%d", 8000)
	// server.Run(run_server)
}
