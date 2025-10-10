package models

import "gorm.io/gorm"

type Carts struct {
	gorm.Model
	UserID uint    `json:"user_id"`
	Items  []Items `gorm:"foreignKey:Cart_ID" json:"items"`
	Status string  `json:"status"`
}

type Items struct {
	gorm.Model
	Cart_ID  uint    `json:"cart_id"`
	BookID   uint    `json:"bookid"`
	Quantity uint    `json:"quantity"`
	Price    float64 `json:"price"`
}
