package models

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	project := Project{}

	project.GenerateToken()

	if len(project.Token) != 30 {
		t.Errorf("Expected project token to have length 30 got length: %d", len(project.Token))
	}
}
