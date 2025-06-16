package main

import (
	"context"
	"fmt"
	"time"

	"trace/internal/models"
	"trace/internal/service"
)

// App struct
type App struct {
	ctx context.Context

	deviceSvc     *service.DeviceService
	connectionSvc *service.ConnectionService
	metricSvc     *service.MetricService
	logSvc        *service.LogService
	alertSvc      *service.AlertService
	authSvc       *service.AuthService
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		deviceSvc:     service.NewDeviceService(),
		connectionSvc: service.NewConnectionService(),
		metricSvc:     service.NewMetricService(),
		logSvc:        service.NewLogService(),
		alertSvc:      service.NewAlertService(),
		authSvc:       service.NewAuthService(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	go a.metricLoop()
}

// metricLoop generates metrics every 15s
func (a *App) metricLoop() {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-a.ctx.Done():
			return
		case <-ticker.C:
			devices, err := a.deviceSvc.List()
			if err != nil {
				a.logSvc.Save(&models.Log{
					Message: fmt.Sprintf("Failed to list devices: %v", err),
					Level:   "error",
				})
				continue
			}

			for _, d := range devices {
				m, err := a.metricSvc.GenerateAndSave(d)
				if err != nil {
					a.logSvc.Save(&models.Log{
						DeviceID: &d.ID,
						Message:  fmt.Sprintf("Failed to generate metrics: %v", err),
						Level:    "error",
					})
					continue
				}

				// Create a copy of the device for updates
				updatedDevice := d

				// update device struct with latest metric
				if m.Name == "cpu" {
					updatedDevice.CPU = int(m.Value)
				} else if m.Name == "memory" {
					updatedDevice.Memory = int(m.Value)
				}

				// update uptime string
				if !updatedDevice.CreatedAt.IsZero() {
					dur := time.Since(updatedDevice.CreatedAt)
					updatedDevice.Uptime = formatDuration(dur)
				}

				if err := a.deviceSvc.Save(&updatedDevice); err != nil {
					a.logSvc.Save(&models.Log{
						DeviceID: &d.ID,
						Message:  fmt.Sprintf("Failed to update device: %v", err),
						Level:    "error",
					})
				}
			}
		}
	}
}

// -------- Devices ---------
func (a *App) ListDevices() ([]models.Device, error) {
	return a.deviceSvc.List()
}

func (a *App) SaveDevice(d models.Device) (models.Device, error) {
	err := a.deviceSvc.Save(&d)
	return d, err
}

func (a *App) DeleteDevice(id uint) error {
	return a.deviceSvc.Delete(id)
}

// GenerateDevice creates a random device, saves it and returns it
func (a *App) GenerateDevice() (models.Device, error) {
	d := a.deviceSvc.GenerateRandomDevice()
	err := a.deviceSvc.Save(&d)
	return d, err
}

// -------- Connections ---------
func (a *App) ListConnections() ([]models.Connection, error) {
	return a.connectionSvc.List()
}

func (a *App) SaveConnection(c models.Connection) (models.Connection, error) {
	err := a.connectionSvc.Save(&c)
	return c, err
}

func (a *App) DeleteConnection(id uint) error {
	return a.connectionSvc.Delete(id)
}

// -------- Metrics ---------
func (a *App) ListDeviceMetrics(deviceID uint) ([]models.Metric, error) {
	return a.metricSvc.ListByDevice(deviceID)
}

func (a *App) ListConnectionMetrics(connID uint) ([]models.Metric, error) {
	return a.metricSvc.ListByConnection(connID)
}

// -------- Logs ---------
func (a *App) ListDeviceLogs(deviceID uint) ([]models.Log, error) {
	return a.logSvc.ListByDevice(deviceID)
}

func (a *App) ListConnectionLogs(connID uint) ([]models.Log, error) {
	return a.logSvc.ListByConnection(connID)
}

// -------- Alerts ---------
func (a *App) ListAlerts() ([]models.Alert, error) {
	return a.alertSvc.List()
}

func (a *App) SaveAlert(al models.Alert) error {
	return a.alertSvc.Save(&al)
}

func (a *App) DeleteAlert(id uint) error {
	return a.alertSvc.Delete(id)
}

// GenerateAlertForDevice analyses device and creates alert if thresholds exceeded
func (a *App) GenerateAlertForDevice(deviceID uint) (models.Alert, error) {
	if deviceID == 0 {
		return models.Alert{}, fmt.Errorf("invalid device ID")
	}

	dev, err := a.deviceSvc.Get(deviceID)
	if err != nil {
		return models.Alert{}, fmt.Errorf("failed to get device: %w", err)
	}

	alert, err := a.alertSvc.GenerateAlertForDevice(dev)
	if err != nil {
		return models.Alert{}, fmt.Errorf("failed to generate alert: %w", err)
	}

	return alert, nil
}

// formatDuration converts duration to a human-readable string like "1d 2h 3m".
func formatDuration(dur time.Duration) string {
	days := int(dur.Hours()) / 24
	hours := int(dur.Hours()) % 24
	minutes := int(dur.Minutes()) % 60
	if days > 0 {
		return fmt.Sprintf("%dd %dh %dm", days, hours, minutes)
	}
	if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	}
	return fmt.Sprintf("%dm", minutes)
}

// -------- Auth ---------
func (a *App) Register(email, password, name string) (*models.User, error) {
	return a.authSvc.Register(email, password, name)
}

func (a *App) Login(email, password string) (string, error) {
	return a.authSvc.Login(email, password)
}

func (a *App) ValidateToken(token string) (*service.TokenClaims, error) {
	return a.authSvc.ValidateToken(token)
}
