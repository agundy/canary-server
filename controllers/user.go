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

	// Hanle error decoding JSON
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("error"))
		return
	}

	user, err := models.CreateUser(&userSignup)

	log.Println("HIT")

	if err != nil {
		log.Println(err)

		w.WriteHeader(500)
		w.Write([]byte("Error creating user"))
		return
	} else {
		log.Println("Created User: ", user.Email)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("success"))
		return
	}

	return
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	type loginUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var user loginUser

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&user)

	if err != nil {
		w.Write([]byte("error"))
		return
	}

	w.Write([]byte("error"))
	return
}
