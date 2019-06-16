package service

import (
	models2 "github.com/kas240sx89/samples/profileService/internal/models"
)

//go:generate moq -out db_mock_test.go . DB

type DB interface {
	GetProfileID(email string) (string, error)
	GetProfile(id string) (*models2.Profile, error)
	CreateProfile(profile *models2.Profile) (*models2.Profile, error)
	UpdateProfile(profile *models2.Profile) (*models2.Profile, error)
	DeleteProfile(id string) error
}
