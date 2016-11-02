package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"

	"github.com/agundy/canary-server/config"
	"github.com/agundy/canary-server/database"
	"github.com/agundy/canary-server/models"
)

// CreateProjectHandler takes a http reuqest containing JSON encoded project
// information and attempts to create a new project in the database with
// this information
func CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	user = context.Get(r, config.RequestUser).(models.User)

	var projectStruct models.Project

	// Obtain project info from JSON
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&projectStruct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding JSON"))
		return
	}

	projectStruct.UserID = user.ID

	// Attempt to create the project in the database
	project, err := models.CreateProject(&projectStruct)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating project"))
		return
	}

	// Attempt to create JSON encoded project info, then send a response
	rs, err := json.Marshal(project)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Marshalling project info to JSON"))
		return
	}

	log.Println("Created Project: ", project.Name)

	w.WriteHeader(http.StatusCreated)
	w.Write(rs)
	return
}

// DeleteProjectHander takes a http request containing a project ID
// and attempts to remove the corresponding project from the database
func DeleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	user = context.Get(r, config.RequestUser).(models.User)

	// Use mux to obtain the ID as an int
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("Bad project ID"))
		return
	}

	// Attempt to delete the project
	result, err := models.DeleteProject(id, user.ID)

	// Send response with success or failure info
	if err != nil {
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

/// GetProjectHandler return a list of the projects a user has
func GetProjectsHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	user = context.Get(r, config.RequestUser).(models.User)
	var projects []models.Project

	database.DB.Where("user_id = ?", user.ID).Find(&projects)
	jsonProjects, err := json.Marshal(projects)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Marshalling project info to JSON"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonProjects)
	return
}

// RegenerateHandler takes an http request containing a project ID
// and attempts to create a new API token for the project and save the
// changes in the database
func RegenerateHandler(w http.ResponseWriter, r *http.Request) {
	// Use mux to obtain the ID as an int
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("Bad project ID"))
		return
	}

	// Check for the project in the database
	project := models.Project{}
	log.Println(id)
	database.DB.Where("id = ?", id).
		First(&project)
	if project.Name == "" {
		log.Println("Project not found")
		w.WriteHeader(404)
		w.Write([]byte("Project not found"))
		return
	}

	// Generate the new token and save the change to the databse
	project.GenerateToken()
	database.DB.Model(&project).Update("token", project.Token)

	// Send response with success info
	log.Println("TOKEN REGENERATION HIT")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(project.Token))
}
