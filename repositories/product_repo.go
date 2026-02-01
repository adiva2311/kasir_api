package repositories

import (
	"errors"
	"kasir_api/models"

	"gorm.io/gorm"
)

type ProductRepo interface {
	GetAllProducts() ([]models.Product, error)
}

type ProductRepoImpl struct {
	DB *gorm.DB
}

// GetAllProducts implements ProductRepo.
func (p *ProductRepoImpl) GetAllProducts() ([]models.Product, error) {
	var products []models.Product

	result := p.DB.Find(&products).Error
	if result != nil {
		return nil, errors.New(result.Error())
	}

	return products, nil
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &ProductRepoImpl{
		DB: db,
	}
}
