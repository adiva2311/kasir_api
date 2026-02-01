package controllers

import (
	"kasir_api/dto"
	"kasir_api/repositories"
	"kasir_api/services"
	"net/http"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type ProductController interface {
	GetAllProducts(c *echo.Context) error
}

type ProductControllerImpl struct {
	ProductService services.ProductService
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
