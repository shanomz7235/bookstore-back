package services

import (
	"errors"
	"github.com/shanomz7235/bookstore-back/models"
	"github.com/shanomz7235/bookstore-back/repositories"
	"github.com/shanomz7235/bookstore-back/utils"
)

func RegisterUser(user *models.User) error {
	println("1", user.Email, "2", user.Name, "3", user.Password)
	if user.Email == "" || user.Password == "" || user.Name == "" || user.Address == "" {
		return errors.New("some fields are  missing")
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return repositories.RegisterUser(user)
}

func GetUser(id uint) (*models.UserResponse, error) {
	user, err := repositories.GetUser(id)
	if err != nil {
		return nil, err
	}
	return convertUserToResponse(user)
}

func GetUsers() ([]models.UserResponse, error) {
	users, err := repositories.GetUsers()
	if err != nil {
		return nil, err
	}

	return convertUsersToResponse(users)
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

	return utils.GenerateJWT(userDB)

}

func convertUsersToResponse(users []models.User) ([]models.UserResponse, error) {
	var userRes []models.UserResponse
	for _, user := range users {
		userRes = append(userRes, models.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Address:   user.Address,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdateAt:  user.UpdatedAt,
		})
	}
	return userRes, nil
}

func convertUserToResponse(user *models.User) (*models.UserResponse, error) {
	newUser := new(models.UserResponse)

	newUser.ID = user.ID
	newUser.Name = user.Name
	newUser.Address = user.Address
	newUser.Email = user.Email
	newUser.CreatedAt = user.CreatedAt
	newUser.UpdateAt = user.UpdatedAt
	return newUser, nil
}
