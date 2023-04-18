package models

import "time"

type User struct {
	Email        string `gorm:"primaryKey"`
	PasswordHash string
	FirstName    string
	LastName     string
	CreatedDate  time.Time
}

// validatePasswordHash validates that the password hashes match
func (u *User) ValidatePasswordHash(hash string) bool {
	return u.PasswordHash == hash
}
