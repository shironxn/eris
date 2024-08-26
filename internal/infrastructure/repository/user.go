package repository

import (
	"fmt"

	"github.com/shironxn/eris/internal/app/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(req model.Register) error
	GetAll() ([]model.User, error)
	GetByID(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Update(req model.UserUpdate, user model.User) error
	Delete(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(req model.Register) error {
	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	return u.db.Create(&user).Error
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
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	fmt.Println(user)
	return &user, nil
}

func (u *userRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) Update(req model.UserUpdate, user model.User) error {
	return u.db.Model(&user).Updates(&req).Error
}

func (u *userRepository) Delete(user *model.User) error {
	return u.db.Delete(&user).Error
}
