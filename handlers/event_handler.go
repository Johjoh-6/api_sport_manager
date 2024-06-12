package handlers

import (
	"handySports/api/models"
	"reflect"

	"github.com/surrealdb/surrealdb.go"
)

// EventHandler handles event-specific requests
type EventHandler struct {
	BaseHandler
}

// NewEventHandler creates a new EventHandler
func NewEventHandler(db *surrealdb.DB) *EventHandler {
	return &EventHandler{
		BaseHandler: BaseHandler{
			DB:         db,
			Collection: "events",
			ModelType:  reflect.TypeOf(models.Event{}),
		},
	}
}
