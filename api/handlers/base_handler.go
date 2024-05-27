package handlers

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/surrealdb/surrealdb.go"
)

// BaseHandler is a generic handler for common CRUD operations
type BaseHandler struct {
	DB         *surrealdb.DB
	Collection string
	ModelType  reflect.Type
}

type PaginationMeta struct {
	TotalItems   int  `json:"total_items"`
	ItemsPerPage int  `json:"items_per_page"`
	CurrentPage  int  `json:"current_page"`
	TotalPages   int  `json:"total_pages"`
	HasNextPage  bool `json:"has_next_page"`
	HasPrevPage  bool `json:"has_prev_page"`
}

// GetAll retrieves all records from the collection
func (h *BaseHandler) GetAll(c *gin.Context) {

	// records, err := h.DB.Select(h.Collection)
	query, err := h.GetQueryAllRecords(c)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error while selecting records: %s", err.Error()))
		return
	}
	records, err := h.DB.Query(query, h.ModelType)
	// TODO remove this
	log.Print(records)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error while selecting records: %s", err.Error()))
		return
	}
	// Create a slice of the appropriate type dynamically
	slice := reflect.New(reflect.SliceOf(h.ModelType)).Interface()
	// UnmarshalRaw for raw query results
	ok, err := surrealdb.UnmarshalRaw(records, slice)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error while unmarshalling records: %s", err.Error()))
		return
	}

	// Check if unmarshalling was successful
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"data": []string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": slice,
	})
}

// GetByID retrieves a single record by its ID
func (h *BaseHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	// Validate the ID
	if !h.CheckId(id) {
		HandleError(c, http.StatusBadRequest, fmt.Errorf("Invalid ID format"))
		return
	}
	query, err := h.GetQueryByID(c)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error while getting record: %s", err.Error()))
		return
	}

	record, err := h.DB.Query(query, h.ModelType)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error while selecting record: %s", err.Error()))
		return
	}
	model := reflect.New(reflect.SliceOf(h.ModelType)).Interface()

	ok, err := surrealdb.UnmarshalRaw(record, model)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error while unmarshalling record: %s", err.Error()))
		return
	}

	if !ok {
		HandleError(c, http.StatusNotFound, fmt.Errorf("Record not found"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": model,
	})
}

// Create creates a new record
func (h *BaseHandler) Create(c *gin.Context) {
	model := reflect.New(h.ModelType).Interface()
	if err := c.ShouldBindJSON(model); err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error while biding json: %s", err.Error()))
		return
	}
	// TODO remove this
	log.Print(model)
	_, err := h.DB.Create(h.Collection, model)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error while creating record: %s", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": model,
	})
}

// Update updates an existing record by its ID
func (h *BaseHandler) Update(c *gin.Context) {
	id := c.Param("id")
	// Validate the ID
	if !h.CheckId(id) {
		HandleError(c, http.StatusBadRequest, fmt.Errorf("Invalid ID format"))
		return
	}
	model := reflect.New(h.ModelType).Interface()
	if err := c.ShouldBindJSON(model); err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error while binding json: %s", err.Error()))
		return
	}

	_, err := h.DB.Update(id, model)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error while updatung record: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": model,
	})
}

// Delete deletes a record by its ID
func (h *BaseHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	_, err := h.DB.Delete(id)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error while deleting record: %s", err.Error()))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": fmt.Sprintf("Record %s deleted", id),
	})
}

// Function to get the all the query and prepare the select query
func (h *BaseHandler) GetQueryAllRecords(c *gin.Context) (string, error) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "20")
	sort := c.DefaultQuery("sort", "+id")
	filter := c.Query("filter")
	fields := c.DefaultQuery("fields", "*")
	allStr := c.DefaultQuery("all", "false")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return "", err
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return "", err
	}
	all, err := strconv.ParseBool(allStr)
	if err != nil {
		return "", err
	}
	filter = strings.ReplaceAll(filter, "(", "")
	filter = strings.ReplaceAll(filter, ")", "")

	// parse sort, filter, fields
	sortFields := strings.Split(sort, ",")

	// prepare sort fields
	for i, field := range sortFields {
		if strings.HasPrefix(field, "+") {
			sortFields[i] = fmt.Sprintf("%s ASC", field[1:])
		} else if strings.HasPrefix(field, "-") {
			sortFields[i] = fmt.Sprintf("%s DESC", field[1:])
		} else {
			sortFields[i] = fmt.Sprintf("%s ASC", field)
		}
	}

	sortFieldsQuery := strings.Join(sortFields, ", ")

	// prepare filter fields

	query := fmt.Sprintf("SELECT %s FROM %s", fields, h.Collection)

	if len(sortFields) > 0 {
		query += fmt.Sprintf(" ORDER BY %s", sortFieldsQuery)
	}

	if filter != "" {
		query += fmt.Sprintf(" WHERE %s", filter)
	}

	if all == true {
		return query, nil
	}

	query += fmt.Sprintf(" LIMIT %d START %d", limit, (page-1)*limit)
	log.Print(query)
	return query, nil
}

func (h *BaseHandler) GetQueryByID(c *gin.Context) (string, error) {
	id := c.Param("id")
	fields := c.DefaultQuery("fields", "*")
	// In SurrealDB the id is the collection name is if the begining of the id.
	query := fmt.Sprintf("SELECT %s FROM %s ", fields, id)
	return query, nil
}

// Ensure the id contains the collection name as the begining
func (h *BaseHandler) CheckId(id string) bool {
	return strings.HasPrefix(id, h.Collection)
}
