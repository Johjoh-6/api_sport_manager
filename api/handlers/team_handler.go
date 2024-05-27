package handlers

import (
	"handySports/api/models"
	"reflect"

	"github.com/surrealdb/surrealdb.go"
)

// TeamHandler handles team-specific requests
type TeamHandler struct {
	BaseHandler
}

// NewTeamHandler creates a new TeamHandler
func NewTeamHandler(db *surrealdb.DB) *TeamHandler {
	return &TeamHandler{
		BaseHandler: BaseHandler{
			DB:         db,
			Collection: "teams",
			ModelType:  reflect.TypeOf(models.Team{}),
		},
	}
}
