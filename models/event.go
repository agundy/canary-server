package models

import (
	"errors"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/agundy/canary-server/database"
)

type Event struct {
	gorm.Model            `json:"-"`
	Host         string    `json:"host"`
	Code         int       `json:"code"`
	Duration     int       `json:"duration"`
	Endpoint     string    `json:"endpont"`
	ProjectID    int       `json:"project_id"`
	ProjectToken string    `json:"token"`
	Timestamp    time.Time `json:"timestamp"`
}

// StoreEvent takes an event object, verifies that the token matches the correct
// ProjectID, then sets the proper fields and stores it in the database
func StoreEvent(e *Event) (newEvent *Event, err error) {
	// Check the event has host, endpoint, and project token
	if e.Host == "" || e.Endpoint == "" || e.ProjectToken == "" {
		log.Println("Event must contain Host" +
			" and Endpoint info as well as an API Token")
		return nil, errors.New("No host, endpoint, or token information")
	}

	// Verify that the provided token matches the specified project
	var targetProject Project
	database.DB.Where("token = ? AND id = ?", e.ProjectToken, e.ProjectID).First(&targetProject)
	log.Println(targetProject.Name, targetProject.ID)
	if targetProject.Name == "" && targetProject.ID == 0 {
		log.Println("No project found with the following ID and Token:", e.ProjectID, e.ProjectToken)
		return nil, errors.New("Project ID and Token do not match")
	}

	// Create the event to be stored
	newEvent = &Event{Host: e.Host, Code: e.Code, Duration: e.Duration,
		Endpoint: e.Endpoint, ProjectID: e.ProjectID,
		ProjectToken: e.ProjectToken, Timestamp: e.Timestamp}

	// Store the new event in the database
	database.DB.Create(&newEvent)

	return newEvent, err
}

func GetEvent(projectID int, eventID int) (e *Event) {
	event := Event{}
	database.DB.Last(&event)
	log.Println(event.Model.ID)
	// Conditions separated for debugging purposes
	if event.Model.ID < 1 {
		return nil
	}
	if int(event.Model.ID) == eventID {
		return nil
	}
	log.Println(event.ProjectID, projectID)
	if event.ProjectID != projectID {
		return nil
	}
	log.Println("HERE")
	return &event
}