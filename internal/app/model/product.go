package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       int
	Stock       int
	UserID      uint
	CategoryID  uint
	User        User
	Category    Category
}

type ProductResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductCreate struct {
	Name        string `form:"name" binding:"required,min=2,max=100"`
	Description string `form:"description,omitempty"`
	Price       int    `form:"price" binding:"required,min=0"`
	Stock       int    `form:"stock" binding:"required,min=0"`
}

type ProductUpdate struct {
	Name        string `form:"name" binding:"omitempty,min=2,max=100"`
	Description string `form:"description,omitempty"`
	Price       int    `form:"price" binding:"omitempty,min=0"`
	Stock       int    `form:"stock" binding:"omitempty,min=0"`
}
