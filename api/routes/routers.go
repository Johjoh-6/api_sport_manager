package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/surrealdb/surrealdb.go"
)

func SetupRoutes(db *surrealdb.DB, baseUrl string, version string) *gin.Engine {

	// Gin router
	r := gin.Default()

	// DEFINE the routes general group here and his version.
	router := r.Group(baseUrl + "/" + version)
	{
		// Users
		SetupUsersRoutes(db, router)
		SetupPlayerRoutes(db, router)
		SetupEventRoutes(db, router)
		SetupEventTypeRoutes(db, router)
		SetupMatchHistoryRoutes(db, router)
		SetupTeamRoutes(db, router)
		SetupSportRoutes(db, router)
		SetupPositionRoutes(db, router)
		SetupBillRoutes(db, router)
	}
	return r
}
