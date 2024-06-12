package handlers

import (
	"handySports/api/models"
	"reflect"

	"github.com/surrealdb/surrealdb.go"
)

// UserHandler handles player-specific requests
type PlayerHandler struct {
	BaseHandler
}

// NewPlayerHandler creates a new PlayerHandler
func NewPlayerHandler(db *surrealdb.DB) *PlayerHandler {
	return &PlayerHandler{
		BaseHandler: BaseHandler{
			DB:         db,
			Collection: "players",
			ModelType:  reflect.TypeOf(models.Player{}),
		},
	}
}
