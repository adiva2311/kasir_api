package controllers

import (
	"kasir_api/dto"
	"kasir_api/models"
	"kasir_api/repositories"
	"kasir_api/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type CategoryController interface {
	CreateCategory(c *echo.Context) error
	GetCategoryByID(c *echo.Context) error
	UpdateCategory(c *echo.Context) error
	DeleteCategory(c *echo.Context) error
	GetAllCategories(c *echo.Context) error
	GetProductByCategoryID(c *echo.Context) error
}

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
}

// GetProductByCategoryID implements CategoryController.
func (r *CategoryControllerImpl) GetProductByCategoryID(c *echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid category ID: " + err.Error(),
		})
	}

	products, err := r.CategoryService.GetProductByCategoryID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get products by category ID: " + err.Error(),
		})
	}

	ApiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Products retrieved successfully -> Status : " + http.StatusText(http.StatusOK),
		Data:    products,
	}

	return c.JSON(http.StatusOK, ApiResponse)
}

// CreateCategory implements CategoryController.
func (r *CategoryControllerImpl) CreateCategory(c *echo.Context) error {
	userPayload := new(dto.CategoryRequest)
	err := c.Bind(userPayload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request payload : " + err.Error(),
		})
	}

	result, err := r.CategoryService.CreateCategory(models.Category{
		Name:        userPayload.Name,
		Description: userPayload.Description,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create category: " + err.Error(),
		})
	}

	ApiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Category created successfully -> Status : " + http.StatusText(http.StatusCreated),
		Data:    result,
	}

	return c.JSON(http.StatusCreated, ApiResponse)

}

// DeleteCategory implements CategoryController.
func (r *CategoryControllerImpl) DeleteCategory(c *echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid category ID: " + err.Error(),
		})
	}

	err = r.CategoryService.DeleteCategory(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to delete category: " + err.Error(),
		})
	}

	ApiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Category deleted successfully -> Status : " + http.StatusText(http.StatusOK),
	}

	return c.JSON(http.StatusOK, ApiResponse)
}

// GetAllCategories implements CategoryController.
func (r *CategoryControllerImpl) GetAllCategories(c *echo.Context) error {
	categories, err := r.CategoryService.GetAllCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get categories: " + err.Error(),
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Categories retrieved successfully -> Status : " + http.StatusText(http.StatusOK),
		Data:    categories,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

// GetCategoryByID implements CategoryController.
func (r *CategoryControllerImpl) GetCategoryByID(c *echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid category ID: " + err.Error(),
		})
	}

	category, err := r.CategoryService.GetCategoryByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get category: " + err.Error(),
		})
	}

	ApiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Category retrieved successfully -> Status : " + http.StatusText(http.StatusOK),
		Data:    category,
	}

	return c.JSON(http.StatusOK, ApiResponse)
}

// UpdateCategory implements CategoryController.
func (r *CategoryControllerImpl) UpdateCategory(c *echo.Context) error {
	userPayload := new(dto.CategoryRequest)
	err := c.Bind(userPayload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request payload : " + err.Error(),
		})
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid category ID: " + err.Error(),
		})
	}

	result, err := r.CategoryService.UpdateCategory(id, models.Category{
		Name:        userPayload.Name,
		Description: userPayload.Description,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to update category: " + err.Error(),
		})
	}

	ApiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Category updated successfully -> Status : " + http.StatusText(http.StatusOK),
		Data:    result,
	}

	return c.JSON(http.StatusOK, ApiResponse)
}

func NewCategoryController(db *gorm.DB) CategoryController {
	categoryService := services.NewCategoryService(repositories.NewCategoryRepository(db))
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}
