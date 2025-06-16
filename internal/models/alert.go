package models

type Alert struct {
    ID       uint   `gorm:"primaryKey" json:"id"`
    Device   string `json:"device"`
    Type     string `json:"type"`
    Severity string `json:"severity"`
    Time     string `json:"time"`
    Status   string `json:"status"`
    Title    string `json:"title"`
    Message  string `json:"message"`
}
