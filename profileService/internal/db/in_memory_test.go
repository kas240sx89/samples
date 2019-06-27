package db

import (
	"github.com/kas240sx89/samples/profileService/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	testId    = "test12345"
	testname  = "SuperTester"
	testEmail = "someone@testing.com"
)

func testSampleProfile() *models.Profile {
	return &models.Profile{
		Id:          testId,
		Username:    testname,
		Email:       testEmail,
		LastUpdated: time.Now(),
		Items:       make([]models.Item, 0),
	}
}

func testSampleDB() *InMemoryDB {
	db := NewInMemoryDB()
	_, _ = db.CreateProfile(testSampleProfile())
	return &db
}

func TestNewInMemoryDB(t *testing.T) {
	db := NewInMemoryDB()
	assert.NotNil(t, db)
}

func TestInMemoryDB_CreateProfile(t *testing.T) {
	db := NewInMemoryDB()
	profile, err := db.CreateProfile(testSampleProfile())
	assert.Nil(t, err)
	assert.NotNil(t, profile)

	//profile stored in db
	dbProfile, ok := db.profileStore[profile.Id]
	assert.True(t, ok)
	assert.Equal(t, &dbProfile, profile)

	//email stored in db
	profileId, ok := db.emailStore[profile.Email]
	assert.True(t, ok)
	assert.Equal(t, profileId, profile.Id)
}

func TestInMemoryDB_GetProfileID(t *testing.T) {
	db := testSampleDB()

	//existing profile
	profileId, err := db.GetProfileID(testEmail)
	assert.Nil(t, err)
	assert.Equal(t, profileId, testId)

	//non-existing profile
	profileId, err = db.GetProfileID("fakeEmail")
	assert.Nil(t, err)
	assert.Equal(t, profileId, "")
}

func TestInMemoryDB_GetProfile(t *testing.T) {
	db := testSampleDB()
	pro, err := db.GetProfile(testId)
	assert.Nil(t, err)
	assert.Equal(t, pro.Id, testId)

	pro, err = db.GetProfile("fakeid")
	assert.NotNil(t, err)
	assert.Nil(t, pro)
}

func TestInMemoryDB_UpdateProfile(t *testing.T) {
	db := testSampleDB()
	pro := testSampleProfile()

	pro.LastUpdated = pro.LastUpdated.Add(-1 * time.Minute)
	failedPro, err := db.UpdateProfile(pro)
	assert.NotNil(t, err)

	failedPro.Info = "test info"
	failedPro.LastUpdated = failedPro.LastUpdated.Add(5 * time.Minute)

	updated, err := db.UpdateProfile(failedPro)
	assert.Nil(t, err)
	assert.Equal(t, failedPro.Info, updated.Info)
}

func TestInMemoryDB_DeleteProfile(t *testing.T) {
	db := testSampleDB()
	err := db.DeleteProfile(testId)
	assert.Nil(t, err)

	_, ok := db.profileStore[testId]
	assert.False(t, ok)

	_, ok = db.emailStore[testEmail]
	assert.False(t, ok)
}
