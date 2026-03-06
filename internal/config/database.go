package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/TheDurodola/go-email-service/internal/data/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {

	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using system environment variables")
    }


	host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASSWORD")
    name := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")


    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
        host, user, pass, name, port)
		
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }
    err = db.AutoMigrate(&models.OutgoingEmail{}, &models.EmailTemplate{})
    if err != nil {
        return nil, fmt.Errorf("failed to migrate database: %v", err)
    }

    sqlDB, err := db.DB()
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetConnMaxLifetime(time.Hour)
	
    return db, err
}