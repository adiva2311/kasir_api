package controllers

import (
	"kasir_api/dto"
	"kasir_api/repositories"
	"kasir_api/services"
	"net/http"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type TransactionController interface {
	CreateTransaction(c *echo.Context) error
}

type TransactionControllerImpl struct {
	TransactionService services.TransactionService
}

// CreateTransaction implements TransactionController.
func (t *TransactionControllerImpl) CreateTransaction(c *echo.Context) error {
	payload := new(dto.CreateTransactionReq)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ApiResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to bind request body: " + err.Error(),
		})
	}

	transaction, err := t.TransactionService.CreateTransaction(*payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create transaction: " + err.Error(),
		})
	}

	apiResponse := dto.ApiResponse{
		Status:  http.StatusOK,
		Message: "Successfully created transaction",
		Data:    transaction,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func NewTransactionController(db *gorm.DB) TransactionController {
	transactionService := services.NewTransactionService(repositories.NewTransactionRepo(db))
	return &TransactionControllerImpl{
		TransactionService: transactionService,
	}
}
