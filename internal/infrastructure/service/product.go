package service

import (
	"github.com/shironxn/eris/internal/app/model"
	"github.com/shironxn/eris/internal/infrastructure/repository"
)

type ProductService interface {
	Create(req model.Product) error
	GetAll() ([]model.Product, error)
	GetByID(id uint) (*model.Product, error)
	Update(req model.Product) error
	Delete(id uint) error
}

type productService struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) ProductService {
	return &productService{
		repository: repository,
	}
}

func (p *productService) Create(req model.Product) error {
	return p.repository.Create(req)
}

func (p *productService) GetAll() ([]model.Product, error) {
	return p.repository.GetAll()
}

func (p *productService) GetByID(id uint) (*model.Product, error) {
	return p.repository.GetByID(id)
}

func (p *productService) Update(req model.Product) error {
	product, err := p.repository.GetByID(req.ID)
	if err != nil {
		return err
	}

	return p.repository.Update(req, *product)
}

func (p *productService) Delete(id uint) error {
	return p.repository.Delete(id)
}
