package repository

import (
	"nws/internal/model"
	"time"

	"gorm.io/gorm"
)

type MetricRepository struct {
	db *gorm.DB
}

func NewMetricRepository(db *gorm.DB) *MetricRepository {
	return &MetricRepository{db: db}
}

func (r *MetricRepository) CreateMetric(metric *model.Metric) (*model.Metric, error) {
	if err := r.db.Create(metric).Error; err != nil {
		return nil, err
	}
	return metric, nil
}

func (r *MetricRepository) GetMetricsByDeviceId(deviceId int) ([]model.Metric, error) {
	var metrics []model.Metric
	if err := r.db.Where("device_id = ?", deviceId).Find(&metrics).Error; err != nil {
		return nil, err
	}
	return metrics, nil
}

func (r *MetricRepository) GetMetricsByDeviceIdAndInterval(deviceId int, interval string, from, to time.Time) ([]model.Metric, error) {
	var metrics []model.Metric
	if err := r.db.Where("device_id = ? AND interval = ? AND created_at BETWEEN ? AND ?",
		deviceId, interval, from, to).Find(&metrics).Error; err != nil {
		return nil, err
	}
	return metrics, nil
}

func (r *MetricRepository) GetLatestMetricsByDeviceId(deviceId int, limit int) ([]model.Metric, error) {
	var metrics []model.Metric
	if err := r.db.Where("device_id = ?", deviceId).
		Order("created_at desc").
		Limit(limit).
		Find(&metrics).Error; err != nil {
		return nil, err
	}
	return metrics, nil
}

func (r *MetricRepository) GetAllMetrics() ([]model.Metric, error) {
	var metrics []model.Metric
	if err := r.db.Find(&metrics).Error; err != nil {
		return nil, err
	}
	return metrics, nil
}

func (r *MetricRepository) DeleteOldMetrics(before time.Time) error {
	return r.db.Where("created_at < ?", before).Delete(&model.Metric{}).Error
}
