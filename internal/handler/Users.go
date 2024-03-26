package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/models"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/service"
)

// UserHandler is a interface for user handler
type UserHandler interface {
	// activity
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	// users
	GetUsers(ctx *gin.Context)
	GetUsersByID(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

// userHandler struct is a struct that implements the UserHandler interface
type userHandler struct {
	userSvc service.UserService
}

// NewUserHandler is a constructor for user handler services
func NewUserHandler(userSvc service.UserService) UserHandler {
	return &userHandler{userSvc: userSvc}
}



// RegisterUser creates a new user in the database
func (h *userHandler) RegisterUser(ctx *gin.Context) {
	user := models.User{}
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
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
		"data":    createdUser,
	})
}

// LoginUser authenticates a user and returns an access token
func (h *userHandler) LoginUser(ctx *gin.Context) {
	loginUser := models.User{}
	err := ctx.ShouldBindJSON(&loginUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := h.userSvc.LoginUser(ctx, loginUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	Token, err := h.userSvc.GenerateToken(ctx, user)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"data":    user,
		"token":   Token,
	})
}

// GetUsers getting all users from database
func (h *userHandler) GetUsers(ctx *gin.Context) {
	users, err := h.userSvc.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
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

// UpdateUser updates a user in the database
func (h *userHandler) UpdateUser(ctx *gin.Context) {
	// get user id from path parameter
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if id == uuid.Nil || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id",
		})
		return
	}
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	
	// update user
	updatedUser, err := h.userSvc.UpdateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to update user: " + err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user updated successfully",
		"data":    updatedUser,
	})
}

// DeleteUser deletes a user from the database
func (h *userHandler) DeleteUser(ctx *gin.Context) {
	// get user id from path parameter
	idStr := ctx.Param("id")
	id, err := uuid.Parse(idStr)
	if id == uuid.Nil || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id",
		})
		return
	}
	err = h.userSvc.DeleteUser(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
}

