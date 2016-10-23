package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/agundy/canary-server/models"
)

func CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	var projectStruct models.Project

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&projectStruct)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("error"))
		return
	}

	project, err := models.CreateProject(&projectStruct)

	log.Println("PROJECT HIT")

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Error creating project"))
		return
	} else {
		log.Println("Created Project: ", project.Name)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("success"))
		return
	}

	return
}
