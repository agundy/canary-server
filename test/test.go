package test

import (
	"github.com/agundy/canary-server/database"
	"github.com/agundy/canary-server/router"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
)

var (
	dbName     = "canary-test"
	testRouter *mux.Router
	url        string
	client     http.Client
)

func init() {
	database.DB = database.InitDB(dbName)

	testRouter = router.NewRouter()

	server := httptest.NewServer(testRouter)

	url = server.URL + "/api/"
}
