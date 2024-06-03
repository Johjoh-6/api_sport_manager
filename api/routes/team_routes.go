package routes

import (
	"handySports/api/api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/surrealdb/surrealdb.go"
)

func SetupTeamRoutes(db *surrealdb.DB, router *gin.RouterGroup) {
	teamHandler := handlers.NewTeamHandler(db)

	// Teams
	teamsRoutes := router.Group("/teams")
	{
		teamsRoutes.GET("/", teamHandler.GetAll)
		teamsRoutes.GET("/:id", teamHandler.GetByID)
		teamsRoutes.POST("/", teamHandler.Create)
		teamsRoutes.PUT("/:id", teamHandler.Update)
		teamsRoutes.DELETE("/:id", teamHandler.Delete)
	}
}
