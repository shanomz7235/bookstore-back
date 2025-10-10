package repositories

import (
	// "fmt"

	"github.com/shanomz7235/bookstore-back/config"
	"github.com/shanomz7235/bookstore-back/models"
	// "errors"
)

func AddToCart(cart []models.Items) error {
	result := config.DB.Create(&cart)
	if result.Error != nil{
		return result.Error
	}
	return nil
}

func GetCartItems(id uint) (*models.Carts, error) {
	var cart models.Carts
	result := config.DB.Preload("Items").
		Where("user_id  = ? AND status = ?", id, "active").
		First(&cart)
	if result.Error != nil{
		return nil, result.Error
	}
	return  &cart, nil
}

func GetCart(id uint) *models.Carts {
	var cart models.Carts
	result := config.DB.
		Where("user_id = ? AND status = ?", id, "active").
		First(&cart)

	if result.Error != nil {
		return nil
	}
	// print("already have cart")
	return &cart
}

func CreateCart(cart *models.Carts) error {
	return config.DB.Create(cart).Error
}