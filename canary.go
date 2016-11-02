package main

import (
	"log"
	"net/http"

	"github.com/agundy/canary-server/config"
	"github.com/agundy/canary-server/database"
	"github.com/agundy/canary-server/models"
	"github.com/agundy/canary-server/router"
)

func initDatabase() {
	var databaseName = config.DatabaseName
	database.DB = database.InitDB(databaseName)

	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Project{})
	database.DB.AutoMigrate(&models.Event{})
}

func main() {
	config.LoadEnv()

	log.Println("Connecting to the Database:", config.DatabaseName)
	initDatabase()

	log.Println("Staring canary-server")
	router := router.NewRouter()

	log.Println("Listening on 0.0.0.0:9090")
	err := http.ListenAndServe("0.0.0.0:9090", router)

	if err != nil {
		log.Fatal("Failed to start serve: %v", err)
	}
}
