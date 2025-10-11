package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Name     string `gorm:"not null" json:"name"`
	Address  string `gorm:"not null" json:"address"`
	Role     string `gorm:"not null;default:user" json:"role"`
}

type LoginUser struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserResponse struct {
	ID        uint
	Name      string
	Address   string
	Email     string
	CreatedAt time.Time
	UpdateAt  time.Time
}
