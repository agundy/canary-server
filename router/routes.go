package router

import (
	"github.com/gorilla/mux"

	"github.com/agundy/canary-server/controllers"
)

// The router is responsible for maintianing API endpoints and passing
// off incoming HTTP requests to their appropriate handler functions.
// NewRouter creates such a router and adds appropriate endpoints.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.
		Methods("POST").
		Path("/api/signup").
		HandlerFunc(controllers.SignUpHandler)
	router.
		Methods("POST").
		Path("/api/login").
		HandlerFunc(controllers.LoginHandler)
	router.
		Methods("POST").
		Path("/api/project").
		HandlerFunc(controllers.CreateProjectHandler)
	router.
		Methods("DELETE").
		Path("/api/project/{id:[0-9]+}").
		HandlerFunc(controllers.DeleteProjectHandler)
	router.
		Methods("PUT").
		Path("/api/project/{id:[0-9]+}/regentoken").
		HandlerFunc(controllers.RegenerateHandler)
	router.
		Methods("PUT").
		Path("/api/project/{id:[0-9]+}/storeevent").
		HandlerFunc(controllers.StoreEventHandler)
	router.
		Methods("POST").
		Path("/api/project/{id:[0-9]+}/event").
		HandlerFunc(controllers.EventHandler)

	return router
}
