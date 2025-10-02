package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string  `gorn:"not null" json:"title"`
	Author string  `gorm:"not null" json:"author"`
	Price  float64 `gorm:"not null" json:"price"`
	Stock  uint    `gorm:"not null" json:"stock"`
}

type BookUpdate struct {
	gorm.Model
	Title  *string  `json:"title"`
	Author *string  `json:"author"`
	Price  *float64 `json:"price"`
	Stock  *uint    `json:"stock"`
}
