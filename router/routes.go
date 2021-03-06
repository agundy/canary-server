package router

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/dgrijalva/jwt-go.v2"

	"github.com/agundy/canary-server/config"
	"github.com/agundy/canary-server/controllers"
	"github.com/agundy/canary-server/database"
	"github.com/agundy/canary-server/models"
)

// AuthMiddleware
func AuthMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get jwt authorization header
		jwtString := r.Header.Get("Authorization")
		token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.ApiSecret), nil
		})
		if err == nil && token.Valid {
			user := models.User{}
			database.DB.Where("email = ?", token.Claims["email"]).Find(&user)
			context.Set(r, config.RequestUser, user)
			log.Println("Authorization context set")
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./canary-client/client/index.html")
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
		Methods("GET").
		Path("/api/user/me").
		Handler(AuthMiddleware(controllers.MeHandler))
	router.
		Methods("GET").
		Path("/api/project").
		Handler(AuthMiddleware(controllers.GetProjectsHandler))
	router.
		Methods("POST").
		Path("/api/project").
		Handler(AuthMiddleware(controllers.CreateProjectHandler))
	router.
		Methods("DELETE").
		Path("/api/project/{id:[0-9]+}").
		Handler(AuthMiddleware(controllers.DeleteProjectHandler))
	router.
		Methods("PUT").
		Path("/api/project/{id:[0-9]+}/regentoken").
		Handler(AuthMiddleware(controllers.RegenerateHandler))
	router.
		Methods("POST").
		Path("/api/event").
		HandlerFunc(controllers.StoreEventHandler)
	router.
		Methods("GET").
		Path("/api/project/{id:[0-9]+}/event").
		Handler(AuthMiddleware(controllers.EventHandler))
	router.
		PathPrefix("/app").
		Handler(http.FileServer(http.Dir("./canary-client/client/")))
	router.
		PathPrefix("/static").
		Handler(http.FileServer(http.Dir("./canary-client/client/")))
	router.
		PathPrefix("/").
		HandlerFunc(indexHandler)

	return router
}
