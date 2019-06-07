package service

import (
	"github.com/kas240sx89/samples/golang/profileService/internal/models"
)

type DB interface {
	GetProfileID(email string) (string, error)
	GetProfile(id string) (*models.Profile, error)
	CreateProfile(profile *models.Profile) (*models.Profile, error)
	UpdateProfile(profile *models.Profile) (*models.Profile, error)
	DeleteProfile(id string) error
}
