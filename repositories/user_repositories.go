package repositories

import (
	"github.com/shanomz7235/bookstore-back/config"
	"github.com/shanomz7235/bookstore-back/models"
)

func RegisterUser(user *models.User) error {
	result := config.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUser(id uint) ( *models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUsers() ([] models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func GetUserByEmail(email string) (*models.LoginUser, error) {
	var user models.LoginUser
	result := config.DB.Model(&models.User{}).
		Select("email", "password", "role", "id").
		Where("email = ?", email).
		First(&user)
    
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}