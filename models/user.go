package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name           string `json:"name"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	HashedPassword []byte `json:"hashed_password"`
}

func (u *User) HashPassword(password string) {
	bytePassword := []byte(password)

	hashed, err := bcrypt.GenerateFromPassword(bytePassword, 0)
	if err != nil {
		fmt.Println(err)
	}

	u.HashedPassword = hashed
}

func (u *User) CheckPassword(password string) bool {
	bytePassword := []byte(password)
	// Check a password; err == nil if password is correct
	err := bcrypt.CompareHashAndPassword(u.HashedPassword, bytePassword)
	return (err == nil)

}
