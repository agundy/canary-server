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
	gorm.Model
	Host         string    `json:"host"`
	Code         int       `json:"code"`
	Duration     int       `json:"duration"`
	Endpoint     string    `json:"endpont"`
	ProjectID    int       `json:"project_id"`
	ProjectToken string    `json:"token"`
	Timestamp    time.Time `json:"timestamp"`
}

func StoreEvent(e *Event) (newEvent *Event, err error) {
	if e.Host == "" || e.ProjectToken == "" {
		log.Println("Event must contain Host" +
			" and Endpoint info as well as an API Token")
	}

	var targetProject Project
	database.DB.Where("token = ?", e.ProjectToken).First(&targetProject)
	if targetProject.Name != "" {
		log.Println("No project found with token: %s", e.ProjectToken)
		return nil, errors.New("No project found with provided token")
	}

	newEvent = &Event{Host: e.Host, Code: e.Code, Duration: e.Duration,
		Endpoint: e.Endpoint, ProjectID: e.ProjectID,
		ProjectToken: e.ProjectToken, Timestamp: e.Timestamp}

	database.DB.Create(&newEvent)

	return newEvent, err
}
