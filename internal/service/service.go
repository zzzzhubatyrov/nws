package service

import (
	"nws/internal/model"
	"nws/internal/repository"
	"time"
)

type AuthServiceInterface interface {
	Login(user map[string]string) (*model.LoginResponse, error)
	Register(user map[string]string) (*model.User, error)
}

type UserServiceInterface interface {
	GetByID(id string) (*model.User, error)
	GetByName(name string) (*model.User, error)
	Update(user *model.User) error
	Delete(id string) error
}

type DeviceServiceInterface interface {
	CreateDevice(device map[string]string) (*model.NetworkDevice, error)
	GetAllDevices() ([]model.NetworkDevice, error)
	GetDeviceById(id int) (*model.NetworkDevice, error)

	CreateConnection(srcIP, destIP string) (*model.Connection, error)
	GetAllConnections() ([]model.Connection, error)
	GetConnectionById(connectionID int) error
	GetConnectionsByDeviceId(sourceDevice int, destDevice int) (*model.Connection, error)
	UpdateConnection(connection *model.Connection) error
	DeleteConnection(connectionID int) error

	TopologyServiceInterface
}

type TopologyServiceInterface interface {
	CreateTopologyElement(deviceID int) (*model.TopologyElement, error)
	GetAllTopology() ([]model.TopologyElement, error)
}

type LogServiceInterface interface {
	SimulateDeviceLog(deviceId int) (*model.NetworkLog, error)
	GetLogsByDeviceId(deviceId int) ([]model.NetworkLog, error)
	GetAllLogs() ([]model.NetworkLog, error)
	GetRecentLogs(limit int) ([]model.NetworkLog, error)
}

type MetricServiceInterface interface {
	SimulateDeviceMetric(deviceId int, interval string) (*model.Metric, error)
	GetMetricsByDeviceId(deviceId int) ([]model.Metric, error)
	GetMetricsByDeviceIdAndInterval(deviceId int, interval string, from, to time.Time) ([]model.Metric, error)
	GetLatestMetricsByDeviceId(deviceId int, limit int) ([]model.Metric, error)
	GetAllMetrics() ([]model.Metric, error)
	CleanupOldMetrics(retentionPeriod time.Duration) error
}

type ReportServiceInterface interface {
	CreateReport(report *model.Reports) error
	GetReport(id int64) (*model.Reports, error)
	GetAllReports() ([]*model.Reports, error)
	UpdateReport(report *model.Reports) error
	DeleteReport(id int64) error
}

type ConfigurationServiceInterface interface {
	Create(config *model.Configuration) error
	Update(config *model.Configuration) error
	Delete(id int) error
	GetByID(id int) (*model.Configuration, error)
	GetByDeviceID(deviceID int) ([]*model.Configuration, error)
	GetActive(deviceID int) (*model.Configuration, error)
	ValidateConfiguration(config *model.Configuration) error
}

type Service struct {
	AuthServiceInterface
	UserServiceInterface
	DeviceServiceInterface
	LogServiceInterface
	MetricServiceInterface
	ReportServiceInterface
	ConfigurationServiceInterface
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthServiceInterface:          NewAuthService(repo.AuthRepositoryInterface),
		UserServiceInterface:          NewUserService(repo.UserRepositoryInterface),
		DeviceServiceInterface:        NewDeviceService(repo.DeviceRepositoryInterface),
		LogServiceInterface:           NewLogService(repo.LogRepositoryInterface, repo.DeviceRepositoryInterface),
		MetricServiceInterface:        NewMetricService(repo.MetricRepositoryInterface, repo.DeviceRepositoryInterface),
		ReportServiceInterface:        NewReportService(repo.ReportRepositoryInterface),
		ConfigurationServiceInterface: NewConfigurationService(repo.ConfigurationRepositoryInterface),
	}
}
