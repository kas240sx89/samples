package models

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestProfileToJSON(t *testing.T) {
	pro := new(Profile)
	pro.Id = "12345"
	pro.Email = "test@something.com"

	proJSON, err := pro.ToJSON()
	assert.Nil(t, err)
	assert.Equal(t, "{\"id\":\"12345\",\"email\":\"test@something.com\",\"LastUpdated\":\"0001-01-01T00:00:00Z\"}", string(proJSON))
}