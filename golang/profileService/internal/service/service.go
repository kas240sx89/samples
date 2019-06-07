package service

import (
	"fmt"
	"github.com/kas240sx89/samples/golang/profileService/internal/models"
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

func (s *Service) GetProfile(id string) (*models.Profile, error) {
	profile, err := s.db.GetProfile(id)
	if err != nil {
		return nil, err
	}
	if profile == nil {
		return nil, fmt.Errorf("profile not found for id: %s", id)
	}

	return profile, nil
}

func (s *Service) GetProfileID(email string) (string, error) {
	id, err := s.db.GetProfileID(email)
	if err != nil {
		return "", err
	}

	if id == "" {
		return id, fmt.Errorf("id not found")
	}

	return id, nil
}

func (s *Service) CreateProfile(profile *models.Profile) (*models.Profile, error) {
	return s.db.CreateProfile(profile)
}

func (s *Service) UpdateProfile(profile *models.Profile) (*models.Profile, error) {
	return s.db.UpdateProfile(profile)
}

func (s *Service) DeleteProfile(id string) error {
	return s.db.DeleteProfile(id)
}
