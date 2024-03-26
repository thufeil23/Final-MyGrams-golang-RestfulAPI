package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/thufeil23/go-digitalent-24/project-mygram/internal/helper"
	"gorm.io/gorm"
)

// User struct is a model
type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"column:id;type:uuid;primary_key;" json:"id"`
	Username string    `gorm:"unique;not null" json:"username"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Password string    `gorm:"not null" json:"password"`
	DoB      time.Time `gorm:"column:dob;not null;type:date" json:"dob"`
	Photos   []Photo   `gorm:"foreignKey:UserID;references:ID" json:"photos"`
	Comments []Comment `gorm:"foreignKey:UserID;references:ID" json:"comments"`
	Socials  []Social  `gorm:"many2many:user_socials" json:"socials"`
}

// BeforeSave is a gorm hook
func (u *User) BeforeSave(db *gorm.DB) (err error) {
	// Hash password
	u.Password, err = helper.HashPassword(u.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	// // Validate input
	// input := User{
	// 	Email:    u.Email,
	// 	Password: u.Password,
	// }
	// // Create a new validator instance
	// validate := validator.New()

	// // Validate the User struct
	// err = validate.Struct(input)
	// if err != nil {
	// 	// Validation failed, handle the error
	// 	return fmt.Errorf("validation failed: %w", err)
	// }

	return
}

// BeforeCreate is a gorm hook
func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	// Generate UUID
	u.ID = uuid.New()
	// Hash password
	u.Password, err = helper.HashPassword(u.Password)
	if err != nil {
		return err
	}

	// Validate hashed password
	// if !helper.CheckPasswordHash(u.Password, u.Password) {
	// 	return fmt.Errorf("hashed password is not valid")
	// }

	// // Validate input
	// input := User{
	// 	Username: u.Username,
	// 	Email:    u.Email,
	// 	Password: u.Password,
	// 	DoB:      u.DoB,
	// }
	
	// // Create a new validator instance
	// validate := validator.New()

	// // Validate the User struct
	// err = validate.Struct(input)
	// if err != nil {
	// 	// Validation failed, handle the error
	// 	return fmt.Errorf("validation failed: %w", err)
	// }

	return
}
