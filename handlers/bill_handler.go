package handlers

import (
	"handySports/api/models"
	"reflect"

	"github.com/surrealdb/surrealdb.go"
)

// BillHandler handles bill-specific requests
type BillHandler struct {
	BaseHandler
}

// NewBillHandler creates a new BillHandler
func NewBillHandler(db *surrealdb.DB) *BillHandler {
	return &BillHandler{
		BaseHandler: BaseHandler{
			DB:         db,
			Collection: "bills",
			ModelType:  reflect.TypeOf(models.Bill{}),
		},
	}
}
