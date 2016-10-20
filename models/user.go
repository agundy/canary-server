package models

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"

	"github.com/agundy/canary-server/database"
)

type User struct {
	gorm.Model
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword []byte `json:"hashed_password"`
}

type UserSignup struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) HashPassword(password string) {
	bytePassword := []byte(password)

	hashed, err := bcrypt.GenerateFromPassword(bytePassword, 0)
	if err != nil {
		log.Println(err)
	}

	u.HashedPassword = hashed
}

func (u *User) CheckPassword(password string) bool {
	bytePassword := []byte(password)
	// Check a password; err == nil if password is correct
	err := bcrypt.CompareHashAndPassword(u.HashedPassword, bytePassword)
	return (err == nil)

}

// CreateUser takes a user object and sets the proper fields
func CreateUser(u *UserSignup) (newUser *User, err error) {
	var queryUser User

	// Check if email is blank or password is blank
	if u.Email == "" {
		log.Println("Email can't be blank")
		return nil, errors.New("Email can't be blank")
	}

	// Check if a user with that email already in database
	database.DB.Where(&User{Email: u.Email}).First(&queryUser)
	// check if query response is empty
	if queryUser.Email != "" {
		log.Println("Found user: ", queryUser)
		return nil, errors.New("User already exists")
	}

	newUser = &User{Name: u.Name, Email: u.Email}

	newUser.HashPassword(u.Password)

	database.DB.Create(&newUser)

	return newUser, err
}
