package models

import "gorm.io/gorm"

type Carts struct {
	gorm.Model
	UserID uint    `json:"user_id"`
	Items  []Items `gorm:"foreignKey:CartID" json:"items"`
	Status string  `json:"status"`
}

type Items struct {
	gorm.Model
	CartID  uint    `json:"cart_id"`
	BookID   uint    `json:"book_id"`
	Quantity uint    `json:"quantity"`
	Price    float64 `json:"price"`
}

type CartResponse struct {
	ID     uint              `json:"id"`
	UserID uint              `json:"user_id"`
	Status string            `json:"status"`
	Items  []ItemResponse `json:"items"`
}

type ItemResponse struct {
	ID       uint    `json:"id"`
	BookID   uint    `json:"book_id"`
	Quantity uint    `json:"quantity"`
	Price    float64 `json:"price"`
}
