package main

import (
	"handySports/api/api/routes"
	"handySports/api/database"

	"log"
)

func init() {
	// Load environment variables from .env file
	loadEnv()

	// get the variable for the database connection
	connectString := GetEnv("DB_CONNECT_STRING", "ws://localhost:8000/rpc")
	username := GetEnv("DB_USERNAME", "root")
	password := GetEnv("DB_PASSWORD", "root")
	namespace := GetEnv("DB_NAMESPACE", "default")
	collection := GetEnv("DB_COLLECTION", "default")

	// Connect to the database
	database.Connect(connectString, username, password, namespace, collection)
}

func main() {
	log.Printf("Woohooo")
	// Load environment variables

	fullUrl := GetEnv("API_URL", "http://localhost:8080/api/v1")
	port := GetEnv("API_PORT", "8080")
	basePath := GetEnv("API_BASE_PATH", "/api")
	apiVersion := GetEnv("API_VERSION", "v1")

	router := routes.SetupRoutes(database.DB, basePath, apiVersion)
	// router.Use(middleware.CORSMiddleware())
	//
	// Start the server
	log.Printf("API version %s", apiVersion)
	log.Printf("Starting server on port %s", port)
	log.Printf("API URL : %s", fullUrl)
	log.Fatal(router.Run(":" + port))

	defer database.DB.Close()
}
