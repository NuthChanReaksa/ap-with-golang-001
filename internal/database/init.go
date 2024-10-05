package database

import (
	"homework1/internal/models"
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func InitDB(dsn string) *gorm.DB {

    db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    // Perform automatic migration
	err = db.AutoMigrate(&model.Product{}, &model.User{})
    if err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }

    return db;
}