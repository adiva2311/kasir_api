package repositories

import (
	"errors"
	"kasir_api/models"

	"gorm.io/gorm"
)

type ProductRepo interface {
	CreateProduct(product *models.Product) error
	GetAllProducts() ([]models.Product, error)
	UpdateProduct(id int, product *models.Product) error
	DeleteProduct(id int) error
	GetProductByID(id int) (*models.Product, error)
	GetProductByName(name string) (*models.Product, error)
	SearchProductsByName(name string) ([]models.Product, error)
}

type ProductRepoImpl struct {
	DB *gorm.DB
}

// SearchProductsByName implements ProductRepo.
func (p *ProductRepoImpl) SearchProductsByName(name string) ([]models.Product, error) {
	var products []models.Product

	result := p.DB.Where("name ILIKE ?", "%"+name+"%").Find(&products)
	if result.Error != nil {
		return nil, errors.New("product not found")
	}

	return products, nil
}

// CreateProduct implements ProductRepo.
func (p *ProductRepoImpl) CreateProduct(product *models.Product) error {
	result := p.DB.Create(product).Error
	if result != nil {
		return errors.New(result.Error())
	}

	return nil
}

// DeleteProduct implements ProductRepo.
func (p *ProductRepoImpl) DeleteProduct(id int) error {
	result := p.DB.Where("id = ?", id).Delete(&models.Product{})
	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}

	return result.Error
}

// GetProductByID implements ProductRepo.
func (p *ProductRepoImpl) GetProductByID(id int) (*models.Product, error) {
	var product models.Product

	// result := p.DB.Where("id = ?", id).First(&product)
	result := p.DB.Preload("Category").Where("id = ?", id).First(&product)
	if result.Error != nil {
		return &models.Product{}, errors.New("product not found")
	}

	return &product, nil
}

// GetProductByName implements ProductRepo.
func (p *ProductRepoImpl) GetProductByName(name string) (*models.Product, error) {
	var product *models.Product

	result := p.DB.Where("name = ?", name).First(&product)
	if result.Error != nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

// UpdateProduct implements ProductRepo.
func (p *ProductRepoImpl) UpdateProduct(id int, product *models.Product) error {
	result := p.DB.Where("id = ?", id).Updates(product)
	if result.Error != nil {
		return errors.New(result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
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
