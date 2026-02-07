package controllers

import (
	"kasir_api/dto"
	"kasir_api/models"
	"kasir_api/repositories"
	"kasir_api/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type ProductController interface {
	GetAllProducts(c *echo.Context) error
	GetProductByID(c *echo.Context) error
	CreateProduct(c *echo.Context) error
	UpdateProduct(c *echo.Context) error
	DeleteProduct(c *echo.Context) error
	SearchProductsByName(c *echo.Context) error
}

type ProductControllerImpl struct {
	ProductService services.ProductService
}

// SearchProductsByName implements ProductController.
func (p *ProductControllerImpl) SearchProductsByName(c *echo.Context) error {
	name := c.QueryParam("name")
	products, err := p.ProductService.SearchProductsByName(name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to search products: " + err.Error(),
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Successfully searched products",
		Data:    products,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

// CreateProduct implements ProductController.
func (p *ProductControllerImpl) CreateProduct(c *echo.Context) error {
	userPayload := new(models.Product)
	err := c.Bind(userPayload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to bind request body : " + err.Error(),
		})
	}

	product, err := p.ProductService.CreateProduct(&models.Product{
		Name:       strings.ToTitle(userPayload.Name),
		Price:      userPayload.Price,
		Stock:      userPayload.Stock,
		CategoryID: userPayload.CategoryID,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create product: " + err.Error(),
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Successfully created product",
		Data:    product,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

// DeleteProduct implements ProductController.
func (p *ProductControllerImpl) DeleteProduct(c *echo.Context) error {
	// Get ID from URL param
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid product ID",
		})
	}

	err = p.ProductService.DeleteProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to delete product: " + err.Error(),
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Successfully deleted product",
	}

	return c.JSON(http.StatusOK, apiResponse)
}

// UpdateProduct implements ProductController.
func (p *ProductControllerImpl) UpdateProduct(c *echo.Context) error {
	// Get ID from URL param
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid product ID",
		})
	}

	userPayload := new(dto.CreateUpdateProductResponse)
	if err := c.Bind(userPayload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to bind request body",
		})
	}

	product, err := p.ProductService.UpdateProduct(id, &models.Product{
		Name:       strings.ToTitle(userPayload.Name),
		Price:      userPayload.Price,
		Stock:      userPayload.Stock,
		CategoryID: userPayload.CategoryID,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to update product: " + err.Error(),
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Successfully updated product",
		Data:    product,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

// GetProductByID implements ProductController.
func (p *ProductControllerImpl) GetProductByID(c *echo.Context) error {
	// Get ID from URL param
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid product ID -> Status : " + http.StatusText(http.StatusBadRequest),
		})
	}

	product, err := p.ProductService.GetProductByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ApiResponse{
			Status:  http.StatusNotFound,
			Message: "Product not found -> Status : " + http.StatusText(http.StatusNotFound),
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Successfully retrieved product -> Status : " + http.StatusText(http.StatusOK),
		Data:    product,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

// GetAllProducts implements ProductController.
func (p *ProductControllerImpl) GetAllProducts(c *echo.Context) error {
	products, err := p.ProductService.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get products: " + err.Error(),
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Successfully retrieved products -> Status : " + http.StatusText(http.StatusOK),
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
