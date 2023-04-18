package main

import (
	"errors"
	"time"

	"github.com/tnaucoin/zentask/authentication-service/pkg/db"
	"github.com/tnaucoin/zentask/authentication-service/pkg/models"
	"github.com/tnaucoin/zentask/authentication-service/pkg/token"
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
	DBHandler *db.Handler
}

func Init(db *db.Handler) authenticationService {
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
	user, err := as.DBHandler.FindUser(email)
	if err != nil {
		return "", ErrUserNotFound
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

// addUserObject creates a new unique user in the store
func (as *authenticationService) addUserObject(email string, username string, passwordHash string, firstname string, lastname string, role int) (models.User, error) {
	var user models.User

	// check to see if this email already exists
	userExists, err := as.DBHandler.CheckIfUserExists(email)
	if err != nil {
		return models.User{}, err
	}

	if userExists {
		return models.User{}, ErrUserExists
	}

	// create the user and add them to the DB
	user = models.User{
		Email:        email,
		PasswordHash: passwordHash,
		FirstName:    firstname,
		LastName:     lastname,
		CreatedDate:  time.Now(),
	}

	result, err := as.DBHandler.CreateUser(user)

	if err != nil {
		return models.User{}, err
	}

	return result, nil
}
