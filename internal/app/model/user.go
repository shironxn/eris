package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Products []Product
}

type Login struct {
	Email    string
	Password string
}

type Register struct {
	Name     string
	Email    string
	Password string
}
