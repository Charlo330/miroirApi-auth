package main

// Importing http and mux package
import (
	"miroirapiauth/controllers"
	"miroirapiauth/initializers"
	"miroirapiauth/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
	initializers.SeedData()
}

func main() {
	r := gin.Default()

	r.POST("/auth/register", controllers.Register)

	r.POST("/auth/login", controllers.Login)

	r.GET("/auth/validate", middleware.RequireAuth, controllers.Validate)

	r.GET("/auth/ws", controllers.GetLoginId)

	r.POST("/auth/desktopLogin", middleware.RequireAuth, controllers.DesktopLogin)

	r.Run(":9888")
}
