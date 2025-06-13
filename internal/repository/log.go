package repository

import (
	"nws/internal/model"

	"gorm.io/gorm"
)

type LogRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) *LogRepository {
	return &LogRepository{db: db}
}

func (r *LogRepository) CreateLog(log *model.NetworkLog) (*model.NetworkLog, error) {
	if err := r.db.Create(log).Error; err != nil {
		return nil, err
	}
	return log, nil
}

func (r *LogRepository) GetLogsByDeviceId(deviceId int) ([]model.NetworkLog, error) {
	var logs []model.NetworkLog
	if err := r.db.Where("device_id = ?", deviceId).Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}

func (r *LogRepository) GetAllLogs() ([]model.NetworkLog, error) {
	var logs []model.NetworkLog
	if err := r.db.Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}

func (r *LogRepository) GetRecentLogs(limit int) ([]model.NetworkLog, error) {
	var logs []model.NetworkLog
	if err := r.db.Order("timestamp desc").Limit(limit).Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}
