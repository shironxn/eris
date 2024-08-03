package repository

import (
	"github.com/shironxn/eris/internal/app/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(req model.Product) error
	GetAll() ([]model.Product, error)
	GetByID(id uint) (*model.Product, error)
	Update(req model.Product, product model.Product) error
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (p *productRepository) Create(req model.Product) error {
	return p.db.Create(&req).Error
}

func (p *productRepository) GetAll() ([]model.Product, error) {
	var products []model.Product
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *productRepository) GetByID(id uint) (*model.Product, error) {
	var product model.Product
	if err := p.db.Find(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productRepository) Update(req model.Product, product model.Product) error {
	return p.db.Model(&product).Updates(&req).Error
}

func (p *productRepository) Delete(id uint) error {
	return p.db.Delete(&model.Product{}, id).Error
}
