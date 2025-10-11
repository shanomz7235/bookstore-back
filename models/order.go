package models

import (
	"time"

	// "github.com/shanomz7235/bookstore-back/models"
	"gorm.io/gorm"
)

type OrderItem struct {
	ID       uint `gorm:"primaryKey"`
	BookID   uint
	OrderID  uint
	Quantity uint
	Price    float64
}

type Order struct {
	gorm.Model
	UserID uint        `gorm:"foreignKey:UserID"`
	User   User        `gorm:"foreignKey:UserID"`
	Items  []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	Total  float64     `json:"total"`
	Status string      `json:"status"` //  paid, shipping, shipped
}

type OrderResponse struct {
	ID          uint                `json:"id"`
	UserID      uint                `json:"user_id"`
	UserName    string              `json:"user_name"`
	UserAddress string              `json:"user_address"`
	Total       float64             `json:"total"`
	Status      string              `json:"status"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	Items       []OrderItemResponse `json:"items"`
}

type OrderItemResponse struct {
	BookID   uint    `json:"book_id"`
	Quantity uint    `json:"quantity"`
	Price    float64 `json:"price"`
}
