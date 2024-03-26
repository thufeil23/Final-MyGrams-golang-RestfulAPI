package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/handler"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/infrastructure"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/models"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/repository"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/router"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err)
	}

	// Auto-migrate tables
	db, err := infrastructure.NewGormConnect()
	dbConnect, err := infrastructure.ConnectDB()
	dbConnect.AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.Social{})

	r := gin.Default()
	userGroup := r.Group("/users")

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userRouter := router.NewUserRouter(userGroup, userHandler)

	userRouter.MountRoutes()
	r.Run(os.Getenv("SERVER_ADDRESS"))
}
