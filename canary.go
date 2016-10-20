package main

import (
	"log"
	"net/http"

	"github.com/agundy/canary-server/database"
	"github.com/agundy/canary-server/router"
)

func main() {
	log.Println("Connecting to the Database")

	var databaseName = "canary.db"
	database.DB = database.InitDB(databaseName)

	log.Println("Staring canary-server")
	router := router.NewRouter()
	log.Println("Listening on 0.0.0.0:9090")
	err := http.ListenAndServe("0.0.0.0:9090", router)

	if err != nil {
		log.Fatal("Failed to start serve: %v", err)
	}
}