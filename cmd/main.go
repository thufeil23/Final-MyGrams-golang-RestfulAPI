package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/infrastructure"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err)
	}

	dbConnect, err := infrastructure.NewGormConnect()
	if err != nil {
		fmt.Printf("Error connecting to database: %s\n", err)
	}

	// Auto-migrate tables
	db := dbConnect.GetConnection()

	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.Social{})

	r := gin.Default()
	r.Run(os.Getenv("SERVER_ADDRESS"))
}
