package models

type Connection struct {
    ID          uint   `gorm:"primaryKey" json:"id"`
    Source      uint   `json:"source"`
    Target      uint   `json:"target"`
    Status      string `json:"status"`
    Bandwidth   string `json:"bandwidth"`
    Latency     float64 `json:"latency"`
    Type        string `json:"type"`
    Name        string `json:"name"`
    Description string `json:"description"`
}
