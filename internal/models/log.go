package models

import "time"

type Log struct {
    ID           uint      `gorm:"primaryKey" json:"id"`
    DeviceID     *uint     `json:"deviceId,omitempty"`
    ConnectionID *uint     `json:"connectionId,omitempty"`
    Level        string    `json:"level"` // info/warn/error
    Message      string    `json:"message"`
    Time         time.Time `json:"time"`
}
