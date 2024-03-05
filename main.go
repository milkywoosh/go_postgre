package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/milkyway/gin_beginer/initializer"
)

var (
	server *gin.Engine
	// KENAPA HARUS PAKE POINTER TYPE????
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
	server = gin.Default()

}

func main() {
	fmt.Println("test")

	_, err := initializer.LoadConfig(".") // why 2x init ?
	if err != nil {
		log.Fatal("could not load environment variable", err)
	}

	// corsConfig := cors.DefaultConfig()

	// trial
	corsConfig := cors.Default()

	// port 8000 apa ? 3000 apa ?
	// corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClienOrigin}
	// corsConfig.AllowCredentials = true // kalo false??
	// corsConfig.AllowCredentials = false // kalo false??

	// Default
	// server.Use(cors.New(corsConfig))
	// Trial
	server.Use(corsConfig)

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Gin"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	// Default => deal with firewall
	// log.Fatal(server.Run(":" + config.ServerPort))
	// Trial => byPass firewall checking
	run_server := fmt.Sprintf("127.0.0.1:%d", 8000)
	server.Run(run_server)
}
