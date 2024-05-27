package database

import (
	"log"

	"github.com/surrealdb/surrealdb.go"
)

var DB *surrealdb.DB

func Connect(connectString string, username string, password string, namespace string, collection string) {
	var err error
	DB, err = surrealdb.New(connectString)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
	_, err = DB.Signin(map[string]string{
		"user": username,
		"pass": password,
	})
	if err != nil {
		log.Fatalf("Error signing in: %s", err)
	}

	if _, err = DB.Use(namespace, collection); err != nil {
		log.Fatalf("Error using database: %s", err)
	}

	log.Printf("Connected to db with namespace %s and collection %s", namespace, collection)
}
