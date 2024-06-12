package handlers

import (
	"handySports/api/models"
	"reflect"

	"github.com/surrealdb/surrealdb.go"
)

// UserHandler handles user-specific requests
type UserHandler struct {
	BaseHandler
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(db *surrealdb.DB) *UserHandler {
	return &UserHandler{
		BaseHandler: BaseHandler{
			DB:         db,
			Collection: "users",
			ModelType:  reflect.TypeOf(models.User{}),
		},
	}
}
