package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Social struct is a model
type Social struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Name   string    `gorm:"not null" json:"name"`
	URL    string    `gorm:"not null" json:"url"`
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
}

// BeforeCreate is a gorm hook
func (s *Social) BeforeCreate(db *gorm.DB) (err error) {
	// Generate UUID
	s.ID = uuid.New()

	return
}
