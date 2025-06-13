package service

import (
	"fmt"
	"math"
	"math/rand"
	"nws/internal/model"
	"nws/internal/repository"
	"time"
)

type MetricService struct {
	repo       repository.MetricRepositoryInterface
	deviceRepo repository.DeviceRepositoryInterface
}

func NewMetricService(repo repository.MetricRepositoryInterface, deviceRepo repository.DeviceRepositoryInterface) *MetricService {
	return &MetricService{
		repo:       repo,
		deviceRepo: deviceRepo,
	}
}

// SimulateDeviceMetric создает симулированные метрики для устройства
func (s *MetricService) SimulateDeviceMetric(deviceId int, interval string) (*model.Metric, error) {
	device, err := s.deviceRepo.GetDeviceById(deviceId)
	if err != nil {
		return nil, fmt.Errorf("device not found: %w", err)
	}

	// Получаем последние метрики устройства для плавного изменения
	lastMetrics, err := s.repo.GetLatestMetricsByDeviceId(deviceId, 1)
	var lastMetric *model.Metric
	if err == nil && len(lastMetrics) > 0 {
		lastMetric = &lastMetrics[0]
	}

	// Генерируем новые значения метрик
	var cpu, memory, network, uptime int
	if lastMetric != nil {
		// Если есть предыдущие метрики, генерируем новые значения с учетом предыдущих
		cpu = generateNextValue(lastMetric.CPU, 5, 0, 100)
		memory = generateNextValue(lastMetric.Memory, 3, 0, 100)
		network = generateNextValue(lastMetric.Network, 10, 0, 100)
		uptime = lastMetric.Uptime + int(time.Since(lastMetric.CreatedAt).Minutes())
	} else {
		// Если предыдущих метрик нет, генерируем начальные значения
		cpu = rand.Intn(60) + 20     // 20-80%
		memory = rand.Intn(50) + 30  // 30-80%
		network = rand.Intn(70) + 10 // 10-80%
		uptime = 0
	}

	// Определяем статус на основе метрик
	status := "normal"
	if cpu > 90 || memory > 90 || network > 90 {
		status = "critical"
	} else if cpu > 75 || memory > 75 || network > 75 {
		status = "warning"
	}

	metric := &model.Metric{
		DeviceID:  device.ID,
		CPU:       cpu,
		Memory:    memory,
		Network:   network,
		Uptime:    uptime,
		Interval:  interval,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.repo.CreateMetric(metric)
}

// generateNextValue генерирует следующее значение метрики с учетом предыдущего
func generateNextValue(previous, maxChange int, min, max int) int {
	// Генерируем случайное изменение
	change := rand.Float64()*float64(maxChange*2) - float64(maxChange)

	// Добавляем синусоидальное колебание для более реалистичного поведения
	oscillation := math.Sin(float64(time.Now().Unix())/300) * float64(maxChange/2)

	// Применяем изменение и колебание
	newValue := float64(previous) + change + oscillation

	// Ограничиваем значение в пределах min-max
	if newValue < float64(min) {
		newValue = float64(min)
	}
	if newValue > float64(max) {
		newValue = float64(max)
	}

	return int(newValue)
}

// GetMetricsByDeviceId получает все метрики для конкретного устройства
func (s *MetricService) GetMetricsByDeviceId(deviceId int) ([]model.Metric, error) {
	return s.repo.GetMetricsByDeviceId(deviceId)
}

// GetMetricsByDeviceIdAndInterval получает метрики устройства за период с определенным интервалом
func (s *MetricService) GetMetricsByDeviceIdAndInterval(deviceId int, interval string, from, to time.Time) ([]model.Metric, error) {
	return s.repo.GetMetricsByDeviceIdAndInterval(deviceId, interval, from, to)
}

// GetLatestMetricsByDeviceId получает последние метрики устройства
func (s *MetricService) GetLatestMetricsByDeviceId(deviceId int, limit int) ([]model.Metric, error) {
	return s.repo.GetLatestMetricsByDeviceId(deviceId, limit)
}

// GetAllMetrics получает все метрики
func (s *MetricService) GetAllMetrics() ([]model.Metric, error) {
	return s.repo.GetAllMetrics()
}

// CleanupOldMetrics удаляет старые метрики
func (s *MetricService) CleanupOldMetrics(retentionPeriod time.Duration) error {
	cutoff := time.Now().Add(-retentionPeriod)
	return s.repo.DeleteOldMetrics(cutoff)
}
