package routes

import (
	"handySports/api/handlers"
	"handySports/api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/surrealdb/surrealdb.go"
)

func SetupAuthRoutes(db *surrealdb.DB, router *gin.RouterGroup) {
	authHandler := handlers.NewAuthHandler(db)

	// Auths
	authsRoutes := router.Group("/auths")
	{
		authsRoutes.POST("/signin", authHandler.Registration)
		authsRoutes.POST("/login", authHandler.Login)
		authsRoutes.POST("/connect", authHandler.Authenticate)
		authsRoutes.POST("/logout", authHandler.Logout)
		authsRoutes.GET("/", middleware.AuthMiddleware(db), authHandler.GetInfo)

	}
}
