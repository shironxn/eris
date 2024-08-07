package repository

import (
	"github.com/shironxn/eris/internal/app/model"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(req model.CategoryCreate) error
	GetAll() ([]model.Category, error)
	GetByID(id uint) (*model.Category, error)
	Update(req model.CategoryUpdate, Category model.Category) error
	Delete(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (p *categoryRepository) Create(req model.CategoryCreate) error {
	category := model.Category{
		Name: req.Name,
	}
	return p.db.Create(&category).Error
}

func (p *categoryRepository) GetAll() ([]model.Category, error) {
	var categorys []model.Category
	if err := p.db.Find(&categorys).Error; err != nil {
		return nil, err
	}
	return categorys, nil
}

func (p *categoryRepository) GetByID(id uint) (*model.Category, error) {
	var category model.Category
	if err := p.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (p *categoryRepository) Update(req model.CategoryUpdate, category model.Category) error {
	return p.db.Model(&category).Updates(&req).Error
}

func (p *categoryRepository) Delete(id uint) error {
	return p.db.Delete(&model.Category{}, id).Error
}
