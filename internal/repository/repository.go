package repository

import (
	"nws/internal/model"
	"time"

	"gorm.io/gorm"
)

type AuthRepositoryInterface interface {
	Login(user *model.User) (*model.User, error)
	Register(user *model.User) (*model.User, error)
}

type UserRepositoryInterface interface {
	GetByID(id string) (*model.User, error)
	GetByName(name string) (*model.User, error)
	Update(user *model.User) error
	Delete(id string) error
}

type DeviceRepositoryInterface interface {
	CreateDevice(device *model.NetworkDevice) (*model.NetworkDevice, error)
	GetAllDevices() ([]model.NetworkDevice, error)
	GetDeviceById(id int) (*model.NetworkDevice, error)

	CreateConnection(connection *model.Connection) (*model.Connection, error)
	GetAllConnections() ([]model.Connection, error)
	GetConnectionById(connectionID int) error
	GetConnectionsByDeviceId(sourceDevice int, destDevice int) (*model.Connection, error)
	UpdateConnection(connection *model.Connection) error
	DeleteConnection(connectionID int) error

	TopologyRepositoryInterface
}

type ReportRepositoryInterface interface {
	Create(report *model.Reports) error
	GetByID(id int64) (*model.Reports, error)
	GetAll() ([]*model.Reports, error)
	Update(report *model.Reports) error
	Delete(id int64) error
}

type TopologyRepositoryInterface interface {
	CreateTopologyElement(topologyElement *model.TopologyElement) (*model.TopologyElement, error)
	GetAllTopology() ([]model.TopologyElement, error)
}

type LogRepositoryInterface interface {
	CreateLog(log *model.NetworkLog) (*model.NetworkLog, error)
	GetLogsByDeviceId(deviceId int) ([]model.NetworkLog, error)
	GetAllLogs() ([]model.NetworkLog, error)
	GetRecentLogs(limit int) ([]model.NetworkLog, error)
}

type MetricRepositoryInterface interface {
	CreateMetric(metric *model.Metric) (*model.Metric, error)
	GetMetricsByDeviceId(deviceId int) ([]model.Metric, error)
	GetMetricsByDeviceIdAndInterval(deviceId int, interval string, from, to time.Time) ([]model.Metric, error)
	GetLatestMetricsByDeviceId(deviceId int, limit int) ([]model.Metric, error)
	GetAllMetrics() ([]model.Metric, error)
	DeleteOldMetrics(before time.Time) error
}

type ConfigurationRepositoryInterface interface {
	Create(config *model.Configuration) error
	Update(config *model.Configuration) error
	Delete(id int) error
	GetByID(id int) (*model.Configuration, error)
	GetByDeviceID(deviceID int) ([]*model.Configuration, error)
	GetActive(deviceID int) (*model.Configuration, error)
}
type Repository struct {
	AuthRepositoryInterface
	UserRepositoryInterface
	DeviceRepositoryInterface
	LogRepositoryInterface
	MetricRepositoryInterface
	ReportRepositoryInterface
	ConfigurationRepositoryInterface
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		AuthRepositoryInterface:          NewAuthRepo(db),
		UserRepositoryInterface:          NewUserRepository(db),
		DeviceRepositoryInterface:        NewDeviceRepository(db),
		LogRepositoryInterface:           NewLogRepository(db),
		MetricRepositoryInterface:        NewMetricRepository(db),
		ReportRepositoryInterface:        NewReportRepository(db),
		ConfigurationRepositoryInterface: NewConfigurationRepo(db),
	}
}
