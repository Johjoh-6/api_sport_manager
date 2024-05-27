package handlers

import (
	"github.com/gin-gonic/gin"
)

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Message string `json:"message"`
}

// HandleError is a utility function to handle errors
func HandleError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, ErrorResponse{Message: err.Error()})
	c.Abort()
}
