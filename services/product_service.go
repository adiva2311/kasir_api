package services

import (
	"errors"
	"kasir_api/dto"
	"kasir_api/models"
	"kasir_api/repositories"
)

type ProductService interface {
	CreateProduct(productReq *models.Product) (*dto.CreateUpdateProductResponse, error)
	UpdateProduct(id int, productReq *models.Product) (*dto.CreateUpdateProductResponse, error)
	GetAllProducts() ([]dto.ProductResponse, error)
	DeleteProduct(id int) error
	GetProductByID(id int) (*dto.GetProductDetailResponse, error)
	SearchProductsByName(name string) ([]dto.GetProductDetailResponse, error)
}

type ProductServiceImpl struct {
	ProductRepo repositories.ProductRepo
}

// SearchProductsByName implements ProductService.
func (p *ProductServiceImpl) SearchProductsByName(name string) ([]dto.GetProductDetailResponse, error) {
	products, err := p.ProductRepo.SearchProductsByName(name)
	if err != nil {
		return nil, errors.New("failed to search products: " + err.Error())
	}

	var result []dto.GetProductDetailResponse
	for _, product := range products {
		result = append(result, dto.GetProductDetailResponse{
			ID:       product.ID,
			Name:     product.Name,
			Price:    product.Price,
			Stock:    product.Stock,
			Category: product.Category.Name,
		})
	}

	return result, nil
}

// CreateProduct implements ProductService.
func (p *ProductServiceImpl) CreateProduct(productReq *models.Product) (*dto.CreateUpdateProductResponse, error) {
	// Check if product with the same name already exists
	findProduct, err := p.ProductRepo.GetProductByName(productReq.Name)
	if err == nil {
		return &dto.CreateUpdateProductResponse{}, errors.New("product with name " + findProduct.Name + " already exists")
	}

	request := &models.Product{
		Name:       productReq.Name,
		Price:      productReq.Price,
		Stock:      productReq.Stock,
		CategoryID: productReq.CategoryID,
	}

	err = p.ProductRepo.CreateProduct(request)
	if err != nil {
		return &dto.CreateUpdateProductResponse{}, errors.New("failed to create product: " + err.Error())
	}

	return dto.ToCreateUpdateProductResponse(request), nil
}

// DeleteProduct implements ProductService.
func (p *ProductServiceImpl) DeleteProduct(id int) error {
	err := p.ProductRepo.DeleteProduct(id)
	if err != nil {
		return errors.New("failed to delete product: " + err.Error())
	}
	return nil
}

// UpdateProduct implements ProductService.
func (p *ProductServiceImpl) UpdateProduct(id int, productReq *models.Product) (*dto.CreateUpdateProductResponse, error) {
	// Check if product with the same name already exists
	findProduct, err := p.ProductRepo.GetProductByName(productReq.Name)
	if err == nil {
		return &dto.CreateUpdateProductResponse{}, errors.New("product with name " + findProduct.Name + " already exists")
	}

	request := &models.Product{
		Name:       productReq.Name,
		Price:      productReq.Price,
		Stock:      productReq.Stock,
		CategoryID: productReq.CategoryID,
	}

	err = p.ProductRepo.UpdateProduct(id, request)
	if err != nil {
		return &dto.CreateUpdateProductResponse{}, errors.New("failed to update product: " + err.Error())
	}

	return dto.ToCreateUpdateProductResponse(request), nil
}

// GetProductByID implements ProductService.
func (p *ProductServiceImpl) GetProductByID(id int) (*dto.GetProductDetailResponse, error) {
	product, err := p.ProductRepo.GetProductByID(id)
	if err != nil {
		return nil, errors.New("failed to get product: " + err.Error())
	}

	result := dto.GetProductDetailResponse{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Category: product.Category.Name,
	}

	return &result, nil
}

// GetAllProducts implements ProductService.
func (p *ProductServiceImpl) GetAllProducts() ([]dto.ProductResponse, error) {
	products, err := p.ProductRepo.GetAllProducts()
	if err != nil {
		return nil, errors.New("failed to get products: " + err.Error())
	}

	var Products []dto.ProductResponse
	for _, product := range products {
		Products = append(Products, dto.ProductResponse{
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
