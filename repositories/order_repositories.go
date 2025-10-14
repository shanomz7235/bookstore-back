package repositories

import (
	// "fmt"

	"errors"

	"github.com/shanomz7235/bookstore-back/config"
	"github.com/shanomz7235/bookstore-back/models"
	"gorm.io/gorm"
)

func CreateOrder(order *models.Order) error {
	result := config.DB.Create(order)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetOrders(userID uint) ([]models.Order, error) {
	var orders []models.Order
	result := config.DB.
		Preload("Items", func(db *gorm.DB) *gorm.DB {
			return db.Order("id ASC")
		}).
		Preload("User").
		Order("id ASC").
		Where("user_id = ?", userID).
		Order("id ASC").
		Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func GetOrderByID(orderID uint) (*models.Order, error) {
	var order models.Order
	result := config.DB.First(&order, orderID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func UpdateOrderStatus(order *models.Order) error {
	result := config.DB.Model(&models.Order{}).
		Where("id = ?", order.ID).
		Update("status", order.Status)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	result := config.DB.
		Preload("Items", func(db *gorm.DB) *gorm.DB {
			return db.Order("id ASC")
		}).
		Preload("User").
		Order("id ASC").
		Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}
