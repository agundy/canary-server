package models

import (
	// "fmt"
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "math/rand"

	"github.com/agundy/canary-server/database"
)

type Project struct {
	gorm.Model
	Name   string `json:"name"`
	UserID uint   `gorm:"index"`
	Token  string `gorm:"index"`
}

func (p *Project) GenerateToken() {

}

// Something I found on StackOverflow... should look into actual RNG in Go
// func randSeq(n int) string {
// 	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
//     b := make([]rune, n)
//     for i := range b {
//         b[i] = letters[rand.Intn(len(letters))]
//     }
//     return string(b)
// }

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