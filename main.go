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
	PersonController      controllers.PersonController
	PersonRouteController routes.PersonRouteController
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

	PersonController = controllers.NewPersonController(initializer.DB)
	PersonRouteController = routes.NewRoutePersonController(PersonController)

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
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClienOrigin}
	corsConfig.AllowCredentials = false

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	PersonRouteController.PersonRoute(router)

	// Default => deal with firewall
	run_server := fmt.Sprintf("127.0.0.1:%s", config.ServerPort)
	log.Fatal(server.Run(run_server))

	// Trial => byPass firewall checking
	// run_server := fmt.Sprintf("127.0.0.1:%d", 8000)
	// server.Run(run_server)
}
