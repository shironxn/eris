package service

import (
	"github.com/shironxn/eris/internal/app/model"
	"github.com/shironxn/eris/internal/infrastructure/repository"
	"github.com/shironxn/eris/internal/infrastructure/util"
)

type UserService interface {
	Login(req model.Login) error
	Register(req model.Register) error
	GetAll() ([]model.User, error)
	GetByID(id uint) (*model.User, error)
	Update(req model.User) error
	Delete(id uint) error
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (u *userService) Login(req model.Login) error {
	user, err := u.repository.GetByEmail(req.Email)
	if err != nil {
		return err
	}

	if err := util.ComparePassword(req.Password, []byte(user.Password)); err != nil {
		return err
	}

	return nil
}

func (u *userService) Register(req model.Register) error {
	password, err := util.HashPassword(req.Password)
	if err != nil {
		return err
	}

	req.Password = string(password)

	return u.repository.Create(req)
}

func (u *userService) GetAll() ([]model.User, error) {
	return u.repository.GetAll()
}

func (u *userService) GetByID(id uint) (*model.User, error) {
	return u.repository.GetByID(id)
}

func (u *userService) Update(req model.User) error {
	user, err := u.repository.GetByID(req.ID)
	if err != nil {
		return err
	}

	return u.repository.Update(req, *user)
}

func (u *userService) Delete(id uint) error {
	return u.repository.Delete(id)
}
