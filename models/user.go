package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Name     string `gorm:"not null" json:"name"`
	Role     string `gorm:"not null;default:user" json:"role"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
