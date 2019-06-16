package service

import (
	"fmt"
	models2 "github.com/kas240sx89/samples/profileService/internal/models"
)

type Service struct {
	db DB
}

func New(database DB) *Service {

	return &Service{
		db: database,
	}
}

func (s *Service) HealthCheck() string {
	return "service is ok"
}

func (s *Service) GetProfile(id string) (*models2.Profile, error) {
	profile, err := GetProfile(id)
	if err != nil {
		return nil, err
	}
	if profile == nil {
		return nil, fmt.Errorf("profile not found for id: %s", id)
	}

	return profile, nil
}

func (s *Service) GetProfileID(email string) (string, error) {
	id, err := GetProfileID(email)
	if err != nil {
		return "", err
	}

	if id == "" {
		return id, fmt.Errorf("id not found")
	}

	return id, nil
}

func (s *Service) CreateProfile(profile *models2.Profile) (*models2.Profile, error) {
	return CreateProfile(profile)
}

func (s *Service) UpdateProfile(profile *models2.Profile) (*models2.Profile, error) {
	return UpdateProfile(profile)
}

func (s *Service) DeleteProfile(id string) error {
	return DeleteProfile(id)
}
