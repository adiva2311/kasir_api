package services

import (
	"errors"
	"kasir_api/dto"
	"kasir_api/models"
	"kasir_api/repositories"
)

type CategoryService interface {
	CreateCategory(req models.Category) (dto.CategoryResponse, error)
	UpdateCategory(id int, req models.Category) (dto.UpdateCategoryResponse, error)
	DeleteCategory(id int) error
	GetCategoryByID(id int) (dto.CategoryResponse, error)
	GetAllCategories() ([]dto.CategoryResponse, error)
	GetProductByCategoryID(categoryID int) ([]dto.ProductResponse, error)
}

type CategoryServiceImpl struct {
	CategoryRepo repositories.CategoryRepository
}

// GetProductByCategoryID implements CategoryService.
func (c *CategoryServiceImpl) GetProductByCategoryID(categoryID int) ([]dto.ProductResponse, error) {
	products, err := c.CategoryRepo.GetProductByCategoryID(categoryID)
	if err != nil {
		return nil, errors.New("failed to get products by category ID: " + err.Error())
	}

	var ProductCategoryResponse []dto.ProductResponse
	for _, product := range products {
		ProductCategoryResponse = append(ProductCategoryResponse, *dto.ToProductResponse(&product))
	}
	return ProductCategoryResponse, nil
}

// CreateCategory implements CategoryService.
func (c *CategoryServiceImpl) CreateCategory(req models.Category) (dto.CategoryResponse, error) {
	// Check if category already exists
	findCategory, err := c.CategoryRepo.FindByName(req.Name)
	if err == nil && findCategory != nil {
		return dto.CategoryResponse{}, errors.New("category already exists: " + findCategory.Name)
	}

	categoryRequest := models.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	err = c.CategoryRepo.CreateCategory(categoryRequest)
	if err != nil {
		return dto.CategoryResponse{}, errors.New("failed to create category: " + err.Error())
	}
	return dto.ToCategoryResponse(&categoryRequest), nil
}

// DeleteCategory implements CategoryService.
func (c *CategoryServiceImpl) DeleteCategory(id int) error {
	err := c.CategoryRepo.DeleteCategory(id)
	if err != nil {
		return errors.New("failed to delete category: " + err.Error())
	}
	return nil
}

// GetAllCategories implements CategoryService.
func (c *CategoryServiceImpl) GetAllCategories() ([]dto.CategoryResponse, error) {
	categories, err := c.CategoryRepo.GetAllCategories()
	if err != nil {
		return nil, errors.New("failed to get categories: " + err.Error())
	}

	var CategoryResponse []dto.CategoryResponse
	for _, category := range categories {
		CategoryResponse = append(CategoryResponse, dto.ToCategoryResponse(&category))
	}
	return CategoryResponse, nil
}

// GetCategoryByID implements CategoryService.
func (c *CategoryServiceImpl) GetCategoryByID(id int) (dto.CategoryResponse, error) {
	category, err := c.CategoryRepo.GetCategoryByID(id)
	if err != nil {
		return dto.CategoryResponse{}, errors.New("failed to get category: " + err.Error())
	}
	return dto.ToCategoryResponse(category), nil
}

// UpdateCategory implements CategoryService.
func (c *CategoryServiceImpl) UpdateCategory(id int, req models.Category) (dto.UpdateCategoryResponse, error) {
	requestCategory := models.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	err := c.CategoryRepo.UpdateCategory(id, requestCategory)
	if err != nil {
		return dto.UpdateCategoryResponse{}, errors.New("failed to update category: " + err.Error())
	}
	return dto.ToUpdateCategoryResponse(&requestCategory), nil
}

func NewCategoryService(categoryRepo repositories.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepo: categoryRepo,
	}
}
