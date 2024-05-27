package handlers

import (
	"handySports/api/models"
	"reflect"

	"github.com/surrealdb/surrealdb.go"
)

// MatchHistoryHandler handles match_history specific requests
type MatchHistoryHandler struct {
	BaseHandler
}

// NewMatchHistoryHandler creates a new MatchHistoryHandler
func NewMatchHistoryHandler(db *surrealdb.DB) *MatchHistoryHandler {
	return &MatchHistoryHandler{
		BaseHandler: BaseHandler{
			DB:         db,
			Collection: "match_history",
			ModelType:  reflect.TypeOf(models.MatchHistory{}),
		},
	}
}
