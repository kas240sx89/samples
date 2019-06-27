package service

import (
	"github.com/kas240sx89/samples/profileService/internal/db"
	"github.com/kas240sx89/samples/profileService/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//TODO rewrite tests using DB mocks

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

func TestGetProfileID(t *testing.T) {
	svc := testService()
	id := "12345"
	email := "test@testing.com"
	pro := new(models.Profile)
	pro.Id = id
	pro.Email = email

	svc.CreateProfile(pro)

	profileID, err := svc.GetProfileID(email)
	assert.Nil(t, err)
	assert.Equal(t, id, profileID)

	profileID, err = svc.GetProfileID("fake")
	assert.NotNil(t, err)
	assert.Equal(t, "", profileID)

}

func TestUpdateProfileID(t *testing.T) {
	svc := testService()
	id := "12345"
	email := "test@testing.com"
	pro := new(models.Profile)
	pro.Id = id
	pro.Email = email

	_, err := svc.CreateProfile(pro)
	assert.Nil(t, err)

	pro.LastUpdated = pro.LastUpdated.Add(-1 * time.Minute)
	failedPro, err := svc.UpdateProfile(pro)
	assert.NotNil(t, err)

	failedPro.Info = "test info"
	failedPro.LastUpdated = failedPro.LastUpdated.Add(5 * time.Minute)

	updated, err := svc.UpdateProfile(failedPro)
	assert.Nil(t, err)
	assert.Equal(t, failedPro.Info, updated.Info)
}

func TestDeleteProfile(t *testing.T) {
	svc := testService()
	id := "12345"
	email := "test@testing.com"
	pro := new(models.Profile)
	pro.Id = id
	pro.Email = email

	svc.CreateProfile(pro)

	err := svc.DeleteProfile(pro.Id)
	assert.Nil(t, err)

	newProfile, err := svc.GetProfile(pro.Id)
	assert.NotNil(t, err)
	assert.Nil(t, newProfile)

	profileID, err := svc.GetProfileID(pro.Email)
	assert.NotNil(t, err)
	assert.Equal(t, "", profileID)
}
