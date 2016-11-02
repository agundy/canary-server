package models

import (
	"errors"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/dgrijalva/jwt-go.v2"

	"github.com/agundy/canary-server/config"
	"github.com/agundy/canary-server/database"
)

type User struct {
	gorm.Model
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword []byte `json:"-"`
}

type UserSignup struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// HashPassword takes a new user's password and hashes it, so that it may be
// securely stored in the database
func (u *User) HashPassword(password string) {
	bytePassword := []byte(password)

	hashed, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}

	u.HashedPassword = hashed
}

// CheckPassword takes a submitted password and checks that its hash value
// matches the one stored in the database for the corresponding user
func (u *User) CheckPassword(password string) bool {
	bytePassword := []byte(password)
	// Check a password; err == nil if password is correct
	err := bcrypt.CompareHashAndPassword(u.HashedPassword, bytePassword)
	return (err == nil)
}

func (u *User) CheckAuthToken() (tokenString string) {
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims["email"] = u.Email
	token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the JWT with the server secret
	tokenString, _ = token.SignedString(config.ApiSecret)

	return tokenString
}

func (u *User) GetAuthToken() (tokenString string) {
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims["email"] = u.Email
	token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the JWT with the server secret
	tokenString, err := token.SignedString([]byte(config.ApiSecret))

	if err != nil {
		log.Println("Error: ", err)
	}

	return tokenString
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

// LoginUser takes a UserSinup struct with an email and a submitted password,
// then tries to find a correpsonding user and verfy the password
func LoginUser(u *UserSignup) (*User, error) {
	log.Println("Login attempt for user: ", u.Email)

	// try to find a user with the email given in the database
	user := User{}
	database.DB.Where("email = ?", u.Email).Find(&user)

	if user.Email == "" {
		log.Println("User not found")
		return nil, errors.New("User not found")
	}

	// Verify the password given is correct
	correctPassword := user.CheckPassword(u.Password)

	if correctPassword {
		return &user, nil
	}

	return nil, errors.New("Incorrect password")
}
