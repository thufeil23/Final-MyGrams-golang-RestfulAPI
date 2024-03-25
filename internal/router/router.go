package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/handler"
)

// UserRouter is a interface for user router
type UserRouter interface {
	MountRoutes()
}

// userRouter struct is a struct that implements the UserRouter interface
type userRouter struct {
	ur       *gin.RouterGroup
	handler handler.UserHandler
}

// NewUserRouter is a constructor for user router
func NewUserRouter(ur *gin.RouterGroup, handler handler.UserHandler) UserRouter {
	return &userRouter{ur: ur, handler: handler}
}

// MountRoutes is a method for user router
func (u *userRouter) MountRoutes() {
	u.ur.GET("/users/:id", u.handler.GetUsersByID)
	u.ur.POST("/users", u.handler.RegisterUser)
	u.ur.GET("/users", u.handler.GetUsers)
}