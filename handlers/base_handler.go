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
	pageValue, err := h.GetQueryAllRecords(c)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error while selecting records: %s", err.Error()))
		return
	}

	records, err := h.DB.Query(pageValue.query, h.ModelType)
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
			"meta": PaginationMeta{
				TotalItems:   0,
				ItemsPerPage: 0,
				CurrentPage:  0,
				TotalPages:   0,
				HasNextPage:  false,
				HasPrevPage:  false,
			},
		})
		return
	}

	// Get the total number of items
	sqlCount := fmt.Sprintf("SELECT * FROM count((SELECT * FROM %s));", h.Collection)
	countRecords, err := h.DB.Query(sqlCount, nil)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error while counting records: %s", err.Error()))
		return
	}
	count, err := surrealdb.SmartUnmarshal[[]int](countRecords, nil)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error while unmarshalling count: %s", err.Error()))
		return
	}
	// return an array of 1 int [1] extracte the value
	countValue := reflect.ValueOf(count[0]).Interface().(int)

	c.JSON(http.StatusOK, gin.H{
		"data": slice,
		"meta": PaginationMeta{
			TotalItems: countValue,
			// TotalItems:   0,
			ItemsPerPage: pageValue.limit,
			CurrentPage:  pageValue.page,
			TotalPages:   int(countValue / pageValue.limit),
			HasNextPage:  pageValue.page < int(countValue/pageValue.limit),
			HasPrevPage:  pageValue.page > 1,
		},
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
	result, err := h.DB.Create(h.Collection, model)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error while creating record: %s", err.Error()))
		return
	}
	// Type assertion to convert result to a slice of interfaces
	records, ok := result.([]interface{})
	if !ok {
		HandleError(c, http.StatusBadRequest, fmt.Errorf("Error while creating record"))
		return
	}

	// Check if the result contains any data
	if len(records) == 0 {
		HandleError(c, http.StatusBadRequest, fmt.Errorf("Error while creating record"))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Record created",
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

	updates := make(map[string]interface{})
	if err := c.ShouldBindJSON(&updates); err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error while binding json: %s", err.Error()))
		return
	}
	// Check if the record exists
	exist, err := h.DB.Select(id)
	if err != nil {
		HandleError(c, http.StatusNotFound, fmt.Errorf("Record not found"))
		return
	}

	if exist == nil {
		HandleError(c, http.StatusNotFound, fmt.Errorf("Record not found"))
		return
	}

	// Only update the fields that are present in the JSON request
	var query string = fmt.Sprintf("UPDATE %s SET ", id)
	for key, value := range updates {
		log.Printf("k: %s , val : %s", key, value)
		query += fmt.Sprintf("%s = $%s, ", key, key)
	}
	// Remove the trailing comma
	query = query[:len(query)-2]
	query = fmt.Sprintf("%s;", query)
	log.Print(query)
	up, err := h.DB.Query(query, updates)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error while updating record: %s", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": up,
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

	c.JSON(http.StatusNoContent, nil)
}

type PageValue struct {
	query string
	page  int
	limit int
}

// Function to get the all the query and prepare the select query
func (h *BaseHandler) GetQueryAllRecords(c *gin.Context) (PageValue, error) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "20")
	sort := c.DefaultQuery("sort", "+id")
	filter := c.Query("filter")
	fields := c.DefaultQuery("fields", "*")
	allStr := c.DefaultQuery("all", "false")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return PageValue{
			query: "",
			page:  0,
			limit: 0,
		}, err
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return PageValue{
			query: "",
			page:  0,
			limit: 0,
		}, err
	}
	all, err := strconv.ParseBool(allStr)
	if err != nil {
		return PageValue{
			query: "",
			page:  0,
			limit: 0,
		}, err
	}
	filter = strings.ReplaceAll(filter, "(", "")
	filter = strings.ReplaceAll(filter, ")", "")

	// parse sort, filter, fields
	sortFields := strings.Split(sort, ",")

	// prepare sort fields
	for i, field := range sortFields {
		if !h.isValidFieldName(field) {
			sortFields = append(sortFields[:i], sortFields[i+1:]...)
			break
		}
		if strings.HasPrefix(field, "+") {
			sortFields[i] = fmt.Sprintf("%s ASC", field[1:])
		} else if strings.HasPrefix(field, "-") {
			sortFields[i] = fmt.Sprintf("%s DESC", field[1:])
		} else {
			sortFields[i] = fmt.Sprintf("%s ASC", field)
		}
	}

	sortFieldsQuery := strings.Join(sortFields, ", ")

	// prepare fields
	if fields != "*" {
		fieldsPrepare := strings.Split(fields, ",")
		var fieldsQuery []string
		idPresent := false
		for _, field := range fieldsPrepare {
			field = strings.TrimSpace(field)
			if h.isValidFieldName(field) {
				fieldsQuery = append(fieldsQuery, field)
				if field == "id" {
					idPresent = true
				}
			}
		}
		fields = strings.Join(fieldsQuery, ", ")
		// add id to fields if it's not already present
		if !idPresent {
			fields = fmt.Sprintf("id, %s", fields)
		}
	}

	query := fmt.Sprintf("SELECT %s FROM %s", fields, h.Collection)

	if len(sortFields) > 0 {
		query += fmt.Sprintf(" ORDER BY %s", sortFieldsQuery)
	}

	if filter != "" {
		query += fmt.Sprintf(" WHERE %s", filter)
	}

	if all == true {
		return PageValue{
			query: query,
			page:  page,
			limit: limit,
		}, nil
	}

	query += fmt.Sprintf(" LIMIT %d START %d ;", limit, (page-1)*limit)
	log.Print(query)
	return PageValue{
		query: query,
		page:  page,
		limit: limit,
	}, nil
}

func (h *BaseHandler) GetQueryByID(c *gin.Context) (string, error) {
	id := c.Param("id")
	fields := c.DefaultQuery("fields", "*")
	// In SurrealDB the id is the collection name is if the begining of the id.
	query := fmt.Sprintf("SELECT %s FROM %s ", fields, id)
	return query, nil
}

func (h *BaseHandler) isValidFieldName(field string) bool {
	// Get the number of fields in the collection
	numFields := h.ModelType.NumField()
	log.Print(numFields)

	// Loop over the fields
	for i := 0; i < numFields; i++ {
		// Get the field name
		fieldName := h.ModelType.Field(i).Tag.Get("json")
		log.Print(fieldName)
		// remove the omitempty tag
		fieldName = strings.Split(fieldName, ",")[0]

		// Check if the field name matches the input field name
		if fieldName == field {
			return true
		}
	}

	// If no matching field name is found, return false
	return false
}

// Ensure the id contains the collection name as the begining
func (h *BaseHandler) CheckId(id string) bool {
	return strings.HasPrefix(id, h.Collection)
}

func (h *BaseHandler) GetFieldNames() []string {
	// Get the number of fields in the collection
	numFields := h.ModelType.NumField()

	// Create a slice to hold the field names
	fieldNames := make([]string, numFields)

	// Loop over the fields
	for i := 0; i < numFields; i++ {
		// Get the field name
		fieldName := h.ModelType.Field(i).Tag.Get("json")
		// remove the omitempty tag
		fieldName = strings.Split(fieldName, ",")[0]
		// Add the field name to the slice
		fieldNames[i] = fieldName
	}

	// Return the field names
	return fieldNames
}
