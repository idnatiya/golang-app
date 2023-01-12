package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"idnatiya.com/golang-app/cmd/models"
	"idnatiya.com/golang-app/web"
)

func main() {
	// load environtment variable
	godotenv.Load()

	route := gin.Default()

	// define super admin routes
	web.DefineSuperadminRoutes(route)
	// define application route
	web.DefineWebRoutes(route)
	// init to connect database
	models.ConnectDatabase()
	// serve application
	route.Use(cors.Default())
	route.Run(":" + os.Getenv("APP_PORT"))
}
