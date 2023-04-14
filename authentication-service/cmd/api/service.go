package main

import (
	"errors"
	"time"
)

var ErrUserExists = errors.New("user already exists")
var ErrUserNotFound = errors.New("user does not exist")

type user struct {
	email        string
	username     string
	passwordHash string
	firstName    string
	lastName     string
	createdDate  time.Time
	role         int
}

type AuthenticationService interface {
	Signup(email string, username string, passwordHash string, firstname string, lastname string) (user, error)
	Signin(email string, passwordHash string) (string, error)
}

type authenticationService struct{}

// Temp data store
var userList = make([]user, 0)

func (as *authenticationService) Signup(email string, username string, passwordHash string, firstname string, lastname string) (user, error) {
	userObject, err := as.addUserObject(email, username, passwordHash, firstname, lastname, 0)
	if err != nil {
		return user{}, err
	}
	return userObject, nil
}

func (as *authenticationService) Signin(email, passwordHash string) (string, error) {
	return "pass", nil
}

// getUserObject returns the found user object from store
func (as *authenticationService) getUserObject(email string) (user, error) {
	for _, user := range userList {
		if user.email == email {
			return user, ErrUserNotFound
		}
	}
	return user{}, nil
}

// addUserObject creates a new unique user in the store
func (as *authenticationService) addUserObject(email string, username string, passwordHash string, firstname string, lastname string, role int) (user, error) {
	newUser := user{
		email:        email,
		username:     username,
		passwordHash: passwordHash,
		firstName:    firstname,
		lastName:     lastname,
		createdDate:  time.Now(),
		role:         role,
	}
	for _, u := range userList {
		if u.email == email {
			return user{}, ErrUserExists
		}
	}
	userList = append(userList, newUser)
	return newUser, nil
}

// validatePasswordHash validates that the password hashes match
func (u *user) validatePasswordHash(hash string) bool {
	return u.passwordHash == hash
}
