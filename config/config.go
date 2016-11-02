package config

import (
	"log"
	"os"
)

var (
	ApiSecret    = "SECRET"
	DatabaseName = "canary.db"
)

type key int

const RequestUser key = 0

// loadEnv gets secrets and other variables from the environment
func LoadEnv() {
	log.Println("Loading Environment")
	databaseName := os.Getenv("CANARY_DATABASE_NAME")
	if databaseName != "" {
		DatabaseName = databaseName
	}
	envApiSecret := os.Getenv("CANARY_API_SECRET")
	if envApiSecret != "" {
		ApiSecret = envApiSecret
	}
}
