package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/agundy/canary-server/models"
)

// StoreEventHandler takes a http request containing JSON encoded Event
// information and attempts to create a new Event in the database with
// this information
func StoreEventHandler(w http.ResponseWriter, r *http.Request) {
	var incEvent models.Event
	log.Println("Processing event: ", r.Body)

	// Obtain Event info from JSON
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&incEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding JSON"))
		return
	}

	// Attempt to store the Event in the database
	event, err := models.StoreEvent(&incEvent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Created new event: ", event.ID)

	// Send an awknowledge response
	w.WriteHeader(http.StatusCreated)
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
	log.Println(event)
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