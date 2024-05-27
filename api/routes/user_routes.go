package routes

import (
	"handySports/api/api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/surrealdb/surrealdb.go"
)

func SetupUsersRoutes(db *surrealdb.DB, router *gin.RouterGroup) {
	userHandler := handlers.NewUserHandler(db)

	// Users
	usersRoutes := router.Group("/users")
	{
		usersRoutes.GET("/", userHandler.GetAll)
		usersRoutes.GET("/:id", userHandler.GetByID)
		usersRoutes.POST("/", userHandler.Create)
		usersRoutes.PUT("/:id", userHandler.Update)
		usersRoutes.DELETE("/:id", userHandler.Delete)
	}
}
