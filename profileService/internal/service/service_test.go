package service

import (
	db2 "github.com/kas240sx89/samples/profileService/internal/db"
	models2 "github.com/kas240sx89/samples/profileService/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//TODO rewrite tests using DB mocks

func TestNewService(t *testing.T) {

	database := db2.NewInMemoryDB()
	svc := New(&database)
	assert.NotNil(t, svc)
}

func testService() *Service {
	database := db2.NewInMemoryDB()
	return New(&database)
}

func TestHealthCheck(t *testing.T) {
	svc := testService()
	assert.Equal(t, "service is ok", HealthCheck())
}

func TestCreateProfile(t *testing.T) {
	svc := testService()

	pro := new(models2.Profile)
	pro.Id = "12345"

	profile, err := CreateProfile(pro)
	assert.Nil(t, err)
	assert.NotNil(t, profile)
}

func TestGetProfileExists(t *testing.T) {
	svc := testService()
	id := "12345"
	pro := new(models2.Profile)
	pro.Id = id

	CreateProfile(pro)

	profile, err := GetProfile(id)
	assert.Nil(t, err)
	assert.NotNil(t, profile)
	assert.Equal(t, id, profile.Id)
}

func TestGetProfileNonExist(t *testing.T) {
	svc := testService()
	profile, err := GetProfile("not real")
	assert.NotNil(t, err)
	assert.Nil(t, profile)
}

func TestGetProfileID(t *testing.T) {
	svc := testService()
	id := "12345"
	email := "test@testing.com"
	pro := new(models2.Profile)
	pro.Id = id
	pro.Email = email

	CreateProfile(pro)

	profileID, err := GetProfileID(email)
	assert.Nil(t, err)
	assert.Equal(t, id, profileID)

	profileID, err = GetProfileID("fake")
	assert.NotNil(t, err)
	assert.Equal(t, "", profileID)

}

func TestUpdateProfileID(t *testing.T) {
	svc := testService()
	id := "12345"
	email := "test@testing.com"
	pro := new(models2.Profile)
	pro.Id = id
	pro.Email = email

	_, err := CreateProfile(pro)
	assert.Nil(t, err)

	pro.LastUpdated = pro.LastUpdated.Add(-1 * time.Minute)
	failedPro, err := UpdateProfile(pro)
	assert.NotNil(t, err)

	failedPro.Info = "test info"
	failedPro.LastUpdated = failedPro.LastUpdated.Add(5 * time.Minute)

	updated, err := UpdateProfile(failedPro)
	assert.Nil(t, err)
	assert.Equal(t, failedPro.Info, updated.Info)
}

func TestDeleteProfile(t *testing.T) {
	svc := testService()
	id := "12345"
	email := "test@testing.com"
	pro := new(models2.Profile)
	pro.Id = id
	pro.Email = email

	CreateProfile(pro)

	err := DeleteProfile(pro.Id)
	assert.Nil(t, err)

	newProfile, err := GetProfile(pro.Id)
	assert.NotNil(t, err)
	assert.Nil(t, newProfile)

	profileID, err := GetProfileID(pro.Email)
	assert.NotNil(t, err)
	assert.Equal(t, "", profileID)
}
