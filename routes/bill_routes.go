package routes

import (
	"handySports/api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/surrealdb/surrealdb.go"
)

func SetupBillRoutes(db *surrealdb.DB, router *gin.RouterGroup) {
	billHandler := handlers.NewBillHandler(db)

	// Bills
	billsRoutes := router.Group("/bills")
	{
		billsRoutes.GET("/", billHandler.GetAll)
		billsRoutes.GET("/:id", billHandler.GetByID)
		billsRoutes.POST("/", billHandler.Create)
		billsRoutes.PUT("/:id", billHandler.Update)
		billsRoutes.DELETE("/:id", billHandler.Delete)
	}
}
