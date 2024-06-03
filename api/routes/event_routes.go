package routes

import (
	"handySports/api/api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/surrealdb/surrealdb.go"
)

func SetupEventRoutes(db *surrealdb.DB, router *gin.RouterGroup) {
	eventHandler := handlers.NewEventHandler(db)

	// Events
	eventsRoutes := router.Group("/events")
	{
		eventsRoutes.GET("/", eventHandler.GetAll)
		eventsRoutes.GET("/:id", eventHandler.GetByID)
		eventsRoutes.POST("/", eventHandler.Create)
		eventsRoutes.PUT("/:id", eventHandler.Update)
		eventsRoutes.DELETE("/:id", eventHandler.Delete)
	}
}
