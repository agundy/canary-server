package models

import (
	"testing"

	//"github.com/agundy/canary-server/models"
)

func TestUserHashPassword(t *testing.T) {
	user := User{}
	var password = "Password"

	user.hashPassword(password)

	if len(user.HashedPassword) == 0 {
		t.Error("hashed password was not set")
	}
}

func TestUserCheckPasswordSuccess(t *testing.T) {
	user := User{}
	var password = "Password"

	user.hashPassword(password)

	if !user.checkPassword(password) {
		t.Error("user password was incorrect")
	}
}

func TestUserCheckPasswordFailure(t *testing.T) {
	user := User{}
	var password = "Password"

	user.hashPassword(password)

	if user.checkPassword(password + " fail") {
		t.Error("Incorrect password was accepted as correct")
	}
}
