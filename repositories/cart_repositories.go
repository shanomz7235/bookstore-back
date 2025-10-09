package repositories

import (
	// "fmt"

	"github.com/shanomz7235/bookstore-back/config"
	"github.com/shanomz7235/bookstore-back/models"
	// "errors"
)

func AddToCart(cart []models.CartItem) error {
	result := config.DB.Create(&cart)
	if result.Error != nil{
		return result.Error
	}
	return nil
}

func GetCartItems() ([]models.CartItem, error) {
	var cart []models.CartItem
	result := config.DB.Find(&cart)
	if result.Error != nil{
		return nil, result.Error
	}
	return  cart, nil
}

// func SaveCart(cartItems []models.CartItem, userID uint) error {
// 	tx := config.DB.Begin()

// 	cart := models.Carts{
// 		UserID: userID,
// 	}
	

// 	if err := tx.Create(&cart).Error; err != nil{
// 		tx.Rollback()
// 		return fmt.Errorf("failed to create cart: %w", err)
// 	}
// 	print("create cart")

// 	for i := range cartItems{
// 		cartItems[i].CartID = cart.ID
// 	}

// 	if err := tx.Create(&cartItems).Error; err != nil{
// 		tx.Rollback()
// 		return fmt.Errorf("failed to create cart items: %w", err)
// 	}


// 	return tx.Commit().Error
// }
