package handlers

import (
	"handySports/api/models"
	"reflect"

	"github.com/surrealdb/surrealdb.go"
)

// EventTypeHandler handles event_type-specific requests
type EventTypeHandler struct {
	BaseHandler
}

// NewEventTypeHandler creates a new EventTypeHandler
func NewEventTypeHandler(db *surrealdb.DB) *EventTypeHandler {
	return &EventTypeHandler{
		BaseHandler: BaseHandler{
			DB:         db,
			Collection: "event_type",
			ModelType:  reflect.TypeOf(models.EventType{}),
		},
	}
}
