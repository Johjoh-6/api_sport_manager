package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/surrealdb/surrealdb.go"
)

// AuthHandler is a struct that holds the database connection.
type AuthHandler struct {
	DB        *surrealdb.DB
	Namespace string
	Database  string
	Scope     string
}

type AuthRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthHandler(db *surrealdb.DB) *AuthHandler {
	namespace := getEnv("DB_NAMESPACE")
	database := getEnv("DB_COLLECTION")
	scope := getEnv("USER_SCOPE")
	return &AuthHandler{
		DB:        db,
		Namespace: namespace,
		Database:  database,
		Scope:     scope,
	}
}

// start registration for passkey auth.
func (a *AuthHandler) Registration(c *gin.Context) {
	var registerRequestBody AuthRequestBody
	err := c.BindJSON(&registerRequestBody)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error parsing json: %s", err.Error()))
		return
	}
	user, err := a.DB.Signup(map[string]string{
		"NS":       a.Namespace,
		"DB":       a.Database,
		"SC":       a.Scope,
		"email":    registerRequestBody.Email,
		"password": registerRequestBody.Password,
	})
	if err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error register in: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": user,
	})

}

func (a *AuthHandler) Login(c *gin.Context) {
	var loginRequestBody AuthRequestBody
	err := c.BindJSON(&loginRequestBody)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error parsing json: %s", err.Error()))
		return
	}

	log.Printf("Logging in user: %s", loginRequestBody.Email)
	user, err := a.DB.Signin(map[string]string{
		"NS":       a.Namespace,
		"DB":       a.Database,
		"SC":       a.Scope,
		"email":    loginRequestBody.Email,
		"password": loginRequestBody.Password,
	})
	if err != nil {
		log.Printf(err.Error())
		HandleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": user,
	})
}

func (a *AuthHandler) Authenticate(c *gin.Context) {
	auth := a.GetToken(c)
	_, err := a.DB.Authenticate(auth)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"connected": true,
	})
}

func (a *AuthHandler) Logout(c *gin.Context) {
	_, err := a.DB.Invalidate()
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (a *AuthHandler) GetInfo(c *gin.Context) {
	sql := "RETURN $auth;"
	res, err := a.DB.Query(sql, nil)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, err)
		return
	}

	log.Printf("Query result: %s", res)
	// UnmarshalRaw for raw query results
	user, err := surrealdb.SmartUnmarshal[string](res, nil)
	if err != nil {
		HandleError(c, http.StatusInternalServerError, fmt.Errorf("Error while unmarshalling records: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"connected": user,
	})
}

func (a *AuthHandler) GetToken(c *gin.Context) string {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		return ""
	}
	// remove the "Bearer " prefix
	auth = auth[7:]
	return auth
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return ""
}
