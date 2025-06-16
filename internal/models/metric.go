package models

import "time"

type Metric struct {
    ID           uint      `gorm:"primaryKey" json:"id"`
    DeviceID     *uint     `json:"deviceId,omitempty"`
    ConnectionID *uint     `json:"connectionId,omitempty"`
    Name         string    `json:"name"`
    Value        float64   `json:"value"`
    Time         time.Time `json:"time"`
}
