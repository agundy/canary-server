package models

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"math/rand"
	"time"

	"github.com/agundy/canary-server/database"
)

type Project struct {
	gorm.Model
	Name   string `json:"name"`
	UserID uint   `json:"user_id";gorm:"index"`
	Token  string `gorm:"index"`
}

// GenerateToken sets a new token for a project by randomly generating a 30
// character alphanumeric sequence
func (p *Project) GenerateToken() {
	// Use seed based on time and projectID
	rand.Seed(time.Now().UTC().UnixNano() + int64(p.UserID))
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 30)

	// Generate each character
	for i := 0; i < 30; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	// Run with it
	p.Token = string(result)
}

// CreateProject takes a project object and sets the proper fields
func CreateProject(p *Project) (newProject *Project, err error) {
	var queryProject Project

	// Check if project name is blank
	if p.Name == "" {
		log.Println("Project name can't be blank")
		return nil, errors.New("Project name can't be blank")
	}

	// Check if a project with given name is already in database
	database.DB.Where(&Project{Name: p.Name}).First(&queryProject)
	if queryProject.Name != "" {
		log.Println("Found project: ", queryProject)
		return nil, errors.New("Project already exists")
	}

	newProject = &Project{Name: p.Name, UserID: p.UserID}
	newProject.GenerateToken()

	database.DB.Create(&newProject)

	return newProject, err
}
