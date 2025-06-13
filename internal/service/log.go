package service

import (
	"fmt"
	"nws/internal/model"
	"nws/internal/repository"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

type LogService struct {
	repo       repository.LogRepositoryInterface
	deviceRepo repository.DeviceRepositoryInterface
}

func NewLogService(repo repository.LogRepositoryInterface, deviceRepo repository.DeviceRepositoryInterface) *LogService {
	return &LogService{
		repo:       repo,
		deviceRepo: deviceRepo,
	}
}

// SimulateDeviceLog создает симулированный лог для устройства
func (s *LogService) SimulateDeviceLog(deviceId int) (*model.NetworkLog, error) {
	device, err := s.deviceRepo.GetDeviceById(deviceId)
	if err != nil {
		return nil, fmt.Errorf("device not found: %w", err)
	}

	actions := []string{
		"CONNECT",
		"DISCONNECT",
		"CONFIG_CHANGE",
		"STATUS_CHECK",
		"INTERFACE_UP",
		"INTERFACE_DOWN",
		"SECURITY_ALERT",
		"PERFORMANCE_WARNING",
		"SYSTEM_REBOOT",
		"FIRMWARE_UPDATE",
	}

	messageTemplates := map[string][]string{
		"CONNECT":             {"Established connection from %s", "New connection detected from %s", "Device connected successfully from %s"},
		"DISCONNECT":          {"Connection terminated from %s", "Device disconnected from %s", "Lost connection with %s"},
		"CONFIG_CHANGE":       {"Configuration updated: %s", "Settings modified: %s", "System configuration changed: %s"},
		"STATUS_CHECK":        {"Health check completed: %s", "Status verification: %s", "System status: %s"},
		"INTERFACE_UP":        {"Network interface %s is up", "Port %s activated", "Interface %s connected"},
		"INTERFACE_DOWN":      {"Network interface %s is down", "Port %s deactivated", "Interface %s disconnected"},
		"SECURITY_ALERT":      {"Security warning: %s", "Potential security breach: %s", "Security notification: %s"},
		"PERFORMANCE_WARNING": {"Performance issue detected: %s", "System performance warning: %s", "Resource usage alert: %s"},
		"SYSTEM_REBOOT":       {"System reboot initiated: %s", "Restarting device: %s", "System restart: %s"},
		"FIRMWARE_UPDATE":     {"Firmware update started: %s", "System upgrade in progress: %s", "Updating firmware: %s"},
	}

	action := actions[gofakeit.IntRange(0, len(actions)-1)]
	messages := messageTemplates[action]
	messageTemplate := messages[gofakeit.IntRange(0, len(messages)-1)]

	var messageDetails string
	switch action {
	case "CONNECT", "DISCONNECT":
		messageDetails = gofakeit.IPv4Address()
	case "CONFIG_CHANGE":
		configs := []string{"network settings", "security policy", "routing table", "access control", "QoS parameters"}
		messageDetails = configs[gofakeit.IntRange(0, len(configs)-1)]
	case "STATUS_CHECK":
		status := []string{"all systems normal", "minor issues detected", "performance optimal", "requires attention"}
		messageDetails = status[gofakeit.IntRange(0, len(status)-1)]
	case "INTERFACE_UP", "INTERFACE_DOWN":
		messageDetails = fmt.Sprintf("eth%d", gofakeit.IntRange(0, 4))
	case "SECURITY_ALERT":
		alerts := []string{"unusual traffic pattern", "multiple failed login attempts", "port scan detected", "suspicious activity"}
		messageDetails = alerts[gofakeit.IntRange(0, len(alerts)-1)]
	case "PERFORMANCE_WARNING":
		warnings := []string{"high CPU usage", "memory usage above threshold", "network congestion", "disk space low"}
		messageDetails = warnings[gofakeit.IntRange(0, len(warnings)-1)]
	case "SYSTEM_REBOOT":
		reasons := []string{"scheduled maintenance", "system update", "error recovery", "manual restart"}
		messageDetails = reasons[gofakeit.IntRange(0, len(reasons)-1)]
	case "FIRMWARE_UPDATE":
		messageDetails = fmt.Sprintf("version %s.%s.%s",
			gofakeit.IntRange(1, 5),
			gofakeit.IntRange(0, 9),
			gofakeit.IntRange(0, 99))
	}

	message := fmt.Sprintf(messageTemplate, messageDetails)

	log := &model.NetworkLog{
		DeviceID:  device.ID,
		Timestamp: time.Now(),
		Action:    action,
		Message:   message,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.repo.CreateLog(log)
}

// GetLogsByDeviceId получает все логи для конкретного устройства
func (s *LogService) GetLogsByDeviceId(deviceId int) ([]model.NetworkLog, error) {
	return s.repo.GetLogsByDeviceId(deviceId)
}

// GetAllLogs получает все логи
func (s *LogService) GetAllLogs() ([]model.NetworkLog, error) {
	return s.repo.GetAllLogs()
}

// GetRecentLogs получает последние логи с указанным лимитом
func (s *LogService) GetRecentLogs(limit int) ([]model.NetworkLog, error) {
	return s.repo.GetRecentLogs(limit)
}
