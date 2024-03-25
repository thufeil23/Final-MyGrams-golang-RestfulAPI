package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/helper"
	"gorm.io/gorm"
)

// User struct is a model
type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Username string    `gorm:"unique;not null" json:"username"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Password string    `gorm:"not null" json:"password"`
	DoB      time.Time `gorm:"not null" json:"dob"`
	Photos   []Photo   `gorm:"foreignKey:UserID;references:ID" json:"photos"`
	Comments []Comment `gorm:"foreignKey:UserID;references:ID" json:"comments"`
	Socials  []Social  `gorm:"many2many:user_socials" json:"socials"`
}


// BeforeCreate is a gorm hook
func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	// Hash password
	u.Password, err = helper.HashPassword(u.Password)
	if err != nil {
		return err
	}
	// Generate UUID
	u.ID = uuid.New()

	// // Validate input
	// input := User{
	// 	Username: u.Username,
	// 	Email:    u.Email,
	// 	Password: string(hashedPassword),
	// 	DoB:      u.DoB,
	// }
	// if _, err = govalidator.ValidateStruct(input); err != nil {
	// 	return
	// }

	return
}
