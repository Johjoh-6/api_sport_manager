package routes

import (
	"handySports/api/api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/surrealdb/surrealdb.go"
)

func SetupMatchHistoryRoutes(db *surrealdb.DB, router *gin.RouterGroup) {
	matchHistoryHandler := handlers.NewMatchHistoryHandler(db)

	// MatchHistorys
	matchHistorysRoutes := router.Group("/match-history")
	{
		matchHistorysRoutes.GET("/", matchHistoryHandler.GetAll)
		matchHistorysRoutes.GET("/:id", matchHistoryHandler.GetByID)
		matchHistorysRoutes.POST("/", matchHistoryHandler.Create)
		matchHistorysRoutes.PUT("/:id", matchHistoryHandler.Update)
		matchHistorysRoutes.DELETE("/:id", matchHistoryHandler.Delete)
	}
}
