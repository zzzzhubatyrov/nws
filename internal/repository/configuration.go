package repository

import (
	"errors"
	"nws/internal/model"

	"gorm.io/gorm"
)

type ConfigurationRepo struct {
	db *gorm.DB
}

func NewConfigurationRepo(db *gorm.DB) *ConfigurationRepo {
	return &ConfigurationRepo{
		db: db,
	}
}

func (r *ConfigurationRepo) Create(config *model.Configuration) error {
	return r.db.Create(config).Error
}

func (r *ConfigurationRepo) Update(config *model.Configuration) error {
	return r.db.Save(config).Error
}

func (r *ConfigurationRepo) Delete(id int) error {
	return r.db.Delete(&model.Configuration{}, id).Error
}

func (r *ConfigurationRepo) GetByID(id int) (*model.Configuration, error) {
	var config model.Configuration
	if err := r.db.First(&config, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &config, nil
}

func (r *ConfigurationRepo) GetByDeviceID(deviceID int) ([]*model.Configuration, error) {
	var configs []*model.Configuration
	if err := r.db.Where("device_id = ?", deviceID).Find(&configs).Error; err != nil {
		return nil, err
	}
	return configs, nil
}

func (r *ConfigurationRepo) GetActive(deviceID int) (*model.Configuration, error) {
	var config model.Configuration
	if err := r.db.Where("device_id = ? AND is_active = ?", deviceID, true).First(&config).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &config, nil
}
