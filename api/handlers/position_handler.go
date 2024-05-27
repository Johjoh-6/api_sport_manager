package handlers

import (
	"handySports/api/models"
	"reflect"

	"github.com/surrealdb/surrealdb.go"
)

// PositionHandler handles position-specific requests
type PositionHandler struct {
	BaseHandler
}

// NewPositionHandler creates a new PositionHandler
func NewPositionHandler(db *surrealdb.DB) *PositionHandler {
	return &PositionHandler{
		BaseHandler: BaseHandler{
			DB:         db,
			Collection: "positions",
			ModelType:  reflect.TypeOf(models.Position{}),
		},
	}
}
