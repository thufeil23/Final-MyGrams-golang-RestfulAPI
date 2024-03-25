package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/models"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/repository"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/service"
)

// UserHandler is a interface for user handler
type UserHandler interface {
	RegisterUser(ctx *gin.Context)
	GetUsersByID(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
}

// userHandler struct is a struct that implements the UserHandler interface
type userHandler struct {
	userSvc service.UserService
	userRep repository.UserRepository
}

// NewUserHandler is a constructor for user handler services
func NewUserHandler(userSvc service.UserService) UserHandler {
	return &userHandler{userSvc: userSvc}
}

// NewRepoUserHandler is a constructor for user handler repository
// NewRepoUserHandler is a constructor for user handler repository
func NewRepoUserHandler(userRep repository.UserRepository) repository.UserRepository {
	return userRep
}


// GetUsers getting all users from database
func (h *userHandler) GetUsers(ctx *gin.Context) {
	users, err := h.userRep.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
}


// RegisterUser creates a new user in the database
func (h *userHandler) RegisterUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	createdUser, err := h.userSvc.RegisterUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, createdUser)
}

// GetUsersByID getting user by their id from database
func (h *userHandler) GetUsersByID(ctx *gin.Context) {
	// get user id from path parameter
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if id == uuid.Nil || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id",
		})
		return
	}
	user, err := h.userSvc.GetUserByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, user)
}