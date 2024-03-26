package service

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/helper"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/models"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/repository"
)

// UserService is a interface for user service
type UserService interface {
	// activity
	RegisterUser(ctx context.Context, user models.User) (models.User, error)
	LoginUser(ctx context.Context, loginUser models.User) (models.User, error)
	// token
	GenerateToken(ctx context.Context, user models.User) (token string, err error)
	// users
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) (models.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

// userService is a struct that implements the UserService interface
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService is a constructor for user service
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// RegisterUser creates a new user in the database
func (u *userService) RegisterUser(ctx context.Context, user models.User) (models.User, error) {
	// // parse DOB
	// parsedDOB, err := time.Parse("2006-01-02", user.DoB.String())
	// if err != nil {
	// 	return models.User{}, err
	// }
	// user.DoB = parsedDOB
	
	// check if user already exists
	// existingUser, err := u.userRepo.GetUsersByID(ctx, user.ID)
	// if err != nil {
	// 	return models.User{}, err
	// }
	// if existingUser.ID != uuid.Nil {
	// 	return models.User{}, errors.New("user already exists")
	// }

	// check if email already exists
	// existingEmail, err := u.userRepo.GetUsersByEmail(ctx, user.Email)
	// if err != nil {
	// 	return models.User{}, err
	// }
	// if existingEmail.ID != uuid.Nil {
	// 	return models.User{}, errors.New("email already exists")
	// }

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

// LoginUser checks if the user exists and the password is correct
func (u *userService) LoginUser(ctx context.Context, loginUser models.User) (models.User, error) {
	// check if user by email exists
	user, err := u.userRepo.GetUsersByEmail(ctx, loginUser.Email)
	if err != nil {
		return models.User{}, err
	}
	if user.ID == uuid.Nil {
		return models.User{}, errors.New("Email is not registered")
	}

	// check if password is correct
	isPasswordCorrect := helper.CheckPasswordHash(loginUser.Password, user.Password)
	if !isPasswordCorrect {
		return models.User{}, errors.New("Password is incorrect") 
	}

	return user, nil
}

// GenerateToken generates a JWT token for the user
func (u *userService) GenerateToken(ctx context.Context, user models.User) (token string, err error) {
	now := time.Now()
	// Generate JWT token
	claim := models.StandardClaim{
		Jti: uuid.New().String(),
		Iss: os.Getenv("JWT_ISSUER"),
		Sub: user.ID.String(),
		Aud: os.Getenv("JWT_AUDIENCE"),
		Exp: uint64(now.Add(time.Hour * 24).Unix()),
		Nbf: uint64(now.Unix()),
		Iat: uint64(now.Unix()),
	}

	// Generate token
	userClaims := models.AccessClaim{
		StandardClaim: claim,
		UserID:        user.ID,
		Username:      user.Username,
		DoB:           user.DoB,
	}

	token, err = helper.GenerateToken(userClaims)
	return
}

// GetUsers gets all users from the database
func (u *userService) GetUsers(ctx context.Context) ([]models.User, error) {
	return u.userRepo.GetUsers(ctx)
}

// GetUserByID gets a user by their ID from the database
func (u *userService) GetUserByID(ctx context.Context, id uuid.UUID) (models.User, error) {
	// check if user exists
	user, err := u.userRepo.GetUsersByID(ctx, id)
	if err != nil {
		return models.User{}, err
	}
	if user.ID == uuid.Nil {
		return models.User{}, errors.New("user not found")
	}

	return u.userRepo.GetUsersByID(ctx, id)
}

// UpdateUser updates a user in the database
func (u *userService) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	// check if user exists
	userID := user.ID
	existingUser, err := u.userRepo.GetUsersByID(ctx, userID)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	if existingUser.ID == uuid.Nil {
		return models.User{}, errors.New("user not found")
	}

	// update user
	existingUser.Username = user.Username
	existingUser.Email = user.Email
	existingUser.Password = user.Password
	existingUser.DoB = user.DoB

	updatedUser, err := u.userRepo.UpdateUsers(ctx, existingUser)

	return updatedUser, err
}

// DeleteUser deletes a user from the database
func (u *userService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	// check if user exists
	existingUser, err := u.userRepo.GetUsersByID(ctx, id)
	if err != nil {
		return err
	}
	if existingUser.ID == uuid.Nil {
		return errors.New("user not found")
	}

	return u.userRepo.DeleteUsers(ctx, id)
}