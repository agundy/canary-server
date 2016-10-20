package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

// InitDB intializes the connection to the database and creates tables
func InitDB(databaseName string) *gorm.DB {
	var databaseOptions = "user=gorm dbname=" + databaseName + " sslmode=disable password=mypassword"

	db, err := gorm.Open("postgres", databaseOptions)
	if err != nil {
		log.Fatal("Error Connecting to Database: ", err)
	}
	return db
}

// CloseDB makes sure the database connection is closed
func CloseDB() {
	log.Println("Closing Database")
	DB.Close()
}
