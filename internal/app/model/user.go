package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Products []Product
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserUri struct {
	ID uint `uri:"id" binding:"required"`
}

type Login struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type Register struct {
	Name     string `form:"name" binding:"required,min=4,max=30"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=8,max=30"`
}

type UserUpdate struct {
	ID       uint   `uri:"id" binding:"required"`
	Name     string `form:"name" binding:"omitempty,min=4,max=30"`
	Email    string `form:"email" binding:"omitempty,email"`
	Password string `form:"password" binding:"omitempty,min=8,max=30"`
}
