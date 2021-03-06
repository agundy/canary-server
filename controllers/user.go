package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/context"

	"github.com/agundy/canary-server/config"
	"github.com/agundy/canary-server/models"
)

// SignUpHandler takes a http request containing JSON encoded UserSignUp
// informaiton and attempts to create a new User in the database with this
// informaiton
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var userSignup models.UserSignup

	// Obtain User info from JSON
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&userSignup)

	// Handle error decoding JSON
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding JSON"))
		return
	}

	// Attempt to create the user in the database
	user, err := models.CreateUser(&userSignup)

	// Send an appropriate response
	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating user"))
		return
	} else {
		jsonUser, err := json.Marshal(user)
		if err != nil {
			log.Println(err)
		}

		log.Println("Created User: ", user.Email)

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonUser)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("error"))
	return
}

// LoginHandler takes a http request containing JSON encoded UserSignUp
// information and attempts to log the user in with this database
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginUser models.UserSignup

	// Attempt to read login object
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&loginUser)
	if err != nil {
		log.Println("Error reading body:", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding JSON"))
		return
	}

	// Attempt to login with information
	user, err := models.LoginUser(&loginUser)

	if err != nil {
		// User not in database
		if err.Error() == "User not found" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("User does not exist"))
			return
		}
		// Password does not match given user
		if err.Error() == "Incorrect password" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Incorrect password"))
			return
		}
		// Unexpected issue occured
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unexpected error logging in"))
		return
	}

	log.Println("Logged in User: ", user.Email)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"token\": \"" + user.GetAuthToken() + "\"}"))
	return
}

// MeHandler
func MeHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	user = context.Get(r, config.RequestUser).(models.User)

	userJson, err := json.Marshal(user)
	if err != nil {
		log.Println("Error encoding user", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(userJson)
	return
}
