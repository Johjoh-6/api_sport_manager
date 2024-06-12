package handlers

import (
	"handySports/api/models"
	"reflect"

	"github.com/surrealdb/surrealdb.go"
)

// SportHandler handles sport-specific requests
type SportHandler struct {
	BaseHandler
}

// NewSportHandler creates a new SportHandler
func NewSportHandler(db *surrealdb.DB) *SportHandler {
	return &SportHandler{
		BaseHandler: BaseHandler{
			DB:         db,
			Collection: "sports",
			ModelType:  reflect.TypeOf(models.Sport{}),
		},
	}
}
