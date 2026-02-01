package dto

import "kasir_api/models"

type CategoryResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToCategoryResponse(category *models.Category) CategoryResponse {
	return CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}

type UpdateCategoryResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToUpdateCategoryResponse(category *models.Category) UpdateCategoryResponse {
	return UpdateCategoryResponse{
		Name:        category.Name,
		Description: category.Description,
	}
}
