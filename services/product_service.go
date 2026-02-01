package services

import (
	"errors"
	"kasir_api/models"
	"kasir_api/repositories"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
}

type ProductServiceImpl struct {
	ProductRepo repositories.ProductRepo
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
