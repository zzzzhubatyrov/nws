package models

import "time"

type Device struct {
    CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
    ID       uint   `gorm:"primaryKey" json:"id"`
    Name     string `json:"name"`
    Type     string `json:"type"`
    IP       string `json:"ip"`
    Status   string `json:"status"`
    CPU      int    `json:"cpu"`
    Memory   int    `json:"memory"`
    Uptime   string `json:"uptime"`
    X        int    `json:"x"`
    Y        int    `json:"y"`
    Vendor   string `json:"vendor"`
    Model    string `json:"model"`
}
