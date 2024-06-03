package routes

import (
	"handySports/api/api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/surrealdb/surrealdb.go"
)

func SetupSportRoutes(db *surrealdb.DB, router *gin.RouterGroup) {
	sportHandler := handlers.NewSportHandler(db)

	// Sports
	sportsRoutes := router.Group("/sports")
	{
		sportsRoutes.GET("/", sportHandler.GetAll)
		sportsRoutes.GET("/:id", sportHandler.GetByID)
		sportsRoutes.POST("/", sportHandler.Create)
		sportsRoutes.PUT("/:id", sportHandler.Update)
		sportsRoutes.DELETE("/:id", sportHandler.Delete)
	}
}
