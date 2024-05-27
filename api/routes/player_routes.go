package routes

import (
	"handySports/api/api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/surrealdb/surrealdb.go"
)

func SetupPlayerRoutes(db *surrealdb.DB, router *gin.RouterGroup) {
	playerHandler := handlers.NewPlayerHandler(db)

	// Users
	playersRoutes := router.Group("/players")
	{
		playersRoutes.GET("/", playerHandler.GetAll)
		playersRoutes.GET("/:id", playerHandler.GetByID)
		playersRoutes.POST("/", playerHandler.Create)
		playersRoutes.PUT("/:id", playerHandler.Update)
		playersRoutes.DELETE("/:id", playerHandler.Delete)
	}
}
