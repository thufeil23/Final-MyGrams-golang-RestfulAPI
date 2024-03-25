package repository

import (
	"context"

	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/infrastructure"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/models"
)

// PhotoRepository is a interface for photo repository
type PhotoRepository interface {
	GetPhotos(ctx context.Context) ([]models.Photo, error)
}

// photoRepository is a struct that implements the PhotoRepository interface
type photoRepository struct {
	db *infrastructure.GormConnect
}

// NewPhotoRepository is a constructor for PhotoRepository
func NewPhotoRepository(db *infrastructure.GormConnect) PhotoRepository {
	return &photoRepository{db: db}
}

// GetPhotos getting all photos from database
func (p photoRepository) GetPhotos(ctx context.Context) ([]models.Photo, error) {
	db := p.db.GetConnection()
	photo := []models.Photo{}
	if err := db.Find(&photo).Error; err != nil {
		return nil, err
	}
	return photo, nil
}