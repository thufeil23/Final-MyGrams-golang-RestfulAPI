package infrastructure

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GormConnect is a wrapper for gorm.DB
type GormConnect struct {
	master *gorm.DB
}

// gormConnection is an interface for GormConnect
type gormConnection interface {
	GetConnection() *gorm.DB
}

// NewGormConnect is a constructor for GormConnect
func NewGormConnect() (*GormConnect, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	return &GormConnect{master: db}, nil
}

// ConnectDB is a helper function to connect to database
func ConnectDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	return db, nil
}

// GetConnection is a getter for gorm.DB
func (g *GormConnect) GetConnection() *gorm.DB {
	return g.master
}
