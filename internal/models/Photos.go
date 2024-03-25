package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Photo struct is a model
type Photo struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Title    string    `gorm:"not null" json:"title"`
	PhotoURL string    `gorm:"not null" json:"url"`
	Caption  string    `gorm:"not null" json:"caption"`
	UserID   uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
}

// BeforeCreate is a gorm hook
func (p *Photo) BeforeCreate(db *gorm.DB) (err error) {
	// Generate UUID
	p.ID = uuid.New()

	return
}
