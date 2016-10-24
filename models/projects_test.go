package models_test

import (
	"testing"

	"github.com/agundy/canary-server/models"
)

func TestProjectGenerateToken(t *testing.T) {
	project := models.Project{}

	project.GenerateToken()

	if len(project.Token) != 30 {
		t.Errorf("Expected project token to have length 30 got length: %d", len(project.Token))
	}
}
