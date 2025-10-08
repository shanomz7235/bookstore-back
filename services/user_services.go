package services

import (
	"errors"
	"github.com/shanomz7235/bookstore-back/models"
	"github.com/shanomz7235/bookstore-back/repositories"
	"github.com/shanomz7235/bookstore-back/utils"

)

func RegisterUser(user *models.User) error {
	println("1", user.Email, "2", user.Name, "3", user.Password)
	if user.Email == "" || user.Password == "" || user.Name == "" {
		return errors.New("some fields are  missing")
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return repositories.RegisterUser(user)
}

func GetUser(id uint) (*models.User, error) {
	user, err := repositories.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUsers() ([]models.User, error) {
	return repositories.GetUsers()
}

func LoginUser(user *models.LoginUser) (string, error) {
	if user.Email == "" || user.Password == "" {
		return "", errors.New("email and password cant be empty")
	}

	userDB, err := repositories.GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}

	if err := utils.CheckPassword(userDB.Password, user.Password); err != nil {
		return "", err
	}

	return utils.GenerateJWT(user.Email)

}
