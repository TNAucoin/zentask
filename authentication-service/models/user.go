package models

import "time"

type User struct {
	Email        string
	Username     string
	PasswordHash string
	FirstName    string
	LastName     string
	CreatedDate  time.Time
	Role         int
}

// validatePasswordHash validates that the password hashes match
func (u *User) ValidatePasswordHash(hash string) bool {
	return u.PasswordHash == hash
}
