package service

import (
	"github.com/shironxn/eris/internal/app/model"
	"github.com/shironxn/eris/internal/infrastructure/repository"
)

type CategoryService interface {
	Create(req model.Category) error
	GetAll() ([]model.Category, error)
	GetByID(id uint) (*model.Category, error)
	Update(req model.Category) error
	Delete(id uint) error
}

type categoryService struct {
	repository repository.CategoryRepository
}

func NewCategoryService(repository repository.CategoryRepository) CategoryService {
	return &categoryService{
		repository: repository,
	}
}

func (c *categoryService) Create(req model.Category) error {
	return c.repository.Create(req)
}

func (c *categoryService) GetAll() ([]model.Category, error) {
	return c.repository.GetAll()
}

func (c *categoryService) GetByID(id uint) (*model.Category, error) {
	return c.repository.GetByID(id)
}

func (c *categoryService) Update(req model.Category) error {
	category, err := c.repository.GetByID(req.ID)
	if err != nil {
		return err
	}

	return c.repository.Update(req, *category)
}

func (c *categoryService) Delete(id uint) error {
	return c.repository.Delete(id)
}
