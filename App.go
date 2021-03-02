package main

import (
	"log"
	"time"

	"myshop/config"
	"myshop/database"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

const port string = ":8000"

func init() {
	log.Println("Initializing DB")
	database.InitiateDB()
}

func main() {

	// gin.SetMode(gin.ReleaseMode)
	routeEngine := gin.New()

	routeEngine.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
	config.SetupRoutes(routeEngine)
	routeEngine.Run(port)

}
