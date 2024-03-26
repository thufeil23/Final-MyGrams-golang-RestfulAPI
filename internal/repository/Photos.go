package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/infrastructure"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/models"
	"gorm.io/gorm"
)

// PhotoRepository is a interface for photo repository
type PhotoRepository interface {
	GetPhotosByID(ctx context.Context, id uuid.UUID) (models.Photo, error)
	GetPhotosByUserID(ctx context.Context, userID uuid.UUID) ([]models.Photo, error)
	CreatePhotos(ctx context.Context, photo models.Photo) (models.Photo, error)
	UpdatePhotos(ctx context.Context, photo models.Photo) (models.Photo, error)
	DeletePhotos(ctx context.Context, id uuid.UUID) error
}

// photoRepository is a struct that implements the PhotoRepository interface
type photoRepository struct {
	db *infrastructure.GormConnect
}

// NewPhotoRepository is a constructor for PhotoRepository
func NewPhotoRepository(db *infrastructure.GormConnect) PhotoRepository {
	return &photoRepository{db: db}
}

// GetPhotosByID getting photo by photo id from database
func (p photoRepository) GetPhotosByID(ctx context.Context, id uuid.UUID) (models.Photo, error) {
	db := p.db.GetConnection()
	photo := models.Photo{}
	if err := db.
		Where("id = ?", id).
		Preload("User").
		First(&photo).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Photo{}, errors.New("photo not found")
		}
		return models.Photo{}, err
	}
	return photo, nil
}

// GetPhotosByUserID getting photo by user id from database
func (p photoRepository) GetPhotosByUserID(ctx context.Context, userID uuid.UUID) ([]models.Photo, error) {
	db := p.db.GetConnection()
	photos := []models.Photo{}
	if err := db.
		Where("user_id = ?", userID).
		Preload("User").
		Find(&photos).
		Error; err != nil {
		return nil, errors.New("user does not have any photos")
	}
	return photos, nil
}

// CreatePhotos creates a photo in the database
func (p photoRepository) CreatePhotos(ctx context.Context, photo models.Photo) (models.Photo, error) {
	db := p.db.GetConnection()
	if err := db.
		Create(&photo).
		Error; err != nil {
		return models.Photo{}, errors.New("failed to create photo")
	}
	return photo, nil
}

// UpdatePhotos updates a photo in the database
func (p photoRepository) UpdatePhotos(ctx context.Context, photo models.Photo) (models.Photo, error) {
	db := p.db.GetConnection()
	if err := db.
		Save(&photo).
		Error; err != nil {
		return models.Photo{}, errors.New("failed to update photo")
	}
	return photo, nil
}

// DeletePhotos deletes a photo from the database
func (p photoRepository) DeletePhotos(ctx context.Context, id uuid.UUID) error {
	db := p.db.GetConnection()
	if err := db.
		Where("id = ?", id).
		Delete(&models.Photo{}).
		Error; err != nil {
		return errors.New("failed to delete photo")
	}
	return nil
}