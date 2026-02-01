package controllers

import (
	"kasir_api/dto"
	"kasir_api/repositories"
	"kasir_api/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type ProductController interface {
	GetAllProducts(c *echo.Context) error
	GetProductByID(c *echo.Context) error
}

type ProductControllerImpl struct {
	ProductService services.ProductService
}

// GetProductByID implements ProductController.
func (p *ProductControllerImpl) GetProductByID(c *echo.Context) error {
	// Get ID from URL param
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "Invalid product ID",
		})
	}

	product, err := p.ProductService.GetProductByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ApiResponse{
			Status:  http.StatusText(http.StatusNotFound),
			Message: "Product not found",
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusText(http.StatusOK),
		Message: "Successfully retrieved product",
		Data:    product,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

// GetAllProducts implements ProductController.
func (p *ProductControllerImpl) GetAllProducts(c *echo.Context) error {
	products, err := p.ProductService.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: "Failed to get products: " + err.Error(),
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusText(http.StatusOK),
		Message: "Successfully retrieved products",
		Data:    products,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func NewProductController(db *gorm.DB) ProductController {
	service := services.NewProductService(repositories.NewProductRepo(db))
	return &ProductControllerImpl{
		ProductService: service,
	}
}
