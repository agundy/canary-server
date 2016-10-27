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

// StoreEventHandler takes a htp request containing JSON encoded Event
// information and attempts to create a new Event in the database with 
// this information
func StoreEventHandler(w http.ResponseWriter, r *http.Request) {
	var incEvent models.Event

	// Obtain Event info from JSON
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&incEvent)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Error decoding JSON"))
		return
	}

	// Attempt to store the Event in the database
	Event, err := models.StoreEvent(incEvent)
	if err != nil {

		return
	}

	// Send an awknowledge response
}