package router

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"

	"github.com/agundy/canary-server/config"
	"github.com/agundy/canary-server/controllers"
	"github.com/agundy/canary-server/database"
	"github.com/agundy/canary-server/models"
)

type key int

const User key = 0

func AuthMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get jwt authorization header
		jwtString := r.Header.Get("Authorization")
		token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.ApiSecret), nil
		})
		if err == nil && token.Valid {
			log.Println(token.Claims)
			user := models.User{}
			database.DB.Where("email = ?", token.Claims["email"]).Find(&user)
			log.Println("Query")
			context.Set(r, User, user.ID)
			log.Println("Set context")
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
	})
}

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
		Handler(AuthMiddleware(controllers.CreateProjectHandler))
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
