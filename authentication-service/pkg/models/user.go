package models

import "time"

type User struct {
	ID           uint
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	FirstName    string
	LastName     string
	CreatedAt    time.Time
}

// validatePasswordHash validates that the password hashes match
func (u *User) ValidatePasswordHash(hash string) bool {
	return u.PasswordHash == hash
}
