package service

import (
	"github.com/kas240sx89/samples/profileService/internal/models"
)

//go:generate moq -out db_mock_test.go . DB

type DB interface {
	GetProfileID(email string) (string, error)
	GetProfile(id string) (*models.Profile, error)
	CreateProfile(profile *models.Profile) (*models.Profile, error)
	UpdateProfile(profile *models.Profile) (*models.Profile, error)
	DeleteProfile(id string) error
}
