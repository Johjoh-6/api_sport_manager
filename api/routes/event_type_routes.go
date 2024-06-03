package routes

import (
	"handySports/api/api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/surrealdb/surrealdb.go"
)

func SetupEventTypeRoutes(db *surrealdb.DB, router *gin.RouterGroup) {
	eventTypeHandler := handlers.NewEventTypeHandler(db)

	// EventTypes
	eventTypesRoutes := router.Group("/event-types")
	{
		eventTypesRoutes.GET("/", eventTypeHandler.GetAll)
		eventTypesRoutes.GET("/:id", eventTypeHandler.GetByID)
		eventTypesRoutes.POST("/", eventTypeHandler.Create)
		eventTypesRoutes.PUT("/:id", eventTypeHandler.Update)
		eventTypesRoutes.DELETE("/:id", eventTypeHandler.Delete)
	}
}
