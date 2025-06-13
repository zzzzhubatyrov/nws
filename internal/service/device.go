package service

import (
	"nws/internal/model"
	"nws/internal/repository"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type DeviceService struct {
	repo repository.DeviceRepositoryInterface
}

func NewDeviceService(repo repository.DeviceRepositoryInterface) *DeviceService {
	return &DeviceService{repo: repo}
}

func (s *DeviceService) CreateDevice(device map[string]string) (*model.NetworkDevice, error) {
	newDevice := &model.NetworkDevice{
		Hostname:    device["hostname"],
		IPAddress:   device["ipAddress"],
		Type:        device["type"],
		Vendor:      device["vendor"],
		Model:       device["model"],
		Serial:      gofakeit.UUID(),
		Status:      string(model.UP),
		LastChecked: time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createDevice, err := s.repo.CreateDevice(newDevice)
	if err != nil {
		return nil, err
	}

	_, err = s.CreateTopologyElement(createDevice.ID)
	if err != nil {
		return nil, err
	}

	return createDevice, nil
}

func (s *DeviceService) GetAllDevices() ([]model.NetworkDevice, error) {
	return s.repo.GetAllDevices()
}

func (s *DeviceService) GetDeviceById(id int) (*model.NetworkDevice, error) {
	return s.repo.GetDeviceById(id)
}

func (s *DeviceService) CreateConnection(srcIP, destIP string) (*model.Connection, error) {
	//sourceId, _ := strconv.Atoi()
	//destId, _ := strconv.Atoi()

	newConnection := &model.Connection{
		SourceDeviceIP:      srcIP,
		DestinationDeviceIP: destIP,
		Protocol:            string(model.TCP),
		Port:                int(model.PORT80),
		Status:              string(model.ACTIVE),
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	createConnection, err := s.repo.CreateConnection(newConnection)
	if err != nil {
		return nil, err
	}

	return createConnection, nil
}

func (s *DeviceService) GetAllConnections() ([]model.Connection, error) {
	return s.repo.GetAllConnections()
}

func (s *DeviceService) GetConnectionById(connectionID int) error {
	return s.repo.GetConnectionById(connectionID)
}

func (s *DeviceService) DeleteConnection(connectionID int) error {
	return s.repo.DeleteConnection(connectionID)
}

func (s *DeviceService) UpdateConnection(connection *model.Connection) error {
	return s.repo.UpdateConnection(connection)
}

func (s *DeviceService) GetConnectionsByDeviceId(sourceDevice int, destDevice int) (*model.Connection, error) {
	return s.repo.GetConnectionsByDeviceId(sourceDevice, destDevice)
}

// CreateTopologyElement TODO: ПЕРЕДЕЛЕАТЬ
func (s *DeviceService) CreateTopologyElement(deviceID int) (*model.TopologyElement, error) {
	newTopology := &model.TopologyElement{
		TopologyID:  1,
		DeviceID:    deviceID,
		X:           gofakeit.Float32(),
		Y:           gofakeit.Float32(),
		CustomStyle: "color: #FF0000; size: 30",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createTopology, err := s.repo.CreateTopologyElement(newTopology)
	if err != nil {
		return nil, err
	}

	return createTopology, nil
}

func (s *DeviceService) GetAllTopology() ([]model.TopologyElement, error) {
	return s.repo.GetAllTopology()
}
