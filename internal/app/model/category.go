package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string
}

type CategoryResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryCreate struct {
	Name string `form:"name" binding:"required"`
}

type CategoryUpdate struct {
	Name string `form:"name" binding:"required"`
}
