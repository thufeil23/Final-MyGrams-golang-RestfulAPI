package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/infrastructure"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/models"
	"gorm.io/gorm"
)

// UserRepository is a interface for user repository
type UserRepository interface {
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUsersByID(ctx context.Context, id uuid.UUID) (models.User, error)
	GetUsersByEmail(ctx context.Context, email string) (models.User, error)
	CreateUsers(ctx context.Context, user models.User) (models.User, error)
	UpdateUsers(ctx context.Context, user models.User) (models.User, error)
	DeleteUsers(ctx context.Context, id uuid.UUID) error
}

// userRepository is a struct that implements the UserRepository interface
type userRepository struct {
	db *infrastructure.GormConnect
}

// NewUserRepository is a constructor for UserRepository
func NewUserRepository(db *infrastructure.GormConnect) UserRepository {
	return &userRepository{db: db}
}

// GetUsers getting all users from database
func (u userRepository) GetUsers(ctx context.Context) ([]models.User, error) {
	db := u.db.GetConnection()
	user := []models.User{}
	if err := db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUsersByID getting user by their id from database
func (u userRepository) GetUsersByID(ctx context.Context, id uuid.UUID) (models.User, error) {
	db := u.db.GetConnection()
	user := models.User{}
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.User{}, err
		}
		return models.User{}, err
	}
	return user, nil
}

// GetUsersByEmail getting user by their email from database
func (u userRepository) GetUsersByEmail(ctx context.Context, email string) (models.User, error) {
	db := u.db.GetConnection()
	user := models.User{}
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.User{}, err
		}
		return models.User{}, err
	}
	return user, nil
}

// CreateUsers creates a new user in the database
func (u userRepository) CreateUsers(ctx context.Context, user models.User) (models.User, error) {
	db := u.db.GetConnection()
	if err := db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}


// UpdateUsers updates a user in the database
func (u userRepository) UpdateUsers(ctx context.Context, user models.User) (models.User, error) {
	db := u.db.GetConnection()
	if err := db.Save(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// DeleteUsers deletes a user from the database
func (u userRepository) DeleteUsers(ctx context.Context, id uuid.UUID) error {
	db := u.db.GetConnection()
	if err := db.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}