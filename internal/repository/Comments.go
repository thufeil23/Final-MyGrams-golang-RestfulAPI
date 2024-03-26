package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/infrastructure"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/models"
	"gorm.io/gorm"
)

// CommentRepository is a interface for comment repository
type CommentRepository interface {
	GetCommentsByPhotoID(ctx context.Context, photoID uuid.UUID) ([]models.Comment, error)
	GetCommentsByID(ctx context.Context, id uuid.UUID) (models.Comment, error)
	CreateComments(ctx context.Context, comment models.Comment) (models.Comment, error)
	UpdateComments(ctx context.Context, comment models.Comment) (models.Comment, error)
	DeleteComments(ctx context.Context, id uuid.UUID) error
}

// commentRepository is a struct that implements the CommentRepository interface
type commentRepository struct {
	db *infrastructure.GormConnect
}

// NewCommentRepository is a constructor for CommentRepository
func NewCommentRepository(db *infrastructure.GormConnect) CommentRepository {
	return &commentRepository{db: db}
}

// GetCommentsByPhotoID getting comment by photo id from database
func (c commentRepository) GetCommentsByPhotoID(ctx context.Context, photoID uuid.UUID) ([]models.Comment, error) {
	db := c.db.GetConnection()
	comments := []models.Comment{}
	if err := db.
	Where("photo_id = ?", photoID).
	Preload("User").
	Preload("Photo").
	Find(&comments).
	Error; err != nil {
		return nil, errors.New("No comments found")
	}

	return comments, nil
}

// GetCommentsByID getting comment by comment id from database
func (c commentRepository) GetCommentsByID(ctx context.Context, id uuid.UUID) (models.Comment, error) {
	db := c.db.GetConnection()
	comment := models.Comment{}
	if err := db.
		Where("id = ?", id).
		Preload("User").
		Preload("Photo").
		First(&comment).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Comment{}, errors.New("comment not found")
		}
		return models.Comment{}, err
	}
	return comment, nil
}

// CreateComments creating comment in database
func (c commentRepository) CreateComments(ctx context.Context, comment models.Comment) (models.Comment, error) {
	db := c.db.GetConnection()
	if err := db.
	Create(&comment).
	Error; err != nil {
		return models.Comment{}, errors.New("failed to create comment")
	}
	return comment, nil
}

// UpdateComments updating comment in database
func (c commentRepository) UpdateComments(ctx context.Context, comment models.Comment) (models.Comment, error) {
	db := c.db.GetConnection()
	if err := db.
	Save(&comment).
	Error; err != nil {
		return models.Comment{}, errors.New("failed to update comment")
	}
	return comment, nil
}

// DeleteComments deleting comment from database
func (c commentRepository) DeleteComments(ctx context.Context, id uuid.UUID) error {
	db := c.db.GetConnection()
	if err := db.
	Where("id = ?", id).
	Delete(&models.Comment{}).
	Error; err != nil {
		return errors.New("failed to delete comment")
	}
	return nil
}