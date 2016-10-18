package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/agundy/canary-server/models"
	"github.com/agundy/canary-server/router"
)

func setupDatabase(db *gorm.DB) {
	db.CreateTable(&models.User{})
}

func main() {
	log.Println("Staring canary-server")
	router := router.NewRouter()

	log.Println("Connecting to the Database")

	db, err := gorm.Open("postgres", "user=gorm dbname=canary.db sslmode=disable password=mypassword")
	if err != nil {
		log.Fatal("Error Connecting to Database: ", err)
	}
	defer db.Close()

	setupDatabase(db)

	log.Println("Listening on 0.0.0.0:9090")
	err = http.ListenAndServe("0.0.0.0:9090", router)

	if err != nil {
		log.Fatal("Failed to start serve: %v", err)
	}
}
