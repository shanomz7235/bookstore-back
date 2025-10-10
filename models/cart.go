package models

import "gorm.io/gorm"

type Carts struct {
    gorm.Model
    UserID uint       `json:"user_id"`
    Cart_ID uint
    Items  []Items `gorm:"foreignKey:CartID" json:"items"` 
}

type Items struct {
    gorm.Model
    CartID   uint   `json:"cart_id"` 
    BookID   uint    `json:"book_id"`
    Quantity uint    `json:"quantity"`
    Price    float64 `json:"price"`
}
