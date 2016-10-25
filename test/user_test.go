package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/agundy/canary-server/models"
)

type SignupUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func addUser(userInfo *SignupUser) (resp *http.Response, err error) {
	mJson, _ := json.Marshal(userInfo)
	req, err := http.NewRequest("POST", url+"signup", bytes.NewReader(mJson))
	req.Header.Set("Content-Type", "application/json")
	return client.Do(req)
}

func TestAddUser(t *testing.T) {
	var user models.User

	signupInfo := SignupUser{
		Name:     "Test User",
		Email:    "test@user.com",
		Password: "password",
	}

	t.Log("Server URL:", url)

	res, err := addUser(&signupInfo)

	if err != nil {
		t.Errorf("Error adding user: %s", err)
		t.FailNow()
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Error: expecte %d, received %d", http.StatusCreated, res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &user)

	if err != nil {
		t.Errorf("Error unmarshalling JSON into user. %s %s", body, err)
	}
}
