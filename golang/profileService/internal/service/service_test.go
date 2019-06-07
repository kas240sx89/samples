package service

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/kas240sx89/samples/golang/profileService/internal/db"	
	"github.com/kas240sx89/samples/golang/profileService/internal/models"
)

func TestNewService(t *testing.T) {

	database := db.NewInMemoryDB()
	svc := New(&database)
	assert.NotNil(t, svc)
}

func testService() *Service {
	database := db.NewInMemoryDB()
	return New(&database)
}

func TestHealthCheck(t *testing.T) {
	svc := testService()
	assert.Equal(t, "service is ok", svc.HealthCheck())
}

func TestCreateProfile(t *testing.T) {
	svc := testService()

	pro := new(models.Profile)
	pro.Id = "12345"

	profile, err := svc.CreateProfile(pro)
	assert.Nil(t, err)
	assert.NotNil(t, profile)
}

func TestGetProfileExists(t *testing.T) {
	svc := testService()
	id := "12345"
	pro := new(models.Profile)
	pro.Id = id

	svc.CreateProfile(pro)

	profile, err := svc.GetProfile(id)
	assert.Nil(t, err)
	assert.NotNil(t, profile)
	assert.Equal(t, id, profile.Id)
}

func TestGetProfileNonExist(t *testing.T) {
	svc := testService()
	profile, err := svc.GetProfile("not real")
	assert.NotNil(t, err)
	assert.Nil(t, profile)
}