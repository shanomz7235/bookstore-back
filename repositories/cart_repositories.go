package repositories

import (
	// "fmt"

	"errors"

	"github.com/shanomz7235/bookstore-back/config"
	"github.com/shanomz7235/bookstore-back/models"
	"gorm.io/gorm"
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
	result := config.DB.
		Preload("Items", func(db *gorm.DB) *gorm.DB {
			return db.Order("id ASC")
		}).
		Where("user_id  = ? AND status = ?", id, "active").
		First(&cart)
	if result.Error != nil{
		return nil, result.Error
	}
	return  &cart, nil
}

func CheckCart(id uint) *models.Carts {
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

func UpdateItem(itemID uint, newItem *models.Items, cartId uint)  error{
	result := config.DB.Model(&models.Items{}).
		Where("id = ? AND cart_id = ?", itemID, cartId).
		Updates(newItem)
	if result.Error != nil{
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}

    return  nil

}

func DeleteItem(itemID uint, cartID uint)  error{
	result := config.DB.
		Where("id = ? AND cart_id = ?", itemID, cartID).
		Delete(&models.Items{})
	if result.Error != nil{
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}

    return  nil

}

func UpdateCartStatus(cart *models.Carts) error {
	result := config.DB.Model(&cart).Updates(cart)

    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return errors.New("no rows affected")
    }

    return nil
}