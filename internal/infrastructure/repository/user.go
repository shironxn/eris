package repository

import (
	"github.com/shironxn/eris/internal/app/model"
	"gorm.io/gorm"
)

type User interface {
	Create(req model.Register) error
	GetAll() ([]model.User, error)
	GetByID(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Update(req model.User, user model.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) User {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(req model.Register) error {
	return u.db.Create(&req).Error
}

func (u *userRepository) GetAll() ([]model.User, error) {
	var users []model.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	if err := u.db.Find(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.db.Find(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) Update(req model.User, user model.User) error {
	if err := u.db.Model(&user).Updates(&req).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Delete(id uint) error {
	return u.db.Delete(&model.User{}, id).Error
}
