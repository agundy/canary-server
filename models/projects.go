package models

import (
	"errors"
	"log"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"math/rand"
	"time"

	"github.com/agundy/canary-server/database"
)

type Project struct {
	ID     uint   `json:"id";gorm:"primary_key";`
	Name   string `json:"name"`
	UserID uint   `json:"user_id";gorm:"index"`
	Token  string `gorm:"index"`
}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func MakeToken() (token string) {
	result := make([]byte, 30)
	for i := 0; i < 30; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

// GenerateToken sets a new token for a project by randomly generating a 30
// character alphanumeric sequence
func (p *Project) GenerateToken() {
	// Use seed based on time and projectID
	rand.Seed(time.Now().UTC().UnixNano() + int64(p.UserID))
	var queryProject Project
	isUsed := true
	var result string

	for isUsed {
		result = MakeToken()
		database.DB.Where("token = ?", result).First(&queryProject)
		if queryProject.Name == "" {
			isUsed = false
		}
	}

	// Run with it
	p.Token = string(result)
}

// CreateProject takes a project object, sets the proper fields then
// stores it in the database
func CreateProject(p *Project) (newProject *Project, err error) {
	var queryProject Project

	// Check if project name is blank
	if p.Name == "" {
		log.Println("Project name can't be blank")
		return nil, errors.New("Project name can't be blank")
	}

	// Check if a project with given name is already in database
	database.DB.Where(&Project{Name: p.Name, UserID: p.UserID}).First(&queryProject)
	if queryProject.Name != "" {
		log.Println("Found project: ", queryProject)
		return nil, errors.New("Project already exists")
	}

	// Create a new Project object and generate its API token
	newProject = &Project{Name: p.Name, UserID: p.UserID}
	newProject.GenerateToken()

	//Store the new project in the database
	database.DB.Create(&newProject)

	return newProject, err
}

// DeleteProject takes a project ID and attempts to delete the corresponding
// project from the database
func DeleteProject(id int, userID uint) (result string, err error) {

	// Check that the project is in the database
	project := Project{}
	database.DB.Where("id = ?", id).Find(&project)
	if project.Name == "" {
		log.Println("Project not found")
		return "ERROR", errors.New("Project not found")
	}

	// Make sure the project belongs to the user signed in
	if project.UserID != userID {
		log.Println("Attempt to delete project not belonging to user in session")
		return "ERROR", errors.New("Attempt to delete project not belonging to user in session")
	}
	// Delete the project
	database.DB.Where("id = ?", id).Delete(Project{})
	database.DB.Where("project_id = ?", id).Delete(Event{})
	rs := string("Project ID=") + strconv.Itoa(id) + string(" deleted")
	return rs, err

}
