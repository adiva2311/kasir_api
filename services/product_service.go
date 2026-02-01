package services

import (
	"errors"
	"kasir_api/dto"
	"kasir_api/models"
	"kasir_api/repositories"
)

type ProductService interface {
	// CreateProduct(product *models.Product) (dto.ProductResponse, error)
	// UpdateProduct(id int, product *models.Product) (dto.ProductResponse, error)
	GetAllProducts() ([]models.Product, error)
	// DeleteProduct(id int) error
	GetProductByID(id int) (*dto.ProductResponse, error)
}

type ProductServiceImpl struct {
	ProductRepo repositories.ProductRepo
}

// GetProductByID implements ProductService.
func (p *ProductServiceImpl) GetProductByID(id int) (*dto.ProductResponse, error) {
	product, err := p.ProductRepo.GetProductByID(id)
	if err != nil {
		return nil, errors.New("failed to get product: " + err.Error())
	}

	result := dto.ProductResponse{
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Category: product.Category.Name,
	}

	return &result, nil
}

// GetAllProducts implements ProductService.
func (p *ProductServiceImpl) GetAllProducts() ([]models.Product, error) {
	products, err := p.ProductRepo.GetAllProducts()
	if err != nil {
		return nil, errors.New("failed to get products: " + err.Error())
	}

	var Products []models.Product
	for _, product := range products {
		Products = append(Products, models.Product{
			ID:         product.ID,
			Name:       product.Name,
			Price:      product.Price,
			Stock:      product.Stock,
			CategoryID: product.CategoryID,
		})
	}

	return Products, nil
}

func NewProductService(productRepo repositories.ProductRepo) ProductService {
	return &ProductServiceImpl{
		ProductRepo: productRepo,
	}
}
