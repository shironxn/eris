package model

import "gorm.io/gorm"

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
