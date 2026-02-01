package dto

import "kasir_api/models"

type ProductResponse struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
	CategoryID int     `json:"category_id"`
	Category   string  `json:"category,omitempty"`
}

func ToProductResponse(product *models.Product) *ProductResponse {
	return &ProductResponse{
		ID:         product.ID,
		Name:       product.Name,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryID: product.CategoryID,
		Category:   product.Category.Name,
	}
}
