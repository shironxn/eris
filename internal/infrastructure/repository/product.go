package repository

import (
	"github.com/shironxn/eris/internal/app/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(req model.ProductCreate) error
	GetAll() ([]model.Product, error)
	GetByID(id uint) (*model.Product, error)
	Update(req model.ProductUpdate, product model.Product) error
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

func (p *productRepository) Create(req model.ProductCreate) error {
	product := model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		UserID:      req.UserID,
		CategoryID:  req.CategoryID,
	}
	return p.db.Create(&product).Error
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
	if err := p.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productRepository) Update(req model.ProductUpdate, product model.Product) error {
	return p.db.Model(&product).Updates(&req).Error
}

func (p *productRepository) Delete(id uint) error {
	return p.db.Delete(&model.Product{}, id).Error
}
