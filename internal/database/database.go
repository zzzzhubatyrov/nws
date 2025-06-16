package database

import (
    "log"
    "sync"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var (
    db   *gorm.DB
    once sync.Once
)

// GetDB returns singleton *gorm.DB instance
func GetDB() *gorm.DB {
    once.Do(func() {
        var err error
        db, err = gorm.Open(sqlite.Open("trace.db"), &gorm.Config{})
        if err != nil {
            log.Fatalf("failed to connect database: %v", err)
        }
    })
    return db
}
