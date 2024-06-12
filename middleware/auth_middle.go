package middleware

import (
	"fmt"
	"net/http"

	"handySports/api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/surrealdb/surrealdb.go"
)

// Middleware to authenticate the user
func AuthMiddleware(DB *surrealdb.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := GetToken(c)
		if auth == "" {
			handlers.HandleError(c, http.StatusUnauthorized, fmt.Errorf("Missing authorization token"))
			c.Abort()
			return
		}

		_, err := DB.Authenticate(auth)
		if err != nil {
			handlers.HandleError(c, http.StatusForbidden, fmt.Errorf("Invalid token"))
			c.Abort()
			return
		}

		c.Next()
	}
}
func GetToken(c *gin.Context) string {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		return ""
	}
	// remove the "Bearer " prefix
	auth = auth[7:]
	return auth
}
