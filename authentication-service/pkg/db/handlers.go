package db

import (
	"errors"

	"github.com/tnaucoin/zentask/authentication-service/pkg/models"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{db}
}

func (h *Handler) FindUser(email string) (models.User, error) {
	var user models.User
	if result := h.DB.Where("email = ?", email).First(&user); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil

}

func (h *Handler) CheckIfUserExists(email string) (bool, error) {
	user := models.User{}
	result := h.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func (h *Handler) CreateUser(user models.User) (models.User, error) {
	if result := h.DB.Create(&user); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
