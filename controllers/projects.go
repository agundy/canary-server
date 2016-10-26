package controllers

import (
	"encoding/json"
	"log"	
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/agundy/canary-server/models"
	"github.com/agundy/canary-server/database"

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

	// dec := json.NewDecoder(r.Body)
	// err := dec.Decode(&projectStruct)

	// if err != nil {
	// 	w.WriteHeader(400)
	// 	w.Write([]byte("Error decoding JSON"))
	// 	return
	// }

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("Bad project ID"))
		return
	}

	result, err := models.DeleteProject(id)


	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Error deleting project"))
		return
	} else {
		log.Println("PROJECT DELETE HIT")	
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
		return
	}
	
	return
}

func RegenerateHandler(w http.ResponseWriter, r *http.Request) {
	var projectStruct models.Project

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&projectStruct)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Error decoding JSON"))
		return
	}

	project := models.Project{}
	log.Println(projectStruct.ID)
	database.DB.Where("ID = ?", projectStruct.ID).
			   First(&project)
	if project.Name == "" {
		log.Println("Project not found")
		w.WriteHeader(404)
		w.Write([]byte("Project not found"))
		return
	}

	project.GenerateToken()
	database.DB.Update("token", project.Token)

	// if err != {
	// 	//error handling
	// } else {
		log.Println("TOKEN REGENERATION HIT")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(project.Token))
	// 	//write response
	// }
}
