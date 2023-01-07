package web

import (
	"github.com/gin-gonic/gin"
	"idnatiya.com/golang-app/cmd/controllers"
	"idnatiya.com/golang-app/cmd/middleware"
)

func DefineWebRoutes(routes *gin.Engine) {
	rg := routes.Group("/api/1.0")
	{
		auth := rg.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		authenticatedRoutes := rg.Group("")
		authenticatedRoutes.Use(middleware.Authenticated())

		authenticatedRoutes.GET("/auth/refresh-token", controllers.RefreshToken)

		book := authenticatedRoutes.Group("/books")
		{
			book.GET("/", controllers.FindBooks)
			book.GET("/:bookID", controllers.FindBook)
			book.POST("/", controllers.StoreBook)
			book.PUT("/:bookID", controllers.UpdateBook)
			book.DELETE("/:bookID", controllers.DeleteBook)
		}
	}
}
