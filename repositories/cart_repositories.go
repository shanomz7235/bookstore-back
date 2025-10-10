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

func GetCartItems() ([]models.Items, error) {
	var cart []models.Items
	result := config.DB.Find(&cart)
	if result.Error != nil{
		return nil, result.Error
	}
	return  cart, nil
}

func SaveCart(cartItems []models.Items, userID uint, ) (error) {
    // สร้าง Cart ใหม่
    cart := models.Carts{
        UserID: userID,
        Items:  cartItems,
    }
    
    // บันทึกลง DB (จะบันทึกทั้ง Cart และ CartItem ที่เชื่อมด้วย)
    result := config.DB.Create(&cart)
    if result.Error != nil {
        return result.Error
    }
    
    return nil
}
