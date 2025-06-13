package service

import (
	"errors"
	"nws/internal/model"
	"nws/internal/repository"
)

type ConfigurationSvc struct {
	repo repository.ConfigurationRepositoryInterface
}

func NewConfigurationService(repo repository.ConfigurationRepositoryInterface) *ConfigurationSvc {
	return &ConfigurationSvc{
		repo: repo,
	}
}

func (s *ConfigurationSvc) Create(config *model.Configuration) error {
	if err := s.ValidateConfiguration(config); err != nil {
		return err
	}
	return s.repo.Create(config)
}

func (s *ConfigurationSvc) Update(config *model.Configuration) error {
	if err := s.ValidateConfiguration(config); err != nil {
		return err
	}
	return s.repo.Update(config)
}

func (s *ConfigurationSvc) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *ConfigurationSvc) GetByID(id int) (*model.Configuration, error) {
	return s.repo.GetByID(id)
}

func (s *ConfigurationSvc) GetByDeviceID(deviceID int) ([]*model.Configuration, error) {
	return s.repo.GetByDeviceID(deviceID)
}

func (s *ConfigurationSvc) GetActive(deviceID int) (*model.Configuration, error) {
	return s.repo.GetActive(deviceID)
}

func (s *ConfigurationSvc) ValidateConfiguration(config *model.Configuration) error {
	if config == nil {
		return errors.New("configuration is nil")
	}
	if config.DeviceID <= 0 {
		return errors.New("invalid device ID")
	}
	if config.Version == "" {
		return errors.New("version is required")
	}
	if config.Content == "" {
		return errors.New("configuration content is required")
	}
	return nil
}
