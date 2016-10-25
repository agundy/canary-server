package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/agundy/canary-server/models"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var userSignup models.UserSignup

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&userSignup)

	// Handle error decoding JSON
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("error"))
		return
	}

	user, err := models.CreateUser(&userSignup)

	if err != nil {
		log.Println(err)

		w.WriteHeader(500)
		w.Write([]byte("Error creating user"))
		return
	} else {
		log.Println("Created User: ", user.Email)

		jsonUser, err := json.Marshal(user)
		if err != nil {
			log.Println(err)
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonUser)
		return
	}

	w.WriteHeader(500)
	w.Write([]byte("error"))
	return
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginUser models.UserSignup

	// Attempt to read login object
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&loginUser)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("error"))
		return
	}

	log.Println("Login User: ", loginUser)
	user, err := models.LoginUser(&loginUser)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("error"))
		return
	}

	log.Println("Logged in User: ", user.Email)
	w.Write([]byte("Logged In"))
	return
}
