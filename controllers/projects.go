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
		w.Write([]byte("Error decoding JSON"))
		return
	}

	project, err := models.CreateProject(&projectStruct)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Error creating project"))
		return
	} else {
		rs, err := json.Marshal(project)
		if err != nil {
			log.Print(err)
			w.WriteHeader(500)
			w.Write([]byte("Error creating project"))
			return
		}

		log.Println("Created Project: ", project.Name)

		w.WriteHeader(http.StatusCreated)
		m := []byte("Created project: ")
		m = append(m, rs...)
		w.Write(m)
		return
	}

	return
}

func DeleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	var projectStruct models.Project

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&projectStruct)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Error decoding JSON"))
		return
	}

	result, err := models.DeleteProject(&projectStruct)

	log.Println("PROJECT DELETE HIT")

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Error deleting project"))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
		return
	}
	
	return
}
