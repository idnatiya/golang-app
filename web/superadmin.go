package web

import (
	"github.com/gin-gonic/gin"
	"idnatiya.com/golang-app/cmd/controllers"
	"idnatiya.com/golang-app/cmd/middleware"
)

func DefineSuperadminRoutes(routes *gin.Engine) {
	adminRoute := routes.Group("/api/1.0/admin")
	adminRoute.Use(middleware.Authenticated())

	permissionRouteGroup := adminRoute.Group("/permission")
	{
		permissionRouteGroup.GET("/", controllers.ListPermission)
		permissionRouteGroup.POST("/", controllers.CreatePermission)
		permissionRouteGroup.PUT("/:permissionID", controllers.UpdatePermission)
		permissionRouteGroup.DELETE("/:permissionID", controllers.DeletePermission)
	}

	roleRouteGroup := adminRoute.Group("role")
	{
		roleRouteGroup.GET("/", controllers.ListRole)
		roleRouteGroup.POST("/", controllers.CreateRole)
	}
}
