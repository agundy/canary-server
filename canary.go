package main

import (
	"log"
	"net/http"

	"github.com/agundy/canary-server/database"
	"github.com/agundy/canary-server/models"
	"github.com/agundy/canary-server/router"
)

var (
	ApiSecret = "SECRET"
)

func initDatabase() {
	var databaseName = "canary.db"
	database.DB = database.InitDB(databaseName)

	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Project{})
	database.DB.AutoMigrate(&models.Event{})
}

// loadEnv gets secrets and other variables from the environment
func loadEnv() {
	envApiSecret := os.Getenv("CANARY_API_SECRET")
	if len(envApiSecret) > 0 {
		ApiSecret = envApiSecret
	}
}

func main() {
	log.Println("Connecting to the Database")
	initDatabase()

	log.Println("Staring canary-server")
	router := router.NewRouter()

	log.Println("Listening on 0.0.0.0:9090")
	err := http.ListenAndServe("0.0.0.0:9090", router)

	if err != nil {
		log.Fatal("Failed to start serve: %v", err)
	}
}
