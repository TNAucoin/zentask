package main

import (
	"errors"
	"time"

	"github.com/tnaucoin/zentask/authentication-service/pkg/models"
	"github.com/tnaucoin/zentask/authentication-service/pkg/token"
	"gorm.io/gorm"
)

var ErrUserExists = errors.New("user already exists")
var ErrUserNotFound = errors.New("user does not exist")
var ErrFailedLogin = errors.New("invalid user credentials")

type AuthenticationService interface {
	Signup(email string, username string, passwordHash string, firstname string, lastname string) (models.User, error)
	Signin(email string, passwordHash string) (string, error)
	RefreshToken(token string) (string, error)
}

type authenticationService struct {
	DB *gorm.DB
}

func Init(db *gorm.DB) authenticationService {
	return authenticationService{db}
}

// Temp data store
// TODO: add DB for users
var userList = make([]models.User, 0)

func (as *authenticationService) Signup(email string, username string, passwordHash string, firstname string, lastname string) (models.User, error) {
	userObject, err := as.addUserObject(email, username, passwordHash, firstname, lastname, 0)
	if err != nil {
		return models.User{}, err
	}
	return userObject, nil
}

func (as *authenticationService) Signin(email, passwordHash string) (string, error) {
	var user models.User
	for _, u := range userList {

		if u.Email == email {
			user = u
		}
	}
	// return error if we didn't find the user by email
	if user.Email == "" {
		return "", ErrFailedLogin
	}
	//check password sha matches
	isValid := user.ValidatePasswordHash(passwordHash)

	if !isValid {
		return "", ErrFailedLogin
	}

	// Generate the token
	token, err := token.GenerateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

// RefreshToken given a valid token, this will issue a new refreshed token
func (as *authenticationService) RefreshToken(userToken string) (string, error) {
	t, err := token.RefreshToken(userToken)
	if err != nil {
		return "", err
	}
	return t, nil
}

// getUserObject returns the found user object from store
func (as *authenticationService) getUserObject(email string) (models.User, error) {
	for _, user := range userList {
		if user.Email == email {
			return user, ErrUserNotFound
		}
	}
	return models.User{}, nil
}

// addUserObject creates a new unique user in the store
func (as *authenticationService) addUserObject(email string, username string, passwordHash string, firstname string, lastname string, role int) (models.User, error) {
	var user models.User
	result := as.DB.Where("email = ?", email).First(&user)
	if result.RowsAffected > 0 {
		return models.User{}, ErrUserExists
	}

	user = models.User{
		Email:        email,
		Username:     username,
		PasswordHash: passwordHash,
		FirstName:    firstname,
		LastName:     lastname,
		CreatedDate:  time.Now(),
		Role:         0,
	}

	if result := as.DB.Create(&user); result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
