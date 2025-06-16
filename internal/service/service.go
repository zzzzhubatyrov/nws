package service

import (
	"fmt"
	"math/rand"
	"time"

	"trace/internal/database"
	"trace/internal/models"
	"trace/internal/repository"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// DeviceService manages devices
type DeviceService struct {
	repo *repository.Repository[models.Device]
}

type ConnectionService struct {
	repo *repository.Repository[models.Connection]
}

// MetricService handles metrics
type MetricService struct {
	repo *repository.Repository[models.Metric]
}

type LogService struct {
	repo *repository.Repository[models.Log]
}

// AlertService manages alerts
type AlertService struct {
	repo *repository.Repository[models.Alert]
}

func NewDeviceService() *DeviceService {
	db := database.GetDB()
	db.AutoMigrate(&models.Device{})
	return &DeviceService{repo: repository.New[models.Device](db)}
}

func (s *DeviceService) Get(id uint) (models.Device, error) {
	var d models.Device
	err := s.repo.DB().First(&d, id).Error
	return d, err
}

func (s *DeviceService) List() ([]models.Device, error) {
	var res []models.Device
	err := s.repo.All(&res)
	return res, err
}

func (s *DeviceService) GenerateRandomDevice() models.Device {
	types := []string{"router", "switch", "server"}
	vendors := []string{"Cisco", "Juniper", "Dell", "HP"}
	modelsList := []string{"XR500", "MX480", "PowerEdge", "ProCurve"}

	ip := fmt.Sprintf("10.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256))
	name := fmt.Sprintf("Device-%d", rand.Intn(10000))

	d := models.Device{
		Name:      name,
		Type:      types[rand.Intn(len(types))],
		IP:        ip,
		Status:    "offline",
		CPU:       rand.Intn(20),
		Memory:    rand.Intn(100),
		Vendor:    vendors[rand.Intn(len(vendors))],
		Model:     modelsList[rand.Intn(len(modelsList))],
		CreatedAt: time.Now(),
	}
	// assign coordinates
	d.X = 100 + rand.Intn(500)
	d.Y = 100 + rand.Intn(300)
	return d
}

func (s *DeviceService) Save(d *models.Device) error {
	if d.ID == 0 {
		// default values for new device
		if d.Status == "" {
			d.Status = "offline"
		}
		if d.X == 0 && d.Y == 0 {
			d.X = 100 + rand.Intn(500)
			d.Y = 100 + rand.Intn(300)
		}
		// set creation time if not present
		if d.CreatedAt.IsZero() {
			d.CreatedAt = time.Now()
		}
		// default uptime for new device
		if d.Uptime == "" {
			d.Uptime = "0s"
		}
		// new device – we could generate coordinates or uuid here if needed
		return s.repo.Create(d)
	}
	return s.repo.Update(d)
}

func (s *DeviceService) Delete(id uint) error {
	return s.repo.Delete(id)
}

// Connection service
func NewConnectionService() *ConnectionService {
	db := database.GetDB()
	db.AutoMigrate(&models.Connection{})
	return &ConnectionService{repo: repository.New[models.Connection](db)}
}

func (s *ConnectionService) List() ([]models.Connection, error) {
	var res []models.Connection
	err := s.repo.All(&res)
	return res, err
}

func (s *ConnectionService) Save(c *models.Connection) error {
	if c.ID == 0 {
		return s.repo.Create(c)
	}
	return s.repo.Update(c)
}

func (s *ConnectionService) Delete(id uint) error { return s.repo.Delete(id) }

// Metric service
func NewMetricService() *MetricService {
	db := database.GetDB()
	db.AutoMigrate(&models.Metric{})
	return &MetricService{repo: repository.New[models.Metric](db)}
}

func (s *MetricService) ListByDevice(id uint) ([]models.Metric, error) {
	var res []models.Metric
	err := s.repo.DB().Where("device_id = ?", id).Find(&res).Error
	return res, err
}

func (s *MetricService) ListByConnection(id uint) ([]models.Metric, error) {
	var res []models.Metric
	err := s.repo.DB().Where("connection_id = ?", id).Find(&res).Error
	return res, err
}

func (s *MetricService) Save(m *models.Metric) error {
	if m.ID == 0 {
		return s.repo.Create(m)
	}
	return s.repo.Update(m)
}

// GenerateAndSave creates random CPU & Memory metric for given device
func (s *MetricService) GenerateAndSave(device models.Device) (models.Metric, error) {
	metricNames := []string{"cpu", "memory"}
	name := metricNames[rand.Intn(len(metricNames))]
	val := 0.0
	if name == "cpu" {
		val = float64(rand.Intn(100))
	} else {
		val = float64(rand.Intn(100))
	}
	m := models.Metric{
		DeviceID: &device.ID,
		Name:     name,
		Value:    val,
		Time:     time.Now(),
	}
	err := s.Save(&m)
	return m, err
}

// Log service
func NewLogService() *LogService {
	db := database.GetDB()
	db.AutoMigrate(&models.Log{})
	return &LogService{repo: repository.New[models.Log](db)}
}

func (s *LogService) ListByDevice(id uint) ([]models.Log, error) {
	var res []models.Log
	err := s.repo.DB().Where("device_id = ?", id).Find(&res).Error
	return res, err
}

func (s *LogService) ListByConnection(id uint) ([]models.Log, error) {
	var res []models.Log
	err := s.repo.DB().Where("connection_id = ?", id).Find(&res).Error
	return res, err
}

func (s *LogService) Save(l *models.Log) error {
	if l.ID == 0 {
		return s.repo.Create(l)
	}
	return s.repo.Update(l)
}

// Alert service
func NewAlertService() *AlertService {
	db := database.GetDB()
	db.AutoMigrate(&models.Alert{})
	return &AlertService{repo: repository.New[models.Alert](db)}
}

func (s *AlertService) List() ([]models.Alert, error) {
	var res []models.Alert
	err := s.repo.All(&res)
	return res, err
}

func (s *AlertService) Save(a *models.Alert) error {
	if a.ID == 0 {
		a.Time = time.Now().Format(time.RFC3339)
		return s.repo.Create(a)
	}
	return s.repo.Update(a)
}

// GenerateAlertForDevice creates and saves an alert based on device resource usage
func (s *AlertService) GenerateAlertForDevice(device models.Device) (models.Alert, error) {
	sev := "info"
	if device.CPU > 90 || device.Memory > 90 {
		sev = "critical"
	} else if device.CPU > 70 || device.Memory > 70 {
		sev = "warning"
	}
	alert := models.Alert{
		Device:   device.Name,
		Type:     "resource",
		Severity: sev,
		Status:   "open",
		Title:    fmt.Sprintf("High usage on %s", device.Name),
		Message:  fmt.Sprintf("CPU %d%%, Memory %d%%", device.CPU, device.Memory),
		Time:     time.Now().Format(time.RFC3339),
	}
	err := s.Save(&alert)
	return alert, err
}

func (s *AlertService) Delete(id uint) error { return s.repo.Delete(id) }
