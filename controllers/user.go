package controllers

import (
	"encoding/json"
	"net/http"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	type signupUser struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var user signupUser

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&user)

	if err != nil {
		w.Write([]byte("error"))
		return
	}

	w.Write([]byte("error"))
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
