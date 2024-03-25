package helper

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword bcrypt hash password
func HashPassword(p string) (string, error) {
	salt := 8
	password := []byte(p)

	hash, err := bcrypt.GenerateFromPassword(password, salt)
	if err != nil {
		fmt.Println("Error hashing password: ", err)
		return "", err
	}

	return string(hash), nil
}

// CheckPasswordHash validate password
func CheckPasswordHash(p string, h string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	return err == nil
}
