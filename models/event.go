package models

import (
	"errors"
	"log"
	"time"

	"github.com/agundy/canary-server/database"
)

type Event struct {
	ID        uint      `json:"id";gorm:"primary_key";`
	Host      string    `json:"host"`
	Code      int       `json:"code"`
	Duration  int       `json:"duration"`
	Endpoint  string    `json:"endpoint"`
	ProjectID uint      `json:"project_id"`
	Timestamp time.Time `json:"timestamp"`
}

// StoreEvent takes an event object, verifies that the token matches the correct
// ProjectID, then sets the proper fields and stores it in the database
func StoreEvent(e *Event) (newEvent *Event, err error) {
	// Check the event has host, endpoint, and project token
	if e.Host == "" || e.Endpoint == "" {
		log.Println("Event must contain Host and Endpoint info")
		return nil, errors.New("No host, endpoint, or token information")
	}

	// Verify that the provided token matches the specified project
	var targetProject Project
	database.DB.Where("id = ?", e.ProjectID).First(&targetProject)

	log.Println(targetProject.Name, targetProject.ID)
	if targetProject.Name == "" && targetProject.ID == 0 {
		log.Println("No project found with the following ID and Token:", e.ProjectID)
		return nil, errors.New("Project ID and Token do not match")
	}

	// Create the event to be stored
	newEvent = &Event{Host: e.Host, Code: e.Code, Duration: e.Duration,
		Endpoint: e.Endpoint, ProjectID: e.ProjectID,
		Timestamp: time.Now()}

	// Store the new event in the database
	database.DB.Create(&newEvent)

	return newEvent, err
}

func GetEvent(projectID int, eventID int) (e *Event) {
	event := Event{}
	database.DB.Last(&event)
	// Conditions separated for debugging purposes
	if event.ID < 1 {
		return nil
	}
	if event.ID == uint(eventID) {
		return nil
	}
	if event.ProjectID != uint(projectID) {
		return nil
	}
	return &event
}
