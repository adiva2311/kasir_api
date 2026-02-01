package dto

import "kasir_api/models"

type ProductResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

func ToProductResponse(product *models.Product) *ProductResponse {
	return &ProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}
}

type GetProductDetailResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
	Category string `json:"category,omitempty"`
}

func ToGetProductDetailResponse(product *models.Product) *GetProductDetailResponse {
	return &GetProductDetailResponse{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Category: product.Category.Name,
	}
}

type CreateUpdateProductResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	CategoryID int    `json:"category_id"`
}

func ToCreateUpdateProductResponse(product *models.Product) *CreateUpdateProductResponse {
	return &CreateUpdateProductResponse{
		ID:         product.ID,
		Name:       product.Name,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryID: product.CategoryID,
	}
}
