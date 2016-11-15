package models_test

import (
	"testing"

	"github.com/agundy/canary-server/models"
)

func TestProjectGenerateToken(t *testing.T) {
	token := models.MakeToken()

	if len(token) != 30 {
		t.Errorf("Expected project token to have length 30 got length: %d", len(token))
	}
}
