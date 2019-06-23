package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProfileToJSON(t *testing.T) {
	pro := new(Profile)
	pro.Id = "12345"
	pro.Email = "test@something.com"

	expected := `{"id":"12345","username":"","email":"test@something.com","info":"","items":null,"LastUpdated":"0001-01-01T00:00:00Z"}`
	proJSON, err := pro.ToJSON()
	assert.Nil(t, err)
	assert.Equal(t, expected, string(proJSON))
}
