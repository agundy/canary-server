package router

import (
	"github.com/gorilla/mux"

	"github.com/agundy/canary-server/controllers"
)

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

	return router
}
