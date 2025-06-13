package repository

import (
	"errors"
	"fmt"
	"nws/internal/model"

	"gorm.io/gorm"
)

type DeviceRepository struct {
	db *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) *DeviceRepository {
	return &DeviceRepository{db: db}
}

func (r *DeviceRepository) CreateDevice(device *model.NetworkDevice) (*model.NetworkDevice, error) {
	if err := r.db.Create(&device).Error; err != nil {
		return nil, err
	}

	return device, nil
}

func (r *DeviceRepository) GetAllDevices() ([]model.NetworkDevice, error) {
	var devices []model.NetworkDevice

	if err := r.db.Preload("Connections").Preload("Metrics").
		Preload("NetworkInterfaces").Preload("NetworkLogs").
		Preload("Configurations").Find(&devices).Error; err != nil {
		return nil, err
	}

	return devices, nil
}

func (r *DeviceRepository) GetDeviceById(id int) (*model.NetworkDevice, error) {
	var device model.NetworkDevice

	if err := r.db.Preload("Configurations").Preload("Connections").
		Preload("Metrics").Preload("NetworkInterfaces").Preload("NetworkLogs").
		Where("id = ?", id).First(&device).Error; err != nil {
		return nil, err
	}

	return &device, nil
}

func (r *DeviceRepository) GetAllConnections() ([]model.Connection, error) {
	var connections []model.Connection

	if err := r.db.Find(&connections).Error; err != nil {
		return nil, err
	}

	return connections, nil
}

func (r *DeviceRepository) GetConnectionsByDeviceId(sourceDevice int, destDevice int) (*model.Connection, error) {
	var (
		connection model.Connection
		//device     model.NetworkDevice
	)

	if err := r.db.Model(connection).Where("source_device_id = ? AND destination_device_id = ?", sourceDevice, destDevice).
		First(&connection).Error; err != nil {
		return nil, err
	}

	return &connection, nil
}

func (r *DeviceRepository) CreateConnection(connection *model.Connection) (*model.Connection, error) {
	// Проверяем существование исходного устройства
	var sourceDevice model.NetworkDevice
	if err := r.db.Where("ip_address = ?", connection.SourceDeviceIP).First(&sourceDevice).Error; err != nil {
		return nil, fmt.Errorf("source device with IP %s not found: %w", connection.SourceDeviceIP, err)
	}

	// Проверяем существование целевого устройства
	var destDevice model.NetworkDevice
	if err := r.db.Where("ip_address = ?", connection.DestinationDeviceIP).First(&destDevice).Error; err != nil {
		return nil, fmt.Errorf("destination device with IP %s not found: %w", connection.DestinationDeviceIP, err)
	}

	// Проверяем существование соединения
	var existConnection model.Connection
	err := r.db.Where("source_device_ip = ? AND destination_device_ip = ?",
		connection.SourceDeviceIP, connection.DestinationDeviceIP).
		First(&existConnection).Error

	if err == nil {
		return nil, fmt.Errorf("connection already exists between devices %s and %s",
			connection.SourceDeviceIP, connection.DestinationDeviceIP)
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Создаем соединение
	if err := r.db.Create(&connection).Error; err != nil {
		return nil, err
	}

	// Загружаем связанные данные
	if err := r.db.Preload("SourceDevice").
		Preload("DestinationDevice").
		First(&connection, connection.ID).Error; err != nil {
		return nil, err
	}

	return connection, nil
}

func (r *DeviceRepository) UpdateConnection(connection *model.Connection) error {
	return r.db.Save(connection).Error
}

func (r *DeviceRepository) GetConnectionById(connectionID int) error {
	return r.db.First(&model.Connection{}, connectionID).Error
}

func (r *DeviceRepository) DeleteConnection(connectionID int) error {
	var (
		connection model.Connection
		err        error
	)

	if err = r.db.Where("id = ?", connectionID).First(&connection).Error; err != nil {
		return err
	}

	if err = r.db.Delete(&connection).Error; err != nil {
		return err
	}

	return err
}

func (r *DeviceRepository) CreateTopologyElement(topologyElement *model.TopologyElement) (*model.TopologyElement, error) {
	if err := r.db.Create(&topologyElement).Error; err != nil {
		return nil, err
	}

	return topologyElement, nil
}

func (r *DeviceRepository) GetAllTopology() ([]model.TopologyElement, error) {
	var topology []model.TopologyElement

	if err := r.db.Preload("Device").Preload("Device.Connections").
		Preload("Device.Metrics").Preload("Device.NetworkInterfaces").
		Preload("Device.NetworkLogs").Preload("Device.Configurations").
		Preload("Topology").Find(&topology).Error; err != nil {
		return nil, err
	}

	return topology, nil
}
