package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/agundy/canary-server/database"
	"github.com/agundy/canary-server/router"
	"github.com/agundy/canary-server/models"
	"github.com/gorilla/mux"
)

var (
	dbName     = "canary-test"
	testRouter *mux.Router
	url        string
	client     http.Client
	server 
)

func Init() {
	database.DB = database.InitDB(dbName)

	testRouter = router.NewRouter()

	server := httptest.NewServer(testRouter)

	url = server.URL + "/api/"
}

func Teardown() {
	//drop all tables 
	database.DB.DropTable(&models.User{})
	database.DB.DropTable(&models.Project{})
	database.DB.DropTable(&models.Event{})

	//server.Close()
}

func TestMain(m *testing.M) {
	Init()
	results := m.Run()
	Teardown()
	os.Exit(results)
}

