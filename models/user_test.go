package models_test

import (
	"testing"

	"github.com/agundy/canary-server/models"
)

func TestUserHashPassword(t *testing.T) {
	user := models.User{}
	var password = "Password"

	user.HashPassword(password)

	if len(user.HashedPassword) == 0 {
		t.Error("hashed password was not set")
	}
}

func TestUserCheckPasswordSuccess(t *testing.T) {
	user := models.User{}
	var password = "Password"

	user.HashPassword(password)

	if !user.CheckPassword(password) {
		t.Error("user password was incorrect")
	}
}

func TestUserCheckPasswordFailure(t *testing.T) {
	user := models.User{}
	var password = "Password"

	user.HashPassword(password)

	if user.CheckPassword(password + " fail") {
		t.Error("Incorrect password was accepted as correct")
	}
}
