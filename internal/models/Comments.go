package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Comment struct is a model
type Comment struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Message string    `gorm:"not null" json:"message"`
	UserID  uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	PhotoID uuid.UUID `gorm:"type:uuid;not null" json:"photo_id"`
}

// BeforeCreate is a gorm hook
func (c *Comment) BeforeCreate(db *gorm.DB) (err error) {
	// Generate UUID
	c.ID = uuid.New()

	return
}
