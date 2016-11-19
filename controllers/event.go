package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/agundy/canary-server/database"
	"github.com/agundy/canary-server/models"
)

// StoreEventHandler takes a http request containing JSON encoded Event
// information and attempts to create a new Event in the database with
// this information
func StoreEventHandler(w http.ResponseWriter, r *http.Request) {
	var incEvent models.Event
	log.Println("Processing event: ", r.Body)
	token := r.Header.Get("CANARY_TOKEN")

	project := models.Project{}
	database.DB.Where("token = ?", token).
		First(&project)
	if project.Name == "" {
		log.Println("Could not find project with token")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Token"))
	}

	// Obtain Event info from JSON
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&incEvent)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding JSON"))
		return
	}

	incEvent.ProjectID = project.ID

	// Attempt to store the Event in the database
	event, err := models.StoreEvent(&incEvent)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error saving event"))
		return
	}

	log.Println("Created new event: ", event.ID)

	eventJson, err := json.Marshal(event)
	if err != nil {
		log.Println("Error encoding event", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send an awknowledge response
	w.WriteHeader(http.StatusOK)
	w.Write(eventJson)
	return
}

func EventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectID, err := strconv.Atoi(vars["id"])
	eventID, err := strconv.Atoi(vars["event_id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad project ID"))
		return
	}

	event := models.GetEvent(projectID, eventID)
	if event != nil {
		rs, marshErr := json.Marshal(event)
		if marshErr != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error Marshalling event to JSON"))
			return
		}

		log.Println("Event retrieved")

		w.WriteHeader(http.StatusOK)
		w.Write(rs)
		return
	}
	return
}
