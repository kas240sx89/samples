package db

import (
	"fmt"
	"github.com/kas240sx89/samples/profileService/internal/models"
	"sync"
	"time"
)

type InMemoryDB struct {
	mtx          sync.RWMutex
	profileStore map[string]models.Profile
	emailStore   map[string]string
}

func NewInMemoryDB() InMemoryDB {
	p := make(map[string]models.Profile)
	e := make(map[string]string)
	return InMemoryDB{
		profileStore: p,
		emailStore:   e,
	}
}

//GetProfileID returns the profile id associated with the "email", if none exists a empty string is returned
func (db *InMemoryDB) GetProfileID(email string) (string, error) {
	db.mtx.RLock()
	defer db.mtx.RUnlock()

	email, ok := db.emailStore[email]
	if !ok {
		return "", nil
	}
	return email, nil
}

//GetProfile returns the profile associated with the passed in "id"
func (db *InMemoryDB) GetProfile(id string) (*models.Profile, error) {
	db.mtx.RLock()
	defer db.mtx.RUnlock()

	profile, ok := db.profileStore[id]
	if !ok {
		return nil, fmt.Errorf("unable to find profile")
	}
	return &profile, nil
}

//CreateProfile creates the profile withing the database. if the profile exists or the email exist returns error
func (db *InMemoryDB) CreateProfile(profile *models.Profile) (*models.Profile, error) {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	if _, ok := db.profileStore[profile.Id]; ok {
		return nil, fmt.Errorf("profile already exists")
	}

	if _, ok := db.emailStore[profile.Email]; ok {
		return nil, fmt.Errorf("email already exists")
	}

	profile.LastUpdated = time.Now()
	db.profileStore[profile.Id] = *profile
	db.emailStore[profile.Email] = profile.Id

	return profile, nil
}

//UpdateProfile updates a profile, if profile does not exist returns error
func (db *InMemoryDB) UpdateProfile(profile *models.Profile) (*models.Profile, error) {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	old, ok := db.profileStore[profile.Id]
	if !ok {
		return nil, fmt.Errorf("profile not found")
	}

	if old.Email != profile.Email {
		delete(db.emailStore, old.Email)
		db.emailStore[profile.Email] = profile.Id
	}

	if old.LastUpdated.UnixNano() > profile.LastUpdated.UnixNano() {
		return &old, fmt.Errorf("there is a newer profile")
	}

	profile.LastUpdated = time.Now()
	db.profileStore[profile.Id] = *profile
	return profile, nil
}

//DeleteProfile removes a profile by "id"
func (db *InMemoryDB) DeleteProfile(id string) error {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	profile, ok := db.profileStore[id]
	if !ok {
		return nil
	}

	delete(db.emailStore, profile.Email)
	delete(db.profileStore, profile.Id)
	return nil
}
