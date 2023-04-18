package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tnaucoin/zentask/authentication-service/pkg/models"
)

// TODO: this should be hidden
var jwtKey = []byte("some_secret_key")

var ErrTokenCreationFailed = errors.New("failed to create a jwt")
var ErrInvalidToken = errors.New("invalid token")

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", ErrTokenCreationFailed
	}

	return tokenString, nil

}

func ValidateTokenWithClaims(token string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, ErrInvalidToken
	}
	if !tkn.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}

func RefreshToken(token string) (string, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", ErrInvalidToken
	}
	if !tkn.Valid {
		return "", ErrInvalidToken
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString(jwtKey)
	if err != nil {
		return "", ErrTokenCreationFailed
	}
	return tokenString, nil
}
