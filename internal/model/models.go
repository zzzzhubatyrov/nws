package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `json:"username"` // gorm:"unique"
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type LoginResponse struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

type NetworkDevice struct {
	ID                int                `gorm:"primaryKey"`
	Hostname          string             `json:"hostname"`   // router.example.com or switch.example.com
	IPAddress         string             `json:"ip_address"` // 192.168.1.1
	Type              string             `json:"type"`       // Router or Switch, etc.
	Vendor            string             `json:"vendor"`     // Cisco
	Model             string             `json:"model"`      // Cisco 881
	Serial            string             `json:"serial"`     // 1234567890
	OSVersion         string             `json:"os_version"` // 10.15.7
	Status            string             `json:"status"`     // up or down, etc.
	LastChecked       time.Time          `json:"last_checked"`
	Connections       []Connection       `json:"connections" gorm:"foreignKey:SourceDeviceIP"`
	Metrics           []Metric           `json:"metrics" gorm:"foreignKey:DeviceID"`
	NetworkInterfaces []NetworkInterface `json:"network_interfaces" gorm:"foreignKey:DeviceID"`
	NetworkLogs       []NetworkLog       `json:"network_logs" gorm:"foreignKey:DeviceID"`
	Configurations    []Configuration    `json:"configurations" gorm:"foreignKey:DeviceID"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

type Connection struct {
	ID                  int           `gorm:"primaryKey"`
	SourceDeviceID      string        `json:"source_device_id"`
	DestinationDeviceID string        `json:"destination_device_id "`
	SourceDeviceIP      string        `json:"source_device_ip"`
	DestinationDeviceIP string        `json:"destination_device_ip"`
	Protocol            string        `json:"protocol"`            // TCP, UDP, ICMP, etc.
	Port                int           `json:"port"`                // 80, 443, etc.
	Status              string        `json:"status"`              // active, disabled, etc.
	Latency             int           `json:"latency"`             // 10ms
	PacketLoss          int           `json:"packet_loss"`         // 10%
	PacketsSent         int           `json:"packets_sent"`        // 1000
	PacketsReceived     int           `json:"packets_received"`    // 1000
	PacketsLost         int           `json:"packets_lost"`        // 100
	PacketsReordered    int           `json:"packets_reordered"`   // 10
	PacketsDuplicated   int           `json:"packets_duplicated"`  // 10
	PacketsCorrupted    int           `json:"packets_corrupted"`   // 10
	PacketsReassembled  int           `json:"packets_reassembled"` // 10
	SourceDevice        NetworkDevice `json:"source_device" gorm:"foreignKey:SourceDeviceIP"`
	DestinationDevice   NetworkDevice `json:"destination_device" gorm:"foreignKey:DestinationDeviceIP"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}

type NetworkTopology struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type TopologyElement struct {
	ID          int             `gorm:"primaryKey"`
	TopologyID  int             `json:"topology_id"`
	DeviceID    int             `json:"device_id"`
	X           float32         `json:"x"`
	Y           float32         `json:"y"`
	CustomStyle string          `json:"custom_style"` // {"color": "#FF0000", "size": 30}
	Topology    NetworkTopology `json:"topology" gorm:"foreignKey:TopologyID"`
	Device      NetworkDevice   `json:"device" gorm:"foreignKey:DeviceID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type NetworkLog struct {
	ID        int           `gorm:"primaryKey"`
	DeviceID  int           `json:"device_id"`
	Device    NetworkDevice `json:"device" gorm:"foreignKey:DeviceID"`
	Timestamp time.Time     `json:"timestamp"`
	Action    string        `json:"action"`
	Message   string        `json:"message"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type NetworkInterface struct {
	ID         int           `gorm:"primaryKey"`
	DeviceID   int           `json:"device_id"`
	Name       string        `json:"name"`        // Ethernet1/1
	IPAddress  string        `json:"ip_address"`  // 192.168.1.1
	Subnet     string        `json:"subnet"`      // 255.255.255.0
	Gateway    string        `json:"gateway"`     // 192.168.1.1
	MACAddress string        `json:"mac_address"` // 00:00:00:00:00:00
	Speed      int           `json:"speed"`       // 1000
	Status     string        `json:"status"`      // up or down, etc.
	Device     NetworkDevice `json:"device" gorm:"foreignKey:DeviceID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type Configuration struct {
	ID        int           `gorm:"primaryKey"`
	DeviceID  int           `json:"device_id"`
	Version   string        `json:"version"` // 1.0.0
	Content   string        `json:"content"` // {"interfaces": {"Ethernet1/1": {"ip": "192.168.1.1", "subnet": "255.255.255.0", "gateway": "192.168.1.1", "mac_address": "00:00:00:00:00:00", "speed": 1000, "status": "up"}}}
	Changes   string        `json:"changes"` // Initial configuration, etc.
	IsActive  bool          `json:"is_active"`
	Device    NetworkDevice `json:"device" gorm:"foreignKey:DeviceID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Metric struct {
	ID        int           `gorm:"primaryKey"`
	DeviceID  int           `json:"device_id"`
	CPU       int           `json:"cpu"`      // 100%
	Memory    int           `json:"memory"`   // 100%
	Network   int           `json:"network"`  // 100%
	Uptime    int           `json:"uptime"`   // 20h 10m
	Interval  string        `json:"interval"` // 1s, 5s, 10s, 30s, 1m, 5m, 10m, 30m, 1h, 2h, 4h, 8h, 12h, 1d
	Status    string        `json:"status"`   // up or down, etc.
	Device    NetworkDevice `json:"device" gorm:"foreignKey:DeviceID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Reports struct {
	ID          int           `gorm:"primaryKey"`
	DeviceID    int           `json:"device_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Status      string        `json:"status"`   // pending, in_progress, completed, failed
	Priority    int           `json:"priority"` // 1-5
	Category    string        `json:"category"` // network, security, performance, etc.
	Tags        string        `json:"tags"`     // tag1, tag2, etc.
	Device      NetworkDevice `json:"device" gorm:"foreignKey:DeviceID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// type NetworkTopologyNode struct {
// 	ID         int             `gorm:"primaryKey"`
// 	Name       string          `json:"name"`
// 	TopologyID int             `json:"topology_id"`
// 	Topology   NetworkTopology `json:"topology" gorm:"foreignKey:TopologyID"`
// 	CreatedAt  time.Time
// 	UpdatedAt  time.Time
// }
