package models

import (
	"gorm.io/gorm"
	"time"
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
	UserID uint
	Items  []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	Total  float64     `json:"total"`
	Status string      `json:"status"` // pending, paid, shipped, canceled
}


type OrderResponse struct {
	ID     uint              `json:"id"`
	UserID uint              `json:"user_id"`
	Total  float64           `json:"total"`
	Status string            `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Items  []OrderItemResponse `json:"items"`
}

type OrderItemResponse struct {
	BookID   uint    `json:"book_id"`
	Quantity uint    `json:"quantity"`
	Price    float64 `json:"price"`
}

