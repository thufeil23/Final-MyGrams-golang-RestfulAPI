package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/models"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/repository"
)

// UserServiceInterface is a interface for user service
type UserServiceInterface interface {
	RegisterUser(ctx context.Context, user models.User) (models.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (models.User, error)
}

// UserService is a struct that implements the UserService interface
type UserService struct {
	userRepo repository.UserRepository
}

// NewUserService is a constructor for user service
func NewUserService(userRepo repository.UserRepository) UserServiceInterface {
	return &UserService{userRepo: userRepo}
}

// RegisterUser creates a new user in the database
func (u *UserService) RegisterUser(ctx context.Context, user models.User) (models.User, error) {
	// check if user already exists
	existingUser, err := u.userRepo.GetUsersByID(ctx, user.ID)
	if err != nil {
		return models.User{}, err
	}
	if existingUser.ID != uuid.Nil {
		return models.User{}, errors.New("user already exists")
	}

	// create user
	createdUser, err := u.userRepo.CreateUsers(ctx, user)
	if err != nil {
		return models.User{}, err
	}
	if createdUser.ID == uuid.Nil {
		return models.User{}, errors.New("failed to create user")
	}

	return createdUser, nil
}

// GetUserByID getting user by their id from database
func (u *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (models.User, error) {
	user, err := u.userRepo.GetUsersByID(ctx, id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}