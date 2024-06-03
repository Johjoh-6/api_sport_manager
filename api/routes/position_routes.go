package routes

import (
	"handySports/api/api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/surrealdb/surrealdb.go"
)

func SetupPositionRoutes(db *surrealdb.DB, router *gin.RouterGroup) {
	positionHandler := handlers.NewPositionHandler(db)

	// Positions
	positionsRoutes := router.Group("/positions")
	{
		positionsRoutes.GET("/", positionHandler.GetAll)
		positionsRoutes.GET("/:id", positionHandler.GetByID)
		positionsRoutes.POST("/", positionHandler.Create)
		positionsRoutes.PUT("/:id", positionHandler.Update)
		positionsRoutes.DELETE("/:id", positionHandler.Delete)
	}
}
